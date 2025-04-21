package concurrency

import (
	"context"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestKeyedLocker_SerialExecutionSameKey(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestKeyedLocker_SerialExecutionSameKey")

	locker := NewKeyedLocker[string](100 * time.Millisecond)

	var result []int
	var mu sync.Mutex

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := locker.Do(context.Background(), "key1", func() {
				time.Sleep(10 * time.Millisecond)
				mu.Lock()
				defer mu.Unlock()
				result = append(result, i)
			})

			assert.IsNil(err)
		}(i)
	}
	wg.Wait()

	assert.Equal(5, len(result))
}

func TestKeyedLocker_ParallelExecutionDifferentKeys(t *testing.T) {
	locker := NewKeyedLocker[string](100 * time.Millisecond)

	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := "key" + strconv.Itoa(i)
			locker.Do(context.Background(), key, func() {
				time.Sleep(50 * time.Millisecond)
			})
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)

	if elapsed > 100*time.Millisecond {
		t.Errorf("parallel execution took too long: %s", elapsed)
	}
}

func TestKeyedLocker_ContextTimeout(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestKeyedLocker_ContextTimeout")

	locker := NewKeyedLocker[string](100 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// Lock key before calling
	go func() {
		_ = locker.Do(context.Background(), "key-timeout", func() {
			time.Sleep(50 * time.Millisecond)
		})
	}()

	time.Sleep(1 * time.Millisecond) // ensure lock is acquired first

	err := locker.Do(ctx, "key-timeout", func() {
		t.Error("should not execute")
	})

	assert.IsNotNil(err)
}

func TestKeyedLocker_LockReleaseAfterTTL(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestKeyedLocker_LockReleaseAfterTTL")

	locker := NewKeyedLocker[string](50 * time.Millisecond)

	err := locker.Do(context.Background(), "ttl-key", func() {})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Wait for TTL to pass
	time.Sleep(100 * time.Millisecond)

	err = locker.Do(context.Background(), "ttl-key", func() {})
	assert.IsNil(err)
}

func TestRWKeyedLocker_LockAndUnlock(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestKeyedLocker_LockReleaseAfterTTL")

	locker := NewRWKeyedLocker[string](500 * time.Millisecond)

	var locked bool
	err := locker.Lock(context.Background(), "key1", func() {
		locked = true
	})

	assert.IsNil(err)
	assert.Equal(true, locked)
}

func TestRWKeyedLocker_RLockParallel(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestKeyedLocker_LockReleaseAfterTTL")

	locker := NewRWKeyedLocker[string](1 * time.Second)

	var mu sync.Mutex
	var count int
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := locker.RLock(context.Background(), "shared-key", func() {
				time.Sleep(10 * time.Millisecond)
				mu.Lock()
				count++
				mu.Unlock()
			})

			assert.IsNil(err)
		}()
	}

	wg.Wait()

	assert.Equal(5, count)
}

func TestRWKeyedLocker_LockTimeout(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRWKeyedLocker_LockTimeout")

	locker := NewRWKeyedLocker[string](1 * time.Second)

	start := make(chan struct{})

	go func() {
		locker.Lock(context.Background(), "key-timeout", func() {
			close(start)
			time.Sleep(200 * time.Millisecond)
		})
	}()

	<-start
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := locker.Lock(ctx, "key-timeout", func() {
		t.Error("should not reach here")
	})

	assert.IsNotNil(err)
}

func TestTryKeyedLocker_SimpleLockUnlock(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestTryKeyedLocker_SimpleLockUnlock")

	locker := NewTryKeyedLocker[string]()

	ok := locker.TryLock("key1")
	assert.Equal(true, ok)

	ok = locker.TryLock("key1")
	assert.Equal(false, ok)

	locker.Unlock("key1")

	ok = locker.TryLock("key1")
	assert.Equal(true, ok)

	locker.Unlock("key1")
}

func TestTryKeyedLocker_ParallelTry(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestTryKeyedLocker_ParallelTry")

	locker := NewTryKeyedLocker[string]()

	var wg sync.WaitGroup
	var mu sync.Mutex
	var count int

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ok := locker.TryLock("key" + strconv.Itoa(i))
			mu.Lock()
			if ok {
				count++
			}
			mu.Unlock()
			time.Sleep(10 * time.Millisecond)
			if ok {
				locker.Unlock("key" + strconv.Itoa(i))
			}
		}(i)
	}

	wg.Wait()

	assert.Equal(5, count)
	assert.Equal(0, len(locker.locks))
}

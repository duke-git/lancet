package concurrency

import (
	"context"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
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

func TestKeyedLocker_MutualExclusionUnderContention(t *testing.T) {
	t.Parallel()

	locker := NewKeyedLocker[string](time.Millisecond)

	var holders, violations int32
	var wg sync.WaitGroup

	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 50000; i++ {
				_ = locker.Do(context.Background(), "key", func() {
					if atomic.AddInt32(&holders, 1) > 1 {
						atomic.AddInt32(&violations, 1)
					}
					runtime.Gosched()
					atomic.AddInt32(&holders, -1)
				})
			}
		}()
	}
	wg.Wait()

	if v := atomic.LoadInt32(&violations); v != 0 {
		t.Fatalf("the lock was held by more than one goroutine %d times", v)
	}
}

func TestKeyedLocker_ContextTimeoutWhileRunning(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestKeyedLocker_ContextTimeoutWhileRunning")

	locker := NewKeyedLocker[string](time.Millisecond)

	ctx, cancel := context.WithCancel(context.Background())
	finish := make(chan struct{})

	var running int32
	result := make(chan error)
	go func() {
		result <- locker.Do(ctx, "key", func() {
			atomic.StoreInt32(&running, 1)
			cancel()
			<-finish
			atomic.StoreInt32(&running, 0)
		})
	}()

	// returns although the function still runs
	select {
	case err := <-result:
		assert.Equal(context.Canceled, err)
	case <-time.After(time.Second):
		close(finish)
		t.Fatal("Do waited for the function after the context was canceled")
	}

	time.AfterFunc(50*time.Millisecond, func() { close(finish) })
	var overlapped bool
	err := locker.Do(context.Background(), "key", func() {
		overlapped = atomic.LoadInt32(&running) == 1
	})
	assert.IsNil(err)
	assert.Equal(false, overlapped)
}

func TestRWKeyedLocker_ContextTimeoutWhileRunning(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRWKeyedLocker_ContextTimeoutWhileRunning")

	locker := NewRWKeyedLocker[string](time.Millisecond)

	ctx, cancel := context.WithCancel(context.Background())
	finish := make(chan struct{})

	var running int32
	err := locker.Lock(ctx, "key", func() {
		atomic.StoreInt32(&running, 1)
		cancel()
		<-finish
		atomic.StoreInt32(&running, 0)
	})
	assert.Equal(context.Canceled, err)

	// the TTL passes while the function still runs
	time.Sleep(50 * time.Millisecond)
	time.AfterFunc(50*time.Millisecond, func() { close(finish) })

	var overlapped bool
	err = locker.Lock(context.Background(), "key", func() {
		overlapped = atomic.LoadInt32(&running) == 1
	})
	assert.IsNil(err)
	assert.Equal(false, overlapped)
}

func TestRWKeyedLocker_RemoveAfterTTL(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRWKeyedLocker_RemoveAfterTTL")

	locker := NewRWKeyedLocker[string](10 * time.Millisecond)

	count := func() int {
		locker.mu.Lock()
		defer locker.mu.Unlock()
		return len(locker.locks)
	}

	err := locker.Lock(context.Background(), "key", func() {})
	assert.IsNil(err)

	time.Sleep(200 * time.Millisecond)
	assert.Equal(0, count())
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

func TestTryKeyedLocker_MutualExclusionUnderContention(t *testing.T) {
	t.Parallel()

	locker := NewTryKeyedLocker[string]()

	var holders, violations int32
	var wg sync.WaitGroup

	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 20000; i++ {
				if !locker.TryLock("key") {
					continue
				}
				if atomic.AddInt32(&holders, 1) > 1 {
					atomic.AddInt32(&violations, 1)
				}
				atomic.AddInt32(&holders, -1)
				locker.Unlock("key")
			}
		}()
	}
	wg.Wait()

	if v := atomic.LoadInt32(&violations); v != 0 {
		t.Fatalf("the lock was held by more than one goroutine %d times", v)
	}
}

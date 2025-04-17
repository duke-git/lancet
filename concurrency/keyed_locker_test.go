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

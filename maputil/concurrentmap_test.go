package maputil

import (
	"fmt"
	"sync"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestConcurrentMap_Set_Get(t *testing.T) {
	assert := internal.NewAssert(t, "TestConcurrentMap_Set_Get")

	cm := NewConcurrentMap[string, int](100)

	var wg1 sync.WaitGroup
	wg1.Add(10)

	for i := 0; i < 10; i++ {
		go func(n int) {
			cm.Set(fmt.Sprintf("%d", n), n)
			wg1.Done()
		}(i)
	}
	wg1.Wait()

	var wg2 sync.WaitGroup
	wg2.Add(10)
	for j := 0; j < 10; j++ {
		go func(n int) {
			val, ok := cm.Get(fmt.Sprintf("%d", n))
			assert.Equal(n, val)
			assert.Equal(true, ok)
			wg2.Done()
		}(j)
	}
	wg2.Wait()
}

func TestConcurrentMap_GetOrSet(t *testing.T) {
	assert := internal.NewAssert(t, "TestConcurrentMap_GetOrSet")

	cm := NewConcurrentMap[string, int](100)

	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			val, ok := cm.GetOrSet(fmt.Sprintf("%d", n), n)
			assert.Equal(n, val)
			assert.Equal(false, ok)
			wg.Done()
		}(i)
	}
	wg.Wait()

	for j := 0; j < 5; j++ {
		go func(n int) {
			val, ok := cm.Get(fmt.Sprintf("%d", n))
			assert.Equal(n, val)
			assert.Equal(true, ok)
		}(j)
	}
}

func TestConcurrentMap_Delete(t *testing.T) {
	assert := internal.NewAssert(t, "TestConcurrentMap_Delete")

	cm := NewConcurrentMap[string, int](100)

	var wg1 sync.WaitGroup
	wg1.Add(10)

	for i := 0; i < 10; i++ {
		go func(n int) {
			cm.Set(fmt.Sprintf("%d", n), n)
			wg1.Done()
		}(i)
	}
	wg1.Wait()

	var wg2 sync.WaitGroup
	wg2.Add(10)

	for i := 0; i < 10; i++ {
		go func(n int) {
			cm.Delete(fmt.Sprintf("%d", n))
			wg2.Done()
		}(i)
	}
	wg2.Wait()

	for j := 0; j < 10; j++ {
		go func(n int) {
			_, ok := cm.Get(fmt.Sprintf("%d", n))
			assert.Equal(false, ok)
		}(j)
	}
}

func TestConcurrentMap_GetAndDelete(t *testing.T) {
	assert := internal.NewAssert(t, "TestConcurrentMap_GetAndDelete")

	cm := NewConcurrentMap[string, int](100)

	for i := 0; i < 10; i++ {
		go func(n int) {
			cm.Set(fmt.Sprintf("%d", n), n)
		}(i)
	}

	for j := 0; j < 10; j++ {
		go func(n int) {
			val, ok := cm.GetAndDelete(fmt.Sprintf("%d", n))
			assert.Equal(n, val)
			assert.Equal(true, ok)

			_, ok = cm.Get(fmt.Sprintf("%d", n))
			assert.Equal(false, ok)
		}(j)
	}
}

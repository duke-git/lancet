package algorithm

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestLRUCache(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestLRUCache")

	cache := NewLRUCache[int, int](3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	assert.Equal(3, cache.Len())

	v, ok := cache.Get(1)
	assert.Equal(true, ok)
	assert.Equal(1, v)

	v, ok = cache.Get(2)
	assert.Equal(true, ok)
	assert.Equal(2, v)

	ok = cache.Delete(2)
	assert.Equal(true, ok)

	_, ok = cache.Get(2)
	assert.Equal(false, ok)
}

func TestLRUCacheDeleteLen(t *testing.T) {
	t.Parallel()
	cache := NewLRUCache[int, int](3)
	for key := 1; key <= 3; key++ {
		cache.Put(key, key)
	}

	for _, key := range []int{2, 1, 3} {
		before := cache.Len()
		if !cache.Delete(key) {
			t.Fatalf("Delete(%d) = false, want true", key)
		}
		if got := cache.Len(); got != before-1 {
			t.Fatalf("Len() after Delete(%d) = %d, want %d", key, got, before-1)
		}
		if cache.Delete(key) {
			t.Fatalf("Delete(%d) = true for a missing key", key)
		}
		if got := cache.Len(); got != before-1 {
			t.Fatalf("Len() after missing Delete(%d) = %d, want %d", key, got, before-1)
		}
	}
}

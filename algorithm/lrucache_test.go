package algorithm

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestLRUCache(t *testing.T) {
	t.Parallel()
	asssert := internal.NewAssert(t, "TestLRUCache")

	cache := NewLRUCache[int, int](3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	asssert.Equal(3, cache.Len())

	v, ok := cache.Get(1)
	asssert.Equal(true, ok)
	asssert.Equal(1, v)

	v, ok = cache.Get(2)
	asssert.Equal(true, ok)
	asssert.Equal(2, v)

	ok = cache.Delete(2)
	asssert.Equal(true, ok)

	_, ok = cache.Get(2)
	asssert.Equal(false, ok)
}

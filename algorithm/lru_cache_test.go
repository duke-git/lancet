package algorithm

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestLRUCache(t *testing.T) {
	asssert := internal.NewAssert(t, "TestLRUCache")

	cache := NewLRUCache[int, int](2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	_, ok := cache.Get(0)
	asssert.Equal(false, ok)

	v, ok := cache.Get(1)
	asssert.Equal(true, ok)
	asssert.Equal(1, v)

	v, ok = cache.Get(2)
	asssert.Equal(true, ok)
	asssert.Equal(2, v)

	cache.Put(3, 3)
	v, ok = cache.Get(1)
	asssert.Equal(false, ok)
	asssert.NotEqual(1, v)

	v, ok = cache.Get(3)
	asssert.Equal(true, ok)
	asssert.Equal(3, v)
}

package algorithm

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestLRUCache(t *testing.T) {
	asssert := internal.NewAssert(t, "TestLRUCache")

	cache := NewLRUCache[int, int](2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	_, ok := cache.Get(0)
	asssert.Equal(false, ok)
}

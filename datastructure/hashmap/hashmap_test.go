package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestHashMap_PutAndGet(t *testing.T) {
	assert := internal.NewAssert(t, "TestHashMap_Get")

	hm := NewHashMap()

	hm.Put("abc", 3)
	assert.Equal(3, hm.Get("abc"))
	assert.IsNil(hm.Get("abcd"))

	hm.Put("abc", 4)
	assert.Equal(4, hm.Get("abc"))
}

package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestHashMap_PutAndGet(t *testing.T) {
	assert := internal.NewAssert(t, "TestHashMap_PutAndGet")

	hm := NewHashMap()

	hm.Put("abc", 3)
	assert.Equal(3, hm.Get("abc"))
	assert.IsNil(hm.Get("abcd"))

	hm.Put("abc", 4)
	assert.Equal(4, hm.Get("abc"))
}

func TestHashMap_Resize(t *testing.T) {
	assert := internal.NewAssert(t, "TestHashMap_Resize")

	hm := NewHashMapWithCapacity(3, 3)

	for i := 0; i < 20; i++ {
		hm.Put(i, 10)
	}

	assert.Equal(10, hm.Get(5))
}

func TestHashMap_Delete(t *testing.T) {
	assert := internal.NewAssert(t, "TestHashMap_Delete")

	hm := NewHashMap()

	hm.Put("abc", 3)
	assert.Equal(3, hm.Get("abc"))

	hm.Delete("abc")
	assert.IsNil(hm.Get("abc"))
}

func TestHashMap_Contains(t *testing.T) {
	assert := internal.NewAssert(t, "TestHashMap_Contains")

	hm := NewHashMap()
	assert.Equal(false, hm.Contains("abc"))

	hm.Put("abc", 3)
	assert.Equal(true, hm.Contains("abc"))
}

func TestHashMap_KeysValues(t *testing.T) {
	assert := internal.NewAssert(t, "TestHashMap_KeysValues")

	hm := NewHashMap()

	hm.Put("a", 1)
	hm.Put("b", 2)
	hm.Put("c", 3)

	keys := hm.Keys()
	values := hm.Values()
	t.Log(keys, values)

	assert.Equal(3, len(values))
	assert.Equal(3, len(keys))
}

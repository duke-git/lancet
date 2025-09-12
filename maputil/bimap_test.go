package maputil

import (
	"encoding/json"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestBiMap_Get(t *testing.T) {
	biMap, _ := NewBiMapFromMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	assert := internal.NewAssert(t, "TestBiMap_Get")
	assert.Equal(1, biMap.Value("one"))
	assert.Equal(true, biMap.ContainsKey("one"))
	assert.Equal(2, biMap.Value("two"))
	assert.Equal(3, biMap.Value("three"))
	assert.Equal(0, biMap.Value("four"))
	assert.Equal(false, biMap.ContainsKey("four"))

	assert.Equal("one", biMap.Key(1))
	assert.Equal(true, biMap.ContainsValue(1))
	assert.Equal("two", biMap.Key(2))
	assert.Equal("three", biMap.Key(3))
	assert.Equal("", biMap.Key(4))
	assert.Equal(false, biMap.ContainsValue(4))
}

func TestBiMap_Contains(t *testing.T) {
	biMap, _ := NewBiMapFromMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	assert := internal.NewAssert(t, "TestBiMap_Contains")

	assert.Equal(false, biMap.ContainsKey("zero"))
	assert.Equal(true, biMap.ContainsKey("one"))
	assert.Equal(true, biMap.ContainsKey("two"))
	assert.Equal(true, biMap.ContainsKey("three"))
	assert.Equal(false, biMap.ContainsKey("four"))

	assert.Equal(false, biMap.ContainsValue(0))
	assert.Equal(true, biMap.ContainsValue(1))
	assert.Equal(true, biMap.ContainsValue(2))
	assert.Equal(true, biMap.ContainsValue(3))
	assert.Equal(false, biMap.ContainsValue(4))
}

func TestBiMap_Put(t *testing.T) {
	biMap := NewBiMap[string, int]()

	assert := internal.NewAssert(t, "TestBiMap_Put")

	err := biMap.Put("one", 1)
	assert.IsNil(err)
	assert.Equal(1, biMap.Len())
	err = biMap.Put("one", 2)
	assert.IsNotNil(err)
	assert.Equal(1, biMap.Len())
	err = biMap.Put("two", 1)
	assert.IsNotNil(err)
	err = biMap.Put("two", 2)
	assert.IsNil(err)
	assert.Equal(2, biMap.Len())

	err = biMap.PutForce("one", 2)
	assert.IsNil(err)
	assert.Equal(1, biMap.Len())
	assert.Equal(true, biMap.ContainsKey("one"))
	assert.Equal(true, biMap.ContainsValue(2))
	assert.Equal(false, biMap.ContainsKey("two"))
	assert.Equal(false, biMap.ContainsValue(1))

	_, err = NewBiMapFromMap(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"1":     1,
	})
	assert.IsNotNil(err)
}

func TestBiMap_Remove(t *testing.T) {
	biMap, _ := NewBiMapFromMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
	})

	assert := internal.NewAssert(t, "TestBiMap_Remove")
	biMap.RemoveKey("one")
	assert.Equal(false, biMap.ContainsKey("one"))
	assert.Equal(false, biMap.ContainsValue(1))
	assert.Equal(true, biMap.ContainsKey("two"))
	assert.Equal(true, biMap.ContainsValue(2))
	assert.Equal(5, biMap.Len())

	biMap.RemoveValue(5)
	assert.Equal(false, biMap.ContainsKey("five"))
	assert.Equal(false, biMap.ContainsValue(5))
	assert.Equal(true, biMap.ContainsKey("four"))
	assert.Equal(true, biMap.ContainsValue(6))
	assert.Equal(4, biMap.Len())
	assert.Equal(4, len(biMap.ToMap()))
	assert.Equal(4, len(biMap.Keys()))
	assert.Equal(4, len(biMap.Values()))

	biMap.RemoveKey("seven")
	biMap.RemoveValue(7)
	assert.Equal(4, biMap.Len())

	biMap.Clear()
	assert.Equal(0, biMap.Len())
	assert.Equal(false, biMap.ContainsKey("four"))
	assert.Equal(false, biMap.ContainsValue(6))
	assert.Equal(0, len(biMap.ToMap()))
	assert.Equal(0, len(biMap.Keys()))
	assert.Equal(0, len(biMap.Values()))
}

func TestBiMap_Inverse(t *testing.T) {
	_biMap, _ := NewBiMapFromMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})
	biMap := _biMap.Inverse()

	assert := internal.NewAssert(t, "TestBiMap_Inverse")
	assert.Equal(true, biMap.ContainsKey(1))
	assert.Equal(true, biMap.ContainsValue("one"))
	assert.Equal(3, biMap.Len())
}

func TestBiMap_Json(t *testing.T) {
	_biMap, _ := NewBiMapFromMap[string, int](map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})
	assert := internal.NewAssert(t, "TestBiMap_Inverse")

	j, err := json.Marshal(_biMap)
	assert.IsNil(err)
	biMap := NewBiMap[string, int]()
	err = json.Unmarshal(j, biMap)
	assert.IsNil(err)

	assert.Equal(1, biMap.Value("one"))
	assert.Equal("one", biMap.Key(1))
	assert.Equal(3, biMap.Len())

	biMap = NewBiMap[string, int]()
	err = biMap.UnmarshalJSON([]byte("{;"))
	assert.IsNotNil(err)

}

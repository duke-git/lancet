package maputil

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestOrderedMap_Set_Get(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	assert := internal.NewAssert(t, "TestOrderedMap_Set_Get")

	val, ok := om.Get("a")
	assert.Equal(1, val)
	assert.Equal(true, ok)

	om.Set("a", 2)
	val, _ = om.Get("a")
	assert.Equal(2, val)

	val, ok = om.Get("d")
	assert.Equal(false, ok)
	assert.Equal(0, val)
}

func TestOrderedMap_Front_Back(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	assert := internal.NewAssert(t, "TestOrderedMap_Front_Back")

	frontElement, ok := om.Front()
	assert.Equal("a", frontElement.Key)
	assert.Equal(1, frontElement.Value)
	assert.Equal(true, ok)

	backElement, ok := om.Back()
	assert.Equal("c", backElement.Key)
	assert.Equal(3, backElement.Value)
	assert.Equal(true, ok)
}

func TestOrderedMap_Delete_Clear(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	assert := internal.NewAssert(t, "TestOrderedMap_Delete_Clear")

	assert.Equal(3, om.Len())

	om.Delete("b")
	assert.Equal(2, om.Len())

	om.Clear()
	assert.Equal(0, om.Len())
}

func TestOrderedMap_Keys_Values(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	assert := internal.NewAssert(t, "TestOrderedMap_Keys_Values")

	assert.Equal([]string{"a", "b", "c"}, om.Keys())
	assert.Equal([]int{1, 2, 3}, om.Values())
}

func TestOrderedMap_Contains(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	assert := internal.NewAssert(t, "TestOrderedMap_Contains")

	assert.Equal(true, om.Contains("a"))
	assert.Equal(false, om.Contains("d"))
}

func TestOrderedMap_Eelements(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	assert := internal.NewAssert(t, "TestOrderedMap_Eelements")

	elements := []struct {
		Key   string
		Value int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	assert.Equal(elements, om.Elements())
}

func TestOrderedMap_Range(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)
	om.Set("d", 4)

	assert := internal.NewAssert(t, "TestOrderedMap_Range")

	var keys []string
	om.Range(func(key string, value int) bool {
		keys = append(keys, key)
		return key != "c"
	})

	assert.Equal([]string{"a", "b", "c"}, keys)
}

func TestOrderedMap_Iter(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	assert := internal.NewAssert(t, "TestOrderedMap_Iter")

	var items []struct {
		Key   string
		Value int
	}

	iterCh := om.Iter()

	for item := range iterCh {
		items = append(items, item)
	}

	expected := []struct {
		Key   string
		Value int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	assert.Equal(expected, items)
}

func TestOrderedMap_ReverseIter(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	assert := internal.NewAssert(t, "TestOrderedMap_ReverseIter")

	var items []struct {
		Key   string
		Value int
	}

	iterCh := om.ReverseIter()

	for item := range iterCh {
		items = append(items, item)
	}

	expected := []struct {
		Key   string
		Value int
	}{
		{"c", 3},
		{"b", 2},
		{"a", 1},
	}

	assert.Equal(expected, items)
}

func TestOrderedMap_SortByKey(t *testing.T) {
	assert := internal.NewAssert(t, "TestOrderedMap_SortByKey")

	om := NewOrderedMap[string, int]()

	om.Set("d", 4)
	om.Set("b", 2)
	om.Set("c", 3)
	om.Set("a", 1)

	om.SortByKey(func(a, b string) bool {
		return a < b
	})

	assert.Equal([]string{"a", "b", "c", "d"}, om.Keys())
}

func TestOrderedMap_MarshalJSON(t *testing.T) {
	om := NewOrderedMap[string, int]()

	om.Set("a", 1)
	om.Set("b", 2)
	om.Set("c", 3)

	assert := internal.NewAssert(t, "TestOrderedMap_MarshalJSON")

	jsonBytes, err := om.MarshalJSON()

	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}

	assert.Equal(`{"a":1,"b":2,"c":3}`, string(jsonBytes))
}

func TestOrderedMap_UnmarshalJSON(t *testing.T) {
	assert := internal.NewAssert(t, "TestOrderedMap_UnmarshalJSON")

	jsonStr := `{"a":1,"b":2,"c":3}`

	om := NewOrderedMap[string, int]()
	err := om.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}

	assert.Equal(3, om.Len())
	assert.Equal(true, om.Contains("a"))
	assert.Equal(true, om.Contains("b"))
	assert.Equal(true, om.Contains("c"))
}

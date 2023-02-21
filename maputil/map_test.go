package maputil

import (
	"sort"
	"strconv"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestKeys(t *testing.T) {
	assert := internal.NewAssert(t, "TestKeys")

	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	keys := Keys(m)
	sort.Ints(keys)

	assert.Equal([]int{1, 2, 3, 4, 5}, keys)
}

func TestValues(t *testing.T) {
	assert := internal.NewAssert(t, "TestValues")

	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	values := Values(m)
	sort.Strings(values)

	assert.Equal([]string{"a", "a", "b", "c", "d"}, values)
}

func TestKeysBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestKeysBy")

	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
	}

	keys := KeysBy(m, func(n int) int {
		return n + 1
	})

	sort.Ints(keys)

	assert.Equal([]int{2, 3, 4}, keys)
}

func TestValuesBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestValuesBy")

	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	values := ValuesBy(m, func(v string) string {
		switch v {
		case "a":
			return "a-1"
		case "b":
			return "b-2"
		case "c":
			return "c-3"
		default:
			return ""
		}
	})

	sort.Strings(values)

	assert.Equal([]string{"a-1", "b-2", "c-3"}, values)
}

func TestMerge(t *testing.T) {
	assert := internal.NewAssert(t, "TestMerge")

	m1 := map[int]string{
		1: "a",
		2: "b",
	}
	m2 := map[int]string{
		1: "1",
		3: "2",
	}

	expected := map[int]string{
		1: "1",
		2: "b",
		3: "2",
	}
	acturl := Merge(m1, m2)

	assert.Equal(expected, acturl)
}

func TestForEach(t *testing.T) {
	assert := internal.NewAssert(t, "TestForEach")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	var sum int

	ForEach(m, func(_ string, value int) {
		sum += value
	})

	assert.Equal(10, sum)
}

func TestFilter(t *testing.T) {
	assert := internal.NewAssert(t, "TestFilter")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}

	acturl := Filter(m, isEven)

	assert.Equal(map[string]int{
		"b": 2,
		"d": 4,
	}, acturl)
}

func TestFilterByKeys(t *testing.T) {
	assert := internal.NewAssert(t, "TestFilterByKeys")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	acturl := FilterByKeys(m, []string{"a", "b"})

	assert.Equal(map[string]int{
		"a": 1,
		"b": 2,
	}, acturl)
}

func TestFilterByValues(t *testing.T) {
	assert := internal.NewAssert(t, "TestFilterByValues")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	acturl := FilterByValues(m, []int{3, 4})

	assert.Equal(map[string]int{
		"c": 3,
		"d": 4,
	}, acturl)
}

func TestOmitBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestOmitBy")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}

	acturl := OmitBy(m, isEven)

	assert.Equal(map[string]int{
		"a": 1,
		"c": 3,
		"e": 5,
	}, acturl)
}

func TestOmitByKeys(t *testing.T) {
	assert := internal.NewAssert(t, "TestOmitByKeys")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	acturl := OmitByKeys(m, []string{"a", "b"})

	assert.Equal(map[string]int{
		"c": 3,
		"d": 4,
		"e": 5,
	}, acturl)
}

func TestOmitByValues(t *testing.T) {
	assert := internal.NewAssert(t, "TestOmitByValues")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	acturl := OmitByValues(m, []int{4, 5})

	assert.Equal(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}, acturl)
}

func TestIntersect(t *testing.T) {
	assert := internal.NewAssert(t, "TestIntersect")

	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 6,
		"d": 7,
	}

	m3 := map[string]int{
		"a": 1,
		"b": 9,
		"e": 9,
	}

	assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, Intersect(m1))
	assert.Equal(map[string]int{"a": 1, "b": 2}, Intersect(m1, m2))
	assert.Equal(map[string]int{"a": 1}, Intersect(m1, m2, m3))
}

func TestMinus(t *testing.T) {
	assert := internal.NewAssert(t, "TestMinus")

	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 11,
		"b": 22,
		"d": 33,
	}

	assert.Equal(map[string]int{"c": 3}, Minus(m1, m2))
}

func TestIsDisjoint(t *testing.T) {
	assert := internal.NewAssert(t, "TestMinus")

	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"d": 22,
	}

	assert.Equal(true, IsDisjoint(m1, m2))

	m3 := map[string]int{
		"a": 22,
	}

	assert.Equal(false, IsDisjoint(m1, m3))
}

func TestEntries(t *testing.T) {
	assert := internal.NewAssert(t, "TestEntries")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	result := Entries(m)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value < result[j].Value
	})

	expected := []Entry[string, int]{{Key: "a", Value: 1}, {Key: "b", Value: 2}, {Key: "c", Value: 3}}

	assert.Equal(expected, result)
}

func TestFromEntries(t *testing.T) {
	assert := internal.NewAssert(t, "TestFromEntries")

	result := FromEntries([]Entry[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
		{Key: "c", Value: 3},
	})

	assert.Equal(3, len(result))
	assert.Equal(1, result["a"])
	assert.Equal(2, result["b"])
	assert.Equal(3, result["c"])
}

func TestTransform(t *testing.T) {
	assert := internal.NewAssert(t, "TestTransform")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	result := Transform(m, func(k string, v int) (string, string) {
		return k, strconv.Itoa(v)
	})

	expected := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	assert.Equal(expected, result)
}

func TestMapKeys(t *testing.T) {
	assert := internal.NewAssert(t, "TestMapKeys")

	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	result := MapKeys(m, func(k int, _ string) string {
		return strconv.Itoa(k)
	})

	expected := map[string]string{
		"1": "a",
		"2": "b",
		"3": "c",
	}

	assert.Equal(expected, result)
}

func TestMapValues(t *testing.T) {
	assert := internal.NewAssert(t, "TestMapKeys")

	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	result := MapValues(m, func(k int, v string) string {
		return v + strconv.Itoa(k)
	})

	expected := map[int]string{
		1: "a1",
		2: "b2",
		3: "c3",
	}

	assert.Equal(expected, result)
}

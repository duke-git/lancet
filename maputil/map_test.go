package maputil

import (
	"sort"
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

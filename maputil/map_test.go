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

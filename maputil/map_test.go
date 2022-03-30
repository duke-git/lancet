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

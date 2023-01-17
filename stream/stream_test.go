package stream

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestFromSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestFromSlice")

	stream := FromSlice([]int{1, 2, 3})

	assert.Equal([]int{1, 2, 3}, stream.ToSlice())
}

func TestFromRange(t *testing.T) {
	assert := internal.NewAssert(t, "TestFromSlice")

	s1 := FromRange(1, 5, 1)
	s2 := FromRange(1.1, 5.0, 1.0)

	assert.Equal([]int{1, 2, 3, 4, 5}, s1.ToSlice())
	assert.Equal([]float64{1.1, 2.1, 3.1, 4.1}, s2.ToSlice())
}

package internal

import (
	"testing"
)

func TestAssert(t *testing.T) {
	assert := NewAssert(t, "TestAssert")
	assert.Equal(0, 0)
	assert.NotEqual(1, 0)
	assert.Greater(1, 0)
	assert.GreaterOrEqual(1, 1)
	assert.Less(0, 1)
	assert.LessOrEqual(0, 0)

	assert.Greater(1.1, 0.1)
	assert.Less(0.1, 1.1)

	assert.Equal("abc", "abc")
	assert.NotEqual("abc", "abd")
	assert.Less("abc", "abd")
	assert.Greater("abd", "abc")

	assert.Equal([]int{1, 2, 3}, []int{1, 2, 3})
	assert.NotEqual([]int{1, 2, 3}, []int{1, 2})

	assert.IsNil(nil)
	assert.IsNotNil("abc")

}

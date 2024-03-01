package internal

import (
	"testing"
)

func TestAssert(t *testing.T) {
	assert := NewAssert(t, "TestAssert")

	assert.Equal(0, 0)
	assert.NotEqual(1, 0)
	assert.NotEqual("1", 1)

	var uInt1 uint
	var uInt2 uint
	var uInt8 uint8
	var uInt16 uint16
	var uInt32 uint32
	var uInt64 uint64
	assert.NotEqual(uInt1, uInt8)
	assert.NotEqual(uInt8, uInt16)
	assert.NotEqual(uInt16, uInt32)
	assert.NotEqual(uInt32, uInt64)

	assert.Equal(uInt1, uInt2)

	uInt1 = 1
	uInt2 = 2
	assert.Less(uInt1, uInt2)

	assert.Greater(1, 0)
	assert.GreaterOrEqual(1, 1)
	assert.Less(0, 1)
	assert.LessOrEqual(0, 0)

	assert.Equal(0.1, 0.1)
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

	var valA int = 1
	var valB int64 = 1
	assert.NotEqual(valA, valB)
	assert.EqualValues(valA, valB)

	assert.ShouldBeFalse(false)
	assert.ShouldBeTrue(true)
}

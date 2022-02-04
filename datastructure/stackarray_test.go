package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestStackArray_Push(t *testing.T) {
	assert := internal.NewAssert(t, "TestStackArray_Push")

	stack := NewStackArray[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expected := []int{3, 2, 1}
	values := stack.Data()
	length := stack.Length()

	assert.Equal(expected, values)
	assert.Equal(3, length)
}

func TestStackArray_Pop(t *testing.T) {
	assert := internal.NewAssert(t, "TestStackArray_Pop")

	stack := NewStackArray[int]()
	topItem, err := stack.Pop()
	assert.IsNotNil(err)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topItem, err = stack.Pop()
	assert.IsNil(err)
	assert.Equal(3, *topItem)

	expected := []int{2, 1}
	assert.Equal(expected, stack.Data())
}

func TestStackArray_Peak(t *testing.T) {
	assert := internal.NewAssert(t, "TestStackArray_Peak")

	stack := NewStackArray[int]()
	topItem, err := stack.Peak()
	assert.IsNotNil(err)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topItem, err = stack.Peak()
	assert.IsNil(err)
	assert.Equal(3, *topItem)

	expected := []int{3, 2, 1}
	assert.Equal(expected, stack.Data())
}

func TestStackArray_Empty(t *testing.T) {
	assert := internal.NewAssert(t, "TestStackArray_Empty")

	stack := NewStackArray[int]()
	assert.Equal(true, stack.IsEmpty())
	assert.Equal(0, stack.Length())

	stack.Push(1)
	assert.Equal(false, stack.IsEmpty())
	assert.Equal(1, stack.Length())

	stack.EmptyStack()
	assert.Equal(true, stack.IsEmpty())
	assert.Equal(0, stack.Length())
}

package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestArrayStack_Push(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayStack_Push")

	stack := NewArrayStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expected := []int{3, 2, 1}
	values := stack.Data()
	length := stack.Size()

	assert.Equal(expected, values)
	assert.Equal(3, length)
}

func TestArrayStack_Pop(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayStack_Pop")

	stack := NewArrayStack[int]()
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

func TestArrayStack_Peak(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayStack_Peak")

	stack := NewArrayStack[int]()
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

func TestArrayStack_Clear(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayStack_Clear")

	stack := NewArrayStack[int]()
	assert.Equal(true, stack.IsEmpty())
	assert.Equal(0, stack.Size())

	stack.Push(1)
	assert.Equal(false, stack.IsEmpty())
	assert.Equal(1, stack.Size())

	stack.Clear()
	assert.Equal(true, stack.IsEmpty())
	assert.Equal(0, stack.Size())
}

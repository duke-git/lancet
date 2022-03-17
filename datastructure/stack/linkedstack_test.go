package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestLinkedStack_Push(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinkedStack_Push")

	stack := NewLinkedStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Print()

	expected := []int{3, 2, 1}
	values := stack.Data()
	size := stack.Size()

	assert.Equal(expected, values)
	assert.Equal(3, size)
}

func TestLinkedStack_Pop(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinkedStack_Pop")

	stack := NewLinkedStack[int]()
	topItem, err := stack.Pop()
	assert.IsNotNil(err)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topItem, err = stack.Pop()
	assert.IsNil(err)
	assert.Equal(3, *topItem)

	expected := []int{2, 1}
	stack.Print()
	assert.Equal(expected, stack.Data())
}

func TestLinkedStack_Peak(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinkedStack_Peak")

	stack := NewLinkedStack[int]()
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

func TestLinkedStack_Empty(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinkedStack_Empty")

	stack := NewLinkedStack[int]()
	assert.Equal(true, stack.IsEmpty())
	assert.Equal(0, stack.Size())

	stack.Push(1)
	assert.Equal(false, stack.IsEmpty())
	assert.Equal(1, stack.Size())

	stack.Clear()
	assert.Equal(true, stack.IsEmpty())
	assert.Equal(0, stack.Size())
}

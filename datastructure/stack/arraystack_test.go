package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestArrayStack_Push(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestArrayStack_Push")

	stack := NewArrayStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	values := stack.Data()
	length := stack.Size()

	assert.Equal([]int{3, 2, 1}, values)
	assert.Equal(3, length)
}

func TestArrayStack_Pop(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestArrayStack_Pop")

	stack := NewArrayStack[int]()
	_, err := stack.Pop()
	assert.IsNotNil(err)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topItem, err := stack.Pop()
	assert.IsNil(err)
	assert.Equal(3, *topItem)

	assert.Equal([]int{2, 1}, stack.Data())
}

func TestArrayStack_Peak(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestArrayStack_Peak")

	stack := NewArrayStack[int]()
	_, err := stack.Peak()
	assert.IsNotNil(err)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topItem, err := stack.Peak()
	assert.IsNil(err)
	assert.Equal(3, *topItem)

	assert.Equal([]int{3, 2, 1}, stack.Data())
}

func TestArrayStack_Clear(t *testing.T) {
	t.Parallel()

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

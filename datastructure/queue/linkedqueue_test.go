package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestLinkedQueue_EnQueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinkedQueue_EnQueue")

	queue := NewLinkedQueue[int]()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	queue.Print()

	assert.Equal([]int{1, 2, 3}, queue.Data())
	assert.Equal(3, queue.Size())
}

func TestLinkedQueue_DeQueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinkedQueue_DeQueue")

	queue := NewLinkedQueue[int]()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	val, _ := queue.DeQueue()

	queue.Print()

	assert.Equal([]int{2, 3}, queue.Data())
	assert.Equal(1, *val)
}

func TestLinkedQueue_Front(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinkedQueue_Front")

	queue := NewLinkedQueue[int]()
	val, err := queue.Front()
	assert.IsNotNil(err)

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	val, err = queue.Front()
	assert.Equal(1, *val)
	assert.IsNil(err)
}

func TestLinkedQueue_Back(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinkedQueue_Back")

	queue := NewLinkedQueue[int]()
	val, err := queue.Back()
	assert.IsNotNil(err)

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	val, err = queue.Back()
	assert.Equal(3, *val)
	assert.IsNil(err)
}

func TestLinkedQueue_Clear(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinkedQueue_Back")

	queue := NewLinkedQueue[int]()
	assert.Equal(true, queue.IsEmpty())

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	assert.Equal(false, queue.IsEmpty())

	queue.Clear()
	assert.Equal(true, queue.IsEmpty())
}

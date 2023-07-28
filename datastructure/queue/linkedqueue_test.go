package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestLinkedQueue_Enqueue(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLinkedQueue_Enqueue")

	queue := NewLinkedQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.Equal([]int{1, 2, 3}, queue.Data())
	assert.Equal(3, queue.Size())
}

func TestLinkedQueue_Dequeue(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLinkedQueue_DeQueue")

	queue := NewLinkedQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, _ := queue.Dequeue()

	queue.Print()

	assert.Equal([]int{2, 3}, queue.Data())
	assert.Equal(1, *val)
}

func TestLinkedQueue_Front(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLinkedQueue_Front")

	queue := NewLinkedQueue[int]()
	_, err := queue.Front()
	assert.IsNotNil(err)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, err := queue.Front()
	assert.Equal(1, *val)
	assert.IsNil(err)
}

func TestLinkedQueue_Back(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLinkedQueue_Back")

	queue := NewLinkedQueue[int]()
	_, err := queue.Back()
	assert.IsNotNil(err)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, err := queue.Back()
	assert.Equal(3, *val)
	assert.IsNil(err)
}

func TestLinkedQueue_Clear(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLinkedQueue_Back")

	queue := NewLinkedQueue[int]()
	assert.Equal(true, queue.IsEmpty())

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	assert.Equal(false, queue.IsEmpty())

	queue.Clear()
	assert.Equal(true, queue.IsEmpty())
}

func TestLinkedQueue_Contain(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLinkedQueue_Contain")

	queue := NewLinkedQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.Equal(true, queue.Contain(1))
	assert.Equal(false, queue.Contain(4))
}

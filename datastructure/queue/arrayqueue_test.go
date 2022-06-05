package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestArrayQueue_Enqueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Enqueue")

	queue := NewArrayQueue[int](5)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	expected := []int{1, 2, 3}
	data := queue.Data()
	size := queue.Size()

	queue.Print()

	assert.Equal(expected, data)
	assert.Equal(3, size)
}

func TestArrayQueue_Dequeue(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Dequeue")

	queue := NewArrayQueue[int](4)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, ok := queue.Dequeue()
	assert.Equal(true, ok)

	queue.Print()
	assert.Equal(1, val)
	assert.Equal([]int{2, 3}, queue.Data())
}

func TestArrayQueue_Front(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Front")

	queue := NewArrayQueue[int](4)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val := queue.Front()

	queue.Print()

	assert.Equal(1, val)
	assert.Equal([]int{1, 2, 3}, queue.Data())
}

func TestArrayQueue_Back(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Back")

	queue := NewArrayQueue[int](4)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val := queue.Back()

	queue.Print()

	assert.Equal(3, val)
	assert.Equal([]int{1, 2, 3}, queue.Data())
}

func TestArrayQueue_Contain(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Contain")

	queue := NewArrayQueue[int](4)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.Equal(true, queue.Contain(1))
	assert.Equal(false, queue.Contain(4))
}

func TestArrayQueue_Clear(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Clear")

	queue := NewArrayQueue[int](4)

	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Size())

	queue.Enqueue(1)
	assert.Equal(false, queue.IsEmpty())
	assert.Equal(1, queue.Size())

	queue.Clear()
	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Size())
}

func TestArrayQueue_IsFull(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_IsFull")

	queue := NewArrayQueue[int](3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.Equal(true, queue.IsFull())
}

package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestCircularQueue_Enqueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Enqueue")

	queue := NewCircularQueue[int](6)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	queue.Print()
	assert.Equal([]int{1, 2, 3, 4, 5}, queue.Data())
	assert.Equal(5, queue.Size())

	err := queue.Enqueue(6)
	assert.IsNotNil(err)
}

func TestCircularQueue_Dequeue(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_DeQueue")

	queue := NewCircularQueue[int](6)
	assert.Equal(true, queue.IsEmpty())

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	val, err := queue.Dequeue()
	assert.IsNil(err)

	assert.Equal(1, *val)
	assert.Equal(false, queue.IsFull())

	val, _ = queue.Dequeue()
	queue.Print()
	assert.Equal(2, *val)

	queue.Enqueue(6)
	queue.Print()
	assert.Equal(false, queue.IsFull())
}

func TestCircularQueue_Front(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Front")

	queue := NewCircularQueue[int](6)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	queue.Print()

	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue(6)
	queue.Enqueue(7)

	queue.Print()

	val := queue.Front()
	assert.Equal(3, val)
	assert.Equal(5, queue.Size())
}

func TestCircularQueue_Back(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Back")

	queue := NewCircularQueue[int](6)
	assert.Equal(true, queue.IsEmpty())

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	queue.Print()
	assert.Equal(5, queue.Back())

	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue(6)
	queue.Enqueue(7)

	queue.Print()
	assert.Equal(7, queue.Back())
}

func TestCircularQueue_Contain(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Contain")

	queue := NewCircularQueue[int](2)
	queue.Enqueue(1)
	assert.Equal(true, queue.Contain(1))
	assert.Equal(false, queue.Contain(2))
}

func TestCircularQueue_Clear(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Clear")

	queue := NewCircularQueue[int](3)
	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Size())

	queue.Enqueue(1)
	assert.Equal(false, queue.IsEmpty())
	assert.Equal(1, queue.Size())

	queue.Clear()
	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Size())
}

func TestCircularQueue_Data(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Data")

	queue := NewCircularQueue[int](6)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	queue.Print()
	assert.Equal([]int{1, 2, 3, 4, 5}, queue.Data())

	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue(6)
	queue.Enqueue(7)

	queue.Print()
	assert.Equal([]int{3, 4, 5, 6, 7}, queue.Data())

}

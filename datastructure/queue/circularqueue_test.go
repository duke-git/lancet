package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestCircularQueue_EnQueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_EnQueue")

	queue := NewCircularQueue[int](6)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)

	queue.Print()
	// assert.Equal([]int{1, 2, 3, 4, 5}, queue.Data())
	assert.Equal(5, queue.Length())

	err := queue.EnQueue(6)
	assert.IsNotNil(err)
}

func TestCircularQueue_DeQueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_DeQueue")

	queue := NewCircularQueue[int](6)
	assert.Equal(true, queue.IsEmpty())

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)

	val, err := queue.DeQueue()
	assert.IsNil(err)

	assert.Equal(1, *val)
	assert.Equal(false, queue.IsFull())

	val, _ = queue.DeQueue()
	queue.Print()
	assert.Equal(2, *val)

	queue.EnQueue(6)
	queue.Print()
	assert.Equal(false, queue.IsFull())
}

func TestCircularQueue_Front(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Front")

	queue := NewCircularQueue[int](6)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)

	queue.Print()

	queue.DeQueue()
	queue.DeQueue()
	queue.EnQueue(6)
	queue.EnQueue(7)

	queue.Print()

	val := queue.Front()
	assert.Equal(3, val)
	assert.Equal(5, queue.Length())
}

func TestCircularQueue_Back(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Back")

	queue := NewCircularQueue[int](6)
	assert.Equal(true, queue.IsEmpty())

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)

	queue.Print()
	assert.Equal(5, queue.Back())

	queue.DeQueue()
	queue.DeQueue()
	queue.EnQueue(6)
	queue.EnQueue(7)

	queue.Print()
	assert.Equal(7, queue.Back())
}

func TestCircularQueue_Contain(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Contain")

	queue := NewCircularQueue[int](2)
	queue.EnQueue(1)
	assert.Equal(true, queue.Contain(1))
	assert.Equal(false, queue.Contain(2))
}

func TestCircularQueue_Clear(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Clear")

	queue := NewCircularQueue[int](3)
	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Length())

	queue.EnQueue(1)
	assert.Equal(false, queue.IsEmpty())
	assert.Equal(1, queue.Length())

	queue.Clear()
	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Length())
}

func TestCircularQueue_Data(t *testing.T) {
	assert := internal.NewAssert(t, "TestCircularQueue_Data")

	queue := NewCircularQueue[int](6)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)

	queue.Print()
	assert.Equal([]int{1, 2, 3, 4, 5}, queue.Data())

	queue.DeQueue()
	queue.DeQueue()
	queue.EnQueue(6)
	queue.EnQueue(7)

	queue.Print()
	assert.Equal([]int{3, 4, 5, 6, 7}, queue.Data())

}

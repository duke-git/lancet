package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestCircularQueue_Enqueue(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCircularQueue_Enqueue")

	queue := NewCircularQueue[int](6)

	err := queue.Enqueue(1)
	assert.IsNil(err)

	err = queue.Enqueue(2)
	assert.IsNil(err)

	err = queue.Enqueue(3)
	assert.IsNil(err)

	err = queue.Enqueue(4)
	assert.IsNil(err)

	err = queue.Enqueue(5)
	assert.IsNil(err)

	assert.Equal([]int{1, 2, 3, 4, 5}, queue.Data())
	assert.Equal(5, queue.Size())

	err = queue.Enqueue(6)
	assert.IsNotNil(err)
}

func TestCircularQueue_Dequeue(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCircularQueue_DeQueue")

	queue := NewCircularQueue[int](4)
	assert.Equal(true, queue.IsEmpty())

	err := queue.Enqueue(1)
	assert.IsNil(err)

	err = queue.Enqueue(2)
	assert.IsNil(err)

	err = queue.Enqueue(3)
	assert.IsNil(err)

	val, err := queue.Dequeue()
	assert.IsNil(err)

	assert.Equal(1, *val)
	assert.Equal(false, queue.IsFull())

	val, _ = queue.Dequeue()
	assert.Equal(2, *val)
	assert.Equal(false, queue.IsFull())
}

func TestCircularQueue_Front(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCircularQueue_Front")

	queue := NewCircularQueue[int](6)

	err := queue.Enqueue(1)
	assert.IsNil(err)

	err = queue.Enqueue(2)
	assert.IsNil(err)

	err = queue.Enqueue(3)
	assert.IsNil(err)

	val := queue.Front()
	assert.IsNil(err)
	assert.Equal(1, val)
	assert.Equal(3, queue.Size())
}

func TestCircularQueue_Back(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCircularQueue_Back")

	queue := NewCircularQueue[int](3)
	assert.Equal(true, queue.IsEmpty())

	err := queue.Enqueue(1)
	assert.IsNil(err)

	err = queue.Enqueue(2)
	assert.IsNil(err)

	assert.Equal(2, queue.Back())

	val, _ := queue.Dequeue()
	assert.Equal(1, *val)

	err = queue.Enqueue(3)
	assert.IsNil(err)

	assert.Equal(3, queue.Back())
}

func TestCircularQueue_Contain(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCircularQueue_Contain")

	queue := NewCircularQueue[int](2)
	err := queue.Enqueue(1)
	assert.IsNil(err)

	assert.Equal(true, queue.Contain(1))
	assert.Equal(false, queue.Contain(2))
}

func TestCircularQueue_Clear(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCircularQueue_Clear")

	queue := NewCircularQueue[int](3)
	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Size())

	err := queue.Enqueue(1)
	assert.IsNil(err)

	assert.Equal(false, queue.IsEmpty())
	assert.Equal(1, queue.Size())

	queue.Clear()
	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Size())
}

func TestCircularQueue_Data(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCircularQueue_Data")

	queue := NewCircularQueue[int](3)
	err := queue.Enqueue(1)
	assert.IsNil(err)

	err = queue.Enqueue(2)
	assert.IsNil(err)

	assert.Equal([]int{1, 2}, queue.Data())
}

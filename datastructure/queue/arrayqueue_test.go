package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestArrayQueue_EnQueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_EnQueue")

	queue := NewArrayQueue[int]()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)

	expected := []int{1, 2, 3}
	data := queue.Data()
	size := queue.Size()

	queue.Print()

	assert.Equal(expected, data)
	assert.Equal(3, size)
}

func TestArrayQueue_DeQueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_DeQueue")

	queue := NewArrayQueue(1, 2, 3)

	val, err := queue.DeQueue()
	if err != nil {
		t.Fail()
	}

	queue.Print()
	assert.Equal(1, *val)

	assert.Equal([]int{2, 3}, queue.Data())
}

func TestArrayQueue_Front(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Front")

	queue := NewArrayQueue(1, 2, 3)
	val := queue.Front()

	queue.Print()

	assert.Equal(1, val)
	assert.Equal([]int{1, 2, 3}, queue.Data())
}

func TestArrayQueue_Back(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Back")

	queue := NewArrayQueue(1, 2, 3)
	val := queue.Back()

	queue.Print()

	assert.Equal(3, val)
	assert.Equal([]int{1, 2, 3}, queue.Data())
}

func TestArrayQueue_Contain(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Contain")

	queue := NewArrayQueue(1, 2, 3)

	assert.Equal(true, queue.Contain(1))
	assert.Equal(false, queue.Contain(4))
}

func TestArrayQueue_Clear(t *testing.T) {
	assert := internal.NewAssert(t, "TestArrayQueue_Clear")

	queue := NewArrayQueue[int]()
	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Size())

	queue.EnQueue(1)
	assert.Equal(false, queue.IsEmpty())
	assert.Equal(1, queue.Size())

	queue.Clear()
	assert.Equal(true, queue.IsEmpty())
	assert.Equal(0, queue.Size())
}

package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestDoublyLink_InsertAtFirst(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDoublyLink_InsertAtFirst")

	link := NewDoublyLink[int]()
	link.InsertAtHead(1)
	link.InsertAtHead(2)
	link.InsertAtHead(3)
	link.Print()

	expected := []int{3, 2, 1}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestDoublyLink_InsertAtTail(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDoublyLink_InsertAtTail")

	link := NewDoublyLink[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.Print()

	expected := []int{1, 2, 3}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestDoublyLink_InsertAt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDoublyLink_InsertAt")

	link := NewDoublyLink[int]()
	link.InsertAt(1, 1) //do nothing
	link.InsertAt(0, 1)
	link.InsertAt(1, 2)
	link.InsertAt(2, 4)
	link.InsertAt(2, 3)

	expected := []int{1, 2, 3, 4}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestDoublyLink_DeleteAtHead(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDoublyLink_DeleteAtHead")

	link := NewDoublyLink[int]()
	link.DeleteAtHead()

	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	link.DeleteAtHead()

	expected := []int{2, 3, 4}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestDoublyLink_DeleteAtTail(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDoublyLink_DeleteAtTail")

	link := NewDoublyLink[int]()
	link.DeleteAtTail()

	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	link.DeleteAtTail()

	expected := []int{1, 2, 3}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestDoublyLink_DeleteAt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDoublyLink_DeleteAt")

	link := NewDoublyLink[int]()
	link.DeleteAt(0)

	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)
	link.InsertAtTail(5)

	link.DeleteAt(0)
	assert.Equal([]int{2, 3, 4, 5}, link.Values())

	link.DeleteAt(3)
	assert.Equal([]int{2, 3, 4}, link.Values())

	link.DeleteAt(1)
	assert.Equal(2, link.Size())
	assert.Equal([]int{2, 4}, link.Values())
}

func TestDoublyLink_Reverse(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDoublyLink_Reverse")

	link := NewDoublyLink[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	link.Reverse()
	link.Print()
	assert.Equal([]int{4, 3, 2, 1}, link.Values())
}

func TestDoublyLink_GetMiddleNode(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDoublyLink_GetMiddleNode")

	link := NewDoublyLink[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	middle1 := link.GetMiddleNode()
	assert.Equal(3, middle1.Value)

	link.InsertAtTail(5)
	link.InsertAtTail(6)
	link.InsertAtTail(7)
	middle2 := link.GetMiddleNode()
	assert.Equal(4, middle2.Value)
}

func TestDoublyLink_Clear(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDoublyLink_Clear")

	link := NewDoublyLink[int]()
	assert.Equal(true, link.IsEmpty())
	assert.Equal(0, link.Size())

	link.InsertAtTail(1)
	assert.Equal(false, link.IsEmpty())
	assert.Equal(1, link.Size())

	link.Clear()
	assert.Equal(true, link.IsEmpty())
	assert.Equal(0, link.Size())
}

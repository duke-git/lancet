package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestSinglyLink_InsertAtFirst(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_InsertAtFirst")

	link := NewSinglyLink[int]()
	link.InsertAtHead(1)
	link.InsertAtHead(2)
	link.InsertAtHead(3)
	link.Print()

	expected := []int{3, 2, 1}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestSinglyLink_InsertAtTail(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_InsertAtTail")

	link := NewSinglyLink[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.Print()

	expected := []int{1, 2, 3}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestSinglyLink_InsertAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_InsertAt")

	link := NewSinglyLink[int]()

	link.InsertAt(1, 1) //do nothing

	link.InsertAt(0, 1)
	link.InsertAt(1, 2)
	link.InsertAt(2, 4)
	link.InsertAt(2, 3)

	link.Print()

	expected := []int{1, 2, 3, 4}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestSinglyLink_DeleteAtHead(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_DeleteAtHead")

	link := NewSinglyLink[int]()

	link.DeleteAtHead()

	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	link.DeleteAtHead()
	link.Print()

	expected := []int{2, 3, 4}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestSinglyLink_DeleteAtTail(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_DeleteAtTail")

	link := NewSinglyLink[int]()

	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	link.DeleteAtTail()

	expected := []int{1, 2, 3}
	values := link.Values()

	assert.Equal(expected, values)
}

func TestSinglyLink_DeleteValue(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_DeleteValue")

	link := NewSinglyLink[int]()

	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	link.DeleteValue(2)
	assert.Equal([]int{1, 3, 4}, link.Values())

	link.DeleteValue(1)
	assert.Equal([]int{3, 4}, link.Values())
}

func TestSinglyLink_DeleteAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_DeleteAt")

	link := NewSinglyLink[int]()

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

func TestSinglyLink_Reverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_Reverse")

	link := NewSinglyLink[int]()
	link.InsertAtTail(1)
	link.InsertAtTail(2)
	link.InsertAtTail(3)
	link.InsertAtTail(4)

	link.Reverse()
	link.Print()
	assert.Equal([]int{4, 3, 2, 1}, link.Values())
}

func TestSinglyLink_GetMiddleNode(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_GetMiddleNode")

	link := NewSinglyLink[int]()
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

func TestSinglyLink_Clear(t *testing.T) {
	assert := internal.NewAssert(t, "TestSinglyLink_Clear")

	link := NewSinglyLink[int]()

	assert.Equal(true, link.IsEmpty())
	assert.Equal(0, link.Size())

	link.InsertAtTail(1)
	assert.Equal(false, link.IsEmpty())
	assert.Equal(1, link.Size())

	link.Clear()
	assert.Equal(true, link.IsEmpty())
	assert.Equal(0, link.Size())
}

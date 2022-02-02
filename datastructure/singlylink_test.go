package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/internal"
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

	err := link.InsertAt(1, 1)
	assert.IsNotNil(err)

	err = link.InsertAt(0, 1)
	err = link.InsertAt(1, 2)
	err = link.InsertAt(2, 4)
	err = link.InsertAt(2, 3)

	if err != nil {
		t.FailNow()
	}
	link.Print()

	expected := []int{1, 2, 3, 4}
	values := link.Values()

	assert.Equal(expected, values)

}

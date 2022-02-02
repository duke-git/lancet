package datastructure

import (
	"errors"
	"fmt"
)

// SinglyLink is a linkedlist of node, which have a value and Pre points to previous node, Next points to a next node of the link.
type SinglyLink[T any] struct {
	Head   *LinkNode[T]
	length int
}

// NewSinglyLink return *SinglyLink instance
func NewSinglyLink[T any]() *SinglyLink[T] {
	return &SinglyLink[T]{Head: nil}
}

// Values return slice of all singly linklist node value
func (link *SinglyLink[T]) Values() []T {
	res := []T{}
	current := link.Head
	for current != nil {
		res = append(res, current.Value)
		current = current.Next
	}
	return res
}

// Print all nodes info of a linked list
func (link *SinglyLink[T]) Print() {
	current := link.Head
	info := "[ "
	for current != nil {
		info += fmt.Sprintf("%+v, ", current)
		current = current.Next
	}
	info += " ]"
	fmt.Println(info)
}

// InsertAtHead insert value into singly linklist at head index
func (link *SinglyLink[T]) InsertAtHead(value T) {
	newNode := NewLinkNode(value)
	newNode.Next = link.Head
	link.Head = newNode
	link.length++
}

// InsertAtTail insert value into singly linklist at tail index
func (link *SinglyLink[T]) InsertAtTail(value T) {
	current := link.Head
	if current == nil {
		link.InsertAtHead(value)
		return
	}

	for current.Next != nil {
		current = current.Next
	}

	newNode := NewLinkNode(value)
	newNode.Next = nil
	current.Next = newNode

	link.length++
}

// InsertAt insert value into singly linklist at index
func (link *SinglyLink[T]) InsertAt(index int, value T) error {
	size := link.length
	if index < 0 || index > size {
		return errors.New("param index should between 0 and the length of singly link.")
	}

	if index == 0 {
		link.InsertAtHead(value)
		return nil
	}

	if index == size {
		link.InsertAtTail(value)
		return nil
	}

	i := 0
	current := link.Head

	for current != nil {
		if i == index-1 {
			newNode := NewLinkNode(value)
			newNode.Next = current.Next
			current.Next = newNode
			link.length++

			return nil
		}
		i++
		current = current.Next
	}

	return errors.New("singly link no exist")
}

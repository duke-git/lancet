package datastructure

import (
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/datastructure"
)

// DoublyLink is a linked list. Whose node has a generic Value, Pre pointer points to a previous node of the link, Next pointer points to a next node of the link.
type DoublyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}

// NewDoublyLink return *DoublyLink instance
func NewDoublyLink[T any]() *DoublyLink[T] {
	return &DoublyLink[T]{Head: nil}
}

// InsertAtHead insert value into doubly linklist at head index
func (link *DoublyLink[T]) InsertAtHead(value T) {
	newNode := datastructure.NewLinkNode(value)
	size := link.Size()

	if size == 0 {
		link.Head = newNode
		link.length++
		return
	}

	newNode.Next = link.Head
	newNode.Pre = nil

	link.Head.Pre = newNode
	link.Head = newNode

	link.length++
}

// InsertAtTail insert value into doubly linklist at tail index
func (link *DoublyLink[T]) InsertAtTail(value T) {
	current := link.Head
	if current == nil {
		link.InsertAtHead(value)
		return
	}

	for current.Next != nil {
		current = current.Next
	}

	newNode := datastructure.NewLinkNode(value)
	newNode.Next = nil
	newNode.Pre = current
	current.Next = newNode

	link.length++
}

// InsertAt insert value into doubly linklist at index
func (link *DoublyLink[T]) InsertAt(index int, value T) error {
	size := link.length
	if index < 0 || index > size {
		return errors.New("param index should between 0 and the length of doubly link.")
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
			newNode := datastructure.NewLinkNode(value)
			newNode.Next = current.Next
			newNode.Pre = current

			current.Next = newNode
			link.length++

			return nil
		}
		i++
		current = current.Next
	}

	return errors.New("doubly link list no exist")
}

// DeleteAtHead delete value in doubly linklist at head index
func (link *DoublyLink[T]) DeleteAtHead() error {
	if link.Head == nil {
		return errors.New("doubly link list no exist")
	}
	current := link.Head
	link.Head = current.Next
	link.Head.Pre = nil
	link.length--

	return nil
}

// DeleteAtTail delete value in doubly linklist at tail index
func (link *DoublyLink[T]) DeleteAtTail() error {
	if link.Head == nil {
		return errors.New("doubly link list no exist")
	}
	current := link.Head
	if current.Next == nil {
		return link.DeleteAtHead()
	}

	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
	link.length--
	return nil
}

// DeleteAt delete value in doubly linklist at index
func (link *DoublyLink[T]) DeleteAt(index int) error {
	if link.Head == nil {
		return errors.New("doubly link list no exist")
	}
	current := link.Head
	if current.Next == nil || index == 0 {
		return link.DeleteAtHead()
	}

	if index == link.length-1 {
		return link.DeleteAtTail()
	}

	if index < 0 || index > link.length-1 {
		return errors.New("param index should between 0 and link size -1.")
	}

	i := 0
	for current != nil {
		if i == index-1 {
			current.Next = current.Next.Next
			link.length--
			return nil
		}
		i++
		current = current.Next
	}

	return errors.New("delete error")
}

// Reverse the linked list
func (link *DoublyLink[T]) Reverse() {
	current := link.Head
	var temp *datastructure.LinkNode[T]

	for current != nil {
		temp = current.Pre
		current.Pre = current.Next
		current.Next = temp
		current = current.Pre
	}

	if temp != nil {
		link.Head = temp.Pre
	}
}

// GetMiddleNode return node at middle index of linked list
func (link *DoublyLink[T]) GetMiddleNode() *datastructure.LinkNode[T] {
	if link.Head == nil {
		return nil
	}
	if link.Head.Next == nil {
		return link.Head
	}
	fast := link.Head
	slow := link.Head

	for fast != nil {
		fast = fast.Next

		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		} else {
			return slow
		}
	}
	return slow
}

// Size return the count of doubly linked list
func (link *DoublyLink[T]) Size() int {
	return link.length
}

// Values return slice of all doubly linklist node value
func (link *DoublyLink[T]) Values() []T {
	result := []T{}
	current := link.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// Print all nodes info of a linked list
func (link *DoublyLink[T]) Print() {
	current := link.Head
	info := "[ "
	for current != nil {
		info += fmt.Sprintf("%+v, ", current)
		current = current.Next
	}
	info += " ]"
	fmt.Println(info)
}

// IsEmpty checks if link is empty or not
func (link *DoublyLink[T]) IsEmpty() bool {
	return link.length == 0
}

// Clear all nodes in doubly linklist
func (link *DoublyLink[T]) Clear() {
	link.Head = nil
	link.length = 0
}

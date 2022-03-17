package datastructure

import (
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/datastructure"
)

// SinglyLink is a linked list. Whose node has a Value generics and Next pointer points to a next node of the link.
type SinglyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}

// NewSinglyLink return *SinglyLink instance
func NewSinglyLink[T any]() *SinglyLink[T] {
	return &SinglyLink[T]{Head: nil}
}

// InsertAtHead insert value into singly linklist at head index
func (link *SinglyLink[T]) InsertAtHead(value T) {
	newNode := datastructure.NewLinkNode(value)
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

	newNode := datastructure.NewLinkNode(value)
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
			newNode := datastructure.NewLinkNode(value)
			newNode.Next = current.Next
			current.Next = newNode
			link.length++

			return nil
		}
		i++
		current = current.Next
	}

	return errors.New("singly link list no exist")
}

// DeleteAtHead delete value in singly linklist at head index
func (link *SinglyLink[T]) DeleteAtHead() error {
	if link.Head == nil {
		return errors.New("singly link list no exist")
	}
	current := link.Head
	link.Head = current.Next
	link.length--

	return nil
}

// DeleteAtTail delete value in singly linklist at tail index
func (link *SinglyLink[T]) DeleteAtTail() error {
	if link.Head == nil {
		return errors.New("singly link list no exist")
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

// DeleteAt delete value in singly linklist at index
func (link *SinglyLink[T]) DeleteAt(index int) error {
	if link.Head == nil {
		return errors.New("singly link list no exist")
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
func (link *SinglyLink[T]) Reverse() {
	var pre, next *datastructure.LinkNode[T]

	current := link.Head

	for current != nil {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}

	link.Head = pre
}

// GetMiddleNode return node at middle index of linked list
func (link *SinglyLink[T]) GetMiddleNode() *datastructure.LinkNode[T] {
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

// Size return the count of singly linked list
func (link *SinglyLink[T]) Size() int {
	return link.length
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

// IsEmpty checks if link is empty or not
func (link *SinglyLink[T]) IsEmpty() bool {
	return link.length == 0
}

// Clear checks if link is empty or not
func (link *SinglyLink[T]) Clear() {
	link.Head = nil
	link.length = 0
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

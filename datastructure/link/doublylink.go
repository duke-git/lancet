package datastructure

import (
	"fmt"

	"github.com/duke-git/lancet/v2/datastructure"
)

// DoublyLink is a linked list. Whose node has a generic Value, Pre pointer points to a previous node of the dl, Next pointer points to a next node of the dl.
type DoublyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}

// NewDoublyLink return *DoublyLink instance
func NewDoublyLink[T any]() *DoublyLink[T] {
	return &DoublyLink[T]{Head: nil}
}

// InsertAtHead insert value into doubly linklist at head index
func (dl *DoublyLink[T]) InsertAtHead(value T) {
	newNode := datastructure.NewLinkNode(value)
	size := dl.Size()

	if size == 0 {
		dl.Head = newNode
		dl.length++
		return
	}

	newNode.Next = dl.Head
	newNode.Pre = nil

	dl.Head.Pre = newNode
	dl.Head = newNode

	dl.length++
}

// InsertAtTail insert value into doubly linklist at tail index
func (dl *DoublyLink[T]) InsertAtTail(value T) {
	current := dl.Head
	if current == nil {
		dl.InsertAtHead(value)
		return
	}

	for current.Next != nil {
		current = current.Next
	}

	newNode := datastructure.NewLinkNode(value)
	newNode.Next = nil
	newNode.Pre = current
	current.Next = newNode

	dl.length++
}

// InsertAt insert value into doubly linklist at index
// param `index` should between [0, length], if index do not meet the conditions, do nothing
func (dl *DoublyLink[T]) InsertAt(index int, value T) {
	size := dl.length
	if index < 0 || index > size {
		return
	}

	if index == 0 {
		dl.InsertAtHead(value)
		return
	}

	if index == size {
		dl.InsertAtTail(value)
		return
	}

	i := 0
	current := dl.Head

	for current != nil {
		if i == index-1 {
			newNode := datastructure.NewLinkNode(value)
			newNode.Next = current.Next
			newNode.Pre = current

			current.Next = newNode
			dl.length++

			return
		}
		i++
		current = current.Next
	}
}

// DeleteAtHead delete value in doubly linklist at head index
func (dl *DoublyLink[T]) DeleteAtHead() {
	if dl.Head == nil {
		return
	}

	current := dl.Head
	dl.Head = current.Next
	dl.Head.Pre = nil
	dl.length--
}

// DeleteAtTail delete value in doubly linklist at tail
func (dl *DoublyLink[T]) DeleteAtTail() {
	if dl.Head == nil {
		return
	}

	current := dl.Head
	if current.Next == nil {
		dl.DeleteAtHead()
	}

	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
	dl.length--
}

// DeleteAt delete value in doubly linklist at index
// param `index` should be [0, len(DoublyLink)-1]
func (dl *DoublyLink[T]) DeleteAt(index int) {
	if dl.Head == nil {
		return
	}

	current := dl.Head
	if current.Next == nil || index == 0 {
		dl.DeleteAtHead()
	}

	if index == dl.length-1 {
		dl.DeleteAtTail()
	}

	if index < 0 || index > dl.length-1 {
		return
	}

	i := 0
	for current != nil {
		if i == index-1 {
			current.Next = current.Next.Next
			dl.length--
			return
		}
		i++
		current = current.Next
	}
}

// Reverse the linked list
func (dl *DoublyLink[T]) Reverse() {
	current := dl.Head
	var temp *datastructure.LinkNode[T]

	for current != nil {
		temp = current.Pre
		current.Pre = current.Next
		current.Next = temp
		current = current.Pre
	}

	if temp != nil {
		dl.Head = temp.Pre
	}
}

// GetMiddleNode return node at middle index of linked list
func (dl *DoublyLink[T]) GetMiddleNode() *datastructure.LinkNode[T] {
	if dl.Head == nil {
		return nil
	}
	if dl.Head.Next == nil {
		return dl.Head
	}
	fast := dl.Head
	slow := dl.Head

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
func (dl *DoublyLink[T]) Size() int {
	return dl.length
}

// Values return slice of all doubly linklist node value
func (dl *DoublyLink[T]) Values() []T {
	result := []T{}
	current := dl.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// Print all nodes info of a linked list
func (dl *DoublyLink[T]) Print() {
	current := dl.Head
	info := "[ "
	for current != nil {
		info += fmt.Sprintf("%+v, ", current)
		current = current.Next
	}
	info += " ]"
	fmt.Println(info)
}

// IsEmpty checks if dl is empty or not
func (dl *DoublyLink[T]) IsEmpty() bool {
	return dl.length == 0
}

// Clear all nodes in doubly linklist
func (dl *DoublyLink[T]) Clear() {
	dl.Head = nil
	dl.length = 0
}

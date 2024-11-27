// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure contains some data structure. Link structure contains SinglyLink and DoublyLink.
package datastructure

import (
	"fmt"
	"reflect"

	"github.com/duke-git/lancet/v2/datastructure"
)

// SinglyLink is a linked list. Whose node has a Value generics and Next pointer points to a next node of the sl.
type SinglyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}

// NewSinglyLink return *SinglyLink instance
func NewSinglyLink[T any]() *SinglyLink[T] {
	return &SinglyLink[T]{Head: nil}
}

// InsertAtHead insert value into singly linklist at head index
func (sl *SinglyLink[T]) InsertAtHead(value T) {
	newNode := datastructure.NewLinkNode(value)
	newNode.Next = sl.Head
	sl.Head = newNode
	sl.length++
}

// InsertAtTail insert value into singly linklist at tail index
func (sl *SinglyLink[T]) InsertAtTail(value T) {
	current := sl.Head
	if current == nil {
		sl.InsertAtHead(value)
		return
	}

	for current.Next != nil {
		current = current.Next
	}

	newNode := datastructure.NewLinkNode(value)
	newNode.Next = nil
	current.Next = newNode

	sl.length++
}

// InsertAt insert value into singly linklist at index
// param `index` should between [0, len(SinglyLink)], if index do not meet the conditions, do nothing
func (sl *SinglyLink[T]) InsertAt(index int, value T) {
	size := sl.length
	if index < 0 || index > size {
		return
	}

	if index == 0 {
		sl.InsertAtHead(value)
		return
	}

	if index == size {
		sl.InsertAtTail(value)
		return
	}

	i := 0
	current := sl.Head

	for current != nil {
		if i == index-1 {
			newNode := datastructure.NewLinkNode(value)
			newNode.Next = current.Next
			current.Next = newNode
			sl.length++
			return
		}
		i++
		current = current.Next
	}
}

// DeleteAtHead delete value in singly linklist at head index
func (sl *SinglyLink[T]) DeleteAtHead() {
	if sl.Head == nil {
		return
	}

	current := sl.Head
	sl.Head = current.Next
	sl.length--
}

// DeleteAtTail delete value in singly linklist at tail
func (sl *SinglyLink[T]) DeleteAtTail() {
	if sl.Head == nil {
		return
	}

	current := sl.Head
	if current.Next == nil {
		sl.DeleteAtHead()
	}

	for current.Next.Next != nil {
		current = current.Next
	}

	current.Next = nil
	sl.length--
}

// DeleteAt delete value in singly linklist at index
// param `index` should be [0, len(SinglyLink)-1]
func (sl *SinglyLink[T]) DeleteAt(index int) {
	if sl.Head == nil {
		return
	}
	current := sl.Head
	if current.Next == nil || index == 0 {
		sl.DeleteAtHead()
	}

	if index == sl.length-1 {
		sl.DeleteAtTail()
	}

	if index < 0 || index > sl.length-1 {
		return
	}

	i := 0
	for current != nil {
		if i == index-1 {
			current.Next = current.Next.Next
			sl.length--
			return
		}
		i++
		current = current.Next
	}
}

// DeleteValue delete value in singly linklist
func (sl *SinglyLink[T]) DeleteValue(value T) {
	if sl.Head == nil {
		return
	}
	dummyHead := datastructure.NewLinkNode(value)
	dummyHead.Next = sl.Head
	current := dummyHead

	for current.Next != nil {
		if reflect.DeepEqual(current.Next.Value, value) {
			current.Next = current.Next.Next
			sl.length--
		} else {
			current = current.Next
		}
	}

	sl.Head = dummyHead.Next
}

// Reverse the linked list
func (sl *SinglyLink[T]) Reverse() {
	var pre, next *datastructure.LinkNode[T]

	current := sl.Head

	for current != nil {
		next = current.Next
		current.Next = pre
		pre = current
		current = next
	}

	sl.Head = pre
}

// GetMiddleNode return node at middle index of linked list
func (sl *SinglyLink[T]) GetMiddleNode() *datastructure.LinkNode[T] {
	if sl.Head == nil {
		return nil
	}
	if sl.Head.Next == nil {
		return sl.Head
	}
	fast := sl.Head
	slow := sl.Head

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
func (sl *SinglyLink[T]) Size() int {
	return sl.length
}

// Values return slice of all singly linklist node value
func (sl *SinglyLink[T]) Values() []T {
	result := make([]T, 0, sl.length)
	current := sl.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// IsEmpty checks if sl is empty or not
func (sl *SinglyLink[T]) IsEmpty() bool {
	return sl.length == 0
}

// Clear all the node in singly linklist
func (sl *SinglyLink[T]) Clear() {
	sl.Head = nil
	sl.length = 0
}

// Print all nodes info of a linked list
func (sl *SinglyLink[T]) Print() {
	current := sl.Head
	info := "[ "
	for current != nil {
		info += fmt.Sprintf("%+v, ", current)
		current = current.Next
	}
	info += " ]"
	fmt.Println(info)
}

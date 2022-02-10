package datastructure

import (
	"errors"
	"fmt"
)

// LinkedStack implements stack with link list
type LinkedStack[T any] struct {
	top    *StackNode[T]
	length int
}

// NewLinkedStack return a empty LinkedStack pointer
func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{top: nil, length: 0}
}

// Data return stack data
func (s *LinkedStack[T]) Data() []T {
	res := []T{}
	current := s.top

	for current != nil {
		res = append(res, current.Value)
		current = current.Next
	}
	return res
}

// Length return length of stack data
func (s *LinkedStack[T]) Length() int {
	return s.length
}

// IsEmpty checks if stack is empty or not
func (s *LinkedStack[T]) IsEmpty() bool {
	return s.length == 0
}

// Push element into stack
func (s *LinkedStack[T]) Push(value T) {
	newNode := NewStackNode(value)
	top := s.top
	if top == nil {
		s.top = newNode
	} else {
		newNode.Next = top
		s.top = newNode
	}

	s.length++
}

// Pop delete the top element of stack then return it, if stack is empty, return nil and error
func (s *LinkedStack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	top := s.top
	s.top = s.top.Next
	s.length--

	return &top.Value, nil
}

// Peak return the top element of stack then return it
func (s *LinkedStack[T]) Peak() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.top.Value, nil
}

// Clear clear the stack data
func (s *LinkedStack[T]) Clear() {
	s.top = nil
	s.length = 0
}

// Print all nodes info of stack link
func (s *LinkedStack[T]) Print() {
	current := s.top
	info := "[ "
	for current != nil {
		info += fmt.Sprintf("%+v, ", current)
		current = current.Next
	}
	info += " ]"
	fmt.Println(info)
}

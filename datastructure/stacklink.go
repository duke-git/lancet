package datastructure

import (
	"errors"
	"fmt"
)

// StackLink is a linear table, implemented with slice
type StackLink[T any] struct {
	top    *StackNode[T]
	length int
}

// NewStackLink return a empty StackLink pointer
func NewStackLink[T any]() *StackLink[T] {
	return &StackLink[T]{top: nil, length: 0}
}

// Data return stack data
func (s *StackLink[T]) Data() []T {
	res := []T{}
	current := s.top

	for current != nil {
		res = append(res, current.Value)
		current = current.Next
	}
	return res
}

// Length return length of stack data
func (s *StackLink[T]) Length() int {
	return s.length
}

// IsEmpty checks if stack is empty or not
func (s *StackLink[T]) IsEmpty() bool {
	return s.length == 0
}

// Push element into stack
func (s *StackLink[T]) Push(value T) {
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
func (s *StackLink[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	top := s.top
	s.top = s.top.Next
	s.length--

	return &top.Value, nil
}

// Peak return the top element of stack then return it
func (s *StackLink[T]) Peak() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.top.Value, nil
}

// EmptyStack clear the stack data
func (s *StackLink[T]) EmptyStack() {
	s.top = nil
	s.length = 0
}

// Print all nodes info of stack link
func (s *StackLink[T]) Print() {
	current := s.top
	info := "[ "
	for current != nil {
		info += fmt.Sprintf("%+v, ", current)
		current = current.Next
	}
	info += " ]"
	fmt.Println(info)
}

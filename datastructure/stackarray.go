package datastructure

import "errors"

// StackArray implements stack with slice
type StackArray[T any] struct {
	data   []T
	length int
}

// NewStackArray return a empty StackArray pointer
func NewStackArray[T any]() *StackArray[T] {
	return &StackArray[T]{data: []T{}, length: 0}
}

// Data return stack data
func (s *StackArray[T]) Data() []T {
	return s.data
}

// Length return length of stack data
func (s *StackArray[T]) Length() int {
	return s.length
}

// IsEmpty checks if stack is empty or not
func (s *StackArray[T]) IsEmpty() bool {
	return s.length == 0
}

// Push element into stack
func (s *StackArray[T]) Push(value T) {
	s.data = append([]T{value}, s.data...)
	s.length++
}

// Pop delete the top element of stack then return it, if stack is empty, return nil and error
func (s *StackArray[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	topItem := s.data[0]
	s.data = s.data[1:]
	s.length--

	return &topItem, nil
}

// Peak return the top element of stack then return it
func (s *StackArray[T]) Peak() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.data[0], nil
}

// EmptyStack clear the stack data
func (s *StackArray[T]) EmptyStack() {
	s.data = []T{}
	s.length = 0
}

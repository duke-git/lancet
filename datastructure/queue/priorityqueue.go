package datastructure

import "github.com/duke-git/lancet/v2/lancetconstraints"

// PriorityQueue is a priority queue implemented by binary heap tree
// type T should implements Compare function in lancetconstraints.Comparator interface.
type PriorityQueue[T any] struct {
	items      []T
	size       int
	comparator lancetconstraints.Comparator
}

// NewPriorityQueue return a pointer of PriorityQueue
// param `comparator` is used to compare valuse in the heap
func NewPriorityQueue[T any](capacity int, comparator lancetconstraints.Comparator) *NewPriorityQueue[T] {
	return &NewPriorityQueue[T]{
		items:      make([]T, capacity+1),
		size:       0,
		comparator: comparator,
	}
}

// IsEmpty checks if the queue is empty or not
func (q *PriorityQueue[T]) IsEmpty() bool {
	return q.size == 0
}

// IsFull checks if the queue capacity is full or not
func (q *PriorityQueue[T]) IsFull() bool {
	return q.size == len(q.items)
}

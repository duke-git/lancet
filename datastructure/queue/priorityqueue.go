package datastructure

import (
	"errors"

	"github.com/duke-git/lancet/v2/lancetconstraints"
)

// PriorityQueue is a priority queue implemented by binary heap tree
// type T should implements Compare function in lancetconstraints.Comparator interface.
type PriorityQueue[T any] struct {
	items      []T
	size       int
	comparator lancetconstraints.Comparator
}

// NewPriorityQueue return a pointer of PriorityQueue
// param `comparator` is used to compare valuse in the heap
func NewPriorityQueue[T any](capacity int, comparator lancetconstraints.Comparator) *PriorityQueue[T] {
	return &PriorityQueue[T]{
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
	return q.size == len(q.items)-1
}

// Enqueue insert value into queue
func (q *PriorityQueue[T]) Enqueue(val T) error {
	if q.IsFull() {
		return errors.New("queue is already full.")
	}
	q.size++
	q.items[q.size] = val
	q.swim(q.size)
	return nil
}

// swim when child's key is larger than parent's key, exchange them.
func (q *PriorityQueue[T]) swim(index int) {
	for index > 1 && q.comparator.Compare(index/2, index) < 0 {
		q.swap(index, index/2)
		index = index / 2
	}
}

// swap the two values at index i and j
func (q *PriorityQueue[T]) swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

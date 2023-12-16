// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure contains some data structure.
// Queue structure contains ArrayQueue, LinkedQueue, CircularQueue, and PriorityQueue.
package datastructure

import (
	"errors"

	"github.com/duke-git/lancet/v2/constraints"
)

// PriorityQueue is a priority queue implemented by binary heap tree
// type T should implements Compare function in constraints.Comparator interface.
type PriorityQueue[T any] struct {
	items      []T
	size       int
	comparator constraints.Comparator
}

// NewPriorityQueue return a pointer of PriorityQueue
// param `comparator` is used to compare values in the queue
func NewPriorityQueue[T any](capacity int, comparator constraints.Comparator) *PriorityQueue[T] {
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

// Size get number of items in the queue
func (q *PriorityQueue[T]) Size() int {
	return q.size
}

// IsFull checks if the queue capacity is full or not
func (q *PriorityQueue[T]) IsFull() bool {
	return q.size == len(q.items)-1
}

// Data return a slice of queue data
func (q *PriorityQueue[T]) Data() []T {
	data := make([]T, q.size)
	for i := 1; i < q.size+1; i++ {
		data[i-1] = q.items[i]
	}

	return data
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

// Dequeue delete and return max value in queue
func (q *PriorityQueue[T]) Dequeue() (T, bool) {
	var val T
	if q.IsEmpty() {
		return val, false
	}

	max := q.items[1]

	q.swap(1, q.size)
	q.size--
	q.sink(1)

	//set zero value for rest values of the queue
	q.items[q.size+1] = val

	return max, true
}

// swim when child's key is larger than parent's key, exchange them.
func (q *PriorityQueue[T]) swim(index int) {
	for index > 1 && q.comparator.Compare(q.items[index/2], q.items[index]) < 0 {
		q.swap(index, index/2)
		index = index / 2
	}
}

// sink when parent's key smaller than child's key, exchange parent's key with larger child's key.
func (q *PriorityQueue[T]) sink(index int) {

	for 2*index <= q.size {
		j := 2 * index

		// get larger child node index
		if j < q.size && q.comparator.Compare(q.items[j], q.items[j+1]) < 0 {
			j++
		}
		// if parent larger than child, stop
		if !(q.comparator.Compare(q.items[index], q.items[j]) < 0) {
			break
		}

		q.swap(index, j)
		index = j
	}
}

// swap the two values at index i and j
func (q *PriorityQueue[T]) swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

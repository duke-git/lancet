// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure contains some data structure.
// Queue structure contains ArrayQueue, LinkedQueue, CircularQueue, and PriorityQueue.
package datastructure

import (
	"fmt"
	"reflect"
)

// ArrayQueue implements queue with slice
type ArrayQueue[T any] struct {
	data     []T
	head     int
	tail     int
	capacity int
	size     int
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		data:     make([]T, 0, capacity),
		head:     0,
		tail:     0,
		capacity: capacity,
		size:     0,
	}
}

// Data return slice of queue data
func (q *ArrayQueue[T]) Data() []T {
	items := []T{}
	for i := q.head; i < q.tail; i++ {
		items = append(items, q.data[i])
	}
	return items
}

// Size return number of elements in queue
func (q *ArrayQueue[T]) Size() int {
	return q.size
}

// IsEmpty checks if queue is empty or not
func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.size == 0
}

// IsFull checks if queue is full or not
func (q *ArrayQueue[T]) IsFull() bool {
	return q.size == q.capacity
}

// Front return front value of queue
func (q *ArrayQueue[T]) Front() T {
	return q.data[0]
}

// Back return back value of queue
func (q *ArrayQueue[T]) Back() T {
	return q.data[q.size-1]
}

// EnQueue put element into queue
func (q *ArrayQueue[T]) Enqueue(item T) bool {
	if q.tail < q.capacity {
		q.data = append(q.data, item)
		// q.tail++
		q.data[q.tail] = item
	} else {
		//upgrade
		if q.head > 0 {
			for i := 0; i < q.tail-q.head; i++ {
				q.data[i] = q.data[i+q.head]
			}
			q.tail -= q.head
			q.head = 0
		} else {
			if q.capacity < 65536 {
				if q.capacity == 0 {
					q.capacity = 1
				}
				q.capacity *= 2
			} else {
				q.capacity += 2 ^ 16
			}

			tmp := make([]T, q.capacity, q.capacity)
			copy(tmp, q.data)
			q.data = tmp
		}

		q.data[q.tail] = item
	}

	q.tail++
	q.size++

	return true
}

// DeQueue remove head element of queue and return it, if queue is empty, return nil and error
func (q *ArrayQueue[T]) Dequeue() (T, bool) {
	var item T
	if q.size == 0 {
		return item, false
	}

	item = q.data[q.head]
	q.head++

	if q.head >= 1024 || q.head*2 > q.tail {
		q.capacity -= q.head
		q.tail -= q.head
		tmp := make([]T, q.capacity, q.capacity)
		copy(tmp, q.data[q.head:])
		q.data = tmp
		q.head = 0
	}

	q.size--
	return item, true
}

// Clear the queue data
func (q *ArrayQueue[T]) Clear() {
	capacity := q.capacity
	q.data = make([]T, 0, capacity)
	q.head = 0
	q.tail = 0
	q.size = 0
	q.capacity = capacity
}

// Contain checks if the value is in queue or not
func (q *ArrayQueue[T]) Contain(value T) bool {
	for _, v := range q.data {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}
	return false
}

// Print queue data
func (q *ArrayQueue[T]) Print() {
	info := "["
	for i := q.head; i < q.tail; i++ {
		info += fmt.Sprintf("%+v, ", q.data[i])
	}
	info += "]"
	fmt.Println(info)
}

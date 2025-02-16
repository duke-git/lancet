// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure contains some data structure.
// Queue structure contains ArrayQueue, LinkedQueue, CircularQueue, and PriorityQueue.
package datastructure

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/duke-git/lancet/v2/datastructure"
)

// LinkedQueue implements queue with link list
type LinkedQueue[T any] struct {
	head   *datastructure.QueueNode[T]
	tail   *datastructure.QueueNode[T]
	length int
}

// NewLinkedQueue return a empty LinkedQueue pointer
func NewLinkedQueue[T any]() *LinkedQueue[T] {
	return &LinkedQueue[T]{head: nil, tail: nil, length: 0}
}

// Data return slice of queue data
func (q *LinkedQueue[T]) Data() []T {
	res := make([]T, 0, q.length)
	current := q.head

	for current != nil {
		res = append(res, current.Value)
		current = current.Next
	}
	return res
}

// Size return length of queue data
func (q *LinkedQueue[T]) Size() int {
	return q.length
}

// IsEmpty checks if queue is empty or not
func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.length == 0
}

// Enqueue put element into queue
func (q *LinkedQueue[T]) Enqueue(value T) {
	newNode := datastructure.NewQueueNode(value)

	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.Next = newNode
		q.tail = newNode
	}
	q.length++
}

// Dequeue delete head element of queue then return it, if queue is empty, return nil and error
func (q *LinkedQueue[T]) Dequeue() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	head := q.head
	q.head = q.head.Next
	q.length--

	return &head.Value, nil
}

// Front return front value of queue
func (q *LinkedQueue[T]) Front() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return &q.head.Value, nil
}

// Back return back value of queue
func (q *LinkedQueue[T]) Back() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return &q.tail.Value, nil
}

// Clear clear the queue data
func (q *LinkedQueue[T]) Clear() {
	q.head = nil
	q.tail = nil
	q.length = 0
}

// Print all nodes info of queue link
func (q *LinkedQueue[T]) Print() {
	current := q.head
	info := "[ "
	for current != nil {
		info += fmt.Sprintf("%+v, ", current)
		current = current.Next
	}
	info += " ]"
	fmt.Println(info)
}

// Contain checks if the value is in queue or not
func (q *LinkedQueue[T]) Contain(value T) bool {
	current := q.head
	for current != nil {
		if reflect.DeepEqual(current.Value, value) {
			return true
		}
		current = current.Next
	}
	return false
}

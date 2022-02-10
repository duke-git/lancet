package datastructure

import (
	"errors"
	"fmt"
	"reflect"
)

// ArrayQueue implements queue with slice
type ArrayQueue[T any] struct {
	data   []T
	length int
}

// NewArrayQueue return a empty ArrayQueue pointer
func NewArrayQueue[T any](values ...T) *ArrayQueue[T] {
	data := make([]T, len(values))
	for i, v := range values {
		data[i] = v
	}
	return &ArrayQueue[T]{data: data, length: len(values)}
}

// Data return queue data
func (q *ArrayQueue[T]) Data() []T {
	return q.data
}

// Size return length of queue data
func (q *ArrayQueue[T]) Size() int {
	return q.length
}

// IsEmpty checks if queue is empty or not
func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.length == 0
}

// Front return front value of queue
func (q *ArrayQueue[T]) Front() T {
	return q.data[0]
}

// Back return back value of queue
func (q *ArrayQueue[T]) Back() T {
	return q.data[q.length-1]
}

// EnQueue put element into queue
func (q *ArrayQueue[T]) EnQueue(value T) {
	q.data = append(q.data, value)
	q.length++
}

// DeQueue get the head element of queue, if queue is empty, return nil and error
func (q *ArrayQueue[T]) DeQueue() (*T, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	headItem := q.data[0]
	q.data = q.data[1:]
	q.length--

	return &headItem, nil
}

// Clear the queue data
func (q *ArrayQueue[T]) Clear() {
	q.data = []T{}
	q.length = 0
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
	for _, v := range q.data {
		info += fmt.Sprintf("%+v, ", v)
	}
	info += "]"
	fmt.Println(info)
}

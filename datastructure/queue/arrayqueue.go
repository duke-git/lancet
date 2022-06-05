package datastructure

import (
	"fmt"
	"reflect"
)

// ArrayQueue implements queue with slice
type ArrayQueue[T any] struct {
	items    []T
	head     int
	tail     int
	capacity int
	size     int
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		items:    make([]T, 0, capacity),
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
		items = append(items, q.items[i])
	}
	return items
}

// Size return length of queue data
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
	return q.items[0]
}

// Back return back value of queue
func (q *ArrayQueue[T]) Back() T {
	return q.items[q.size-1]
}

// EnQueue put element into queue
func (q *ArrayQueue[T]) Enqueue(item T) bool {
	if q.head == 0 && q.tail == q.capacity {
		return false
	} else if q.head != 0 && q.tail == q.capacity {
		for i := q.head; i < q.tail; i++ {
			q.items[i-q.head] = q.items[i]
		}
		q.tail = q.tail - q.head
		q.head = 0
	}

	q.items = append(q.items, item)
	q.tail++
	q.size++
	return true
}

// DeQueue remove head element of queue and return it, if queue is empty, return nil and error
func (q *ArrayQueue[T]) Dequeue() (T, bool) {
	var item T
	if q.head == q.tail {
		return item, false
	}
	item = q.items[q.head]
	q.head++
	q.size--
	return item, true
}

// Clear the queue data
func (q *ArrayQueue[T]) Clear() {
	capacity := q.capacity
	q.items = make([]T, 0, capacity)
	q.head = 0
	q.tail = 0
	q.size = 0
	q.capacity = capacity
}

// Contain checks if the value is in queue or not
func (q *ArrayQueue[T]) Contain(value T) bool {
	for _, v := range q.items {
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
		info += fmt.Sprintf("%+v, ", q.items[i])
	}
	info += "]"
	fmt.Println(info)
}

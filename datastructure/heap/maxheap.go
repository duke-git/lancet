// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure implements some data structure. eg. list, linklist, stack, queue, tree, graph.
package datastructure

import (
	"github.com/duke-git/lancet/v2/lancetconstraints"
)

// MaxHeap implements a binary max heap
// type T should implements Compare function in lancetconstraints.Comparator interface.
type MaxHeap[T any] struct {
	data       []T
	comparator lancetconstraints.Comparator
}

// NewMaxHeap returns a MaxHeap instance with the given comparator.
func NewMaxHeap[T any](comparator lancetconstraints.Comparator) *MaxHeap[T] {
	return &MaxHeap[T]{
		data:       make([]T, 0),
		comparator: comparator,
	}
}

// heapifyUp heapify the data from bottom to up
func (h *MaxHeap[T]) heapifyUp(i int) {
	for h.comparator.Compare(h.data[parentIndex(i)], h.data[i]) < 0 {
		h.swap(parentIndex(i), i)
		i = parentIndex(i)
	}
}

// parentIndex get parent index of the given index
func parentIndex(i int) int {
	return (i - 1) / 2
}

// leftChildIndex get left child index of the given index
func leftChildIndex(i int) int {
	return 2*i + 1
}

// rightChildIndex get right child index of the given index
func rightChildIndex(i int) int {
	return 2*i + 2
}

// swap two elements in the heap
func (h *MaxHeap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

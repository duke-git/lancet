package datastructure

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func TestMaxHeap_BuildMaxHeap(t *testing.T) {
	assert := internal.NewAssert(t, "TestMaxHeap_BuildMaxHeap")

	values := []int{6, 5, 2, 4, 7, 10, 12, 1, 3, 8, 9, 11}
	heap := BuildMaxHeap(values, &intComparator{})

	expected := []int{12, 9, 11, 4, 8, 10, 7, 1, 3, 5, 6, 2}
	assert.Equal(expected, heap.data)

	assert.Equal(12, heap.Size())

	heap.PrintStructure()
}

func TestMaxHeap_Push(t *testing.T) {
	assert := internal.NewAssert(t, "TestMaxHeap_Push")

	heap := NewMaxHeap[int](&intComparator{})
	values := []int{6, 5, 2, 4, 7, 10, 12, 1, 3, 8, 9, 11}

	for _, v := range values {
		heap.Push(v)
	}

	expected := []int{12, 9, 11, 4, 8, 10, 7, 1, 3, 5, 6, 2}
	assert.Equal(expected, heap.data)

	assert.Equal(12, heap.Size())

	heap.PrintStructure()
}

func TestMaxHeap_Pop(t *testing.T) {
	assert := internal.NewAssert(t, "TestMaxHeap_Pop")

	heap := NewMaxHeap[int](&intComparator{})

	_, ok := heap.Pop()
	assert.Equal(false, ok)

	values := []int{6, 5, 2, 4, 7, 10, 12, 1, 3, 8, 9, 11}
	for _, v := range values {
		heap.Push(v)
	}

	val, ok := heap.Pop()
	assert.Equal(12, val)
	assert.Equal(true, ok)
	assert.Equal(11, heap.Size())
}

func TestMaxHeap_Peek(t *testing.T) {
	assert := internal.NewAssert(t, "TestMaxHeap_Peek")

	heap := NewMaxHeap[int](&intComparator{})

	_, ok := heap.Peek()
	assert.Equal(false, ok)

	values := []int{6, 5, 2, 4, 7, 10, 12, 1, 3, 8, 9, 11}
	for _, v := range values {
		heap.Push(v)
	}

	val, ok := heap.Peek()
	assert.Equal(12, val)
	assert.Equal(true, ok)

	assert.Equal(12, heap.Size())
}

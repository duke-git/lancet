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
func TestPriorityQueue_Enqueue(t *testing.T) {
	assert := internal.NewAssert(t, "TestPriorityQueue_Enqueue")

	comparator := &intComparator{}
	pq := NewPriorityQueue[int](10, comparator)

	assert.Equal(true, pq.IsEmpty())
	assert.Equal(false, pq.IsFull())

	for i := 1; i < 11; i++ {
		pq.Enqueue(i)
	}

	assert.Equal(true, pq.IsFull())

	queueData := pq.Data()
	assert.Equal([]int{10, 9, 6, 7, 8, 2, 5, 1, 4, 3}, queueData)

}

func TestPriorityQueue_Dequeue(t *testing.T) {
	assert := internal.NewAssert(t, "TestPriorityQueue_Dequeue")

	comparator := &intComparator{}
	pq := NewPriorityQueue[int](10, comparator)

	_, ok := pq.Dequeue()
	assert.Equal(false, ok)

	for i := 1; i < 11; i++ {
		pq.Enqueue(i)
	}

	val, ok := pq.Dequeue()
	assert.Equal(true, ok)
	assert.Equal(10, val)

	assert.Equal([]int{9, 8, 6, 7, 3, 2, 5, 1, 4}, pq.Data())
}

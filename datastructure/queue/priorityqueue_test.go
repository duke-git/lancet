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
	pq := NewPriorityQueue[int](3, comparator)

	assert.Equal(true, pq.IsEmpty())
	assert.Equal(false, pq.IsFull())

	err := pq.Enqueue(1)
	assert.IsNil(err)

	err = pq.Enqueue(2)
	assert.IsNil(err)

	err = pq.Enqueue(3)
	assert.IsNil(err)

	assert.Equal(true, pq.IsFull())

	queueData := pq.Data()
	assert.Equal([]int{3, 1, 2}, queueData)

}

func TestPriorityQueue_Dequeue(t *testing.T) {
	assert := internal.NewAssert(t, "TestPriorityQueue_Dequeue")

	comparator := &intComparator{}
	pq := NewPriorityQueue[int](3, comparator)

	_, ok := pq.Dequeue()
	assert.Equal(false, ok)

	err := pq.Enqueue(1)
	assert.IsNil(err)

	err = pq.Enqueue(2)
	assert.IsNil(err)

	err = pq.Enqueue(3)
	assert.IsNil(err)

	assert.Equal(3, pq.Size())

	val, ok := pq.Dequeue()
	assert.Equal(true, ok)
	assert.Equal(3, val)
}

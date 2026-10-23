package datastructure

import "testing"

func TestCircularQueueReuseAfterClear(t *testing.T) {
	queue := NewCircularQueue[int](3)
	queue.Enqueue(1)
	queue.Clear()
	if err := queue.Enqueue(2); err != nil {
		t.Fatal(err)
	}
	if queue.Size() != 1 || queue.Front() != 2 {
		t.Error("cleared queue could not be reused")
	}
	zero := NewCircularQueue[int](0)
	if !zero.IsFull() || zero.Enqueue(1) == nil {
		t.Error("zero-capacity queue should reject an element")
	}
}

package datastructure

import "testing"

func TestQueueContainActiveValues(t *testing.T) {
	array := NewArrayQueue[int](5)
	array.Enqueue(1)
	array.Enqueue(2)
	array.Dequeue()
	if array.Contain(1) || !array.Contain(2) {
		t.Error("array queue includes a dequeued value")
	}
	circular := NewCircularQueue[int](4)
	if circular.Contain(0) {
		t.Error("empty circular queue contains a zero value")
	}
	circular.Enqueue(1)
	circular.Enqueue(2)
	circular.Dequeue()
	circular.Enqueue(3)
	circular.Enqueue(4)
	if circular.Contain(0) || circular.Contain(1) || !circular.Contain(4) {
		t.Error("circular queue did not restrict its search to active values")
	}
}

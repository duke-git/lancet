package datastructure

// Node is a linkedlist node, which have a value and Pre points to previous node, Next points to a next node of the link.
type Node[T any] struct {
	Value T
	Pre   *Node[T]
	Next  *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value, nil, nil}
}

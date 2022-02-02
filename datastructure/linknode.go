package datastructure

// LinkNode is a linkedlist node, which have a value and Pre points to previous node, Next points to a next node of the link.
type LinkNode[T any] struct {
	Value T
	Pre   *LinkNode[T]
	Next  *LinkNode[T]
}

func NewLinkNode[T any](value T) *LinkNode[T] {
	return &LinkNode[T]{value, nil, nil}
}


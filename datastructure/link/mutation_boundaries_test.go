package datastructure

import (
	"reflect"
	"testing"
)

type mutableIntLink interface {
	InsertAtHead(int)
	InsertAtTail(int)
	DeleteAtHead()
	DeleteAtTail()
	DeleteAt(int)
	Reverse()
	Values() []int
	Size() int
	IsEmpty() bool
}

func TestLinkDeletionBoundaries(t *testing.T) {
	for name, newLink := range map[string]func() mutableIntLink{
		"singly": func() mutableIntLink { return NewSinglyLink[int]() },
		"doubly": func() mutableIntLink { return NewDoublyLink[int]() },
	} {
		t.Run(name, func(t *testing.T) {
			for _, tt := range []struct {
				name   string
				size   int
				delete func(mutableIntLink)
				want   []int
			}{
				{"empty head", 0, func(l mutableIntLink) { l.DeleteAtHead() }, []int{}},
				{"empty tail", 0, func(l mutableIntLink) { l.DeleteAtTail() }, []int{}},
				{"empty index", 0, func(l mutableIntLink) { l.DeleteAt(0) }, []int{}},
				{"singleton head", 1, func(l mutableIntLink) { l.DeleteAtHead() }, []int{}},
				{"singleton tail", 1, func(l mutableIntLink) { l.DeleteAtTail() }, []int{}},
				{"singleton index", 1, func(l mutableIntLink) { l.DeleteAt(0) }, []int{}},
				{"singleton negative index", 1, func(l mutableIntLink) { l.DeleteAt(-1) }, []int{1}},
				{"singleton large index", 1, func(l mutableIntLink) { l.DeleteAt(1) }, []int{1}},
				{"two item head index", 2, func(l mutableIntLink) { l.DeleteAt(0) }, []int{2}},
				{"two item tail index", 2, func(l mutableIntLink) { l.DeleteAt(1) }, []int{1}},
				{"middle index", 4, func(l mutableIntLink) { l.DeleteAt(1) }, []int{1, 3, 4}},
				{"tail index", 4, func(l mutableIntLink) { l.DeleteAt(3) }, []int{1, 2, 3}},
				{"negative index", 4, func(l mutableIntLink) { l.DeleteAt(-1) }, []int{1, 2, 3, 4}},
				{"large index", 4, func(l mutableIntLink) { l.DeleteAt(4) }, []int{1, 2, 3, 4}},
			} {
				t.Run(tt.name, func(t *testing.T) {
					defer func() {
						if value := recover(); value != nil {
							t.Errorf("list mutation panicked: %v", value)
						}
					}()
					l := newLink()
					for i := 1; i <= tt.size; i++ {
						l.InsertAtTail(i)
					}
					tt.delete(l)
					if got := l.Values(); !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("Values() = %v, want %v", got, tt.want)
					}
					if l.Size() != len(tt.want) || l.IsEmpty() != (len(tt.want) == 0) {
						t.Fatalf("Size() = %d, IsEmpty() = %v, want length %d", l.Size(), l.IsEmpty(), len(tt.want))
					}
					if doubly, ok := l.(*DoublyLink[int]); ok {
						current := doubly.Head
						for i := 0; current != nil; i++ {
							if i >= l.Size() {
								t.Fatal("cycle in list after deletion")
							}
							if current.Next != nil && current.Next.Pre != current {
								t.Fatal("deletion did not preserve reciprocal links")
							}
							current = current.Next
						}
					}
					l.Reverse()
					wantReverse := make([]int, len(tt.want))
					for i, value := range tt.want {
						wantReverse[len(tt.want)-1-i] = value
					}
					if got := l.Values(); !reflect.DeepEqual(got, wantReverse) {
						t.Fatalf("Reverse Values() = %v, want %v", got, wantReverse)
					}
					l.Reverse()
					l.InsertAtHead(9)
					if got := l.Values(); !reflect.DeepEqual(got, append([]int{9}, tt.want...)) {
						t.Fatalf("reused list Values() = %v", got)
					}
				})
			}
		})
	}
}

func TestDoublyLinkMiddleMutationBacklinks(t *testing.T) {
	l := NewDoublyLink[int]()
	l.InsertAtTail(1)
	l.InsertAtTail(3)
	l.InsertAt(1, 2)
	if l.Head.Next.Next.Pre != l.Head.Next {
		t.Fatal("middle insertion did not update the next node's Pre link")
	}
	l.Reverse()
	if got := l.Values(); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Fatalf("Reverse() after insertion = %v", got)
	}
	l.DeleteAt(1)
	if l.Head.Next.Pre != l.Head {
		t.Fatal("middle deletion did not update the next node's Pre link")
	}
	l.Reverse()
	if got := l.Values(); !reflect.DeepEqual(got, []int{1, 3}) {
		t.Fatalf("Reverse() after deletion = %v", got)
	}
}

package iterator

import "testing"

func TestTakeHasNextExhaustedSource(t *testing.T) {
	it := Take[int](FromSlice([]int{1}), 3)
	if value, ok := it.Next(); !ok || value != 1 {
		t.Fatal("missing first value")
	}
	if it.HasNext() {
		t.Error("exhausted source still reports an element")
	}
}

func TestFilterHasNextMatches(t *testing.T) {
	it := Filter[int](FromSlice([]int{1, 2, 3}), func(value int) bool { return value%2 == 0 })
	if !it.HasNext() || !it.HasNext() {
		t.Fatal("missing matching value")
	}
	if value, ok := it.Next(); !ok || value != 2 {
		t.Fatal("HasNext consumed the matching value")
	}
	if it.HasNext() {
		t.Error("unmatched tail reported as another element")
	}
}

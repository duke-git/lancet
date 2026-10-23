package algorithm

import "testing"

func TestBinaryIterativeSearchEmpty(t *testing.T) {
	for _, source := range [][]int{nil, {}} {
		if got := BinaryIterativeSearch(source, 1, 0, 0, &intComparator{}); got != -1 {
			t.Errorf("got %d, want -1", got)
		}
	}
}

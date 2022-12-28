package algorithm

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestLinearSearch(t *testing.T) {
	asssert := internal.NewAssert(t, "TestLinearSearch")

	numbers := []int{3, 4, 5, 3, 2, 1}
	equalFunc := func(a, b int) bool {
		return a == b
	}

	asssert.Equal(0, LinearSearch(numbers, 3, equalFunc))
	asssert.Equal(-1, LinearSearch(numbers, 6, equalFunc))
}

func TestBinarySearch(t *testing.T) {
	asssert := internal.NewAssert(t, "TestBinarySearch")

	sortedNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	comparator := &intComparator{}

	asssert.Equal(4, BinarySearch(sortedNumbers, 5, 0, len(sortedNumbers)-1, comparator))
	asssert.Equal(-1, BinarySearch(sortedNumbers, 9, 0, len(sortedNumbers)-1, comparator))
}

func TestBinaryIterativeSearch(t *testing.T) {
	asssert := internal.NewAssert(t, "TestBinaryIterativeSearch")

	sortedNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	comparator := &intComparator{}

	asssert.Equal(4, BinaryIterativeSearch(sortedNumbers, 5, 0, len(sortedNumbers)-1, comparator))
	asssert.Equal(-1, BinaryIterativeSearch(sortedNumbers, 9, 0, len(sortedNumbers)-1, comparator))
}

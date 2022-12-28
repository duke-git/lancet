package algorithm

import "fmt"

func ExampleLinearSearch() {
	numbers := []int{3, 4, 5, 3, 2, 1}

	equalFunc := func(a, b int) bool {
		return a == b
	}

	result1 := LinearSearch(numbers, 3, equalFunc)
	result2 := LinearSearch(numbers, 6, equalFunc)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 0
	// -1
}

func ExampleBinarySearch() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	comparator := &intComparator{}

	result1 := BinarySearch(numbers, 5, 0, len(numbers)-1, comparator)
	result2 := BinarySearch(numbers, 9, 0, len(numbers)-1, comparator)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 4
	// -1
}

func ExampleBinaryIterativeSearch() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
	comparator := &intComparator{}

	result1 := BinaryIterativeSearch(numbers, 5, 0, len(numbers)-1, comparator)
	result2 := BinaryIterativeSearch(numbers, 9, 0, len(numbers)-1, comparator)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 4
	// -1
}

package algorithm

import "fmt"

func ExampleBubbleSort() {
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}

	BubbleSort(numbers, comparator)

	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleCountSort() {
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}

	sortedNumber := CountSort(numbers, comparator)

	fmt.Println(numbers)
	fmt.Println(sortedNumber)
	// Output:
	// [2 1 5 3 6 4]
	// [1 2 3 4 5 6]
}

func ExampleHeapSort() {
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}

	HeapSort(numbers, comparator)

	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleMergeSort() {
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}

	MergeSort(numbers, comparator)

	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleInsertionSort() {
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}

	InsertionSort(numbers, comparator)

	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleSelectionSort() {
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}

	SelectionSort(numbers, comparator)

	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleShellSort() {
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}

	ShellSort(numbers, comparator)

	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleQuickSort() {
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}

	QuickSort(numbers, comparator)

	fmt.Println(numbers)
	// Output:
	// [1 2 3 4 5 6]
}

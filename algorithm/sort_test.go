package algorithm

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

// People test mock data
type people struct {
	Name string
	Age  int
}

// PeopleAageComparator sort people slice by age field
type peopleAgeComparator struct{}

// Compare implements github.com/duke-git/lancet/v2/lancetconstraints/constraints.go/Comparator
func (pc *peopleAgeComparator) Compare(v1 any, v2 any) int {
	p1, _ := v1.(people)
	p2, _ := v2.(people)

	//ascending order
	if p1.Age < p2.Age {
		return -1
	} else if p1.Age > p2.Age {
		return 1
	}
	return 0
}

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	//ascending order
	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func TestBubbleSortForStructSlice(t *testing.T) {
	asssert := internal.NewAssert(t, "TestBubbleSortForStructSlice")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &peopleAgeComparator{}
	BubbleSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	asssert.Equal(expected, actual)
}

func TestBubbleSortForIntSlice(t *testing.T) {
	asssert := internal.NewAssert(t, "TestBubbleSortForIntSlice")

	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}
	BubbleSort(numbers, comparator)

	asssert.Equal([]int{1, 2, 3, 4, 5, 6}, numbers)
}

func TestInsertionSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestInsertionSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &peopleAgeComparator{}
	InsertionSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	asssert.Equal(expected, actual)
}

func TestSelectionSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestSelectionSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &peopleAgeComparator{}
	SelectionSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	asssert.Equal(expected, actual)
}

func TestShellSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestShellSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &peopleAgeComparator{}
	ShellSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	asssert.Equal(expected, actual)
}

func TestQuickSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestQuickSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &peopleAgeComparator{}
	QuickSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	asssert.Equal(expected, actual)
}

func TestHeapSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestHeapSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &peopleAgeComparator{}
	HeapSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	asssert.Equal(expected, actual)
}

func TestMergeSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestMergeSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &peopleAgeComparator{}
	MergeSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	asssert.Equal(expected, actual)
}

func TestCountSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestCountSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &peopleAgeComparator{}
	sortedPeopleByAge := CountSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

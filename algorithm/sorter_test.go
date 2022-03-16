package algorithm

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/internal"
)

// People test mock data
type people struct {
	Name string
	Age  int
}

// PeopleAageComparator sort people slice by age field
type peopleAgeComparator struct{}

// Compare implements github.com/duke-git/lancet/lancetconstraints/constraints.go/Comparator
func (pc *peopleAgeComparator) Compare(v1 interface{}, v2 interface{}) int {
	p1, _ := v1.(people)
	p2, _ := v2.(people)

	//ascending order
	if p1.Age < p2.Age {
		return -1
	} else if p1.Age > p2.Age {
		return 1
	}
	return 0

	//decending order
	// if p1.Age > p2.Age {
	// 	return -1
	// } else if p1.Age < p2.Age {
	// 	return 1
	// }
}

var peoples = []people{
	{Name: "a", Age: 20},
	{Name: "b", Age: 10},
	{Name: "c", Age: 17},
	{Name: "d", Age: 8},
	{Name: "e", Age: 28},
}

type intComparator struct{}

func (c *intComparator) Compare(v1 interface{}, v2 interface{}) int {
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

var intSlice = []int{2, 1, 5, 3, 6, 4}

func TestBubbleSortForStructSlice(t *testing.T) {
	asssert := internal.NewAssert(t, "TestBubbleSortForStructSlice")

	comparator := &peopleAgeComparator{}
	sortedPeopleByAge := BubbleSort(peoples, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

func TestBubbleSortForIntSlice(t *testing.T) {
	asssert := internal.NewAssert(t, "TestBubbleSortForIntSlice")

	comparator := &intComparator{}
	sortedInt := BubbleSort(intSlice, comparator)
	expected := []int{1, 2, 3, 4, 5, 6}

	asssert.Equal(expected, sortedInt)
}

func TestInsertionSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestInsertionSort")

	comparator := &peopleAgeComparator{}
	sortedPeopleByAge := SelectionSort(peoples, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

func TestSelectionSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestSelectionSort")

	comparator := &peopleAgeComparator{}
	sortedPeopleByAge := SelectionSort(peoples, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

func TestShellSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestShellSort")

	comparator := &peopleAgeComparator{}
	sortedPeopleByAge := ShellSort(peoples, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

func TestQuickSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestQuickSort")

	comparator := &peopleAgeComparator{}
	sortedPeopleByAge := QuickSort(peoples, 0, len(peoples)-1, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

func TestHeapSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestHeapSort")

	comparator := &peopleAgeComparator{}
	sortedPeopleByAge := HeapSort(peoples, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

func TestMergeSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestMergeSort")

	comparator := &peopleAgeComparator{}
	sortedPeopleByAge := MergeSort(peoples, 0, len(peoples)-1, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

func TestCountSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestCountSort")

	comparator := &peopleAgeComparator{}
	sortedPeopleByAge := CountSort(peoples, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

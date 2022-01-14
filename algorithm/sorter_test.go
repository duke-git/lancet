package algorithm

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/internal"
)

// People test mock data
type People struct {
	Name string
	Age  int
}

// PeopleAageComparator sort people slice by age field
type PeopleAageComparator struct{}

// Compare implements github.com/duke-git/lancet/lancetconstraints/constraints.go/Comparator
func (pc *PeopleAageComparator) Compare(v1 interface{}, v2 interface{}) int {
	p1, _ := v1.(People)
	p2, _ := v2.(People)

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

var peoples = []People{
	{Name: "a", Age: 20},
	{Name: "b", Age: 10},
	{Name: "c", Age: 17},
	{Name: "d", Age: 8},
	{Name: "e", Age: 28},
}

func TestSelectionSort(t *testing.T) {
	asssert := internal.NewAssert(t, "TestSelectionSort")
	comparator := &PeopleAageComparator{}
	sortedPeopleByAge := SelectionSort(peoples, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	asssert.Equal(expected, actual)
}

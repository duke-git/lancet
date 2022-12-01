package slice

import (
	"fmt"
	"reflect"

	"github.com/duke-git/lancet/v2/lancetconstraints"
)

// sliceValue return the reflect value of a slice
func sliceValue(slice any) reflect.Value {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid slice type, value of type %T", slice))
	}
	return v
}

// sliceElemType get slice element type
func sliceElemType(reflectType reflect.Type) reflect.Type {
	for {
		if reflectType.Kind() != reflect.Slice {
			return reflectType
		}

		reflectType = reflectType.Elem()
	}
}

func quickSort[T lancetconstraints.Ordered](slice []T, lowIndex, highIndex int, order string) {
	if lowIndex < highIndex {
		p := partition(slice, lowIndex, highIndex, order)
		quickSort(slice, lowIndex, p-1, order)
		quickSort(slice, p+1, highIndex, order)
	}
}

// partition split slice into two parts for quick sort
func partition[T lancetconstraints.Ordered](slice []T, lowIndex, highIndex int, order string) int {
	p := slice[highIndex]
	i := lowIndex

	for j := lowIndex; j < highIndex; j++ {
		if order == "desc" {
			if slice[j] > p {
				swap(slice, i, j)
				i++
			}
		} else {
			if slice[j] < p {
				swap(slice, i, j)
				i++
			}
		}
	}

	swap(slice, i, highIndex)

	return i
}

// swap two slice value at index i and j
func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

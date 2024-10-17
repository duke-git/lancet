package slice

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

// resultChunk is used to store the intermediate results of UniqueByConcurrent.
// It is defined separately to be compatible with versions of go up to 1.20.
type resultChunk[T comparable] struct {
	index int
	data  []T
}

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

func quickSort[T constraints.Ordered](slice []T, lowIndex, highIndex int, order string) {
	if lowIndex < highIndex {
		p := partitionOrderedSlice(slice, lowIndex, highIndex, order)
		quickSort(slice, lowIndex, p-1, order)
		quickSort(slice, p+1, highIndex, order)
	}
}

// partitionOrderedSlice split ordered slice into two parts for quick sort
func partitionOrderedSlice[T constraints.Ordered](slice []T, lowIndex, highIndex int, order string) int {
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

func quickSortBy[T any](slice []T, lowIndex, highIndex int, less func(a, b T) bool) {
	if lowIndex < highIndex {
		p := partitionAnySlice(slice, lowIndex, highIndex, less)
		quickSortBy(slice, lowIndex, p-1, less)
		quickSortBy(slice, p+1, highIndex, less)
	}
}

// partitionAnySlice split any slice into two parts for quick sort
func partitionAnySlice[T any](slice []T, lowIndex, highIndex int, less func(a, b T) bool) int {
	p := slice[highIndex]
	i := lowIndex

	for j := lowIndex; j < highIndex; j++ {

		if less(slice[j], p) {
			swap(slice, i, j)
			i++
		}
	}

	swap(slice, i, highIndex)

	return i
}

// swap two slice value at index i and j
func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

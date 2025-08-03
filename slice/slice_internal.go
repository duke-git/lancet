package slice

import (
	"fmt"
	"math/bits"
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

// `repeat` returns a new slice that repeats the provided slice the given number of
// times. The result has length and capacity (len(x) * count). The result is never nil.
// Repeat panics if count is negative or if the result of (len(x) * count)  overflows.
//
// repeat has been provided in the standard lib within the package `slices` under the
// name Repeat since GO version 1.21 onwards. As lancet commits to compatibility with GO
// 1.18 onwards, we implement the functionality as an internal function.
func repeat[S ~[]E, E any](x S, count int) S {
	if count < 0 {
		panic("count cannot be negative")
	}

	const maxInt = ^uint(0) >> 1
	hi, lo := bits.Mul(uint(len(x)), uint(count))
	if hi > 0 || lo > maxInt {
		panic("the result of (len(x) * count) overflows")
	}

	newslice := make(S, int(lo)) // lo = len(x) * count
	n := copy(newslice, x)
	for n < len(newslice) {
		n += copy(newslice[n:], newslice[:n])
	}
	return newslice
}

// concat returns a new slice concatenating the passed in slices.
//
// concat has been provided in the standard lib within the package `slices` under the
// name Concat since GO version 1.21 onwards. As lancet commits to compatibility with GO
// 1.18 onwards, we implement the functionality as an internal function.
func concat[S ~[]E, E any](slices ...S) S {
	size := 0
	for _, s := range slices {
		size += len(s)
		if size < 0 {
			panic("len out of range")
		}
	}
	// Use Grow, not make, to round up to the size class:
	// the extra space is otherwise unused and helps
	// callers that append a few elements to the result.
	newslice := grow[S](nil, size)
	for _, s := range slices {
		newslice = append(newslice, s...)
	}
	return newslice
}

// grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, grow panics.
//
// grow has been provided in the standard lib within the package `slices` under the
// name Grow since GO version 1.21 onwards. As lancet commits to compatibility with GO
// 1.18 onwards, we implement the functionality as an internal function.
func grow[S ~[]E, E any](s S, n int) S {
	if n < 0 {
		panic("cannot be negative")
	}
	if n -= cap(s) - len(s); n > 0 {
		// This expression allocates only once.
		s = append(s[:cap(s)], make([]E, n)...)[:len(s)]
	}
	return s
}

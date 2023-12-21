// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

package algorithm

import "github.com/duke-git/lancet/v2/lancetconstraints"

// BubbleSort applys the bubble sort algorithm to sort the collection, will change the original collection data.
// Play: https://go.dev/play/p/GNdv7Jg2Taj
func BubbleSort[T any](slice []T, comparator lancetconstraints.Comparator) {
	for i := 0; i < len(slice); i++ {
		breakTag := false
		for j := 0; j < len(slice)-1-i; j++ {
			isCurrGreatThanNext := comparator.Compare(slice[j], slice[j+1]) == 1
			if isCurrGreatThanNext {
				swap(slice, j, j+1)
				breakTag = true
			}
		}
		if !breakTag {
			break
		}
	}
}

// InsertionSort applys the insertion sort algorithm to sort the collection, will change the original collection data.
// Play: https://go.dev/play/p/G5LJiWgJJW6
func InsertionSort[T any](slice []T, comparator lancetconstraints.Comparator) {
	for i := 0; i < len(slice); i++ {
		for j := i; j > 0; j-- {
			isPreLessThanCurrent := comparator.Compare(slice[j], slice[j-1]) == -1
			if isPreLessThanCurrent {
				swap(slice, j, j-1)
			} else {
				break
			}
		}
	}
}

// SelectionSort applys the selection sort algorithm to sort the collection, will change the original collection data.
// Play: https://go.dev/play/p/oXovbkekayS
func SelectionSort[T any](slice []T, comparator lancetconstraints.Comparator) {
	for i := 0; i < len(slice); i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if comparator.Compare(slice[j], slice[min]) == -1 {
				min = j
			}
		}
		swap(slice, i, min)
	}
}

// ShellSort applys the shell sort algorithm to sort the collection, will change the original collection data.
// Play: https://go.dev/play/p/3ibkszpJEu3
func ShellSort[T any](slice []T, comparator lancetconstraints.Comparator) {
	size := len(slice)

	gap := 1
	for gap < size/3 {
		gap = 3*gap + 1
	}

	for gap >= 1 {
		for i := gap; i < size; i++ {
			for j := i; j >= gap && comparator.Compare(slice[j], slice[j-gap]) == -1; j -= gap {
				swap(slice, j, j-gap)
			}
		}
		gap = gap / 3
	}
}

// QuickSort quick sorting for slice, lowIndex is 0 and highIndex is len(slice)-1.
// Play: https://go.dev/play/p/7Y7c1Elk3ax
func QuickSort[T any](slice []T, comparator lancetconstraints.Comparator) {
	quickSort(slice, 0, len(slice)-1, comparator)
}

func quickSort[T any](slice []T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) {
	if lowIndex < highIndex {
		p := partition(slice, lowIndex, highIndex, comparator)
		quickSort(slice, lowIndex, p-1, comparator)
		quickSort(slice, p+1, highIndex, comparator)
	}
}

// partition split slice into two parts
func partition[T any](slice []T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) int {
	p := slice[highIndex]
	i := lowIndex
	for j := lowIndex; j < highIndex; j++ {
		if comparator.Compare(slice[j], p) == -1 { //slice[j] < p
			swap(slice, i, j)
			i++
		}
	}

	swap(slice, i, highIndex)

	return i
}

// HeapSort applys the heap sort algorithm to sort the collection, will change the original collection data.
// Play: https://go.dev/play/p/u6Iwa1VZS_f
func HeapSort[T any](slice []T, comparator lancetconstraints.Comparator) {
	size := len(slice)

	for i := size/2 - 1; i >= 0; i-- {
		sift(slice, i, size-1, comparator)
	}
	for j := size - 1; j > 0; j-- {
		swap(slice, 0, j)
		sift(slice, 0, j-1, comparator)
	}
}

func sift[T any](slice []T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) {
	i := lowIndex
	j := 2*i + 1

	temp := slice[i]
	for j <= highIndex {
		if j < highIndex && comparator.Compare(slice[j], slice[j+1]) == -1 { //slice[j] < slice[j+1]
			j++
		}
		if comparator.Compare(temp, slice[j]) == -1 { //tmp < slice[j]
			slice[i] = slice[j]
			i = j
			j = 2*i + 1
		} else {
			break
		}
	}
	slice[i] = temp
}

// MergeSort applys the merge sort algorithm to sort the collection, will change the original collection data.
// Play: https://go.dev/play/p/ydinn9YzUJn
func MergeSort[T any](slice []T, comparator lancetconstraints.Comparator) {
	mergeSort(slice, 0, len(slice)-1, comparator)
}

func mergeSort[T any](slice []T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) {
	if lowIndex < highIndex {
		mid := (lowIndex + highIndex) / 2
		mergeSort(slice, lowIndex, mid, comparator)
		mergeSort(slice, mid+1, highIndex, comparator)
		merge(slice, lowIndex, mid, highIndex, comparator)
	}
}

func merge[T any](slice []T, lowIndex, midIndex, highIndex int, comparator lancetconstraints.Comparator) {
	i := lowIndex
	j := midIndex + 1
	temp := []T{}

	for i <= midIndex && j <= highIndex {
		//slice[i] < slice[j]
		if comparator.Compare(slice[i], slice[j]) == -1 {
			temp = append(temp, slice[i])
			i++
		} else {
			temp = append(temp, slice[j])
			j++
		}
	}

	if i <= midIndex {
		temp = append(temp, slice[i:midIndex+1]...)
	} else {
		temp = append(temp, slice[j:highIndex+1]...)
	}

	for k := 0; k < len(temp); k++ {
		slice[lowIndex+k] = temp[k]
	}
}

// CountSort applys the count sort algorithm to sort the collection, don't change the original collection data.
// Play: https://go.dev/play/p/tB-Umgm0DrP
func CountSort[T any](slice []T, comparator lancetconstraints.Comparator) []T {
	size := len(slice)
	out := make([]T, size)

	for i := 0; i < size; i++ {
		count := 0
		for j := 0; j < size; j++ {
			//slice[i] > slice[j]
			if comparator.Compare(slice[i], slice[j]) == 1 {
				count++
			}
		}
		out[count] = slice[i]
	}

	return out
}

// swap two slice value at index i and j
func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

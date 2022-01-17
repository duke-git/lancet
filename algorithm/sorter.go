// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package algorithm contain some basic algorithm functions. eg. sort, search, list, linklist, stack, queue, tree, graph. TODO
package algorithm

import "github.com/duke-git/lancet/lancetconstraints"

// BubbleSort use bubble to sort slice.
func BubbleSort[T any](slice []T, comparator lancetconstraints.Comparator) []T {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			isCurrGreatThanNext := comparator.Compare(slice[j], slice[j+1]) == 1
			if isCurrGreatThanNext {
				swap(slice, j, j+1)
			}
		}
	}
	return slice
}

// InsertionSort use insertion to sort slice.
func InsertionSort[T any](slice []T, comparator lancetconstraints.Comparator) []T {
	size := len(slice)
	if size <= 1 {
		return slice
	}

	for i := 1; i < size; i++ {
		currentItem := slice[i]
		preIndex := i - 1
		preItem := slice[preIndex]

		isPreLessThanCurrent := comparator.Compare(preItem, currentItem) == -1
		for preIndex >= 0 && isPreLessThanCurrent {
			slice[preIndex+1] = slice[preIndex]
			preIndex--
		}

		slice[preIndex+1] = preItem
	}

	return slice
}

// SelectionSort use selection to sort slice.
func SelectionSort[T any](slice []T, comparator lancetconstraints.Comparator) []T {
	for i := 0; i < len(slice); i++ {
		min := i

		for j := i + 1; j < len(slice); j++ {
			if comparator.Compare(slice[j], slice[min]) == -1 {
				min = j
			}
		}
		swap(slice, i, min)
	}
	return slice
}

// ShellSort shell sort slice.
func ShellSort[T any](slice []T, comparator lancetconstraints.Comparator) []T {
	size := len(slice)
	if size <= 1 {
		return slice
	}

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
		gap /= 3
	}

	return slice
}

// QuickSort quick sorting for slice, lowIndex is 0 and highIndex is len(slice)-1
func QuickSort[T any](slice []T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) []T {
	if lowIndex < highIndex {
		p := partition(slice, lowIndex, highIndex, comparator)
		QuickSort(slice, lowIndex, p-1, comparator)
		QuickSort(slice, p+1, highIndex, comparator)
	}

	return slice
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

// HeapSort use heap to sort slice
func HeapSort[T any](slice []T, comparator lancetconstraints.Comparator) []T {
	size := len(slice)

	for i := size/2 - 1; i >= 0; i-- {
		sift(slice, i, size-1, comparator)
	}
	for j := size - 1; j > 0; j-- {
		swap(slice, 0, j)
		sift(slice, 0, j-1, comparator)
	}

	return slice
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

// MergeSort merge sorting for slice
func MergeSort[T any](slice []T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) []T {
	if lowIndex < highIndex {
		mid := (lowIndex + highIndex) / 2
		MergeSort(slice, lowIndex, mid, comparator)
		MergeSort(slice, mid+1, highIndex, comparator)
		merge(slice, lowIndex, mid, highIndex, comparator)
	}

	return slice
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

// CountSort use count sorting for slice
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

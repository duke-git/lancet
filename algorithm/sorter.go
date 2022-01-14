// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package algorithm contain some algorithm functions. eg. sort, find, list, linklist, stack, queue, tree, graph. TODO
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

// func QuickSort[T comparable](slice []T, low, high int) []T {
// 	if low < high {
// 		var p int
// 		slice, p = partitionForQuickSort(slice, low, high)
// 		slice = quickSort(slice, low, p-1)
// 		slice = quickSort(slice, p+1, high)
// 	}

// 	return slice
// }

// swap two slice value at index i and j
func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// func partition[T comparable](slice []T, low, high int) ([]T, int) {
// 	p := slice[high]
// 	i := low
// 	for j := low; j < high; j++ {
// 		//???, error: comparable don't support operator <
// 		if slice[j] < p {
// 			slice[i], slice[j] = slice[j], slice[i]
// 			i++
// 		}
// 	}
// 	slice[i], slice[high] = slice[high], slice[i]

// 	return slice, i
// }

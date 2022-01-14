// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package algorithm contain some algorithm functions. eg. sort, find, list, linklist, stack, queue, tree, graph. TODO
package algorithm

import "github.com/duke-git/lancet/lancetconstraints"

// SelectionSort use selection to sort slice,
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

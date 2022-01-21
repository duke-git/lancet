// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package algorithm contain some basic algorithm functions. eg. sort, search, list, linklist, stack, queue, tree, graph. TODO
package algorithm

import "github.com/duke-git/lancet/lancetconstraints"

// Search algorithms see https://github.com/TheAlgorithms/Go/tree/master/search

// LinearSearch Simple linear search algorithm that iterates over all elements of an slice
// If a target is found, the index of the target is returned. Else the function return -1
func LinearSearch[T any](slice []T, target T, comparator lancetconstraints.Comparator) int {
	for i, v := range slice {
		if comparator.Compare(v, target) == 0 {
			return i
		}
	}
	return -1
}

// BinarySearch search for target within a sorted slice, recursive call itself.
// If a target is found, the index of the target is returned. Else the function return -1
func BinarySearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) int {
	if highIndex < lowIndex || len(sortedSlice) == 0 {
		return -1
	}

	midIndex := int(lowIndex + (highIndex-lowIndex)/2)
	isMidValGreatTarget := comparator.Compare(sortedSlice[midIndex], target) == 1
	isMidValLessTarget := comparator.Compare(sortedSlice[midIndex], target) == -1

	if isMidValGreatTarget {
		return BinarySearch(sortedSlice, target, lowIndex, midIndex-1, comparator)
	} else if isMidValLessTarget {
		return BinarySearch(sortedSlice, target, midIndex+1, highIndex, comparator)
	}

	return midIndex
}

// BinaryIterativeSearch search for target within a sorted slice.
// If a target is found, the index of the target is returned. Else the function return -1
func BinaryIterativeSearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) int {
	startIndex := lowIndex
	endIndex := highIndex

	var midIndex int
	for startIndex <= endIndex {
		midIndex = int(startIndex + (endIndex-startIndex)/2)
		isMidValGreatTarget := comparator.Compare(sortedSlice[midIndex], target) == 1
		isMidValLessTarget := comparator.Compare(sortedSlice[midIndex], target) == -1

		if isMidValGreatTarget {
			endIndex = midIndex - 1
		} else if isMidValLessTarget {
			startIndex = midIndex + 1
		} else {
			return midIndex
		}
	}
	return -1
}

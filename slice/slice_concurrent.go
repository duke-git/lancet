// Copyright 2024 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

package slice

import (
	"runtime"
	"sync"
)

// MapConcurrent applies the iteratee function to each item in the slice concurrently.
// Play: todo
func MapConcurrent[T any, U any](slice []T, iteratee func(index int, item T) U, numThreads int) []U {
	result := make([]U, len(slice))
	var wg sync.WaitGroup

	workerChan := make(chan struct{}, numThreads)

	for index, item := range slice {
		wg.Add(1)

		workerChan <- struct{}{}

		go func(i int, v T) {
			defer wg.Done()

			result[i] = iteratee(i, v)

			<-workerChan
		}(index, item)
	}

	wg.Wait()

	return result
}

// FilterConcurrent applies the provided filter function `predicate` to each element of the input slice concurrently.
// Play: todo
func FilterConcurrent[T any](slice []T, predicate func(index int, item T) bool, numThreads int) []T {
	result := make([]T, 0)
	var wg sync.WaitGroup

	workerChan := make(chan struct{}, numThreads)

	for index, item := range slice {
		wg.Add(1)

		workerChan <- struct{}{}

		go func(i int, v T) {
			defer wg.Done()

			if predicate(i, v) {
				result = append(result, v)
			}

			<-workerChan
		}(index, item)
	}

	wg.Wait()

	return result
}

// UniqueByParallel removes duplicate elements from the slice by parallel
// The comparator function is used to compare the elements
// The numThreads parameter specifies the number of threads to use
// If numThreads is less than or equal to 0, it will be set to 1
// The comparator function should return true if the two elements are equal
// Play: todo
func UniqueByConcurrent[T comparable](slice []T, comparator func(item T, other T) bool, numThreads int) []T {
	if numThreads <= 0 {
		numThreads = 1
	} else if numThreads > len(slice) {
		numThreads = len(slice)
	}

	maxThreads := runtime.NumCPU()
	if numThreads > maxThreads {
		numThreads = maxThreads
	}

	removeDuplicate := func(items []T, comparator func(item T, other T) bool) []T {
		var result []T
		for _, item := range items {
			seen := false
			for _, r := range result {
				if comparator(item, r) {
					seen = true
					break
				}
			}
			if !seen {
				result = append(result, item)
			}
		}
		return result
	}

	chunkSize := (len(slice) + numThreads - 1) / numThreads

	chunks := make([][]T, 0, numThreads)
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	type resultChunk struct {
		index int
		data  []T
	}
	resultCh := make(chan resultChunk, numThreads)
	var wg sync.WaitGroup

	for i, chunk := range chunks {
		wg.Add(1)
		go func(index int, chunk []T) {
			defer wg.Done()
			resultCh <- resultChunk{index, removeDuplicate(chunk, comparator)}
		}(i, chunk)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	results := make([][]T, len(chunks))
	for r := range resultCh {
		results[r.index] = r.data
	}

	result := []T{}
	seen := make(map[T]bool)

	for _, chunk := range results {
		for _, item := range chunk {
			if !seen[item] {
				seen[item] = true
				result = append(result, item)
			}
		}

	}

	return result
}

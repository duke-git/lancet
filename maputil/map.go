// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package maputil includes some functions to manipulate map.
package maputil

// Keys returns a slice of the map's keys
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// Values returns a slice of the map's values
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

// Merge maps, next key will overwrite previous key
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	res := make(map[K]V, 0)

	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}

// ForEach executes iteratee funcation for every key and value pair in map
func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V)) {
	for k, v := range m {
		iteratee(k, v)
	}
}

// Filter iterates over map, return a new map contains all key and value pairs pass the predicate function
func Filter[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	res := make(map[K]V)

	for k, v := range m {
		if predicate(k, v) {
			res[k] = v
		}
	}
	return res
}

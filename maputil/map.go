// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package maputil includes some functions to manipulate map.
package maputil

import "reflect"

// Keys returns a slice of the map's keys.
// Play: https://go.dev/play/p/xNB5bTb97Wd
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	var i int
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// Values returns a slice of the map's values.
// Play: https://go.dev/play/p/CBKdUc5FTW6
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	var i int
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}

// Merge maps, next key will overwrite previous key.
// Play: https://go.dev/play/p/H95LENF1uB-
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V, 0)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// ForEach executes iteratee funcation for every key and value pair in map.
// Play: https://go.dev/play/p/OaThj6iNVXK
func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V)) {
	for k, v := range m {
		iteratee(k, v)
	}
}

// Filter iterates over map, return a new map contains all key and value pairs pass the predicate function.
// Play: https://go.dev/play/p/fSvF3wxuNG7
func Filter[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// Intersect iterates over maps, return a new map of key and value pairs in all given maps.
// Play: https://go.dev/play/p/Zld0oj3sjcC
func Intersect[K comparable, V any](maps ...map[K]V) map[K]V {
	if len(maps) == 0 {
		return map[K]V{}
	}
	if len(maps) == 1 {
		return maps[0]
	}

	var result map[K]V

	reducer := func(m1, m2 map[K]V) map[K]V {
		m := make(map[K]V)
		for k, v1 := range m1 {
			if v2, ok := m2[k]; ok && reflect.DeepEqual(v1, v2) {
				m[k] = v1
			}
		}
		return m
	}

	reduceMaps := make([]map[K]V, 2)
	result = reducer(maps[0], maps[1])

	for i := 2; i < len(maps); i++ {
		reduceMaps[0] = result
		reduceMaps[1] = maps[i]
		result = reducer(reduceMaps[0], reduceMaps[1])
	}

	return result
}

// Minus creates an map of whose key in mapA but not in mapB.
// Play: https://go.dev/play/p/3u5U9K7YZ9m
func Minus[K comparable, V any](mapA, mapB map[K]V) map[K]V {
	result := make(map[K]V)

	for k, v := range mapA {
		if _, ok := mapB[k]; !ok {
			result[k] = v
		}
	}
	return result
}

// IsDisjoint two map are disjoint if they have no keys in common.
// Play: https://go.dev/play/p/N9qgYg_Ho6f
func IsDisjoint[K comparable, V any](mapA, mapB map[K]V) bool {
	for k := range mapA {
		if _, ok := mapB[k]; ok {
			return false
		}
	}
	return true
}

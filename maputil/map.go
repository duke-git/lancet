// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package maputil includes some functions to manipulate map.
package maputil

import (
	"reflect"

	"github.com/duke-git/lancet/v2/slice"
)

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

// KeysBy creates a slice whose element is the result of function mapper invoked by every map's key.
// Play: https://go.dev/play/p/hI371iB8Up8
func KeysBy[K comparable, V any, T any](m map[K]V, mapper func(item K) T) []T {
	keys := make([]T, 0, len(m))

	for k := range m {
		keys = append(keys, mapper(k))
	}

	return keys
}

// ValuesBy creates a slice whose element is the result of function mapper invoked by every map's value.
// Play: https://go.dev/play/p/sg9-oRidh8f
func ValuesBy[K comparable, V any, T any](m map[K]V, mapper func(item V) T) []T {
	keys := make([]T, 0, len(m))

	for _, v := range m {
		keys = append(keys, mapper(v))
	}

	return keys
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

// FilterByKeys iterates over map, return a new map whose keys are all given keys.
// Play: https://go.dev/play/p/7ov6BJHbVqh
func FilterByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if slice.Contain(keys, k) {
			result[k] = v
		}
	}
	return result
}

// FilterByValues iterates over map, return a new map whose values are all given values.
// Play: https://go.dev/play/p/P3-9MdcXegR
func FilterByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if slice.Contain(values, v) {
			result[k] = v
		}
	}
	return result
}

// OmitBy is the opposite of Filter, removes all the map elements for which the predicate function returns true.
// Play: https://go.dev/play/p/YJM4Hj5hNwm
func OmitBy[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if !predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// OmitByKeys the opposite of FilterByKeys, extracts all the map elements which keys are not omitted.
// Play: https://go.dev/play/p/jXGrWDBfSRp
func OmitByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if !slice.Contain(keys, k) {
			result[k] = v
		}
	}
	return result
}

// OmitByValues the opposite of FilterByValues. remov all elements whose value are in the give slice.
// Play: https://go.dev/play/p/XB7Y10uw20_U
func OmitByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if !slice.Contain(values, v) {
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

// Minus creates a map of whose key in mapA but not in mapB.
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

// Entry is a key/value pairs.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Entries transforms a map into array of key/value pairs.
// Play: https://go.dev/play/p/Ltb11LNcElY
func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))

	for k, v := range m {
		entries = append(entries, Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}

	return entries
}

// FromEntries creates a map based on a slice of key/value pairs
// Play: https://go.dev/play/p/fTdu4sCNjQO
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	result := make(map[K]V, len(entries))

	for _, v := range entries {
		result[v.Key] = v.Value
	}

	return result
}

// Transform a map to another type map.
// Play: https://go.dev/play/p/P6ovfToM3zj
func Transform[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, iteratee func(key K1, value V1) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(m))

	for k1, v1 := range m {
		k2, v2 := iteratee(k1, v1)
		result[k2] = v2
	}

	return result
}

// MapKeys transforms a map to other type map by manipulating it's keys.
// Play: https://go.dev/play/p/8scDxWeBDKd
func MapKeys[K comparable, V any, T comparable](m map[K]V, iteratee func(key K, value V) T) map[T]V {
	result := make(map[T]V, len(m))

	for k, v := range m {
		result[iteratee(k, v)] = v
	}

	return result
}

// MapValues transforms a map to other type map by manipulating it's values.
// Play: https://go.dev/play/p/g92aY3fc7Iw
func MapValues[K comparable, V any, T any](m map[K]V, iteratee func(key K, value V) T) map[K]T {
	result := make(map[K]T, len(m))

	for k, v := range m {
		result[k] = iteratee(k, v)
	}

	return result
}

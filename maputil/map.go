// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package maputil includes some functions to manipulate map.
package maputil

// Keys returns a slice of the map's keys
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// Values returns a slice of the map's values
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

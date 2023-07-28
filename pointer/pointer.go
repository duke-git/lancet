// Copyright 2023 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package pointer contains some util functions to operate go pointer.
package pointer

import "reflect"

// Of returns a pointer to the value `v`.
// Play: https://go.dev/play/p/HFd70x4DrMj
func Of[T any](v T) *T {
	return &v
}

// Unwrap returns the value from the pointer.
// Play: https://go.dev/play/p/cgeu3g7cjWb
func Unwrap[T any](p *T) T {
	return *p
}

// UnwarpOr returns the value from the pointer or fallback if the pointer is nil.
// Play: https://go.dev/play/p/mmNaLC38W8C
func UnwarpOr[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}
	return *p
}

// UnwarpOrDefault returns the value from the pointer or the default value if the pointer is nil.
// Play: https://go.dev/play/p/ZnGIHf8_o4E
func UnwarpOrDefault[T any](p *T) T {
	var v T

	if p == nil {
		return v
	}
	return *p
}

// ExtractPointer returns the underlying value by the given interface type
// Play: https://go.dev/play/p/D7HFjeWU2ZP
func ExtractPointer(value any) any {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	if t.Kind() != reflect.Pointer {
		return value
	}
	return ExtractPointer(v.Elem().Interface())
}

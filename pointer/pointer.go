// Copyright 2023 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package pointer contains some util functions to operate go pointer.
package pointer

import (
	"reflect"
)

// Of returns a pointer to the value `v`.
// Play: https://go.dev/play/p/HFd70x4DrMj
func Of[T any](v T) *T {
	if IsNil(v) {
		return nil
	}
	return &v
}

// Unwrap returns the value from the pointer.
//
//	Play: https://go.dev/play/p/cgeu3g7cjWb
//	Deprecated: Please use UnwrapOr
func Unwrap[T any](p *T) T {
	return *p
}

// UnwarpOr returns the value from the pointer or fallback if the pointer is nil.
//
//	Play: https://go.dev/play/p/mmNaLC38W8C
//	Deprecated: Please use UnwrapOr
func UnwarpOr[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}
	return *p
}

// UnwarpOrDefault returns the value from the pointer or the default value if the pointer is nil.
//
//	Play: https://go.dev/play/p/ZnGIHf8_o4E
//	Deprecated: Please use UnwrapOr
func UnwarpOrDefault[T any](p *T) T {
	var v T

	if p == nil {
		return v
	}
	return *p
}

// UnwrapOr returns the value from the pointer or fallback if the pointer is nil.
func UnwrapOr[T any](p *T, fallback ...T) T {
	if !IsNil(p) {
		return *p
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	var t T
	return t
}

// ExtractPointer returns the underlying value by the given interface type
// Play: https://go.dev/play/p/D7HFjeWU2ZP
func ExtractPointer(value any) any {
	if IsNil(value) {
		return value
	}
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	if t.Kind() != reflect.Pointer {
		return value
	}

	if v.Elem().IsValid() {
		return ExtractPointer(v.Elem().Interface())
	}

	return nil
}

// IsNil returns true if the given interface is nil or the underlying value is nil.
func IsNil(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}

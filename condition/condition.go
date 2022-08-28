// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package condition contains some functions for conditional judgment. eg. And, Or, TernaryOperator ...
// The implementation of this package refers to the implementation of carlmjohnson's truthy package, you may find more
// useful information in truthy(https://github.com/carlmjohnson/truthy), thanks carlmjohnson.
package condition

import "reflect"

// Bool returns the truthy value of anything.
// If the value's type has a Bool() bool method, the method is called and returned.
// If the type has an IsZero() bool method, the opposite value is returned.
// Slices and maps are truthy if they have a length greater than zero.
// All other types are truthy if they are not their zero value.
func Bool[T any](value T) bool {
	switch m := any(value).(type) {
	case interface{ Bool() bool }:
		return m.Bool()
	case interface{ IsZero() bool }:
		return !m.IsZero()
	}
	return reflectValue(&value)
}

func reflectValue(vp any) bool {
	switch rv := reflect.ValueOf(vp).Elem(); rv.Kind() {
	case reflect.Map, reflect.Slice:
		return rv.Len() != 0
	default:
		is := rv.IsZero()
		return !is
	}
}

// TernaryOperator if true return trueValue else return falseValue
func TernaryOperator[T any](isTrue bool, trueValue T, falseValue T) T {
	if isTrue {
		return trueValue
	} else {
		return falseValue
	}
}

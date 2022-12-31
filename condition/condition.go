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
// Play: https://go.dev/play/p/ETzeDJRSvhm
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

// And returns true if both a and b are truthy.
// Play: https://go.dev/play/p/W1SSUmt6pvr
func And[T, U any](a T, b U) bool {
	return Bool(a) && Bool(b)
}

// Or returns false iff neither a nor b is truthy.
// Play: https://go.dev/play/p/UlQTxHaeEkq
func Or[T, U any](a T, b U) bool {
	return Bool(a) || Bool(b)
}

// Xor returns true iff a or b but not both is truthy.
// Play: https://go.dev/play/p/gObZrW7ZbG8
func Xor[T, U any](a T, b U) bool {
	valA := Bool(a)
	valB := Bool(b)
	return (valA || valB) && valA != valB
}

// Nor returns true iff neither a nor b is truthy.
// Play: https://go.dev/play/p/g2j08F_zZky
func Nor[T, U any](a T, b U) bool {
	return !(Bool(a) || Bool(b))
}

// Xnor returns true iff both a and b or neither a nor b are truthy.
// Play: https://go.dev/play/p/OuDB9g51643
func Xnor[T, U any](a T, b U) bool {
	valA := Bool(a)
	valB := Bool(b)
	return (valA && valB) || (!valA && !valB)
}

// Nand returns false if both a and b are truthy.
// Play: https://go.dev/play/p/vSRMLxLIbq8
func Nand[T, U any](a T, b U) bool {
	return !Bool(a) || !Bool(b)
}

// TernaryOperator checks the value of param `isTrue`, if true return ifValue else return elseValue.
// Play: https://go.dev/play/p/ElllPZY0guT
func TernaryOperator[T, U any](isTrue T, ifValue U, elseValue U) U {
	if Bool(isTrue) {
		return ifValue
	} else {
		return elseValue
	}
}

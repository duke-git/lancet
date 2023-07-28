// Copyright 2023 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package compare provides a lightweight comparison function on any type.
// reference: https://github.com/stretchr/testify
package compare

import (
	"reflect"
	"time"

	"github.com/duke-git/lancet/v2/convertor"
)

// operator type
const (
	equal          = "eq"
	lessThan       = "lt"
	greaterThan    = "gt"
	lessOrEqual    = "le"
	greaterOrEqual = "ge"
)

var (
	timeType  = reflect.TypeOf(time.Time{})
	bytesType = reflect.TypeOf([]byte{})
)

// Equal checks if two values are equal or not. (check both type and value)
// Play: https://go.dev/play/p/wmVxR-to4lz
func Equal(left, right any) bool {
	return compareValue(equal, left, right)
}

// EqualValue checks if two values are equal or not. (check value only)
// Play: https://go.dev/play/p/fxnna_LLD9u
func EqualValue(left, right any) bool {
	ls, rs := convertor.ToString(left), convertor.ToString(right)
	return ls == rs
}

// LessThan checks if value `left` less than value `right`.
// Play: https://go.dev/play/p/cYh7FQQj0ne
func LessThan(left, right any) bool {
	return compareValue(lessThan, left, right)
}

// GreaterThan checks if value `left` greater than value `right`.
// Play: https://go.dev/play/p/9-NYDFZmIMp
func GreaterThan(left, right any) bool {
	return compareValue(greaterThan, left, right)
}

// LessOrEqual checks if value `left` less than or equal to value `right`.
// Play: https://go.dev/play/p/e4T_scwoQzp
func LessOrEqual(left, right any) bool {
	return compareValue(lessOrEqual, left, right)
}

// GreaterOrEqual checks if value `left` greater than or equal to value `right`.
// Play: https://go.dev/play/p/vx8mP0U8DFk
func GreaterOrEqual(left, right any) bool {
	return compareValue(greaterOrEqual, left, right)
}

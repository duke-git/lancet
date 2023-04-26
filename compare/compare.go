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
func Equal(left, right any) bool {
	return compareValue(equal, left, right)
}

// EqualValue checks if two values are equal or not. (check value only)
func EqualValue(left, right any) bool {
	ls, rs := convertor.ToString(left), convertor.ToString(right)
	return ls == rs
}

// LessThan checks if value `left` less than value `right`.
func LessThan(left, right any) bool {
	return compareValue(lessThan, left, right)
}

// GreaterThan checks if value `left` greater than value `right`.
func GreaterThan(left, right any) bool {
	return compareValue(greaterThan, left, right)
}

// LessOrEqual checks if value `left` less than or equal to value `right`.
func LessOrEqual(left, right any) bool {
	return compareValue(lessOrEqual, left, right)
}

// GreaterOrEqual checks if value `left` greater than or equal to value `right`.
func GreaterOrEqual(left, right any) bool {
	return compareValue(greaterOrEqual, left, right)
}

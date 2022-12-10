// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package formatter implements some functions to format string, struct.
package formatter

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// Comma add comma to a number value by every 3 numbers from right. ahead by symbol char.
// if value is invalid number string eg "aa", return empty string
// Comma("12345", "$") => "$12,345", Comma(12345, "$") => "$12,345"
func Comma[T constraints.Float | constraints.Integer | string](value T, symbol string) string {
	s, err := numberToString(value)
	if err != nil {
		return ""
	}

	dotIndex := strings.Index(s, ".")
	if dotIndex != -1 {
		return symbol + commaString(s[:dotIndex]) + s[dotIndex:]
	}

	return symbol + commaString(s)
}

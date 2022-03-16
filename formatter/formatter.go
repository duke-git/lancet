// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package formatter implements some functions to format string, struct.
package formatter

import "strings"

// Comma add comma to number by every 3 numbers from right. ahead by symbol char
func Comma(v any, symbol string) string {
	s := numString(v)
	dotIndex := strings.Index(s, ".")
	if dotIndex != -1 {
		return symbol + commaString(s[:dotIndex]) + s[dotIndex:]
	}
	return symbol + commaString(s)
}

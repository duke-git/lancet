// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package formatter implements some functions to format string, struct.
package formatter

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// Comma add comma to a number value by every 3 numbers from right. ahead by symbol char.
// if value is invalid number string eg "aa", return empty string
// Comma("12345", "$") => "$12,345", Comma(12345, "$") => "$12,345"
// Play: https://go.dev/play/p/eRD5k2vzUVX
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

func commaString(s string) string {
	if len(s) <= 3 {
		return s
	}
	return commaString(s[:len(s)-3]) + "," + commaString(s[len(s)-3:])
}

func numberToString(value any) (string, error) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fmt.Sprintf("%v", value), nil

		// todo: need to handle 12345678.9 => 1.23456789e+07
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%v", value), nil

	case reflect.String:
		{
			sv := fmt.Sprintf("%v", value)
			if strings.Contains(sv, ".") {
				_, err := strconv.ParseFloat(sv, 64)
				if err != nil {
					return "", err
				}
				return sv, nil
			} else {
				_, err := strconv.ParseInt(sv, 10, 64)
				if err != nil {
					return "", nil
				}
				return sv, nil
			}
		}
	default:
		return "", nil
	}
}

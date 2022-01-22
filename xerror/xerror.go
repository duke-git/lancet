// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package xerror implements helpers for errors
package xerror

// Unwrap if err is nil then it returns a valid value
// If err is not nil, Unwrap panics with err.
func Unwrap[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

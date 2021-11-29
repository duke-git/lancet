// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package utils implements is for internal use.
package utils

import (
	"fmt"
	"testing"
)

// LogFailedTestInfo log test failed info for internal use
func LogFailedTestInfo(t *testing.T, testCase, input, expected, result interface{}) {
	errInfo := fmt.Sprintf("Test case %v:  input is %+v, expected %v, but result is %v", testCase, input, expected, result)
	t.Error(errInfo)
}

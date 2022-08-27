// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package condition contains some functions for conditional judgment. eg. And, Or, TernaryOperator ...
package condition

// TernaryOperator if true return trueValue else return falseValue
func TernaryOperator[T any](isTrue bool, trueValue T, falseValue T) T {
	if isTrue {
		return trueValue
	} else {
		return falseValue
	}
}

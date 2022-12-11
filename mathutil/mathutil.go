// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package mathutil implements some functions for math calculation.
package mathutil

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

// Exponent calculate x^n
func Exponent(x, n int64) int64 {
	if n == 0 {
		return 1
	}

	t := Exponent(x, n/2)

	if n%2 == 1 {
		return t * t * x
	}

	return t * t
}

// Fibonacci calculate fibonacci number before n
func Fibonacci(first, second, n int) int {
	if n <= 0 {
		return 0
	}
	if n < 3 {
		return 1
	} else if n == 3 {
		return first + second
	} else {
		return Fibonacci(second, first+second, n-1)
	}
}

// Factorial calculate x!
func Factorial(x uint) uint {
	var f uint = 1
	for ; x > 1; x-- {
		f *= x
	}
	return f
}

// Percent calculate the percentage of val to total
func Percent(val, total float64, n int) float64 {
	if total == 0 {
		return float64(0)
	}
	tmp := val / total * 100
	result := RoundToFloat(tmp, n)

	return result
}

// RoundToString round up to n decimal places
func RoundToString(x float64, n int) string {
	tmp := math.Pow(10.0, float64(n))
	x *= tmp
	x = math.Round(x)
	result := strconv.FormatFloat(x/tmp, 'f', n, 64)
	return result
}

// RoundToFloat round up to n decimal places
func RoundToFloat(x float64, n int) float64 {
	tmp := math.Pow(10.0, float64(n))
	x *= tmp
	x = math.Round(x)
	return x / tmp
}

// TruncRound round off n decimal places
func TruncRound(x float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n+1)+"f", x)
	temp := strings.Split(floatStr, ".")
	var newFloat string
	if len(temp) < 2 || n >= len(temp[1]) {
		newFloat = floatStr
	} else {
		newFloat = temp[0] + "." + temp[1][:n]
	}
	result, _ := strconv.ParseFloat(newFloat, 64)
	return result
}

// Max return max value of params
func Max[T constraints.Integer | constraints.Float](numbers ...T) T {
	max := numbers[0]

	for _, v := range numbers {
		if max < v {
			max = v
		}
	}

	return max
}

// MaxBy search the maximum value of a slice using the given comparator function.
func MaxBy[T any](slice []T, comparator func(T, T) bool) T {
	var max T

	if len(slice) == 0 {
		return max
	}

	max = slice[0]

	for i := 1; i < len(slice); i++ {
		val := slice[i]

		if comparator(val, max) {
			max = val
		}
	}

	return max
}

// Min return min value of params
func Min[T constraints.Integer | constraints.Float](numbers ...T) T {
	min := numbers[0]

	for _, v := range numbers {
		if min > v {
			min = v
		}
	}

	return min
}

// MinBy search the minimum value of a slice using the given comparator function.
func MinBy[T any](slice []T, comparator func(T, T) bool) T {
	var min T

	if len(slice) == 0 {
		return min
	}

	min = slice[0]

	for i := 1; i < len(slice); i++ {
		val := slice[i]

		if comparator(val, min) {
			min = val
		}
	}

	return min
}

// Average return average value of numbers
func Average[T constraints.Integer | constraints.Float](numbers ...T) T {
	var sum T
	n := T(len(numbers))

	for _, v := range numbers {
		sum += v
	}
	return sum / n
}

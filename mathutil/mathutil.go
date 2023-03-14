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

// Exponent calculate x^n.
// Play: https://go.dev/play/p/uF3HGNPk8wr
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

// Fibonacci calculate fibonacci number before n.
// Play: https://go.dev/play/p/IscseUNMuUc
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

// Factorial calculate x!.
// Play: https://go.dev/play/p/tt6LdOK67Nx
func Factorial(x uint) uint {
	var f uint = 1
	for ; x > 1; x-- {
		f *= x
	}
	return f
}

// Percent calculate the percentage of value to total.
// Play: https://go.dev/play/p/QQM9B13coSP
func Percent(val, total float64, n int) float64 {
	if total == 0 {
		return float64(0)
	}
	tmp := val / total
	result := RoundToFloat(tmp, n)

	return result
}

// RoundToString round up to n decimal places.
// Play: https://go.dev/play/p/kZwpBRAcllO
func RoundToString(x float64, n int) string {
	tmp := math.Pow(10.0, float64(n))
	x *= tmp
	x = math.Round(x)
	result := strconv.FormatFloat(x/tmp, 'f', n, 64)
	return result
}

// RoundToFloat round up to n decimal places.
// Play: https://go.dev/play/p/ghyb528JRJL
func RoundToFloat(x float64, n int) float64 {
	tmp := math.Pow(10.0, float64(n))
	x *= tmp
	x = math.Round(x)
	return x / tmp
}

// TruncRound round off n decimal places.
// Play: https://go.dev/play/p/aumarSHIGzP
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

// Max return max value of numbers.
// Play: https://go.dev/play/p/cN8DHI0rTkH
func Max[T constraints.Integer | constraints.Float](numbers ...T) T {
	max := numbers[0]

	for _, v := range numbers {
		if max < v {
			max = v
		}
	}

	return max
}

// MaxBy return the maximum value of a slice using the given comparator function.
// Play: https://go.dev/play/p/pbe2MT-7DV2
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

// Min return min value of numbers.
// Play: https://go.dev/play/p/21BER_mlGUj
func Min[T constraints.Integer | constraints.Float](numbers ...T) T {
	min := numbers[0]

	for _, v := range numbers {
		if min > v {
			min = v
		}
	}

	return min
}

// MinBy return the minimum value of a slice using the given comparator function.
// Play: https://go.dev/play/p/XuJDKrDdglW
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

// Average return average value of numbers.
// Play: https://go.dev/play/p/Vv7LBwER-pz
func Average[T constraints.Integer | constraints.Float](numbers ...T) T {
	var sum T
	n := T(len(numbers))

	for _, v := range numbers {
		sum += v
	}
	return sum / n
}

// Range creates a slice of numbers from start with specified count, element step is 1.
// Play: todo
func Range[T constraints.Integer | constraints.Float](start T, count int) []T {
	size := count
	if count < 0 {
		size = -count
	}

	result := make([]T, size)

	for i, j := 0, start; i < size; i, j = i+1, j+1 {
		result[i] = j
	}

	return result
}

// RangeWithStep creates a slice of numbers from start to end with specified step.
// Play: todo
func RangeWithStep[T constraints.Integer | constraints.Float](start, end, step T) []T {
	result := []T{}

	if start >= end || step == 0 {
		return result
	}

	for i := start; i < end; i += step {
		result = append(result, i)
	}

	return result
}

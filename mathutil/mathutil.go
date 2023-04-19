// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package mathutil implements some functions for math calculation.
package mathutil

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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
	res := RoundToFloat(tmp, n)

	return res
}

// RoundToString round up to n decimal places
func RoundToString(x float64, n int) string {
	tmp := math.Pow(10.0, float64(n))
	x *= tmp
	x = math.Round(x)
	res := strconv.FormatFloat(x/tmp, 'f', n, 64)
	return res
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
	res, _ := strconv.ParseFloat(newFloat, 64)
	return res
}

// AngleToRadian converts angle value to radian value.
func AngleToRadian(angle float64) float64 {
	radian := angle * (math.Pi / 180)
	return radian
}

// RadianToAngle converts radian value to angle value.
func RadianToAngle(radian float64) float64 {
	angle := radian * (180 / math.Pi)
	return angle
}

// PointDistance get two points distance.
func PointDistance(x1, y1, x2, y2 float64) float64 {
	a := x1 - x2
	b := y1 - y2
	c := math.Pow(a, 2) + math.Pow(b, 2)

	return math.Sqrt(c)
}

// IsPrimes checks if number is prime number.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

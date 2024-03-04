package mathutil

import (
	"fmt"
	"math"
)

func ExampleExponent() {
	result1 := Exponent(10, 0)
	result2 := Exponent(10, 1)
	result3 := Exponent(10, 2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 1
	// 10
	// 100
}

func ExampleFibonacci() {
	result1 := Fibonacci(1, 1, 1)
	result2 := Fibonacci(1, 1, 2)
	result3 := Fibonacci(1, 1, 5)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 1
	// 1
	// 5
}

func ExampleFactorial() {
	result1 := Factorial(1)
	result2 := Factorial(2)
	result3 := Factorial(3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 1
	// 2
	// 6
}

func ExamplePercent() {
	result1 := Percent(1, 2, 2)
	result2 := Percent(0.1, 0.3, 2)
	result3 := Percent(-30305, 408420, 2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 50
	// 33.33
	// -7.42
}

func ExampleRoundToFloat() {
	result1 := RoundToFloat(0.124, 2)
	result2 := RoundToFloat(0.125, 2)
	result3 := RoundToFloat(0.125, 3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 0.12
	// 0.13
	// 0.125
}

func ExampleRoundToString() {
	result1 := RoundToString(0.124, 2)
	result2 := RoundToString(0.125, 2)
	result3 := RoundToString(0.125, 3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 0.12
	// 0.13
	// 0.125
}

func ExampleTruncRound() {
	result1 := TruncRound(0.124, 2)
	result2 := TruncRound(0.125, 2)
	result3 := TruncRound(0.125, 3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 0.12
	// 0.12
	// 0.125
}

func ExampleCeilToFloat() {
	result1 := CeilToFloat(3.14159, 1)
	result2 := CeilToFloat(3.14159, 2)
	result3 := CeilToFloat(5, 4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 3.2
	// 3.15
	// 5
}

func ExampleCeilToString() {
	result1 := CeilToString(3.14159, 1)
	result2 := CeilToString(3.14159, 2)
	result3 := CeilToString(5, 4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 3.2
	// 3.15
	// 5.0000
}

func ExampleFloorToFloat() {
	result1 := FloorToFloat(3.14159, 1)
	result2 := FloorToFloat(3.14159, 2)
	result3 := FloorToFloat(5, 4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 3.1
	// 3.14
	// 5
}

func ExampleFloorToString() {
	result1 := FloorToString(3.14159, 1)
	result2 := FloorToString(3.14159, 2)
	result3 := FloorToString(5, 4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 3.1
	// 3.14
	// 5.0000
}

func ExampleAverage() {
	result1 := Average(1, 2)
	result2 := RoundToFloat(Average(1.2, 1.4), 1)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 1
	// 1.3
}

func ExampleSum() {
	result1 := Sum(1, 2)
	result2 := Sum(0.1, float64(1))

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 3
	// 1.1
}

func ExampleMax() {
	result1 := Max(1, 2, 3)
	result2 := Max(1.2, 1.4, 1.1, 1.4)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 3
	// 1.4
}

func ExampleMaxBy() {
	result1 := MaxBy([]string{"a", "ab", "abc"}, func(v1, v2 string) bool {
		return len(v1) > len(v2)
	})

	result2 := MaxBy([]string{"abd", "abc", "ab"}, func(v1, v2 string) bool {
		return len(v1) > len(v2)
	})

	result3 := MaxBy([]string{}, func(v1, v2 string) bool {
		return len(v1) > len(v2)
	})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// abc
	// abd
	//
}

func ExampleMin() {
	result1 := Min(1, 2, 3)
	result2 := Min(1.2, 1.4, 1.1, 1.4)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 1
	// 1.1
}

func ExampleMinBy() {
	result1 := MinBy([]string{"a", "ab", "abc"}, func(v1, v2 string) bool {
		return len(v1) < len(v2)
	})

	result2 := MinBy([]string{"ab", "ac", "abc"}, func(v1, v2 string) bool {
		return len(v1) < len(v2)
	})

	result3 := MinBy([]string{}, func(v1, v2 string) bool {
		return len(v1) < len(v2)
	})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// a
	// ab
	//
}

func ExampleRange() {
	result1 := Range(1, 4)
	result2 := Range(1, -4)
	result3 := Range(-4, 4)
	result4 := Range(1.0, 4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// [1 2 3 4]
	// [1 2 3 4]
	// [-4 -3 -2 -1]
	// [1 2 3 4]
}

func ExampleRangeWithStep() {
	result1 := RangeWithStep(1, 4, 1)
	result2 := RangeWithStep(1, -1, 0)
	result3 := RangeWithStep(-4, 1, 2)
	result4 := RangeWithStep(1.0, 4.0, 1.1)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// [1 2 3]
	// []
	// [-4 -2 0]
	// [1 2.1 3.2]
}

func ExampleAngleToRadian() {
	result1 := AngleToRadian(45)
	result2 := AngleToRadian(90)
	result3 := AngleToRadian(180)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 0.7853981633974483
	// 1.5707963267948966
	// 3.141592653589793
}

func ExampleRadianToAngle() {
	result1 := RadianToAngle(math.Pi)
	result2 := RadianToAngle(math.Pi / 2)
	result3 := RadianToAngle(math.Pi / 4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 180
	// 90
	// 45
}

func ExamplePointDistance() {
	result1 := PointDistance(1, 1, 4, 5)

	fmt.Println(result1)

	// Output:
	// 5
}

func ExampleIsPrime() {
	result1 := IsPrime(-1)
	result2 := IsPrime(0)
	result3 := IsPrime(1)
	result4 := IsPrime(2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// false
	// false
	// false
	// true
}

func ExampleGCD() {
	result1 := GCD(1, 1)
	result2 := GCD(1, -1)
	result3 := GCD(-1, 1)
	result4 := GCD(-1, -1)
	result5 := GCD(3, 6, 9)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// 1
	// 1
	// -1
	// -1
	// 3
}

func ExampleLCM() {
	result1 := LCM(1, 1)
	result2 := LCM(1, 2)
	result3 := LCM(3, 6, 9)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 1
	// 2
	// 18
}

func ExampleCos() {
	result1 := Cos(0)
	result2 := Cos(90)
	result3 := Cos(180)
	result4 := Cos(math.Pi)
	result5 := Cos(math.Pi / 2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// 1
	// -0.447
	// -0.598
	// -1
	// 0
}

func ExampleSin() {
	result1 := Sin(0)
	result2 := Sin(90)
	result3 := Sin(180)
	result4 := Sin(math.Pi)
	result5 := Sin(math.Pi / 2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// 0
	// 0.894
	// -0.801
	// 0
	// 1
}

func ExampleLog() {
	result1 := Log(8, 2)
	result2 := TruncRound(Log(5, 2), 2)
	result3 := TruncRound(Log(27, 3), 0)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 3
	// 2.32
	// 3
}

func ExampleAbs() {
	result1 := Abs(-1)
	result2 := Abs(-0.1)
	result3 := Abs(float32(0.2))

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 1
	// 0.1
	// 0.2
}

func ExampleDiv() {
	result1 := Div(9, 4)
	result2 := Div(1, 2)
	result3 := Div(0, 666)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	// Output:
	// 2.25
	// 0.5
	// 0
}

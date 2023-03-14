package mathutil

import "fmt"

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

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 0.5
	// 0.33
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

func ExampleAverage() {
	result1 := Average(1, 2)
	result2 := RoundToFloat(Average(1.2, 1.4), 1)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 1
	// 1.3
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

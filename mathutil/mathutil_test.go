package mathutil

import (
	"math"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestExponent(t *testing.T) {
	assert := internal.NewAssert(t, "TestExponent")

	assert.Equal(int64(1), Exponent(10, 0))
	assert.Equal(int64(10), Exponent(10, 1))
	assert.Equal(int64(100), Exponent(10, 2))
}

func TestFibonacci(t *testing.T) {
	assert := internal.NewAssert(t, "TestFibonacci")

	assert.Equal(0, Fibonacci(1, 1, 0))
	assert.Equal(1, Fibonacci(1, 1, 1))
	assert.Equal(1, Fibonacci(1, 1, 2))
	assert.Equal(5, Fibonacci(1, 1, 5))
}

func TestFactorial(t *testing.T) {
	assert := internal.NewAssert(t, "TestFactorial")

	assert.Equal(uint(1), Factorial(0))
	assert.Equal(uint(1), Factorial(1))
	assert.Equal(uint(2), Factorial(2))
	assert.Equal(uint(6), Factorial(3))
}

func TestPercent(t *testing.T) {
	assert := internal.NewAssert(t, "TestPercent")

	assert.Equal(0.5, Percent(1, 2, 2))
	assert.Equal(0.33, Percent(0.1, 0.3, 2))
}

func TestRoundToFloat(t *testing.T) {
	assert := internal.NewAssert(t, "TestRoundToFloat")

	assert.Equal(RoundToFloat(0, 0), float64(0))
	assert.Equal(RoundToFloat(0, 1), float64(0))
	assert.Equal(RoundToFloat(0.124, 2), float64(0.12))
	assert.Equal(RoundToFloat(0.125, 2), float64(0.13))
	assert.Equal(RoundToFloat(0.125, 3), float64(0.125))
	assert.Equal(RoundToFloat(33.33333, 2), float64(33.33))
}

func TestRoundToString(t *testing.T) {
	assert := internal.NewAssert(t, "TestRoundToString")

	assert.Equal(RoundToString(0, 0), "0")
	assert.Equal(RoundToString(0, 1), "0.0")
	assert.Equal(RoundToString(0.124, 2), "0.12")
	assert.Equal(RoundToString(0.125, 2), "0.13")
	assert.Equal(RoundToString(0.125, 3), "0.125")
}

func TestTruncRound(t *testing.T) {
	assert := internal.NewAssert(t, "TestTruncRound")

	assert.Equal(TruncRound(0, 0), float64(0))
	assert.Equal(TruncRound(0, 1), float64(0))
	assert.Equal(TruncRound(0.124, 2), float64(0.12))
	assert.Equal(TruncRound(0.125, 2), float64(0.12))
	assert.Equal(TruncRound(0.125, 3), float64(0.125))
	assert.Equal(TruncRound(33.33333, 2), float64(33.33))
}

func TestAverage(t *testing.T) {
	assert := internal.NewAssert(t, "TestAverage")

	assert.Equal(Average(0, 0), 0)
	assert.Equal(Average(1, 1), 1)
	avg := Average(1.2, 1.4)
	assert.Equal(1.3, RoundToFloat(avg, 1))
}

func TestMax(t *testing.T) {
	assert := internal.NewAssert(t, "TestMax")

	assert.Equal(Max(0, 0), 0)
	assert.Equal(Max(1, 2, 3), 3)
	assert.Equal(Max(1.2, 1.4, 1.1, 1.4), 1.4)

	type Integer int
	assert.Equal(Max(Integer(1), Integer(0)), Integer(1))
}

func TestMaxBy(t *testing.T) {
	assert := internal.NewAssert(t, "MaxBy")

	res1 := MaxBy([]string{"a", "ab", "abc"}, func(v1, v2 string) bool {
		return len(v1) > len(v2)
	})
	assert.Equal("abc", res1)

	res2 := MaxBy([]string{"abd", "abc", "ab"}, func(v1, v2 string) bool {
		return len(v1) > len(v2)
	})
	assert.Equal("abd", res2)

	res3 := MaxBy([]string{}, func(v1, v2 string) bool {
		return len(v1) > len(v2)
	})
	assert.Equal("", res3)
}

func TestMin(t *testing.T) {
	assert := internal.NewAssert(t, "TestMin")

	assert.Equal(Min(0, 0), 0)
	assert.Equal(Min(1, 2, 3), 1)
	assert.Equal(Min(1.2, 1.4, 1.1, 1.4), 1.1)
}

func TestMinBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestMinBy")

	res1 := MinBy([]string{"a", "ab", "abc"}, func(v1, v2 string) bool {
		return len(v1) < len(v2)
	})
	assert.Equal("a", res1)

	res2 := MinBy([]string{"ab", "ac", "abc"}, func(v1, v2 string) bool {
		return len(v1) < len(v2)
	})
	assert.Equal("ab", res2)

	res3 := MinBy([]string{}, func(v1, v2 string) bool {
		return len(v1) < len(v2)
	})
	assert.Equal("", res3)
}

func TestRange(t *testing.T) {
	assert := internal.NewAssert(t, "TestRange")

	result1 := Range(1, 4)
	result2 := Range(1, -4)
	result3 := Range(-4, 4)
	result4 := Range(1.0, 4)
	result5 := Range(1, 0)

	assert.Equal([]int{1, 2, 3, 4}, result1)
	assert.Equal([]int{1, 2, 3, 4}, result2)
	assert.Equal([]int{-4, -3, -2, -1}, result3)
	assert.Equal([]float64{1.0, 2.0, 3.0, 4.0}, result4)
	assert.Equal([]int{}, result5)
}

func TestRangeWithStep(t *testing.T) {
	assert := internal.NewAssert(t, "TestRangeWithStep")

	result1 := RangeWithStep(1, 4, 1)
	result2 := RangeWithStep(1, -1, 0)
	result3 := RangeWithStep(-4, 1, 2)
	result4 := RangeWithStep(1.0, 4.0, 1.1)

	assert.Equal([]int{1, 2, 3}, result1)
	assert.Equal([]int{}, result2)
	assert.Equal([]int{-4, -2, 0}, result3)
	assert.Equal([]float64{1.0, 2.1, 3.2}, result4)
}

func TestAngleToRadian(t *testing.T) {
	assert := internal.NewAssert(t, "TestAngleToRadian")

	result1 := AngleToRadian(45)
	result2 := AngleToRadian(90)
	result3 := AngleToRadian(180)

	assert.Equal(0.7853981633974483, result1)
	assert.Equal(1.5707963267948966, result2)
	assert.Equal(3.141592653589793, result3)
}

func TestRadianToAngle(t *testing.T) {
	assert := internal.NewAssert(t, "TestAngleToRadian")

	result1 := RadianToAngle(math.Pi)
	result2 := RadianToAngle(math.Pi / 2)
	result3 := RadianToAngle(math.Pi / 4)

	assert.Equal(float64(180), result1)
	assert.Equal(float64(90), result2)
	assert.Equal(float64(45), result3)
}

func TestPointDistance(t *testing.T) {
	assert := internal.NewAssert(t, "TestPointDistance")

	result1 := PointDistance(1, 1, 4, 5)

	assert.Equal(float64(5), result1)
}

func TestIsPrime(t *testing.T) {
	assert := internal.NewAssert(t, "TestIsPrime")

	assert.Equal(false, IsPrime(-1))
	assert.Equal(false, IsPrime(0))
	assert.Equal(false, IsPrime(1))
	assert.Equal(true, IsPrime(2))
	assert.Equal(true, IsPrime(3))
	assert.Equal(false, IsPrime(4))
}

func TestGCD(t *testing.T) {
	assert := internal.NewAssert(t, "TestGCD")

	assert.Equal(1, GCD(1, 1))
	assert.Equal(1, GCD(1, -1))
	assert.Equal(-1, GCD(-1, 1))
	assert.Equal(-1, GCD(-1, -1))

	assert.Equal(1, GCD(1, 0))
	assert.Equal(1, GCD(0, 1))
	assert.Equal(-1, GCD(-1, 0))
	assert.Equal(-1, GCD(0, -1))

	assert.Equal(1, GCD(1, -2))
	assert.Equal(1, GCD(-2, 1))
	assert.Equal(-1, GCD(-1, 2))
	assert.Equal(-1, GCD(2, -1))

	assert.Equal(-1, GCD(-1, -2))
	assert.Equal(-1, GCD(-2, -1))

	assert.Equal(-9, GCD(-27, -36))
	assert.Equal(3, GCD(3, 6, 9))
}

func TestLCM(t *testing.T) {
	assert := internal.NewAssert(t, "TestLCM")

	assert.Equal(1, LCM(1))
	assert.Equal(-1, LCM(-1))
	assert.Equal(-1, LCM(1, -1))
	assert.Equal(1, LCM(-1, 1))
	assert.Equal(1, LCM(1, 1))
	assert.Equal(-1, LCM(-1, -1))
	assert.Equal(2, LCM(1, 2))
	assert.Equal(18, LCM(3, 6, 9))
}

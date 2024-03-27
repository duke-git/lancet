package mathutil

import (
	"math"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestExponent(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestExponent")

	assert.Equal(int64(1), Exponent(10, 0))
	assert.Equal(int64(10), Exponent(10, 1))
	assert.Equal(int64(100), Exponent(10, 2))
}

func TestFibonacci(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFibonacci")

	assert.Equal(0, Fibonacci(1, 1, 0))
	assert.Equal(1, Fibonacci(1, 1, 1))
	assert.Equal(1, Fibonacci(1, 1, 2))
	assert.Equal(5, Fibonacci(1, 1, 5))
}

func TestFactorial(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFactorial")

	assert.Equal(uint(1), Factorial(0))
	assert.Equal(uint(1), Factorial(1))
	assert.Equal(uint(2), Factorial(2))
	assert.Equal(uint(6), Factorial(3))
}

func TestPercent(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPercent")

	assert.EqualValues(50, Percent(1, 2, 2))
	assert.EqualValues(33.33, Percent(0.1, 0.3, 2))
	assert.EqualValues(-7.42, Percent(-30305, 408420, 2))
}

func TestRoundToFloat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRoundToFloat")

	assert.Equal(float64(0), RoundToFloat(0, 0))
	assert.Equal(float64(0), RoundToFloat(0, 1))
	assert.Equal(0.12, RoundToFloat(0.124, 2))
	assert.Equal(0.13, RoundToFloat(0.125, 2))
	assert.Equal(0.125, RoundToFloat(0.125, 3))
	assert.Equal(33.33, RoundToFloat(33.33333, 2))
	assert.Equal(0.59, RoundToFloat(0.585, 2))
	assert.Equal(63.59, RoundToFloat(63.585, 2))
	assert.Equal(64.59, RoundToFloat(64.585, 2))
}

func TestRoundToString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRoundToString")

	assert.Equal("0", RoundToString(0, 0))
	assert.Equal("0.0", RoundToString(0, 1))
	assert.Equal("0.12", RoundToString(0.124, 2))
	assert.Equal("0.13", RoundToString(0.125, 2))
	assert.Equal("0.125", RoundToString(0.125, 3))
	assert.Equal("54.321", RoundToString(54.321, 3))
	assert.Equal("17.000", RoundToString(17, 3))
	assert.Equal("0.59", RoundToString(0.585, 2))
	assert.Equal("63.59", RoundToString(63.585, 2))
	assert.Equal("64.59", RoundToString(64.585, 2))
}

func TestTruncRound(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTruncRound")

	assert.Equal(float64(0), TruncRound(float64(0), 0))
	assert.Equal(float64(0), TruncRound(float64(0), 1))
	assert.Equal(float32(0), TruncRound(float32(0), 0))
	assert.Equal(float32(0), TruncRound(float32(0), 1))
	assert.Equal(0, TruncRound(0, 0))
	assert.Equal(uint64(0), TruncRound(uint64(0), 1))
	assert.Equal(0.12, TruncRound(0.124, 2))
	assert.Equal(0.12, TruncRound(0.125, 2))
	assert.Equal(0.125, TruncRound(0.125, 3))
	assert.Equal(33.33, TruncRound(33.33333, 2))
}

func TestFloorToFloat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFloorToFloat")

	assert.Equal(3.14, FloorToFloat(3.14159, 2))
	assert.Equal(3.141, FloorToFloat(3.14159, 3))
	assert.Equal(5.0, FloorToFloat(5, 4))
	assert.Equal(2.0, FloorToFloat(9/4, 2))
}

func TestFloorToString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFloorToString")

	assert.Equal("3.14", FloorToString(3.14159, 2))
	assert.Equal("3.141", FloorToString(3.14159, 3))
	assert.Equal("5.0000", FloorToString(5, 4))
}

func TestCeilToFloat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCeilToFloat")

	assert.Equal(3.15, CeilToFloat(3.14159, 2))
	assert.Equal(3.142, CeilToFloat(3.14159, 3))
	assert.Equal(5.0, CeilToFloat(5, 4))
	assert.Equal(2.25, CeilToFloat(float32(9)/float32(4), 2))
	assert.Equal(0.15, CeilToFloat(float64(1)/float64(7), 2))
}

func TestCeilToString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCeilToFloat")

	assert.Equal("3.15", CeilToString(3.14159, 2))
	assert.Equal("3.142", CeilToString(3.14159, 3))
	assert.Equal("5.0000", CeilToString(5, 4))
	assert.Equal("2.25", CeilToString(float32(9)/float32(4), 2))
	assert.Equal("0.15", CeilToString(float64(1)/float64(7), 2))
}

func TestAverage(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAverage")

	assert.Equal(0, Average(0, 0))
	assert.Equal(1, Average(1, 1))

	avg := Average(1.2, 1.4)
	assert.Equal(1.3, RoundToFloat(avg, 1))
}

func TestSum(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSum")

	assert.Equal(1, Sum(0, 1))
	assert.Equal(1.1, Sum(0.1, float64(1)))
}

func TestMax(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMax")

	assert.Equal(0, Max(0, 0))
	assert.Equal(3, Max(1, 2, 3))
	assert.Equal(1.4, Max(1.2, 1.4, 1.1, 1.4))

	type Integer int
	assert.Equal(Integer(1), Max(Integer(1), Integer(0)))
}

func TestMaxBy(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

	assert := internal.NewAssert(t, "TestMin")

	assert.Equal(0, Min(0, 0))
	assert.Equal(1, Min(1, 2, 3))
	assert.Equal(1.1, Min(1.2, 1.4, 1.1, 1.4))
}

func TestMinBy(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

	assert := internal.NewAssert(t, "TestAngleToRadian")

	result1 := AngleToRadian(45)
	result2 := AngleToRadian(90)
	result3 := AngleToRadian(180)

	assert.Equal(0.7853981633974483, result1)
	assert.Equal(1.5707963267948966, result2)
	assert.Equal(3.141592653589793, result3)
}

func TestRadianToAngle(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAngleToRadian")

	result1 := RadianToAngle(math.Pi)
	result2 := RadianToAngle(math.Pi / 2)
	result3 := RadianToAngle(math.Pi / 4)

	assert.Equal(float64(180), result1)
	assert.Equal(float64(90), result2)
	assert.Equal(float64(45), result3)
}

func TestPointDistance(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPointDistance")

	result1 := PointDistance(1, 1, 4, 5)

	assert.Equal(float64(5), result1)
}

func TestIsPrime(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsPrime")

	assert.Equal(false, IsPrime(-1))
	assert.Equal(false, IsPrime(0))
	assert.Equal(false, IsPrime(1))
	assert.Equal(true, IsPrime(2))
	assert.Equal(true, IsPrime(3))
	assert.Equal(false, IsPrime(4))
}

func TestGCD(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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

func TestCos(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCos")

	assert.EqualValues(1, Cos(0))
	assert.EqualValues(-0.447, Cos(90))
	assert.EqualValues(-0.598, Cos(180))
	assert.EqualValues(-1, Cos(math.Pi))
	assert.EqualValues(0, Cos(math.Pi/2))
}

func TestSin(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSin")

	assert.EqualValues(0, Sin(0))
	assert.EqualValues(0.894, Sin(90))
	assert.EqualValues(-0.801, Sin(180))
	assert.EqualValues(0, Sin(math.Pi))
	assert.EqualValues(1, Sin(math.Pi/2))
}

func TestLog(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLog")

	assert.EqualValues(3, Log(8, 2))
	assert.EqualValues(3, TruncRound(Log(27, 3), 0))
	assert.EqualValues(2.32, TruncRound(Log(5, 2), 2))
}

func TestAbs(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAbs")

	assert.Equal(0, Abs(0))
	assert.Equal(1, Abs(-1))

	assert.Equal(0.1, Abs(-0.1))

	assert.Equal(int64(1), Abs(int64(-1)))
	assert.Equal(float32(1), Abs(float32(-1)))
	assert.Equal(math.Inf(1), Abs(math.Inf(-1)))
}

func TestDiv(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDiv")

	assert.Equal(float64(2), Div(4, 2))
	assert.Equal(2.5, Div(5, 2))
	assert.Equal(0.5, Div(1, 2))
	assert.Equal(0.5, Div(1, 2))
	assert.Equal(math.Inf(1), Div(8, 0))
	assert.Equal(math.Inf(-1), Div(-8, 0))
	assert.Equal(true, math.IsNaN(Div(0, 0)))
}

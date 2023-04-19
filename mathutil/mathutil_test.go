package mathutil

import (
	"math"
	"testing"

	"github.com/duke-git/lancet/internal"
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

	assert.Equal(float64(50), Percent(1, 2, 2))
	assert.Equal(float64(33.33), Percent(0.1, 0.3, 2))
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

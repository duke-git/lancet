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

func TestCos(t *testing.T) {
	assert := internal.NewAssert(t, "TestCos")

	assert.EqualValues(1, Cos(0))
	assert.EqualValues(-0.447, Cos(90))
	assert.EqualValues(-0.598, Cos(180))
	assert.EqualValues(-1, Cos(math.Pi))
	assert.EqualValues(0, Cos(math.Pi/2))
}

func TestSin(t *testing.T) {
	assert := internal.NewAssert(t, "TestSin")

	assert.EqualValues(0, Sin(0))
	assert.EqualValues(0.894, Sin(90))
	assert.EqualValues(-0.801, Sin(180))
	assert.EqualValues(0, Sin(math.Pi))
	assert.EqualValues(1, Sin(math.Pi/2))
}

func TestLog(t *testing.T) {
	assert := internal.NewAssert(t, "TestLog")

	assert.EqualValues(3, Log(8, 2))
	assert.EqualValues(3, TruncRound(Log(27, 3), 0))
	assert.EqualValues(2.32, TruncRound(Log(5, 2), 2))
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
	assert.Equal(0.15, CeilToFloat(float64(1)/float64(7), 2))
}

func TestCeilToString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCeilToFloat")

	assert.Equal("3.15", CeilToString(3.14159, 2))
	assert.Equal("3.142", CeilToString(3.14159, 3))
	assert.Equal("5.0000", CeilToString(5, 4))
	assert.Equal("0.15", CeilToString(float64(1)/float64(7), 2))
}

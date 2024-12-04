# Mathutil

Package mathutil implements some functions for math calculation.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/mathutil/mathutil.go](https://github.com/duke-git/lancet/blob/main/mathutil/mathutil.go)

<div STYLE="page-break-after: always;"></div>

## Example:

```go
import (
    "github.com/duke-git/lancet/v2/mathutil"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [Average](#Average)
-   [Exponent](#Exponent)
-   [Fibonacci](#Fibonacci)
-   [Factorial](#Factorial)
-   [Max](#Max)
-   [MaxBy](#MaxBy)
-   [Min](#Min)
-   [MinBy](#MaxBy)
-   [Percent](#Percent)
-   [RoundToFloat](#RoundToFloat)
-   [RoundToString](#RoundToString)
-   [TruncRound](#TruncRound)
-   [CeilToFloat](#CeilToFloat)
-   [CeilToString](#CeilToString)
-   [FloorToFloat](#FloorToFloat)
-   [FloorToString](#FloorToString)
-   [Range](#Range)
-   [RangeWithStep](#RangeWithStep)
-   [AngleToRadian](#AngleToRadian)
-   [RadianToAngle](#RadianToAngle)
-   [PointDistance](#PointDistance)
-   [IsPrime](#IsPrime)
-   [GCD](#GCD)
-   [LCM](#LCM)
-   [Cos](#Cos)
-   [Sin](#Sin)
-   [Log](#Log)
-   [Sum](#Sum)
-   [Abs](#Abs)
-   [Div](#Div)
-   [Variance](#Variance)
-   [StdDev](#StdDev)
-   [Permutation](#Permutation)
-   [Combination](#Combination)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Average">Average</span>

<p>Return average value of numbers. Maybe call RoundToFloat to round result.</p>

<b>Signature:</b>

```go
func Average[T constraints.Integer | constraints.Float](numbers ...T) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/Vv7LBwER-pz)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Average(1, 2)

    avg := mathutil.Average(1.2, 1.4)
    result2 := mathutil.RoundToFloat(avg, 1)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 1.5
    // 1.3
}
```

### <span id="Exponent">Exponent</span>

<p>Calculate x to the nth power.</p>

<b>Signature:</b>

```go
func Exponent(x, n int64) int64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/uF3HGNPk8wr)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Exponent(10, 0)
    result2 := mathutil.Exponent(10, 1)
    result3 := mathutil.Exponent(10, 2)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 1
    // 10
    // 100
}
```

### <span id="Fibonacci">Fibonacci</span>

<p>Calculate the nth number of fibonacci sequence.</p>

<b>Signature:</b>

```go
func Fibonacci(first, second, n int) int
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/IscseUNMuUc)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Fibonacci(1, 1, 1)
    result2 := mathutil.Fibonacci(1, 1, 2)
    result3 := mathutil.Fibonacci(1, 1, 5)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 1
    // 1
    // 5
}
```

### <span id="Factorial">Factorial</span>

<p>Calculate the factorial of x.</p>

<b>Signature:</b>

```go
func Factorial(x uint) uint
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/tt6LdOK67Nx)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Factorial(1)
    result2 := mathutil.Factorial(2)
    result3 := mathutil.Factorial(3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 1
    // 2
    // 6
}
```

### <span id="Max">Max</span>

<p>Return max value of numbers.</p>

<b>Signature:</b>

```go
func Max[T constraints.Integer | constraints.Float](numbers ...T) T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/cN8DHI0rTkH)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Max(1, 2, 3)
    result2 := mathutil.Max(1.2, 1.4, 1.1, 1.4)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 3
    // 1.4
}
```

### <span id="MaxBy">MaxBy</span>

<p>Return the maximum value of a slice using the given comparator function.</p>

<b>Signature:</b>

```go
func MaxBy[T any](slice []T, comparator func(T, T) bool) T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/pbe2MT-7DV2)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.MaxBy([]string{"a", "ab", "abc"}, func(v1, v2 string) bool {
        return len(v1) > len(v2)
    })

    result2 := mathutil.MaxBy([]string{"abd", "abc", "ab"}, func(v1, v2 string) bool {
        return len(v1) > len(v2)
    })

    result3 := mathutil.MaxBy([]string{}, func(v1, v2 string) bool {
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
```

### <span id="Min">Min</span>

<p>Return the minimum value of numbers.</p>

<b>Signature:</b>

```go
func Min[T constraints.Integer | constraints.Float](numbers ...T) T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/21BER_mlGUj)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Min(1, 2, 3)
    result2 := mathutil.Min(1.2, 1.4, 1.1, 1.4)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 1
    // 1.1
}
```

### <span id="MinBy">MinBy</span>

<p>Return the minimum value of a slice using the given comparator function.</p>

<b>Signature:</b>

```go
func MinBy[T any](slice []T, comparator func(T, T) bool) T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/N9qgYg_Ho6f)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.MinBy([]string{"a", "ab", "abc"}, func(v1, v2 string) bool {
        return len(v1) < len(v2)
    })

    result2 := mathutil.MinBy([]string{"ab", "ac", "abc"}, func(v1, v2 string) bool {
        return len(v1) < len(v2)
    })

    result3 := mathutil.MinBy([]string{}, func(v1, v2 string) bool {
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
```

### <span id="Percent">Percent</span>

<p>calculate the percentage of val to total, retain n decimal places.</p>

<b>Signature:</b>

```go
func Percent(val, total float64, n int) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/s0NdFCtwuyd)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Percent(1, 2, 2)
    result2 := mathutil.Percent(0.1, 0.3, 2)
    result3 := mathutil.Percent(-30305, 408420, 2)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 50
    // 33.33
    // -7.42
}
```

### <span id="RoundToFloat">RoundToFloat</span>

<p>Round float up to n decimal places.</p>

<b>Signature:</b>

```go
func RoundToFloat[T constraints.Float | constraints.Integer](x T, n int) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/ghyb528JRJL)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.RoundToFloat(0.124, 2)
    result2 := mathutil.RoundToFloat(0.125, 2)
    result3 := mathutil.RoundToFloat(0.125, 3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 0.12
    // 0.13
    // 0.125
}
```

### <span id="RoundToString">RoundToString</span>

<p>Round float up to n decimal places. will return string.</p>

<b>Signature:</b>

```go
func RoundToString[T constraints.Float | constraints.Integer](x T, n int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/kZwpBRAcllO)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.RoundToString(0.124, 2)
    result2 := mathutil.RoundToString(0.125, 2)
    result3 := mathutil.RoundToString(0.125, 3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 0.12
    // 0.13
    // 0.125
}
```

### <span id="TruncRound">TruncRound</span>

<p>Round float off n decimal places.</p>

<b>Signature:</b>

```go
func TruncRound[T constraints.Float | constraints.Integer](x T, n int) T 
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/aumarSHIGzP)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.TruncRound(0.124, 2)
    result2 := mathutil.TruncRound(0.125, 2)
    result3 := mathutil.TruncRound(0.125, 3)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 0.12
    // 0.12
    // 0.125
}
```

### <span id="CeilToFloat">CeilToFloat</span>

<p>Round float up n decimal places.</p>

<b>Signature:</b>

```go
func CeilToFloat[T constraints.Float | constraints.Integer](x T, n int) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/8hOeSADZPCo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.CeilToFloat(3.14159, 1)
    result2 := mathutil.CeilToFloat(3.14159, 2)
    result3 := mathutil.CeilToFloat(5, 4)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 3.2
    // 3.15
    // 5
}
```

### <span id="CeilToString">CeilToString</span>

<p>Round float up n decimal places.</p>

<b>Signature:</b>

```go
func CeilToString[T constraints.Float | constraints.Integer](x T, n int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/wy5bYEyUKKG)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.CeilToString(3.14159, 1)
    result2 := mathutil.CeilToString(3.14159, 2)
    result3 := mathutil.CeilToString(5, 4)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 3.2
    // 3.15
    // 5.0000
}
```

### <span id="FloorToFloat">FloorToFloat</span>

<p>Round float down n decimal places.</p>

<b>Signature:</b>

```go
func FloorToFloat[T constraints.Float | constraints.Integer](x T, n int) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/vbCBrQHZEED)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.FloorToFloat(3.14159, 1)
    result2 := mathutil.FloorToFloat(3.14159, 2)
    result3 := mathutil.FloorToFloat(5, 4)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 3.1
    // 3.14
    // 5
}
```

### <span id="FloorToString">FloorToString</span>

<p>Round float down n decimal places.</p>

<b>Signature:</b>

```go
func FloorToString[T constraints.Float | constraints.Integer](x T, n int) string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/Qk9KPd2IdDb)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.FloorToString(3.14159, 1)
    result2 := mathutil.FloorToString(3.14159, 2)
    result3 := mathutil.FloorToString(5, 4)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 3.1
    // 3.14
    // 5.0000
}
```

### <span id="Range">Range</span>

<p>Creates a slice of numbers from start with specified count, element step is 1.</p>

<b>Signature:</b>

```go
func Range[T constraints.Integer | constraints.Float](start T, count int) []T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/9ke2opxa8ZP)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Range(1, 4)
    result2 := mathutil.Range(1, -4)
    result3 := mathutil.Range(-4, 4)
    result4 := mathutil.Range(1.0, 4)

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
```

### <span id="RangeWithStep">RangeWithStep</span>

<p>Creates a slice of numbers from start to end with specified step.</p>

<b>Signature:</b>

```go
func RangeWithStep[T constraints.Integer | constraints.Float](start, end, step T) []T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/akLWz0EqOSM)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.RangeWithStep(1, 4, 1)
    result2 := mathutil.RangeWithStep(1, -1, 0)
    result3 := mathutil.RangeWithStep(-4, 1, 2)
    result4 := mathutil.RangeWithStep(1.0, 4.0, 1.1)

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
```

### <span id="AngleToRadian">AngleToRadian</span>

<p>Converts angle value to radian value.</p>

<b>Signature:</b>

```go
func AngleToRadian(angle float64) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/CIvlICqrHql)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.AngleToRadian(45)
    result2 := mathutil.AngleToRadian(90)
    result3 := mathutil.AngleToRadian(180)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 0.7853981633974483
    // 1.5707963267948966
    // 3.141592653589793
}
```

### <span id="RadianToAngle">RadianToAngle</span>

<p>Converts radian value to angle value.</p>

<b>Signature:</b>

```go
func RadianToAngle(radian float64) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/dQtmOTUOMgi)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.RadianToAngle(math.Pi)
    result2 := mathutil.RadianToAngle(math.Pi / 2)
    result3 := mathutil.RadianToAngle(math.Pi / 4)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 180
    // 90
    // 45
}
```

### <span id="PointDistance">PointDistance</span>

<p>Caculates two points distance.</p>

<b>Signature:</b>

```go
func PointDistance(x1, y1, x2, y2 float64) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/RrG4JIaziM8)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.PointDistance(1, 1, 4, 5)

    fmt.Println(result1)

    // Output:
    // 5
}
```

### <span id="IsPrime">IsPrime</span>

<p>Checks if number is prime number.</p>

<b>Signature:</b>

```go
func IsPrime(n int) bool
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/Rdd8UTHZJ7u)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.IsPrime(-1)
    result2 := mathutil.IsPrime(0)
    result3 := mathutil.IsPrime(1)
    result4 := mathutil.IsPrime(2)

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
```

### <span id="GCD">GCD</span>

<p>Return greatest common divisor (GCD) of integers.</p>

<b>Signature:</b>

```go
func GCD[T constraints.Integer](integers ...T) T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/CiEceLSoAKB)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.GCD(1, 1)
    result2 := mathutil.GCD(1, -1)
    result3 := mathutil.GCD(-1, 1)
    result4 := mathutil.GCD(-1, -1)
    result5 := mathutil.GCD(3, 6, 9)

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
```

### <span id="LCM">LCM</span>

<p>Return Least Common Multiple (LCM) of integers.</p>

<b>Signature:</b>

```go
func LCM[T constraints.Integer](integers ...T) T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/EjcZxfY7G_g)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.LCM(1, 1)
    result2 := mathutil.LCM(1, 2)
    result3 := mathutil.LCM(3, 6, 9)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 1
    // 2
    // 18
}
```

### <span id="Cos">Cos</span>

<p>Returns the cosine of the radian argument.</p>

<b>Signature:</b>

```go
func Cos(radian float64, precision ...int) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/Sm89LoIfvFq)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Cos(0)
    result2 := mathutil.Cos(90)
    result3 := mathutil.Cos(180)
    result4 := mathutil.Cos(math.Pi)
    result5 := mathutil.Cos(math.Pi / 2)

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
```

### <span id="Sin">Sin</span>

<p>Returns the sine of the radian argument.</p>

<b>Signature:</b>

```go
func Sin(radian float64, precision ...int) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/TWMQlMywDsP)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Sin(0)
    result2 := mathutil.Sin(90)
    result3 := mathutil.Sin(180)
    result4 := mathutil.Sin(math.Pi)
    result5 := mathutil.Sin(math.Pi / 2)

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
```

### <span id="Log">Log</span>

<p>Returns the logarithm of base n.</p>

<b>Signature:</b>

```go
func Log(n, base float64) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/_d4bi8oyhat)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Log(8, 2)
    result2 := mathutil.TruncRound(mathutil.Log(5, 2), 2)
    result3 := mathutil.TruncRound(mathutil.Log(27, 3), 0)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 3
    // 2.32
    // 3
}
```

### <span id="Sum">Sum</span>

<p>Returns sum of passed numbers.</p>

<b>Signature:</b>

```go
func Sum[T constraints.Integer | constraints.Float](numbers ...T) T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/1To2ImAMJA7)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Sum(1, 2)
    result2 := mathutil.Sum(0.1, float64(1))

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 3
    // 1.1
}
```

### <span id="Abs">Abs</span>

<p>Returns the absolute value of x.</p>

<b>Signature:</b>

```go
func Abs[T constraints.Integer | constraints.Float](x T) T
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/fsyBh1Os-1d)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Abs(-1)
    result2 := mathutil.Abs(-0.1)
    result3 := mathutil.Abs(float32(0.2))

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)

    // Output:
    // 1
    // 0.1
    // 0.2
}
```

### <span id="Div">Div</span>

<p>Returns the result of x divided by y.</p>

<b>Signature:</b>

```go
func Div[T constraints.Float | constraints.Integer](x T, y T) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/WLxDdGXXYat)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Div(9, 4)
    result2 := mathutil.Div(1, 2)
    result3 := mathutil.Div(0, 666)

    fmt.Println(result1)
    fmt.Println(result2)
    fmt.Println(result3)
    // Output:
    // 2.25
    // 0.5
    // 0
}
```

### <span id="Variance">Variance</span>

<p>Returns the variance of numbers.</p>

<b>Signature:</b>

```go
func Variance[T constraints.Float | constraints.Integer](numbers []T) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/uHuV4YgXf8F)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Variance([]int{1, 2, 3, 4, 5})
    result2 := mathutil.Variance([]float64{1.1, 2.2, 3.3, 4.4, 5.5})

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 2
    // 2.42
}
```

### <span id="StdDev">StdDev</span>

<p>Returns the standard deviation of numbers.</p>

<b>Signature:</b>

```go
func StdDev[T constraints.Float | constraints.Integer](numbers []T) float64
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/FkNZDXvHD2l)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.TruncRound(mathutil.StdDev([]int{1, 2, 3, 4, 5}), 2)
    result2 := mathutil.TruncRound(mathutil.StdDev([]float64{1.1, 2.2, 3.3, 4.4, 5.5}), 2)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 1.41
    // 1.55
}
```

### <span id="Permutation">Permutation</span>

<p>Calculates P(n, k).</p>

<b>Signature:</b>

```go
func Permutation(n, k uint) uint
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/MgobwH_FOxj)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Permutation(5, 3)
	result2 := mathutil.Permutation(5, 5)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 60
	// 120
}
```

### <span id="Combination">Combination</span>

<p>Calculates C(n, k).</p>

<b>Signature:</b>

```go
func Combination(n, k uint) uint
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/ENFQRDQUFi9)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Combination(5, 3)
	result2 := mathutil.Combination(5, 5)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 10
	// 1
}
```
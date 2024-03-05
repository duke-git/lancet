# Mathutil

Package mathutil implements some functions for math calculation.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/v1/mathutil/mathutil.go](https://github.com/duke-git/lancet/blob/v1/mathutil/mathutil.go)

<div STYLE="page-break-after: always;"></div>

## Example:

```go
import (
    "github.com/duke-git/lancet/mathutil"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [Exponent](#Exponent)
-   [Fibonacci](#Fibonacci)
-   [Factorial](#Factorial)
-   [Percent](#Percent)
-   [RoundToFloat](#RoundToFloat)
-   [RoundToString](#RoundToString)
-   [TruncRound](#TruncRound)
-   [CeilToFloat](#CeilToFloat)
-   [CeilToString](#CeilToString)
-   [FloorToFloat](#FloorToFloat)
-   [FloorToString](#FloorToString)
-   [AngleToRadian](#AngleToRadian)
-   [RadianToAngle](#RadianToAngle)
-   [PointDistance](#PointDistance)
-   [IsPrime](#IsPrime)
-   [GCD](#GCD)
-   [LCM](#LCM)
-   [Cos](#Cos)
-   [Sin](#Sin)
-   [Log](#Log)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Exponent">Exponent</span>

<p>Calculate x to the nth power.</p>

<b>Signature:</b>

```go
func Exponent(x, n int64) int64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
)

func main() {
    fmt.Println(mathutil.Exponent(10, 0)) //1
    fmt.Println(mathutil.Exponent(10, 1)) //10
    fmt.Println(mathutil.Exponent(10, 2)) //100
}
```

### <span id="Fibonacci">Fibonacci</span>

<p>Calculate the nth number of fibonacci sequence.</p>

<b>Signature:</b>

```go
func Fibonacci(first, second, n int) int
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
)

func main() {
    fmt.Println(mathutil.Fibonacci(1, 1, 1)) //1
    fmt.Println(mathutil.Fibonacci(1, 1, 2)) //1
    fmt.Println(mathutil.Fibonacci(1, 1, 3)) //2
    fmt.Println(mathutil.Fibonacci(1, 1, 4)) //3
    fmt.Println(mathutil.Fibonacci(1, 1, 5)) //5
}
```

### <span id="Factorial">Factorial</span>

<p>Calculate the factorial of x.</p>

<b>Signature:</b>

```go
func Factorial(x uint) uint
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
)

func main() {
    fmt.Println(mathutil.Factorial(0)) //1
    fmt.Println(mathutil.Factorial(1)) //1
    fmt.Println(mathutil.Factorial(2)) //2
    fmt.Println(mathutil.Factorial(3)) //6
}
```

### <span id="Percent">Percent</span>

<p>calculate the percentage of val to total, retain n decimal places.</p>

<b>Signature:</b>

```go
func Percent(val, total float64, n int) float64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
)

func main() {
    fmt.Println(mathutil.Percent(1, 2, 2)) //1
    fmt.Println(mathutil.Percent(0.1, 0.3, 2)) //33.33
}
```

### <span id="RoundToFloat">RoundToFloat</span>

<p>Round float up to n decimal places.</p>

<b>Signature:</b>

```go
func RoundToFloat(x float64, n int) float64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
)

func main() {
    fmt.Println(mathutil.RoundToFloat(0, 0)) //0
    fmt.Println(mathutil.RoundToFloat(0, 1)) //0
    fmt.Println(mathutil.RoundToFloat(0.124, 2)) //0.12
    fmt.Println(mathutil.RoundToFloat(0.125, 2)) //0.13
    fmt.Println(mathutil.RoundToFloat(0.125, 3)) //0.125
}
```

### <span id="RoundToString">RoundToString</span>

<p>Round float up to n decimal places. will return string.</p>

<b>Signature:</b>

```go
func RoundToString(x float64, n int) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
)

func main() {
    fmt.Println(mathutil.RoundToString(0, 0)) //"0"
    fmt.Println(mathutil.RoundToString(0, 1)) //"0.0:
    fmt.Println(mathutil.RoundToString(0.124, 2)) //"0.12"
    fmt.Println(mathutil.RoundToString(0.125, 2)) //"0.13"
    fmt.Println(mathutil.RoundToString(0.125, 3)) //"0.125"
}
```

### <span id="TruncRound">TruncRound</span>

<p>Round float off n decimal places.</p>

<b>Signature:</b>

```go
func TruncRound(x float64, n int) float64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
)

func main() {
    fmt.Println(mathutil.TruncRound(0, 0)) //0
    fmt.Println(mathutil.TruncRound(0, 1)) //0
    fmt.Println(mathutil.TruncRound(0.124, 2)) //0.12
    fmt.Println(mathutil.TruncRound(0.125, 2)) //0.12
    fmt.Println(mathutil.TruncRound(0.125, 3)) //0.125
}
```

### <span id="CeilToFloat">CeilToFloat</span>

<p>Round float up n decimal places.</p>

<b>Signature:</b>

```go
func CeilToFloat(x float64, n int) float64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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
func CeilToString(x float64, n int) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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
func CeilToString(x float64, n int) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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
func CeilToString(x float64, n int) string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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

### <span id="AngleToRadian">AngleToRadian</span>

<p>Converts angle value to radian value.</p>

<b>Signature:</b>

```go
func AngleToRadian(angle float64) float64
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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
func GCD(integers ...int) int
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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
func LCM(integers ...int) int
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/mathutil"
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

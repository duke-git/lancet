# Mathutil

mathutil 包实现了一些数学计算的函数.

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/v1/mathutil/mathutil.go](https://github.com/duke-git/lancet/blob/v1/mathutil/mathutil.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/mathutil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [Exponent](#Exponent)
-   [Fibonacci](#Fibonacci)
-   [Factorial](#Factorial)
-   [Percent](#Percent)
-   [RoundToFloat](#RoundToFloat)
-   [RoundToString](#RoundToString)
-   [TruncRound](#TruncRound)
-   [AngleToRadian](#AngleToRadian)
-   [RadianToAngle](#RadianToAngle)
-   [PointDistance](#PointDistance)
-   [IsPrime](#IsPrime)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Exponent">Exponent</span>

<p>指数计算（x的n次方）</p>

<b>函数签名:</b>

```go
func Exponent(x, n int64) int64
```

<b>例子:</b>

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

<p>计算斐波那契数列的第n个数</p>

<b>函数签名:</b>

```go
func Fibonacci(first, second, n int) int
```

<b>例子:</b>

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

<p>计算阶乘</p>

<b>函数签名:</b>

```go
func Factorial(x uint) uint
```

<b>例子:</b>

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

<p>计算百分比，保留n位小数</p>

<b>函数签名:</b>

```go
func Percent(val, total float64, n int) float64
```

<b>例子:</b>

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

<p>四舍五入，保留n位小数</p>

<b>函数签名:</b>

```go
func RoundToFloat(x float64, n int) float64
```

<b>例子:</b>

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

<p>四舍五入，保留n位小数，返回字符串</p>

<b>函数签名:</b>

```go
func RoundToString(x float64, n int) string
```

<b>例子:</b>

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

<p>截短n位小数（不进行四舍五入）</p>

<b>函数签名:</b>

```go
func TruncRound(x float64, n int) float64
```

<b>例子:</b>

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

### <span id="AngleToRadian">AngleToRadian</span>

<p>将角度值转为弧度值</p>

<b>函数签名:</b>

```go
func AngleToRadian(angle float64) float64
```

<b>示例:</b>

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

<p>将弧度值转为角度值</p>

<b>函数签名:</b>

```go
func RadianToAngle(radian float64) float64
```

<b>示例:</b>

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

<p>计算两个坐标点的距离</p>

<b>函数签名:</b>

```go
func PointDistance(x1, y1, x2, y2 float64) float64
```

<b>示例:</b>

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

<p>判断质数。</p>

<b>函数签名:</b>

```go
func IsPrime(n int) bool
```

<b>示例:</b>

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

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

### <span id="CeilToFloat">CeilToFloat</span>

<p>向上舍入（进一法），保留n位小数。</p>

<b>函数签名:</b>

```go
func CeilToFloat(x float64, n int) float64
```

<b>示例:</b>

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

<p>向上舍入（进一法），保留n位小数，返回字符串。</p>

<b>函数签名:</b>

```go
func CeilToString(x float64, n int) string
```

<b>示例:</b>

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

<p>向下舍入（去尾法），保留n位小数。</p>

<b>函数签名:</b>

```go
func CeilToString(x float64, n int) string
```

<b>示例:</b>

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

<p>向下舍入（去尾法），保留n位小数，返回字符串。</p>

<b>函数签名:</b>

```go
func CeilToString(x float64, n int) string
```

<b>示例:</b>

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


### <span id="GCD">GCD</span>

<p>计算整数最大公约数。</p>

<b>函数签名:</b>

```go
func GCD(integers ...int) int
```

<b>示例:</b>

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

<p>计算整数最小公倍数。</p>

<b>函数签名:</b>

```go
func LCM(integers ...int) int
```

<b>示例:</b>

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

<p>计算弧度的余弦值。</p>

<b>函数签名:</b>

```go
func Cos(radian float64, precision ...int) float64
```

<b>示例:</b>

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

<p>计算弧度的正弦值。</p>

<b>函数签名:</b>

```go
func Sin(radian float64, precision ...int) float64
```

<b>示例:</b>

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

<p>计算以base为底n的对数。</p>

<b>函数签名:</b>

```go
func Log(n, base float64) float64
```

<b>示例:</b>

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
# Mathutil

mathutil 包实现了一些数学计算的函数.

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/mathutil/mathutil.go](https://github.com/duke-git/lancet/blob/main/mathutil/mathutil.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/mathutil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

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
-   [T运行cRound](#T运行cRound)
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

## 文档

### <span id="Average">Average</span>

<p>计算平均数. 可能需要对结果调用RoundToFloat方法四舍五入</p>

<b>函数签名:</b>

```go
func Average[T constraints.Integer | constraints.Float](numbers ...T) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/HFd70x4DrMj)</span></b>

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

<p>指数计算（x的n次方）</p>

<b>函数签名:</b>

```go
func Exponent(x, n int64) int64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Vv7LBwER-pz)</span></b>

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

<p>计算斐波那契数列的第n个数</p>

<b>函数签名:</b>

```go
func Fibonacci(first, second, n int) int
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/IscseUNMuUc)</span></b>

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

<p>计算阶乘</p>

<b>函数签名:</b>

```go
func Factorial(x uint) uint
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/tt6LdOK67Nx)</span></b>

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

<p>返回参数中的最大数</p>

<b>函数签名:</b>

```go
func Max[T constraints.Integer | constraints.Float](numbers ...T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/cN8DHI0rTkH)</span></b>

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

<p>使用给定的比较器函数返回切片的最大值</p>

<b>函数签名:</b>

```go
func MaxBy[T any](slice []T, comparator func(T, T) bool) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/pbe2MT-7DV2)</span></b>

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

<p>返回参数中的最小数</p>

<b>函数签名:</b>

```go
func Min[T constraints.Integer | constraints.Float](numbers ...T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/pbe2MT-7DV2)</span></b>

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

<p>使用给定的比较器函数返回切片的最小值</p>

<b>函数签名:</b>

```go
func MinBy[T any](slice []T, comparator func(T, T) bool) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/N9qgYg_Ho6f)</span></b>

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

<p>计算百分比，保留n位小数</p>

<b>函数签名:</b>

```go
func Percent(val, total float64, n int) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/s0NdFCtwuyd)</span></b>

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

<p>四舍五入，保留n位小数</p>

<b>函数签名:</b>

```go
func RoundToFloat[T constraints.Float | constraints.Integer](x T, n int) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ghyb528JRJL)</span></b>

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

<p>四舍五入，保留n位小数，返回字符串</p>

<b>函数签名:</b>

```go
func RoundToString[T constraints.Float | constraints.Integer](x T, n int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/kZwpBRAcllO)</span></b>

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

### <span id="T运行cRound">T运行cRound</span>

<p>截短n位小数（不进行四舍五入）</p>

<b>函数签名:</b>

```go
func T运行cRound[T constraints.Float | constraints.Integer](x T, n int) T 
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/aumarSHIGzP)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.T运行cRound(0.124, 2)
    result2 := mathutil.T运行cRound(0.125, 2)
    result3 := mathutil.T运行cRound(0.125, 3)

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

<p>向上舍入（进一法），保留n位小数。</p>

<b>函数签名:</b>

```go
func CeilToFloat[T constraints.Float | constraints.Integer](x T, n int) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/8hOeSADZPCo)</span></b>

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

<p>向上舍入（进一法），保留n位小数，返回字符串。</p>

<b>函数签名:</b>

```go
func CeilToString[T constraints.Float | constraints.Integer](x T, n int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/wy5bYEyUKKG)</span></b>

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

<p>向下舍入（去尾法），保留n位小数。</p>

<b>函数签名:</b>

```go
func FloorToFloat[T constraints.Float | constraints.Integer](x T, n int) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/vbCBrQHZEED)</span></b>

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

<p>向下舍入（去尾法），保留n位小数，返回字符串。</p>

<b>函数签名:</b>

```go
func FloorToString[T constraints.Float | constraints.Integer](x T, n int) string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Qk9KPd2IdDb)</span></b>

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

<p>根据指定的起始值和数量，创建一个数字切片。</p>

<b>函数签名:</b>

```go
func Range[T constraints.Integer | constraints.Float](start T, count int) []T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/9ke2opxa8ZP)</span></b>

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

<p>根据指定的起始值，结束值，步长，创建一个数字切片。</p>

<b>函数签名:</b>

```go
func RangeWithStep[T constraints.Integer | constraints.Float](start, end, step T) []T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/akLWz0EqOSM)</span></b>

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

<p>将角度值转为弧度值</p>

<b>函数签名:</b>

```go
func AngleToRadian(angle float64) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/CIvlICqrHql)</span></b>

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

<p>将弧度值转为角度值</p>

<b>函数签名:</b>

```go
func RadianToAngle(radian float64) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/dQtmOTUOMgi)</span></b>

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

<p>计算两个坐标点的距离</p>

<b>函数签名:</b>

```go
func PointDistance(x1, y1, x2, y2 float64) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/RrG4JIaziM8)</span></b>

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

<p>判断质数。</p>

<b>函数签名:</b>

```go
func IsPrime(n int) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Rdd8UTHZJ7u)</span></b>

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

<p>计算最大公约数。</p>

<b>函数签名:</b>

```go
func GCD[T constraints.Integer](integers ...T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/CiEceLSoAKB)</span></b>

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

<p>计算最小公倍数。</p>

<b>函数签名:</b>

```go
func LCM[T constraints.Integer](integers ...T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/EjcZxfY7G_g)</span></b>

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

<p>计算弧度的余弦值</p>

<b>函数签名:</b>

```go
func Cos(radian float64, precision ...int) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Sm89LoIfvFq)</span></b>

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

<p>计算弧度的正弦值</p>

<b>函数签名:</b>

```go
func Sin(radian float64, precision ...int) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/TWMQlMywDsP)</span></b>

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

<p>计算以base为底n的对数。</p>

<b>函数签名:</b>

```go
func Log(n, base float64) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/_d4bi8oyhat)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Log(8, 2)
    result2 := mathutil.T运行cRound(mathutil.Log(5, 2), 2)
    result3 := mathutil.T运行cRound(mathutil.Log(27, 3), 0)

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

<p>求传入参数之和。</p>

<b>函数签名:</b>

```go
func Sum[T constraints.Integer | constraints.Float](numbers ...T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/1To2ImAMJA7)</span></b>

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

<p>求绝对值。</p>

<b>函数签名:</b>

```go
func Abs[T constraints.Integer | constraints.Float](x T) T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/fsyBh1Os-1d)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
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
```

### <span id="Div">Div</span>

<p>除法运算。</p>

<b>函数签名:</b>

```go
func Div[T constraints.Float | constraints.Integer](x T, y T) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/WLxDdGXXYat)</span></b>

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

<p>计算方差。</p>

<b>函数签名:</b>

```go
func Variance[T constraints.Float | constraints.Integer](numbers []T) float64
```

<b>示例:<span style="float:right;display:inline-block;">[示例](https://go.dev/play/p/uHuV4YgXf8F)</span></b>

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

<p>计算标准差。</p>

<b>函数签名:</b>

```go
func StdDev[T constraints.Float | constraints.Integer](numbers []T) float64
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/FkNZDXvHD2l)</span></b>

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

<p>计算排列数P(n, k)。</p>

<b>函数签名:</b>

```go
func Permutation(n, k uint) uint
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/MgobwH_FOxj)</span></b>

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

<p>计算组合数C(n, k)。</p>

<b>函数签名:</b>

```go
func Combination(n, k uint) uint
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ENFQRDQUFi9)</span></b>

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
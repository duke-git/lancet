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
-   [TruncRound](#TruncRound)
-   [Range](#Range)
-   [RangeWithStep](#RangeWithStep)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Average">Average</span>

<p>计算平均数. 可能需要对结果调用RoundToFloat方法四舍五入</p>

<b>函数签名:</b>

```go
func Average[T constraints.Integer | constraints.Float](numbers ...T) T
```

<b>示例:</b>

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
    // 1
    // 1.3
}
```

### <span id="Exponent">Exponent</span>

<p>指数计算（x的n次方）</p>

<b>函数签名:</b>

```go
func Exponent(x, n int64) int64
```

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
    result1 := mathutil.Percent(1, 2, 2)
    result2 := mathutil.Percent(0.1, 0.3, 2)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 0.5
    // 0.33
}
```

### <span id="RoundToFloat">RoundToFloat</span>

<p>四舍五入，保留n位小数</p>

<b>函数签名:</b>

```go
func RoundToFloat(x float64, n int) float64
```

<b>示例:</b>

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
func RoundToString(x float64, n int) string
```

<b>示例:</b>

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

<p>截短n位小数（不进行四舍五入）</p>

<b>函数签名:</b>

```go
func TruncRound(x float64, n int) float64
```

<b>示例:</b>

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

### <span id="Range">Range</span>

<p>根据指定的起始值和数量，创建一个数字切片。</p>

<b>函数签名:</b>

```go
func Range[T constraints.Integer | constraints.Float](start T, count int) []T
```

<b>示例:</b>

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

<b>示例:</b>

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

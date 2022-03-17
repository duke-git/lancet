# Mathutil
mathutil包实现了一些数学计算的函数.

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/mathutil/mathutil.go](https://github.com/duke-git/lancet/blob/main/mathutil/mathutil.go)


<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/v2/mathutil"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [Exponent](#Exponent)
- [Fibonacci](#Fibonacci)
- [Factorial](#Factorial)

- [Percent](#Percent)
- [RoundToFloat](#RoundToFloat)
- [RoundToString](#RoundToString)
- [TruncRound](#TruncRound)

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
    "github.com/duke-git/lancet/v2/mathutil"
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
    "github.com/duke-git/lancet/v2/mathutil"
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
    "github.com/duke-git/lancet/v2/mathutil"
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
    "github.com/duke-git/lancet/v2/mathutil"
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
    "github.com/duke-git/lancet/v2/mathutil"
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
    "github.com/duke-git/lancet/v2/mathutil"
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
    "github.com/duke-git/lancet/v2/mathutil"
)

func main() {
	fmt.Println(mathutil.TruncRound(0, 0)) //0
	fmt.Println(mathutil.TruncRound(0, 1)) //0
	fmt.Println(mathutil.TruncRound(0.124, 2)) //0.12
	fmt.Println(mathutil.TruncRound(0.125, 2)) //0.12
	fmt.Println(mathutil.TruncRound(0.125, 3)) //0.125
}
```




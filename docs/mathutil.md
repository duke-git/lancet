# Mathutil
Package mathutil implements some functions for math calculation.

<div STYLE="page-break-after: always;"></div>

## Source:

[https://github.com/duke-git/lancet/blob/main/mathutil/mathutil.go](https://github.com/duke-git/lancet/blob/main/mathutil/mathutil.go)


<div STYLE="page-break-after: always;"></div>

## Example:
```go
import (
    "github.com/duke-git/lancet/mathutil"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
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




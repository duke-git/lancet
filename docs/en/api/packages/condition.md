# Condition
Package condition contains some functions for conditional judgment. eg. And, Or, TernaryOperator... The implementation of this package refers to the implementation of carlmjohnson's truthy package, you may find more useful information in [truthy](https://github.com/carlmjohnson/truthy), thanks carlmjohnson.

<div STYLE="page-break-after: always;"></div>

## Source:

- [https://github.com/duke-git/lancet/blob/main/condition/condition.go](https://github.com/duke-git/lancet/blob/main/condition/condition.go)

<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/v2/condition"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

- [Bool](#Bool)
- [And](#And)
- [Or](#Or)
- [Xor](#Generate)
- [Nor](#Nor)
- [Xnor](#Xnor)
- [Nand](#Nand)
- [Ternary](#Ternary)
- [TernaryOperator<sup>deprecated</sup>](#TernaryOperator)

<div STYLE="page-break-after: always;"></div>

## Documentation


### <span id="Bool">Bool</span>
<p>Returns the truthy value of anything.<br/>
If the value's type has a Bool() bool method, the method is called and returned.<br/>
If the type has an IsZero() bool method, the opposite value is returned.<br/>
Slices and maps are truthy if they have a length greater than zero.<br/>
All other types are truthy if they are not their zero value.</p>

<b>Signature:</b>

```go
func Bool[T any](value T) bool
```
<b>Example:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ETzeDJRSvhm)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
    // bool
    result1 := condition.Bool(false)
    result2 := condition.Bool(true)
    fmt.Println(result1) // false
    fmt.Println(result2) // true

    // integer
    result3 := condition.Bool(0) 
    result4 := condition.Bool(1)
    fmt.Println(result3) // false
    fmt.Println(result4) // true

    // string
    result5 := condition.Bool("")
    result6 := condition.Bool(" ")
    fmt.Println(result5) // false
    fmt.Println(result6) // true

    // slice
    nums := []int{}
    result7 := condition.Bool(nums)

    nums = append(nums, 1, 2)
    result8 := condition.Bool(nums)
    fmt.Println(result7) // false
    fmt.Println(result8) // true

    // struct
    result9 = condition.Bool(struct{}{})
    fmt.Println(result8) // false


    // Output:
    // false
    // true
    // false
    // true
    // false
    // true
    // false
    // true
    // false
}
```


### <span id="And">And</span>
<p>Returns true if both a and b are truthy.</p>

<b>Signature:</b>

```go
func And[T, U any](a T, b U) bool
```
<b>Example:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/W1SSUmt6pvr)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
    fmt.Println(condition.And(0, 1)) // false
    fmt.Println(condition.And(0, "")) // false
    fmt.Println(condition.And(0, "0")) // false
    fmt.Println(condition.And(1, "0")) // true
}
```



### <span id="Or">Or</span>
<p>Returns false if neither a nor b is truthy.</p>

<b>Signature:</b>

```go
func Or[T, U any](a T, b U) bool
```
<b>Example:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/UlQTxHaeEkq)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
    fmt.Println(condition.Or(0, "")) // false
    fmt.Println(condition.Or(0, 1)) // true
    fmt.Println(condition.Or(0, "0")) // true
    fmt.Println(condition.Or(1, "0")) // true
}
```



### <span id="Xor">Xor</span>
<p>Returns true if a or b but not both is truthy.</p>

<b>Signature:</b>

```go
func Xor[T, U any](a T, b U) bool
```
<b>Example:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/gObZrW7ZbG8)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
    fmt.Println(condition.Xor(0, 0)) // false
    fmt.Println(condition.Xor(0, 1)) // true
    fmt.Println(condition.Xor(1, 0)) // true
    fmt.Println(condition.Xor(1, 1)) // false
}
```



### <span id="Nor">Nor</span>
<p>Returns true if neither a nor b is truthy.</p>

<b>Signature:</b>

```go
func Nor[T, U any](a T, b U) bool
```
<b>Example:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/g2j08F_zZky)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
    fmt.Println(condition.Nor(0, 0)) // true
    fmt.Println(condition.Nor(0, 1)) // false
    fmt.Println(condition.Nor(1, 0)) // false
    fmt.Println(condition.Nor(1, 1)) // false
}
```


### <span id="Xnor">Xnor</span>
<p>Returns true if both a and b or neither a nor b are truthy.</p>

<b>Signature:</b>

```go
func Xnor[T, U any](a T, b U) bool
```
<b>Example:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/OuDB9g51643)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
    fmt.Println(condition.Xnor(0, 0)) // true
    fmt.Println(condition.Xnor(0, 1)) // false
    fmt.Println(condition.Xnor(1, 0)) // false
    fmt.Println(condition.Xnor(1, 1)) // true
}
```


### <span id="Nand">Nand</span>
<p>Returns false if both a and b are truthy</p>

<b>Signature:</b>

```go
func Nand[T, U any](a T, b U) bool
```
<b>Example:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/vSRMLxLIbq8)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
    fmt.Println(condition.Nand(0, 0)) // true
    fmt.Println(condition.Nand(0, 1)) // true
    fmt.Println(condition.Nand(1, 0)) // true
    fmt.Println(condition.Nand(1, 1)) // false
}
```



### <span id="Ternary">Ternary</span>
<p>Checks the value of param `isTrue`, if true return ifValue else return elseValue</p>

<b>Signature:</b>

```go
func Ternary[T, U any](isTrue T, ifValue U, elseValue U) U
```
<b>Example:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ElllPZY0guT)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
    conditionTrue := 2 > 1
    result1 := condition.Ternary(conditionTrue, 0, 1)

    conditionFalse := 2 > 3
    result2 := condition.Ternary(conditionFalse, 0, 1)
    
    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 0
    // 1
}
```

### <span id="TernaryOperator">TernaryOperator</span>
<p>Checks the value of param `isTrue`, if true return ifValue else return elseValue</p>

> ⚠️ This function is deprecated. use `Ternary` instead.

<b>Signature:</b>

```go
func TernaryOperator[T, U any](isTrue T, ifValue U, elseValue U) U
```
<b>Example:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ElllPZY0guT)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
    conditionTrue := 2 > 1
    result1 := condition.TernaryOperator(conditionTrue, 0, 1)

    conditionFalse := 2 > 3
    result2 := condition.TernaryOperator(conditionFalse, 0, 1)
    
    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 0
    // 1
}
```






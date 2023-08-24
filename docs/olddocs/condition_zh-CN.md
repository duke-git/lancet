# Condition
condition包含一些用于条件判断的函数。这个包的实现参考了carlmjohnson的truthy包的实现，更多有用的信息可以在[truthy](https://github.com/carlmjohnson/truthy)中找到，感谢carlmjohnson。

<div STYLE="page-break-after: always;"></div>

## 源码:

- [https://github.com/duke-git/lancet/blob/main/condition/condition.go](https://github.com/duke-git/lancet/blob/main/condition/condition.go)

<div STYLE="page-break-after: always;"></div>

## 用法:
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
- [TernaryOperator](#TernaryOperator)

<div STYLE="page-break-after: always;"></div>

## 目录

### <span id="Bool">Bool</span>
<p>返回传入参数的bool值.<br/>
如果出入类型参数含有Bool方法, 会调用该方法并返回<br/>
如果传入类型参数有IsZero方法, 返回IsZero方法返回值的取反<br/>
slices和map的length大于0时，返回true，否则返回false<br/>
其他类型会判断是否是零值</p>

<b>函数签名:</b>

```go
func Bool[T any](value T) bool
```
<b>示例:</b>

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
<p>逻辑且操作，当切仅当a和b都为true时返回true</p>

<b>函数签名:</b>

```go
func And[T, U any](a T, b U) bool
```
<b>示例:</b>

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
<p>逻辑或操作，当切仅当a和b都为false时返回false</p>

<b>函数签名:</b>

```go
func Or[T, U any](a T, b U) bool
```
<b>示例:</b>

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
<p>逻辑异或操作，a和b相同返回false，a和b不相同返回true</p>

<b>函数签名:</b>

```go
func Xor[T, U any](a T, b U) bool
```
<b>示例:</b>

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
<p>异或的取反操作</p>

<b>函数签名:</b>

```go
func Nor[T, U any](a T, b U) bool
```
<b>示例:</b>

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
<p>如果a和b都是真的或a和b均是假的，则返回true。</p>

<b>函数签名:</b>

```go
func Xnor[T, U any](a T, b U) bool
```
<b>示例:</b>

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
<p>如果a和b都为真，返回false，否则返回true</p>

<b>函数签名:</b>

```go
func Nand[T, U any](a T, b U) bool
```
<b>示例:</b>

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

### <span id="TernaryOperator">TernaryOperator</span>
<p>三元运算符</p>

<b>函数签名:</b>

```go
func TernaryOperator[T, U any](isTrue T, ifValue U, elseValue U) U
```
<b>示例:</b>

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







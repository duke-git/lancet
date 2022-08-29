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
- [Nand](#Nand)
- [TernaryOperator](#TernaryOperator)

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
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
	// bool
	fmt.Println(condition.Bool(false)) // false
	fmt.Println(condition.Bool(true)) // true

	// integer
	fmt.Println(condition.Bool(0)) // false
	fmt.Println(condition.Bool(1)) // true

	// float
	fmt.Println(condition.Bool(0.0)) // false
	fmt.Println(condition.Bool(0.1)) // true

	// string
	fmt.Println(condition.Bool("")) // false
	fmt.Println(condition.Bool(" ")) // true
	fmt.Println(condition.Bool("0")) // true

	// slice
	var nums [2]int
	fmt.Println(condition.Bool(nums)) // false
	nums = [2]int{0, 1}
	fmt.Println(condition.Bool(nums)) // true

	// map
	fmt.Println(condition.Bool(map[string]string{})) // false
	fmt.Println(condition.Bool(map[string]string{"a": "a"})) // true

	// struct
	fmt.Println(condition.Bool(struct{}{})) // false
	fmt.Println(condition.Bool(time.Now())) // true
}
```



### <span id="And">And</span>
<p>Returns true if both a and b are truthy.</p>

<b>Signature:</b>

```go
func And[T, U any](a T, b U) bool
```
<b>Example:</b>

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
<p>Returns false iff neither a nor b is truthy.</p>

<b>Signature:</b>

```go
func Or[T, U any](a T, b U) bool
```
<b>Example:</b>

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
<p>Returns true iff a or b but not both is truthy.</p>

<b>Signature:</b>

```go
func Xor[T, U any](a T, b U) bool
```
<b>Example:</b>

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
<p>Returns true iff neither a nor b is truthy.</p>

<b>Signature:</b>

```go
func Nor[T, U any](a T, b U) bool
```
<b>Example:</b>

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
	fmt.Println(condition.Nor(1, 1)) // true
}
```



### <span id="Nand">Nand</span>
<p>Returns false iff both a and b are truthy</p>

<b>Signature:</b>

```go
func Nand[T, U any](a T, b U) bool
```
<b>Example:</b>

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
<p>Checks the value of param `isTrue`, if true return ifValue else return elseValue</p>

<b>Signature:</b>

```go
func TernaryOperator[T, U any](isTrue T, ifValue U, elseValue U) U
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/condition"
)

func main() {
	trueValue := "1"
	falseValue := "0"

	fmt.Println(condition.TernaryOperator(true, trueValue, falseValue)) // "1"
}
```







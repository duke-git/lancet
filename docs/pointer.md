# Pointer

pointer contains some util functions to operate go pointer.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/pointer/pointer.go](https://github.com/duke-git/lancet/blob/main/pointer/pointer.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/pointer"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [Of](#Of)
-   [Unwrap](#Unwrap)
-   [UnwarpOr](#UnwarpOr)
-   [UnwarpOrDefault](#UnwarpOrDefault)
-   [ExtractPointer](#ExtractPointer)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Of">Of</span>

<p>Returns a pointer to the pass value `v`.</p>

<b>Signature:</b>

```go
func Of[T any](v T) *T
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
    result1 := pointer.Of(123)
    result2 := pointer.Of("abc")

    fmt.Println(*result1)
    fmt.Println(*result2)

    // Output:
    // 123
    // abc
}
```


### <span id="Unwrap">Unwrap</span>

<p>Returns the value from the pointer.</p>

<b>Signature:</b>

```go
func Unwrap[T any](p *T) T
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
    a := 123
    b := "abc"

    result1 := pointer.Unwrap(&a)
    result2 := pointer.Unwrap(&b)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 123
    // abc
}
```


### <span id="UnwarpOr">UnwarpOr</span>

<p>Returns the value from the pointer or fallback if the pointer is nil.</p>

<b>Signature:</b>
```go
UnwarpOr[T any](p *T, fallback T) T
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
	a := 123
	b := "abc"

	var c *int
	var d *string

	result1 := pointer.UnwarpOr(&a, 456)
	result2 := pointer.UnwarpOr(&b, "abc")
	result3 := pointer.UnwarpOr(c, 456)
	result4 := pointer.UnwarpOr(d, "def")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 123
	// abc
	// 456
	// def
}
```


### <span id="UnwarpOrDefault">UnwarpOrDefault</span>

<p>Returns the value from the pointer or the default value if the pointer is nil.</p>

<b>Signature:</b>
```go
UnwarpOrDefault[T any](p *T) T
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
	a := 123
	b := "abc"

	var c *int
	var d *string

	result1 := pointer.UnwarpOrDefault(&a)
	result2 := pointer.UnwarpOrDefault(&b)
	result3 := pointer.UnwarpOrDefault(c)
	result4 := pointer.UnwarpOrDefault(d)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 123
	// abc
	// 0
	//
}
```


### <span id="ExtractPointer">ExtractPointer</span>

<p>Returns the underlying value by the given interface type</p>

<b>Signature:</b>

```go
func ExtractPointer(value any) any
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/pointer"
)

func main() {
    a := 1
    b := &a
    c := &b
    d := &c

    result := pointer.ExtractPointer(d)

    fmt.Println(result)

    // Output:
    // 1
}
```
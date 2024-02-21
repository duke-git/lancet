# Optional
Optional is a type that may or may not contain a non-nil value.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/datastructure/optional/optional.go](https://github.com/duke-git/lancet/blob/main/datastructure/optional/optional.go)


<div STYLE="page-break-after: always;"></div>

## Usage
```go
import (
    "github.com/duke-git/lancet/v2/datastructure/optional"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

- [Of](#Of)
- [OfNullable](#OfNullable)
- [Empty](#Empty)
- [IsPresent](#IsPresent)
- [IsEmpty](#IsEmpty)
- [IfPresent](#IfPresent)
- [IfPresentOrElse](#IfPresentOrElse)
- [Get](#Get)
- [OrElse](#OrElse)
- [OrElseGet](#OrElseGet)
- [OrElseThrow](#OrElseThrow)



<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Of">Of</span>
<p>Returns an Optional with a non-nil value.</p>

<b>Signature:</b>

```go
func Of[T any](value T) Optional[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    value := 42
    opt := optional.Of(value)

    fmt.Println(opt.Get())

    // Output:
    // 42
}
```

### <span id="OfNullable">OfNullable</span>
<p>Returns an Optional for a given value, which may be nil.</p>

<b>Signature:</b>

```go
func OfNullable[T any](value *T) Optional[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    var value *int = nil
    opt := optional.OfNullable(value)

    fmt.Println(opt.IsPresent())

    value = new(int)
    *value = 42
    opt = optional.OfNullable(value)

    fmt.Println(opt.IsPresent())


    // Output:
    // false
    // true
}
```


### <span id="Empty">Empty</span>
<p>Returns an empty Optional instance.</p>

<b>Signature:</b>

```go
func Empty[T any]() Optional[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optEmpty := OfNullable.Empty[int]()
    fmt.Println(optEmpty.IsEmpty())

    // Output:
    // true
}
```


### <span id="IsEmpty">IsEmpty</span>
<p>Checks if the Optional is empty.</p>

<b>Signature:</b>

```go
func (o Optional[T]) IsEmpty() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optEmpty := OfNullable.Empty[int]()
    fmt.Println(optEmpty.IsEmpty())

    // Output:
    // true
}
```


### <span id="IsPresent">IsPresent</span>
<p>Checks if there is a value present.</p>

<b>Signature:</b>

```go
func (o Optional[T]) IsPresent() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    var value *int = nil
    opt := optional.OfNullable(value)

    fmt.Println(opt.IsPresent())

    value = new(int)
    *value = 42
    opt = optional.OfNullable(value)

    fmt.Println(opt.IsPresent())


    // Output:
    // false
    // true
}
```


### <span id="IfPresent">IfPresent</span>
<p>Performs the given action with the value if a value is present.</p>

<b>Signature:</b>

```go
func (o Optional[T]) IfPresent(action func(value T))
```
<b>Example:</b>

```go
package main

import (
    "fmt"
     "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    called := false
    action := func(value int) { called = true }

    optEmpty := optional.Empty[int]()
    optEmpty.IfPresent(action)

    fmt.Println(called)

    called = false // Reset for next test
    optWithValue := optional.Of(42)
    optWithValue.IfPresent(action)

    fmt.Println(optWithValue.IsPresent())

    // Output:
    // false
    // true
}
```


### <span id="IfPresentOrElse">IfPresentOrElse</span>
<p>Performs the action with the value if present, otherwise performs the empty-based action.</p>

<b>Signature:</b>

```go
func (o Optional[T]) IfPresentOrElse(action func(value T), emptyAction func())
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    calledWithValue := false
    valueAction := func(value int) { calledWithValue = true }
    emptyAction := func() { t.Errorf("Empty action should not be called when value is present") }

    optWithValue := optional.Of(42)
    optWithValue.IfPresentOrElse(valueAction, emptyAction)

    fmt.Println(calledWithValue)

    calledWithEmpty := false
    valueAction = func(value int) { t.Errorf("Value action should not be called when value is not present") }
    emptyAction = func() { calledWithEmpty = true }

    optEmpty := optional.Empty[int]()
    optEmpty.IfPresentOrElse(valueAction, emptyAction)

    fmt.Println(calledWithEmpty)

    // Output:
    // true
    // true
}
```

### <span id="Get">Get</span>
<p>Returns the value if present, otherwise panics.</p>

<b>Signature:</b>

```go
func (o Optional[T]) Get() T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    value := 42
    opt := optional.Of(value)

    fmt.Println(opt.Get())

    // Output:
    // 42
}
```


### <span id="OrElse">OrElse</span>
<p>Returns the value if present, otherwise returns other.</p>

<b>Signature:</b>

```go
func (o Optional[T]) OrElse(other T) T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optEmpty := optional.Empty[int]()
    val := optEmpty.OrElse(100)
    fmt.Println(val)

    optWithValue := optional.Of(42)
    val = optWithValue.OrElse(100)
    fmt.Println(val)

    // Output:
    // 100
    // 42
}
```


### <span id="OrElseGet">OrElseGet</span>
<p>Returns the value if present, otherwise invokes supplier and returns the result.</p>

<b>Signature:</b>

```go
func (o Optional[T]) OrElseGet(supplier func() T) T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optEmpty := optional.Empty[int]()
    supplier := func() int { return 100 }

    val := optEmpty.OrElseGet(supplier)
    fmt.Println(val)

    // Output:
    // 100
}
```


### <span id="OrElseThrow">OrElseThrow</span>
<p>Returns the value if present, otherwise returns an error.</p>

<b>Signature:</b>

```go
func (o Optional[T]) OrElseThrow(errorSupplier func() error) (T, error)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optEmpty := optional.Empty[int]()
    _, err := optEmpty.OrElseThrow(func() error { return errors.New("no value") })
    
    fmt.Println(err.Error())

    optWithValue := optional.Of(42)
    val, err := optWithValue.OrElseThrow(func() error { return errors.New("no value") })

    fmt.Println(val)
    fmt.Println(err)

    // Output:
    // no value
    // 42
    // nil
}
```
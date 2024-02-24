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
- [FromNillable](#FromNillable)
- [Default](#Default)
- [IsNotNil](#IsNotNil)
- [IsNil](#IsNil)
- [IsNotNil](#IsNotNil)
- [IfNotNilOrElse](#IfNotNilOrElse)
- [Umwarp](#Umwarp)
- [OrElse](#OrElse)
- [OrElseGet](#OrElseGet)
- [OrElseTrigger](#OrElseTrigger)



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

### <span id="FromNillable">FromNillable</span>
<p>Returns an Optional for a given value, which may be nil.</p>

<b>Signature:</b>

```go
func FromNillable[T any](value *T) Optional[T]
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
    opt := optional.FromNillable(value)

    fmt.Println(opt.IsNotNil())

    value = new(int)
    *value = 42
    opt = optional.FromNillable(value)

    fmt.Println(opt.IsNotNil())


    // Output:
    // false
    // true
}
```


### <span id="Default">Default</span>
<p>Returns an default Optional instance.</p>

<b>Signature:</b>

```go
func Default[T any]() Optional[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optDefault := optional.Default[int]()
    fmt.Println(optDefault.IsNil())

    // Output:
    // true
}
```


### <span id="IsNil">IsNil</span>
<p>Checks if the Optional is nil.</p>

<b>Signature:</b>

```go
func (o Optional[T]) IsNil() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optDefault := optional.Default[int]()
    fmt.Println(optDefault.IsNil())

    // Output:
    // true
}
```


### <span id="IsNotNil">IsNotNil</span>
<p>Checks if there is a value not nil.</p>

<b>Signature:</b>

```go
func (o Optional[T]) IsNotNil() bool
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
    opt := optional.FromNillable(value)

    fmt.Println(opt.IsNotNil())

    value = new(int)
    *value = 42
    opt = optional.FromNillable(value)

    fmt.Println(opt.IsNotNil())


    // Output:
    // false
    // true
}
```


### <span id="IfNotNil">IfNotNil</span>
<p>Performs the given action with the value if a value is present.</p>

<b>Signature:</b>

```go
func (o Optional[T]) IfNotNil(action func(value T))
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

    optDefault := optional.Default[int]()
    optDefault.IfNotNil(action)

    fmt.Println(called)

    called = false // Reset for next test
    optWithValue := optional.Of(42)
    optWithValue.IfNotNil(action)

    fmt.Println(optWithValue.IsNotNil())

    // Output:
    // false
    // true
}
```


### <span id="IfNotNilOrElse">IfNotNilOrElse</span>
<p>Performs the action with the value if not nil, otherwise performs the fallback action.</p>

<b>Signature:</b>

```go
func (o Optional[T]) IfNotNilOrElse(action func(value T), fallbackAction func())
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
    optWithValue.IfNotNilOrElse(valueAction, emptyAction)

    fmt.Println(calledWithValue)

    calledWithEmpty := false
    valueAction = func(value int) { t.Errorf("Value action should not be called when value is not present") }
    emptyAction = func() { calledWithEmpty = true }

    optDefault := optional.Default[int]()
    optDefault.IfNotNilOrElse(valueAction, emptyAction)

    fmt.Println(calledWithEmpty)

    // Output:
    // true
    // true
}
```

### <span id="Unwrap">Unwrap</span>
<p>Returns the value if not nil, otherwise panics.</p>

<b>Signature:</b>

```go
func (o Optional[T]) Unwrap() T
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

    fmt.Println(opt.Unwrap())

    // Output:
    // 42
}
```


### <span id="OrElse">OrElse</span>
<p>Returns the value if not nill, otherwise returns other.</p>

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
    optDefault := optional.Default[int]()
    val := optDefault.OrElse(100)
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
<p>Returns the value if not nil, otherwise invokes action and returns the result.</p>

<b>Signature:</b>

```go
func (o Optional[T]) OrElseGet(action func() T) T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optDefault := optional.Default[int]()
    action := func() int { return 100 }

    val := optDefault.OrElseGet(action)
    fmt.Println(val)

    // Output:
    // 100
}
```


### <span id="OrElseTrigger">OrElseTrigger</span>
<p>Returns the value if present, otherwise returns an error.</p>

<b>Signature:</b>

```go
 OrElseTrigger(errorHandler func() error) (T, error)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optDefault := optional.Default[int]()
    _, err := optDefault.OrElseTrigger(func() error { return errors.New("no value") })
    
    fmt.Println(err.Error())

    optWithValue := optional.Of(42)
    val, err := optWithValue.OrElseTrigger(func() error { return errors.New("no value") })

    fmt.Println(val)
    fmt.Println(err)

    // Output:
    // no value
    // 42
    // nil
}
```
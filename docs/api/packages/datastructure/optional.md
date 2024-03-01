# Optional
Optional类型代表一个可选的值，它要么包含一个实际值，要么为空。

<div STYLE="page-break-after: always;"></div>

## 源码

- [https://github.com/duke-git/lancet/blob/main/datastructure/optional/optional.go](https://github.com/duke-git/lancet/blob/main/datastructure/optional/optional.go)


<div STYLE="page-break-after: always;"></div>

## 用法
```go
import (
    "github.com/duke-git/lancet/v2/datastructure/optional"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

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

## 文档

### <span id="Of">Of</span>
<p>返回一个包含非空值的Optional。</p>

<b>函数签名:</b>

```go
func Of[T any](value T) Optional[T]
```
<b>示例:</b>

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
<p>返回一个包含给定值的Optional，该值可能为空 (nil)。</p>

<b>函数签名:</b>

```go
func FromNillable[T any](value *T) Optional[T]
```
<b>示例:</b>

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
<p>返回一个空Optional实例。</p>

<b>函数签名:</b>

```go
func Default[T any]() Optional[T]
```
<b>示例:</b>

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
<p>验证Optional是否为空。</p>

<b>函数签名:</b>

```go
func (o Optional[T]) IsNil() bool
```
<b>示例:</b>

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
<p>检查当前Optional内是否存在值。</p>

<b>函数签名:</b>

```go
func (o Optional[T]) IsNotNil() bool
```
<b>示例:</b>

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
<p>如果值存在，则使用action方法执行给定的操作。</p>

<b>函数签名:</b>

```go
func (o Optional[T]) IfNotNil(action func(value T))
```
<b>示例:</b>

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
<p>根据是否存在值执行相应的操作：有值则执行指定操作，没有值则执行默认操作。</p>

<b>函数签名:</b>

```go
func (o Optional[T]) IfNotNilOrElse(action func(value T), fallbackAction func())
```
<b>示例:</b>

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
<p>如果存在，返回该值，否则引发panic。</p>

<b>函数签名:</b>

```go
func (o Optional[T]) Unwrap() T
```
<b>示例:</b>

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
<p>检查Optional值是否存在，如果存在，则直接返回该值。如果不存在，返回参数other值。</p>

<b>函数签名:</b>

```go
func (o Optional[T]) OrElse(other T) T
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/optional"
)

func main() {
    optDefault := optional.Empty[int]()
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
<p>检查Optional值是否存在，如果存在，则直接返回该值。如果不存在，则调用一个提供的函数 (supplier)，并返回该函数的执行结果。</p>

<b>函数签名:</b>

```go
func (o Optional[T]) OrElseGet(action func() T) T
```
<b>示例:</b>

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
<p>检查Optional值是否存在，如果存在，则直接返回该值，否则返回错误。</p>

<b>函数签名:</b>

```go
 OrElseTrigger(errorHandler func() error) (T, error)
```
<b>示例:</b>

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
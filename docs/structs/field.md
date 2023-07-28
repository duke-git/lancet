# Field

Field is abstract struct field for provide several high level functions

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/structs/field.go](https://github.com/duke-git/lancet/blob/main/structs/field.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/structs"
)
```

<div STYLE="page-break-after: always;"></div>

## Index:

-   [Tag](#Tag)
-   [Name](#Name)
-   [Value](#Value)
-   [Kind](#Kind)
-   [IsEmbedded](#IsEmbedded)
-   [IsExported](#IsExported)
-   [IsZero](#IsZero)
-   [IsSlice](#IsSlice)

> Note：Since `Field` inherits from `Struct`, it also has all the methods of 'Struct', as follows:

-   [ToMap](https://github.com/duke-git/lancet/blob/main/docs/structs/struct.md#ToMap)
-   [Fields](https://github.com/duke-git/lancet/blob/main/docs/structs/struct.md#Fields)
-   [Field](https://github.com/duke-git/lancet/blob/main/docs/structs/struct.md#Field)
-   [IsStruct](https://github.com/duke-git/lancet/blob/main/docs/structs/struct.md#IsStruct)

<div STYLE="page-break-after: always;"></div>

## Documentation:

### <span id="Tag">Tag</span>

<p>Get a `Tag` of the `Field`, `Tag` is a abstract struct field tag</p>

<b>Signature:</b>

```go
func (f *Field) Tag() *Tag
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type Parent struct {
        Name string `json:"name,omitempty"`
    }
    p1 := &Parent{"111"}

    s := structs.New(p1)
    n, _ := s.Field("Name")
    tag := n.Tag()

    fmt.Println(tag.Name)

    // Output:
    // name
}
```

### <span id="Value">Value</span>

<p>Get the `Field` underlying value</p>

<b>Signature:</b>

```go
func (f *Field) Value() any
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type Parent struct {
        Name string `json:"name,omitempty"`
    }
    p1 := &Parent{"111"}

    s := structs.New(p1)
    n, _ := s.Field("Name")

    fmt.Println(n.Value())

    // Output:
    // 111
}
```

### <span id="IsEmbedded">IsEmbedded</span>

<p>Check if the field is an embedded field</p>

<b>Signature:</b>

```go
func (f *Field) IsEmbedded() bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type Parent struct {
        Name string
    }
    type Child struct {
        Parent
        Age int
    }
    c1 := &Child{}
    c1.Name = "111"
    c1.Age = 11

    s := structs.New(c1)
    n, _ := s.Field("Name")
    a, _ := s.Field("Age")

    fmt.Println(n.IsEmbedded())
    fmt.Println(a.IsEmbedded())

    // Output:
    // true
    // false
}
```

### <span id="IsExported">IsExported</span>

<p>Check if the field is exported</p>

<b>Signature:</b>

```go
func (f *Field) IsExported() bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type Parent struct {
        Name string
        age  int
    }
    p1 := &Parent{Name: "11", age: 11}
    s := structs.New(p1)
    n, _ := s.Field("Name")
    a, _ := s.Field("age")

    fmt.Println(n.IsExported())
    fmt.Println(a.IsExported())

    // Output:
    // true
    // false
}
```

### <span id="IsZero">IsZero</span>

<p>Check if the field is zero value</p>

<b>Signature:</b>

```go
func (f *Field) IsZero() bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type Parent struct {
        Name string
        Age  int
    }
    p1 := &Parent{Age: 11}
    s := structs.New(p1)
    n, _ := s.Field("Name")
    a, _ := s.Field("Age")

    fmt.Println(n.IsZero())
    fmt.Println(a.IsZero())

    // Output:
    // true
    // false
}
```

### <span id="Name">Name</span>

<p>Get the field name</p>

<b>Signature:</b>

```go
func (f *Field) Name() string
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type Parent struct {
        Name string
        Age  int
    }
    p1 := &Parent{Age: 11}
    s := structs.New(p1)
    n, _ := s.Field("Name")
    a, _ := s.Field("Age")

    fmt.Println(n.Name())
    fmt.Println(a.Name())

    // Output:
    // Name
    // Age
}
```

### <span id="Kind">Kind</span>

<p>Get the field's kind</p>

<b>Signature:</b>

```go
func (f *Field) Kind() reflect.Kind
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type Parent struct {
        Name string
        Age  int
    }
    p1 := &Parent{Age: 11}
    s := structs.New(p1)
    n, _ := s.Field("Name")
    a, _ := s.Field("Age")

    fmt.Println(n.Kind())
    fmt.Println(a.Kind())

    // Output:
    // string
    // int
}
```

### <span id="IsSlice">IsSlice</span>

<p>Check if the field is a slice</p>

<b>Signature:</b>

```go
func (f *Field) IsSlice() bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type Parent struct {
        Name string
        arr  []int
    }

    p1 := &Parent{arr: []int{1, 2, 3}}
    s := structs.New(p1)
    a, _ := s.Field("arr")

    fmt.Println(a.IsSlice())

    // Output:
    // true
}
```

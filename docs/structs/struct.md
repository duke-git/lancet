# Structs

Struct is abstract struct for provide several high level functions

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/structs/struct.go](https://github.com/duke-git/lancet/blob/main/structs/struct.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/duke-git/lancet/v2/structs"
)
```

<div STYLE="page-break-after: always;"></div>

## Index:

-   [New](#New)
-   [ToMap](#ToMap)
-   [Fields](#Fields)
-   [Field](#Field)
-   [IsStruct](#IsStruct)

<div STYLE="page-break-after: always;"></div>

## Documentation:

### <span id="New">New</span>

<p>The constructor function of the `Struct` </p>

<b>Signature:</b>

```go
func New(value any, tagName ...string) *Struct
```

<b>Example:</b>

```go
package main

import (
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type People struct {
        Name string `json:"name"`
    }
    p1 := &People{Name: "11"}
    s := structs.New(p1)
    // to do something
}
```

### <span id="ToMap">ToMap</span>

<p>convert a valid struct to a map</p>

<b>Signature:</b>

```go
func (s *Struct) ToMap() (map[string]any, error)
```

> In addition, provided a convenient static function ToMap

```go
func ToMap(v any) (map[string]any, error)
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type People struct {
        Name string `json:"name"`
    }
    p1 := &People{Name: "11"}
    // use constructor function
    s1 := structs.New(p1)
    m1, _ := s1.ToMap()

    fmt.Println(m1)

    // use static function
    m2, _ := structs.ToMap(p1)

    fmt.Println(m2)

    // Output:
    // map[name:11]
    // map[name:11]
}
```

### <span id="Fields">Fields</span>

<p>Get all fields of a given struct, that the fields are abstract struct field</p>

<b>Signature:</b>

```go
func (s *Struct) Fields() []*Field
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type People struct {
        Name string `json:"name"`
    }
    p1 := &People{Name: "11"}
    s := structs.New(p1)
    fields := s.Fields()

    fmt.Println(len(fields))

    // Output:
    // 1
}
```

### <span id="Field">Field</span>

<p>Get an abstract field of a struct by given field name </p>

<b>Signature:</b>

```go
func (s *Struct) Field(name string) *Field
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type People struct {
        Name string `json:"name"`
    }
    p1 := &People{Name: "11"}
    s := structs.New(p1)
    f := s.Field("Name")

    fmt.Println(f.Value())

    // Output:
    // 11
}
```

### <span id="IsStruct">IsStruct</span>

<p>Check if the struct is valid</p>

<b>Signature:</b>

```go
func (s *Struct) IsStruct() bool
```

<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type People struct {
        Name string `json:"name"`
    }
    p1 := &People{Name: "11"}
    s := structs.New(p1)

    fmt.Println(s.IsStruct())

    // Output:
    // true
}
```

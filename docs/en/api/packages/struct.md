# Structs

Struct is abstract struct for provide several high level functions

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/structs/struct.go](https://github.com/duke-git/lancet/blob/main/structs/struct.go)

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

-   [New](#New)
-   [ToMap](#ToMap)
-   [Fields](#Fields)
-   [Field](#Field)
-   [IsStruct](#IsStruct)
-   [Tag](#Tag)
-   [Name](#Name)
-   [TypeName](#TypeName)
-   [Value](#Value)
-   [Kind](#Kind)
-   [IsEmbedded](#IsEmbedded)
-   [IsExported](#IsExported)
-   [IsZero](#IsZero)
-   [IsSlice](#IsSlice)
-   [IsTargetType](#IsTargetType)

<div STYLE="page-break-after: always;"></div>

## Documentation:

### <span id="New">New</span>

<p>The constructor function of the `Struct` </p>

<b>Signature:</b>

```go
func New(value any, tagName ...string) *Struct
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/O29l8kk-Z17)</span></b>

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
    
    fmt.Println(s.ToMap())

    // Output:
    // map[name:11] <nil>
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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/qQbLySBgerZ)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/w3Kk_CyVY7D)</span></b>

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
func (s *Struct) Field(name string) (*Field, bool)
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/KocZMSYarza)</span></b>

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
    f, found := s.Field("Name")

    fmt.Println(f.Value())
    fmt.Println(found)

    // Output:
    // 11
    // true
}
```

### <span id="IsStruct">IsStruct</span>

<p>Check if the struct is valid</p>

<b>Signature:</b>

```go
func (s *Struct) IsStruct() bool
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/bU2FSdkbK1C)</span></b>

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

### <span id="Tag">Tag</span>

<p>Get a `Tag` of the `Field`, `Tag` is a abstract struct field tag</p>

<b>Signature:</b>

```go
func (f *Field) Tag() *Tag
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/DVrx5HvvUJr)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/qufYEU2o4Oi)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/wV2PrbYm3Ec)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/csK4AXYaNbJ)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/RzqpGISf87r)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/zfIGlqsatee)</span></b>

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

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/wg4NlcUNG5o)</span></b>

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

### <span id="TypeName">TypeName</span>

<p>Return struct type name.</p>

<b>Signature:</b>

```go
func (s *Struct) TypeName() string
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/todo)</span></b>

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

    p := &Parent{Age: 11}
    s := structs.New(p1)

    fmt.Println(s.TypeName())
    
    // Output: 
    // Parent
}
```

### <span id="IsSlice">IsSlice</span>

<p>Check if the field is a slice</p>

<b>Signature:</b>

```go
func (f *Field) IsSlice() bool
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/MKz4CgBIUrU)</span></b>

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

### <span id="IsTargetType">IsTargetType</span>

<p>check if a struct field type is target type or not</p>

<b>Signature:</b>

```go
func (f *Field) IsTargetType(targetType reflect.Kind) bool
```

<b>Example:<span style="float:right;display:inline-block;">[Run](https://go.dev/play/p/Ig75P-agN39)</span></b>

```go
package main

import (
    "fmt"
    "reflect"
    "github.com/duke-git/lancet/v2/structs"
)

func main() {
    type Parent struct {
        Name string
        arr  []int
    }

    p1 := &Parent{arr: []int{1, 2, 3}}
    s := structs.New(p1)
    n, _ := s.Field("Name")
	a, _ := s.Field("arr")
    
    fmt.Println(n.IsTargetType(reflect.String))
    fmt.Println(a.IsTargetType(reflect.Slice))
    
    // Output: 
    // true
    // true
}
```
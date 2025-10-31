# Structs

structs 包封装了一个抽象的`Struct`结构体，提供了操作`struct`的相关函数

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/structs/struct.go](https://github.com/duke-git/lancet/blob/main/structs/struct.go)

-   [https://github.com/duke-git/lancet/blob/main/structs/field.go](https://github.com/duke-git/lancet/blob/main/structs/field.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/structs"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录:

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

## API 文档:

### <span id="New">New</span>

<p>`Struct`结构体的构造函数</p>

<b>函数签名:</b>

```go
func New(value any, tagName ...string) *Struct
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/O29l8kk-Z17)</span></b>

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
    
    fmt.Println(s.ToMap())

    // Output:
    // map[name:11] <nil>
}
```

### <span id="ToMap">ToMap</span>

<p>将一个合法的struct对象转换为map[string]any</p>

<b>函数签名:</b>

```go
func (s *Struct) ToMap() (map[string]any, error)
```

<b>除此之外，提供一个便捷的静态方法 ToMap</b>

```go
func ToMap(v any) (map[string]any, error)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/qQbLySBgerZ)</span></b>

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
    s1 := structs.New(p1)
    m1, _ := s1.ToMap()

    fmt.Println(m1)

    // 如果不需要Struct更多的方法，可以直接使用ToMap
    m2, _ := structs.ToMap(p1)

    fmt.Println(m2)

    // Output:
    // map[name:11]
    // map[name:11]
}
```

### <span id="Fields">Fields</span>

<p>获取一个struct对象的属性列表</p>

<b>函数签名:</b>

```go
func (s *Struct) Fields() []*Field
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/w3Kk_CyVY7D)</span></b>

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

<p>根据属性名获取一个struct对象的属性</p>

<b>函数签名:</b>

```go
func (s *Struct) Field(name string) (*Field, bool)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/KocZMSYarza)</span></b>

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

<p>判断是否为一个合法的struct对象</p>

<b>函数签名:</b>

```go
func (s *Struct) IsStruct() bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/bU2FSdkbK1C)</span></b>

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

<p>获取`Field`的`Tag`，默认的tag key是json</p>

<b>函数签名:</b>

```go
func (f *Field) Tag() *Tag
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/DVrx5HvvUJr)</span></b>

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

<p>获取`Field`属性的值</p>

<b>函数签名:</b>

```go
func (f *Field) Value() any
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/qufYEU2o4Oi)</span></b>

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

<p>判断属性是否为嵌入</p>

<b>函数签名:</b>

```go
func (f *Field) IsEmbedded() bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/wV2PrbYm3Ec)</span></b>

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

<p>判断属性是否导出</p>

<b>函数签名:</b>

```go
func (f *Field) IsExported() bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/csK4AXYaNbJ)</span></b>

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

<p>判断属性是否为零值</p>

<b>函数签名:</b>

```go
func (f *Field) IsZero() bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/RzqpGISf87r)</span></b>

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

<p>获取属性名</p>

<b>函数签名:</b>

```go
func (f *Field) Name() string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/zfIGlqsatee)</span></b>

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

<p>获取属性Kind</p>

<b>函数签名:</b>

```go
func (f *Field) Kind() reflect.Kind
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/wg4NlcUNG5o)</span></b>

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

<p>获取结构体类型名。</p>

<b>函数签名:</b>

```go
func (s *Struct) TypeName() string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/todo)</span></b>

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

<p>判断属性是否是切片</p>

<b>函数签名:</b>

```go
func (f *Field) IsSlice() bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/MKz4CgBIUrU)</span></b>

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

<p>判断属性是否是目标类型</p>

<b>函数签名:</b>

```go
func (f *Field) IsTargetType(targetType reflect.Kind) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Ig75P-agN39)</span></b>

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
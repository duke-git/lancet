# Field

Field 包封装了一个抽象的`Field`结构体，提供了操作`struct`属性的相关函数

<div STYLE="page-break-after: always;"></div>

## 源码:

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
-   [Tag](#Tag)
-   [Name](#Name)
-   [Value](#Value)
-   [Kind](#Kind)
-   [IsEmbedded](#IsEmbedded)
-   [IsExported](#IsExported)
-   [IsZero](#IsZero)
-   [IsSlice](#IsSlice)

> 注意：由于`Field`继承于`Struct`，所以同样拥有`Struct`所有方法，如下：

-   [ToMap](https://github.com/duke-git/lancet/blob/main/docs/structs/struct_zh-CN.md#ToMap)
-   [Fields](https://github.com/duke-git/lancet/blob/main/docs/structs/struct_zh-CN#Fields)
-   [Field](https://github.com/duke-git/lancet/blob/main/docs/structs/struct_zh-CN#Field)
-   [IsStruct](https://github.com/duke-git/lancet/blob/main/docs/structs/struct_zh-CN#IsStruct)

<div STYLE="page-break-after: always;"></div>

## API 文档:

### <span id="Tag">Tag</span>

<p>获取`Field`的`Tag`，默认的tag key是json</p>

<b>函数签名:</b>

```go
func (f *Field) Tag() *Tag
```

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<b>示例:</b>

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

<p>判断属性是否是切片</p>

<b>函数签名:</b>

```go
func (f *Field) IsSlice() bool
```

<b>示例:</b>

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
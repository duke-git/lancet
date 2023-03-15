# Structs

structs 包封装了一个抽象的`Struct`结构体，提供了操作`struct`的相关函数

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/structs/struct.go](https://github.com/duke-git/lancet/blob/main/structs/struct.go)

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

<div STYLE="page-break-after: always;"></div>

## API 文档:

### <span id="New">New</span>

<p>`Struct`结构体的构造函数</p>

<b>函数签名:</b>

```go
func New(value any, tagName ...string) *Struct
```

<b>示例:</b>

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

<p>将一个合法的struct对象转换为map[string]any</p>

<b>函数签名:</b>

```go
func (s *Struct) ToMap() (map[string]any, error)
```

<b>除此之外，提供一个便捷的静态方法ToMap</b>

```go
func ToMap(v any) (map[string]any, error)
```

<b>示例:</b>

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

<b>示例:</b>

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
func (s *Struct) Field(name string) *Field
```

<b>示例:</b>

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

<p>判断是否为一个合法的struct对象</p>

<b>函数签名:</b>

```go
func (s *Struct) IsStruct() bool
```

<b>示例:</b>

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

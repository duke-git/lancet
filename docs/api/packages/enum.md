# Enum

Enum 实现一个简单枚举工具包。

<div STYLE="page-break-after: always;"></div>

## 源码:

-   [https://github.com/duke-git/lancet/blob/main/enum/enum.go](https://github.com/duke-git/lancet/blob/main/enum/enum.go)

<div STYLE="page-break-after: always;"></div>

## 用法:

```go
import (
    "github.com/duke-git/lancet/v2/enum"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [NewItem](#NewItem)
-   [NewItemsFromPairs](#NewItemsFromPairs)
-   [Value](#Value)
-   [Name](#Name)
-   [Valid](#Valid)
-   [MarshalJSON](#MarshalJSON)
-   [NewRegistry](#NewRegistry)
-   [Add](#Add)
-   [Remove](#Remove)
-   [Update](#Update)
-   [GetByValue](#GetByValue)
-   [GetByName](#GetByName)
-   [Items](#Items)
-   [Contains](#Contains)
-   [Size](#Size)
-   [Range](#Range)
-   [SortedItems](#SortedItems)
-   [Filter](#Filter)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="NewItem">NewItem</span>

<p>创建枚举项。</p>

<b>函数签名:</b>

```go
func NewItem[T comparable](value T, name string) *Item[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/8qNsLw01HD5)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    item1 := enum.NewItem(Active, "Active")
    item2 := enum.NewItem(Inactive, "Inactive")

    fmt.Println(item1.Name(), item1.Value())
    fmt.Println(item2.Name(), item2.Value())

    // Output:
    // Active 1
    // Inactive 2
}
```

### <span id="NewItemsFromPairs">NewItemsFromPairs</span>

<p>从Pair结构体的切片创建枚举项。</p>

<b>函数签名:</b>

```go
func NewItemsFromPairs[T comparable](pairs ...Pair[T]) []*Item[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/xKnoGa7gnev)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    items := enum.NewItemsFromPairs(
        enum.Pair[Status]{Value: Active, Name: "Active"},
        enum.Pair[Status]{Value: Inactive, Name: "Inactive"},
    )

    fmt.Println(items[0].Name(), items[0].Value())
    fmt.Println(items[1].Name(), items[1].Value())

    // Output:
    // Active 1
    // Inactive 2
}
```

### <span id="Value">Value</span>

<p>返回枚举项的值。</p>

<b>函数签名:</b>

```go
func (it *Item[T]) Value() T
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/xKnoGa7gnev)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    items := enum.NewItemsFromPairs(
        enum.Pair[Status]{Value: Active, Name: "Active"},
        enum.Pair[Status]{Value: Inactive, Name: "Inactive"},
    )

    fmt.Println(items[0].Name(), items[0].Value())
    fmt.Println(items[1].Name(), items[1].Value())

    // Output:
    // Active 1
    // Inactive 2
}
```

### <span id="Name">Name</span>

<p>返回枚举项的名称。</p>

<b>函数签名:</b>

```go
func (it *Item[T]) Name() string
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/xKnoGa7gnev)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    items := enum.NewItemsFromPairs(
        enum.Pair[Status]{Value: Active, Name: "Active"},
        enum.Pair[Status]{Value: Inactive, Name: "Inactive"},
    )

    fmt.Println(items[0].Name(), items[0].Value())
    fmt.Println(items[1].Name(), items[1].Value())

    // Output:
    // Active 1
    // Inactive 2
}
```

### <span id="Valid">Valid</span>

<p>检查枚举项是否有效。如果提供了自定义检查函数，将使用该函数验证值。</p>

<b>函数签名:</b>

```go
func (it *Item[T]) Valid(checker ...func(T) bool) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/pA3lYY2VSm3)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    item := enum.NewItem(Active, "Active")
    fmt.Println(item.Valid())

    invalidItem := enum.NewItem(Unknown, "")
    fmt.Println(invalidItem.Valid())

    // Output:
    // true
    // false
}
```

### <span id="MarshalJSON">MarshalJSON</span>

<p>枚举项实现json.Marshaler接口。</p>

<b>函数签名:</b>

```go
func (it *Item[T]) MarshalJSON() ([]byte, error)
func (it *Item[T]) UnmarshalJSON(data []byte) error
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/zIZEdAnneB5)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    item := enum.NewItem(Active, "Active")
    data, _ := item.MarshalJSON()
    fmt.Println(string(data))

    var unmarshaledItem enum.Item[Status]
    _ = unmarshaledItem.UnmarshalJSON(data)
    fmt.Println(unmarshaledItem.Name(), unmarshaledItem.Value())

    // Output:
    // {"name":"Active","value":1}
    // Active 1
}
```

### <span id="NewRegistry">NewRegistry</span>

<p>Registry 定义了一个通用的枚举注册表结构体。</p>

<b>函数签名:</b>

```go
func NewRegistry[T comparable](items ...*Item[T]) *Registry[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ABEXsYfJKMo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item1 := enum.NewItem(Active, "Active")
    item2 := enum.NewItem(Inactive, "Inactive")

    registry.Add(item1, item2)

    if item, found := registry.GetByValue(Active); found {
        fmt.Println("Found by value:", item.Name())
    }

    if item, found := registry.GetByName("Inactive"); found {
        fmt.Println("Found by name:", item.Value())
    }

    // Output:
    // Found by value: Active
    // Found by name: 2
}
```

### <span id="Add">Add</span>

<p>向枚举注册表添加枚举项。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) Add(items ...*Item[T])
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/ABEXsYfJKMo)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item1 := enum.NewItem(Active, "Active")
    item2 := enum.NewItem(Inactive, "Inactive")

    registry.Add(item1, item2)

    if item, found := registry.GetByValue(Active); found {
        fmt.Println("Found by value:", item.Name())
    }

    if item, found := registry.GetByName("Inactive"); found {
        fmt.Println("Found by name:", item.Value())
    }

    // Output:
    // Found by value: Active
    // Found by name: 2
}
```

### <span id="Remove">Remove</span>

<p>在枚举注册表中删除枚举项。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) Remove(value T) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/dSG84wQ3TuC)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item1 := enum.NewItem(Active, "Active")

    registry.Add(item1)
    fmt.Println("Size before removal:", registry.Size())

    removed := registry.Remove(Active)
    fmt.Println("Removed:", removed)
    fmt.Println("Size after removal:", registry.Size())

    // Output:
    // Size before removal: 1
    // Removed: true
    // Size after removal: 0
}
```

### <span id="Update">Update</span>

<p>在枚举注册表中更新枚举项。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) Update(value T, newName string) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/Ol0moT1J9Xl)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item1 := enum.NewItem(Active, "Active")

    registry.Add(item1)
    updated := registry.Update(Active, "Activated")
    fmt.Println("Updated:", updated)

    if item, found := registry.GetByValue(Active); found {
        fmt.Println("New name:", item.Name())
    }

    // Output:
    // Updated: true
    // New name: Activated
}
```

### <span id="GetByValue">GetByValue</span>

<p>在枚举注册表中通过值获取枚举项。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) GetByValue(value T) (*Item[T], bool)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/niJ1U2KlE_m)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item := enum.NewItem(Active, "Active")

    registry.Add(item)

    if item, found := registry.GetByValue(Active); found {
        fmt.Println("Found name by value:", item.Name())
    }

    // Output:
    // Found name by value: Active
}
```

### <span id="GetByName">GetByName</span>

<p>在枚举注册表中通过名称获取枚举项。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) GetByName(name string) (*Item[T], bool)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/49ie_gpqH0m)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item := enum.NewItem(Active, "Active")

    registry.Add(item)

    if item, found := registry.GetByName("Active"); found {
        fmt.Println("Found value by name:", item.Value())
    }

    // Output:
    // Found value by name: 1
}
```

### <span id="Items">Items</span>

<p>返回枚举注册表中的枚举项。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) Items() []*Item[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/lAJFAradbvQ)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item1 := enum.NewItem(Active, "Active")
    item2 := enum.NewItem(Inactive, "Inactive")

    registry.Add(item1, item2)

    for _, item := range registry.Items() {
        fmt.Println(item.Name(), item.Value())
    }

    // Output:
    // Active 1
    // Inactive 2
}
```

### <span id="Contains">Contains</span>

<p>检查注册表中是否存在具有给定值的枚举项。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) Contains(value T) bool
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/_T-lPYkZn2j)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item := enum.NewItem(Active, "Active")
    registry.Add(item)

    fmt.Println(registry.Contains(Active))
    fmt.Println(registry.Contains(Inactive))

    // Output:
    // true
    // false
}
```

### <span id="Size">Size</span>

<p>返回注册表中枚举项的数目。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) Size() int
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/TeDArWhlQe2)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    fmt.Println("Initial size:", registry.Size())

    item1 := enum.NewItem(Active, "Active")
    item2 := enum.NewItem(Inactive, "Inactive")
    registry.Add(item1, item2)

    fmt.Println("Size after adding items:", registry.Size())

    registry.Remove(Active)
    fmt.Println("Size after removing an item:", registry.Size())

    // Output:
    // Initial size: 0
    // Size after adding items: 2
    // Size after removing an item: 1
}
```

### <span id="Range">Range</span>

<p>遍历注册表中的所有枚举项，并应用给定的函数。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) Range(fn func(*Item[T]) bool)
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/GPsZbQbefWN)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item1 := enum.NewItem(Active, "Active")
    item2 := enum.NewItem(Inactive, "Inactive")
    registry.Add(item1, item2)

    registry.Range(func(item *enum.Item[Status]) bool {
        fmt.Println(item.Name(), item.Value())
        return true // continue iteration
    })

    // Output:
    // Active 1
    // Inactive 2
}
```

### <span id="SortedItems">SortedItems</span>

<p>返回按给定比较函数排序的所有枚举项的切片。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) SortedItems(less func(*Item[T], *Item[T]) bool) []*Item[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/tN9RE_m_WEI)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item1 := enum.NewItem(Inactive, "Inactive")
    item2 := enum.NewItem(Active, "Active")
    registry.Add(item1, item2)

    for _, item := range registry.SortedItems(func(i1, i2 *enum.Item[Status]) bool {
        return i1.Value() < i2.Value()
    }) {
        fmt.Println(item.Name(), item.Value())
    }

    // Output:
    // Active 1
    // Inactive 2
}
```

### <span id="Filter">Filter</span>

<p>返回满足给定谓词函数的枚举项切片。</p>

<b>函数签名:</b>

```go
func (r *Registry[T]) Filter(predicate func(*Item[T]) bool) []*Item[T]
```

<b>示例:<span style="float:right;display:inline-block;">[运行](https://go.dev/play/p/uTUpTdcyoCU)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/enum"
)

type Status int

const (
    Unknown Status = iota
    Active
    Inactive
)

func main() {
    registry := enum.NewRegistry[Status]()
    item1 := enum.NewItem(Active, "Active")
    item2 := enum.NewItem(Inactive, "Inactive")
    registry.Add(item1, item2)

    activeItems := registry.Filter(func(item *enum.Item[Status]) bool {
        return item.Value() == Active
    })

    for _, item := range activeItems {
        fmt.Println(item.Name(), item.Value())
    }

    // Output:
    // Active 1
}
```

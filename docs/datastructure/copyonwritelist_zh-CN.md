# CopyOnWriteList
CopyOnWriteList 是一个线程安全的List实现，底层使用go 切片
.在写入时，会复制一份新的切片，写入完成后，再将新的切片赋值给原来的切片.
在读取时，直接读取原来的切片

## 源码
- [https://github.com/duke-git/lancet/blob/main/datastructure/list/copyonwritelist.go](https://github.com/duke-git/lancet/blob/main/datastructure/list/copyonwritelist.go)

## 用法
```go
import (
"github.com/duke-git/lancet/datastructure/list"
)

```

<div STYLE="page-break-after: always;"></div>

## 目录
- [NewCopyOnWriteList](#NewCopyOnWriteList)
- [Size](#Size)
- [Get](#Get)
- [Set](#Set)
- [Remove](#Remove)
- [IndexOf](#IndexOf)
- [LastIndexOf](#LastIndexOf)
- [IsEmpty](#IsEmpty)
- [Contain](#Contain)
- [ValueOf](#ValueOf)
- [Add](#Add)
- [AddAll](#AddAll)
- [AddByIndex](#AddByIndex)
- [DeleteAt](#DeleteAt)
- [DeleteIf](#DeleteIf)
- [DeleteBy](#DeleteBy)
- [DeleteRange](#DeleteRange)
- [Equal](#Equal)


## 文档

### NewCopyOnWriteList
返回一个具有空切片的CopyOnWriteList
```go
type CopyOnWriteList[T any] struct {
    data []T
    lock sync.Locker
}

func NewCopyOnWriteList() *CopyOnWriteList

```
#### 示例
```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main()  {
	l := list.NewCopyOnWriteList([]int{1,2,3})
    fmt.Println(l)
}

```
### Size 
返回CopyOnWriteList的长度
```go
func (l *CopyOnWriteList[T]) Size() int
```
#### 示例
```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    fmt.Println(l.Size())
}

```

### Get
返回列表中指定位置的元素
```go
func (c *CopyOnWriteList[T]) Get(index int) *T
```

#### 示例
```go
package main

import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)

func main()  {
	l := list.NewCopyOnWriteList([]int{1,2,3})
	fmt.Println(l.Get(2))
}

```


### Set
将此列表中指定位置的元素替换为指定元素。
```go
func (c *CopyOnWriteList[T]) Set(index int, e T) (oldValue *T, ok bool)
```

#### 示例
```go
package main

import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)

func main()  {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    fmt.Println(l.Set(2, 4))
}

```
### Remove

### IndexOf
返回列表中值的索引，如果没有找到返回-1
```go
func (c *CopyOnWriteList[T]) IndexOf(e T) int
```
#### 示例
```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    fmt.Println(l.IndexOf(1))
}

```

### LastIndexOf
返回指定元素在此列表中最后出现的索引，如果此列表不包含该元素，则返回-1

```go
func (c *CopyOnWriteList[T]) LastIndexOf(e T) int
```

#### 示例
```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main()  {
	l := list.NewCopyOnWriteList([]int{1,2,3,1})
	fmt.Println(l.LastIndexOf(1))
}

```

### IsEmpty
如果此列表不包含任何元素，则返回true。
```go
func (c *CopyOnWriteList[T]) IsEmpty() bool
```

#### 示例
```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main()  {
	l := list.NewCopyOnWriteList([]int{})
	fmt.Println(l.IsEmpty())
}
```



### Contain
判断CopyOnWriteList是否包含某个元素
```go
func (c *CopyOnWriteList[T]) Contain(e T) bool 
```
#### 示例
```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    fmt.Println(l.Contain(1))
}
```


### ValueOf
返回列表中索引处的值指针
```go
func (c *CopyOnWriteList[T]) ValueOf(index int) []T
```

#### 示例
```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    fmt.Println(l.ValueOf(2))
}

```

### Add
将指定的元素追加到此列表的末尾。
```go
func (c *CopyOnWriteList[T]) Add(e T) bool
```

#### 示例
```go

package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main()  {
	l := list.NewCopyOnWriteList([]int{1,2,3})
	l.Add(4)
	fmt.Println(l.getList())
}

```

### AddAll
将指定集合中的所有元素追加到此列表的末尾
```go
func (c *CopyOnWriteList[T]) AddAll(e []T) bool
```

#### 示例
```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main()  {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    l.AddAll([]int{4,5,6})
    fmt.Println(l.getList())
}

```

### AddByIndex
将指定元素插入此列表中的指定位置

```go
func (c *CopyOnWriteList[T]) AddByIndex(index int, e T) bool
```
#### 示例
```go
package main
import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)
func main()  {
	l := list.NewCopyOnWriteList([]int{1,2,3})
	list.AddByIndex(2, 6)
	fmt.Println(l.getList())
}

```

### DeleteAt
移除此列表中指定位置的元素。
```go
func (c *CopyOnWriteList[T]) DeleteAt(index int) (oldValue *T, ok bool)
```
#### 示例
```go
package main
import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)

func main()  {
	l := list.NewCopyOnWriteList([]int{1,2,3})
	list.DeleteAt(2)
	fmt.Println(l.getList())
}
```

### DeleteIf
从此列表中删除第一个出现的指定元素(如果该元素存在)。
```go
func (c *CopyOnWriteList[T]) DeleteIf(f func(T) bool) (oldValue *T, ok bool)
```
#### 示例
```go
package main
import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)

func main()  {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    list.DeleteIf(func(i int) bool {
        return i == 2
    })
    fmt.Println(l.getList())
}
```

### DeleteBy
从此列表中删除第一个出现的指定元素(如果该元素存在)。
```go
func (c *CopyOnWriteList[T]) DeleteBy(e T) (*T bool)
```
#### 示例
```go
package main
import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)

func main()  {
	l := list.NewCopyOnWriteList([]int{1,2,3})
    list.DeleteBy(2)
    fmt.Println(l.getList())
}
```


### DeleteRange
从该列表中删除索引介于fromIndex(包含)和toIndex(不包含)之间的所有元素。
(左闭右开)
```go
func (c *CopyOnWriteList[T]) DeleteRange(start int, end int) 
```
#### 示例
```go
package main
import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3,4,5,6,7,8,9})
    list.DeleteRange(2, 5)
    fmt.Println(l.getList())
}
```

### Equal
如果指定的对象等于此列表，则返回true

```go
func (c *CopyOnWriteList[T]) Equal(e []T) bool
```
#### 示例
```go
package main
import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3,4,5,6,7,8,9})
    fmt.Println(l.Equal([]int{1,2,3,4,5,6,7,8,9}))
}
```
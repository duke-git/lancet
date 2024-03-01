# CopyOnWriteList

CopyOnWriteList is a thread-safe list implementation that uses go slicing as its base. When writing, a new slice is copied and assigned to the original slice when writing is complete. When reading, the original slice is read directly.

## 源码

-   [https://github.com/duke-git/lancet/blob/main/datastructure/list/copyonwritelist.go](https://github.com/duke-git/lancet/blob/main /datastructure/list/copyonwritelist.go)

## 用法

```go
import (
    "github.com/duke-git/lancet/datastructure/list"
)

```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [NewCopyOnWriteList](#NewCopyOnWriteList)
-   [Size](#Size)
-   [Get](#Get)
-   [Set](#Set)
-   [Remove](#Remove)
-   [IndexOf](#IndexOf)
-   [LastIndexOf](#LastIndexOf)
-   [IndexOfFunc](#IndexOfFunc)
-   [LastIndexOfFunc](#LastIndexOfFunc)
-   [IsEmpty](#IsEmpty)
-   [Contain](#Contain)
-   [ValueOf](#ValueOf)
-   [Add](#Add)
-   [AddAll](#AddAll)
-   [AddByIndex](#AddByIndex)
-   [DeleteAt](#DeleteAt)
-   [DeleteIf](#DeleteIf)
-   [DeleteBy](#DeleteBy)
-   [DeleteRange](#DeleteRange)
-   [Equal](#Equal)

## Documentation

### NewCopyOnWriteList

Returns a CopyOnWriteList with empty slices.

```go
type CopyOnWriteList[T any] struct {
    data []T
    lock sync.
}

func NewCopyOnWriteList() *CopyOnWriteList

```

#### Example

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
	l := list.NewCopyOnWriteList([]int{1,2,3})
    fmt.Println(l)
}
```

### Size

Returns the length of the CopyOnWriteList.

```go
func (l *CopyOnWriteList[T]) Size() int
```

#### Example

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

Returns the element at the specified position in the list

```go
func (c *CopyOnWriteList[T]) Get(index int) *T
```

#### Example

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    fmt.Println(l.Get(2))
}

```

### Set

Replaces the element at the specified position in this list with the specified element.

```go
func (c *CopyOnWriteList[T]) Set(index int, e T) (oldValue *T, ok bool)
```

#### Example

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    fmt.Println(l.Set(2, 4))
}

```

### Remove

### IndexOf

Returns the index of the value in the list, or -1 if not found.

```go
func (c *CopyOnWriteList[T]) IndexOf(e T) int
```

#### Example

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

Returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain that element.

```go
func (c *CopyOnWriteList[T]) LastIndexOf(e T) int
```

#### Example

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3,1})
    fmt.Println(l.LastIndexOf(1))
}

```


### <span id="IndexOfFunc">IndexOfFunc</span>
<p> IndexOfFunc returns the first index satisfying the functional predicate f(v) bool. if not found return -1.</p>

<b>Signature:</b>

```go
func (l *CopyOnWriteList[T]) IndexOfFunc(f func(T) bool) int 
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1, 2, 3})

    fmt.Println(l.IndexOfFunc(func(a int) bool { return a == 1 })) //0
    fmt.Println(l.IndexOfFunc(func(a int) bool { return a == 0 })) //-1
}
```

### <span id="LastIndexOfFunc">LastIndexOfFunc</span>
<p>LastIndexOfFunc returns the index of the last occurrence of the value in this list satisfying the functional predicate f(T) bool. if not found return -1.</p>

<b>Signature:</b>

```go
func (l *CopyOnWriteList[T]) LastIndexOfFunc(f func(T) bool) int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1, 2, 3, 1})

    fmt.Println(l.LastIndexOfFunc(func(a int) bool { return a == 1 })) // 3
    fmt.Println(l.LastIndexOfFunc(func(a int) bool { return a == 0 })) //-1
}
```

### IsEmpty

Returns true if this list does not contain any elements.

```go
func (c *CopyOnWriteList[T]) IsEmpty() bool
```

#### Example

```go
package main

import (
"fmt"
"github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{})
    fmt.Println(l.IsEmpty())
}
```

### Contain

Determines if a CopyOnWriteList contains an element.

```go
func (c *CopyOnWriteList[T]) Contain(e T) bool
```

#### Example

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

Returns a pointer to the value at the index in the list

```go
func (c *CopyOnWriteList[T]) ValueOf(index int) []T
```

#### Example

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

Appends the specified element to the end of the list.

```go
func (c *CopyOnWriteList[T]) Add(e T) bool
```

#### Example

```go

package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
	l := list.NewCopyOnWriteList([]int{1,2,3})
	l.Add(4)
	fmt.Println(l.getList())
}

```

### AddAll

Appends all the elements of the specified collection to the end of this list

```go
func (c *CopyOnWriteList[T]) AddAll(e []T) bool
```

#### Example

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    l.AddAll([]int{4,5,6})
    fmt.Println(l.getList())
}

```

### AddByIndex

Inserts the specified element into the list at the specified position.

```go
func (c *CopyOnWriteList[T]) AddByIndex(index int, e T) bool
```

#### Example

```go
package main
import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)
func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    list.AddByIndex(2, 6)
    fmt.Println(l.getList())
}

```

### DeleteAt

Removes the element at the specified position in this list.

```go
func (c *CopyOnWriteList[T]) DeleteAt(index int) (oldValue *T, ok bool)
```

#### Example

```go
package main
import (
    "fmt"
    "github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    list.DeleteAt(2)
    fmt.Println(l.getList())
}
```

### DeleteIf

Removes the first occurrence of the specified element from this list (if it exists).

```go
func (c *CopyOnWriteList[T]) DeleteIf(func(T) bool) (oldValue *T, ok bool)
```

#### Example

```go
package main
import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)

func main() {
    l := list.NewCopyOnWriteList([]int{1,2,3})
    list.DeleteIf(func(i int) bool {
        return i == 2
    })
    fmt.Println(l.getList())
}
```

### DeleteBy

Deletes the first occurrence of the specified element from this list (if it exists).

```go
func (c *CopyOnWriteList[T]) DeleteBy(e T) (*T bool)
```

#### Example

```go
package main
import (
	"fmt"
	"github.com/duke-git/lancet/datastructure/list"
)

func main() {
	l := list.NewCopyOnWriteList([]int{1,2,3})
    list.DeleteBy(2)
    fmt.Println(l.getList())
}
```

### DeleteRange

Deletes all elements from this list with indexes between fromIndex (included) and toIndex (not included).
(leftCloseRightOpen)

```go
func (c *CopyOnWriteList[T]) DeleteRange(start int, end int)
```

#### Example

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

Returns true if the specified object is equal to this list

```go
func (c *CopyOnWriteList[T]) Equal(e []T) bool
```

#### Example

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

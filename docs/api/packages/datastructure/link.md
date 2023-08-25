# Linklist
Linklist是链表数据结构，它的节点有一个值和一个指向下一个节点的指针。

<div STYLE="page-break-after: always;"></div>

## 源码

- [https://github.com/duke-git/lancet/blob/main/datastructure/link/singlylink.go](https://github.com/duke-git/lancet/blob/main/datastructure/link/singlylink.go)
- [https://github.com/duke-git/lancet/blob/main/datastructure/link/doublylink.go](https://github.com/duke-git/lancet/blob/main/datastructure/link/doublylink.go)


<div STYLE="page-break-after: always;"></div>

## 用法
```go
import (
    link "github.com/duke-git/lancet/v2/datastructure/link"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

### 1. SinglyLink单链表

- [NewSinglyLink](#NewSinglyLink)
- [Values](#SinglyLink_Values)
- [InsertAt](#SinglyLink_InsertAt)
- [InsertAtHead](#SinglyLink_InsertAtHead)
- [InsertAtTail](#SinglyLink_InsertAtTail)
- [DeleteAt](#SinglyLink_DeleteAt)
- [DeleteAtHead](#SinglyLink_DeleteAtHead)
- [DeleteAtTail](#SinglyLink_DeleteAtTail)
- [DeleteValue](#SinglyLink_DeleteValue)
- [Reverse](#SinglyLink_Reverse)
- [GetMiddleNode](#SinglyLink_GetMiddleNode)
- [Size](#SinglyLink_Size)
- [IsEmpty](#SinglyLink_IsEmpty)
- [Clear](#SinglyLink_Clear)
- [Print](#SinglyLink_Print)

### 2. DoublyLink双向链表

- [NewDoublyLink](#NewDoublyLink)
- [Values](#DoublyLink_Values)
- [InsertAt](#DoublyLink_InsertAt)
- [InsertAtHead](#DoublyLink_InsertAtHead)
- [InsertAtTail](#DoublyLink_InsertAtTail)
- [DeleteAt](#DoublyLink_DeleteAt)
- [DeleteAtHead](#DoublyLink_DeleteAtHead)
- [DeleteAtTail](#DoublyLink_DeleteAtTail)
- [Reverse](#DoublyLink_Reverse)
- [GetMiddleNode](#DoublyLink_GetMiddleNode)
- [Size](#DoublyLink_Size)
- [IsEmpty](#DoublyLink_IsEmpty)
- [Clear](#DoublyLink_Clear)
- [Print](#DoublyLink_Print)


<div STYLE="page-break-after: always;"></div>

## 文档

### 1. SinglyLink
SingleLink是单向链表，它的节点有一个值和一个指向链表的下一个节点的指针。

### <span id="NewSinglyLink">NewSinglyLink</span>
<p>创建SinglyLink指针实例</p>

<b>函数签名:</b>

```go
type LinkNode[T any] struct {
	Value T
	Next  *LinkNode[T]
}
type SinglyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}
func NewSinglyLink[T any]() *SinglyLink[T]
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()
    fmt.Println(lk)
}
```



### <span id="SinglyLink_Values">Values</span>
<p>返回链表中所有节点值的切片</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) Values() []T
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```




### <span id="SinglyLink_InsertAt">InsertAt</span>
<p>将值插入到索引处的链表中，索引应大于或等于0且小于或等于链表节点数</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) InsertAt(index int, value T)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAt(1, 1) //do nothing

    lk.InsertAt(0, 1)
    lk.InsertAt(1, 2)
    lk.InsertAt(2, 3)
    lk.InsertAt(2, 4)

    fmt.Println(lk.Values()) //[]int{1, 2, 4, 3}
}
```




### <span id="SinglyLink_InsertAtHead">InsertAtHead</span>
<p>将值插入到链表表头</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) InsertAtHead(value T)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtHead(1)
    lk.InsertAtHead(2)
    lk.InsertAtHead(3)

    fmt.Println(lk.Values()) //[]int{3, 2, 1}
}
```




### <span id="SinglyLink_InsertAtTail">InsertAtTail</span>
<p>将值插入到链表末尾</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) InsertAtTail(value T)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```



### <span id="SinglyLink_DeleteAt">DeleteAt</span>
<p>删除特定索引处的值，索引应大于或等于0且小于或等于链接节点数-1</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) DeleteAt(index int)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    lk.InsertAtTail(4)

    lk.DeleteAt(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```



### <span id="SinglyLink_DeleteAtHead">DeleteAtHead</span>
<p>删除链表头节点</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) DeleteAtHead()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    lk.InsertAtTail(4)

    lk.DeleteAtHead()
    
    fmt.Println(lk.Values()) //[]int{2, 3, 4}
}
```




### <span id="SinglyLink_DeleteAtTail">DeleteAtTail</span>
<p>删除链表末尾节点</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) DeleteAtTail()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.DeleteAtTail()
    
    fmt.Println(lk.Values()) //[]int{1, 2}
}
```



### <span id="SinglyLink_DeleteValue">DeleteValue</span>
<p>删除链表中指定的value值</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) DeleteValue(value T)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.DeleteValue(2)
    fmt.Println(lk.Values()) //[]int{1, 3}
}
```




### <span id="SinglyLink_Reverse">Reverse</span>
<p>反转链表所有节点顺序</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) Reverse()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.Reverse()
    fmt.Println(lk.Values()) //[]int{3, 2, 1}
}
```



### <span id="SinglyLink_GetMiddleNode">GetMiddleNode</span>
<p>获取链表中部节点</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) GetMiddleNode() *datastructure.LinkNode[T]
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    midNode := lk.GetMiddleNode()
    fmt.Println(midNode.Value) //2
}
```



### <span id="SinglyLink_Size">Size</span>
<p>获取链表节点数量</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) Size() int
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Size()) //3
}
```



### <span id="SinglyLink_IsEmpty">IsEmpty</span>
<p>判断链表是否为空</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) IsEmpty() bool
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()
    fmt.Println(lk.IsEmpty()) //true

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.IsEmpty()) //false
}
```



### <span id="SinglyLink_Clear">Clear</span>
<p>清空链表所有节点</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) Clear()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.Clear()

    fmt.Println(lk.Values()) //
}
```



### <span id="SinglyLink_Print">Print</span>
<p>打印链表结构</p>

<b>函数签名:</b>

```go
func (link *SinglyLink[T]) Clear()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewSinglyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    
    lk.Print() //[ &{Value:1 Pre:<nil> Next:0xc0000a4048}, &{Value:2 Pre:<nil> Next:0xc0000a4060}, &{Value:3 Pre:<nil> Next:<nil>} ]
}
```



### 2. DoublyLink
DoublyLink是双向链表，它的节点有一个值，next指针指向下一个节点，pre指针指向前一个节点。

### <span id="NewDoublyLink">NewDoublyLink</span>
<p>创建NewDoublyLink指针实例</p>

<b>函数签名:</b>

```go
type LinkNode[T any] struct {
	Value T
    Pre   *LinkNode[T]
	Next  *LinkNode[T]
}
type DoublyLink[T any] struct {
	Head   *datastructure.LinkNode[T]
	length int
}
func NewDoublyLink[T any]() *DoublyLink[T]
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()
    fmt.Println(lk)
}
```



### <span id="DoublyLink_Values">Values</span>
<p>返回链表中所有节点值的切片</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) Values() []T
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```




### <span id="DoublyLink_InsertAt">InsertAt</span>
<p>将值插入到索引处的链表中，索引应大于或等于0且小于或等于链表节点数</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) InsertAt(index int, value T)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAt(1, 1) //do nothing

    lk.InsertAt(0, 1)
    lk.InsertAt(1, 2)
    lk.InsertAt(2, 3)
    lk.InsertAt(2, 4)

    fmt.Println(lk.Values()) //[]int{1, 2, 4, 3}
}
```




### <span id="DoublyLink_InsertAtHead">InsertAtHead</span>
<p>将值插入到链表表头</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) InsertAtHead(value T)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtHead(1)
    lk.InsertAtHead(2)
    lk.InsertAtHead(3)

    fmt.Println(lk.Values()) //[]int{3, 2, 1}
}
```




### <span id="DoublyLink_InsertAtTail">InsertAtTail</span>
<p>将值插入到链表末尾</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) InsertAtTail(value T)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```



### <span id="DoublyLink_DeleteAt">DeleteAt</span>
<p>删除特定索引处的值，索引应大于或等于0且小于或等于链接节点数-1</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) DeleteAt(index int)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    lk.InsertAtTail(4)

    lk.DeleteAt(3)

    fmt.Println(lk.Values()) //[]int{1, 2, 3}
}
```



### <span id="DoublyLink_DeleteAtHead">DeleteAtHead</span>
<p>删除链表头节点</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) DeleteAtHead()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    lk.InsertAtTail(4)

    lk.DeleteAtHead()
    
    fmt.Println(lk.Values()) //[]int{2, 3, 4}
}
```




### <span id="DoublyLink_DeleteAtTail">DeleteAtTail</span>
<p>删除链表末尾节点</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) DeleteAtTail()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.DeleteAtTail()
    
    fmt.Println(lk.Values()) //[]int{1, 2}
}
```




### <span id="DoublyLink_Reverse">Reverse</span>
<p>反转链表所有节点顺序</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) Reverse()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.Reverse()
    fmt.Println(lk.Values()) //[]int{3, 2, 1}
}
```



### <span id="DoublyLink_GetMiddleNode">GetMiddleNode</span>
<p>获取链表中部节点</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) GetMiddleNode() *datastructure.LinkNode[T]
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    midNode := lk.GetMiddleNode()
    fmt.Println(midNode.Value) //2
}
```



### <span id="DoublyLink_Size">Size</span>
<p>获取链表节点数量</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) Size() int
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.Size()) //3
}
```



### <span id="DoublyLink_IsEmpty">IsEmpty</span>
<p>判断链表是否为空</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) IsEmpty() bool
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()
    fmt.Println(lk.IsEmpty()) //true

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    fmt.Println(lk.IsEmpty()) //false
}
```



### <span id="DoublyLink_Clear">Clear</span>
<p>清空链表所有节点</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) Clear()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)

    lk.Clear()

    fmt.Println(lk.Values()) //
}
```



### <span id="DoublyLink_Print">Print</span>
<p>打印链表结构</p>

<b>函数签名:</b>

```go
func (link *DoublyLink[T]) Clear()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    link "github.com/duke-git/lancet/v2/datastructure/link"
)

func main() {
    lk := link.NewDoublyLink[int]()

    lk.InsertAtTail(1)
    lk.InsertAtTail(2)
    lk.InsertAtTail(3)
    
    lk.Print() //
}
```
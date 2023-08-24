# Stack
栈数据结构，包括ArrayStack（数组栈）和LinkedStack（链表栈）。

<div STYLE="page-break-after: always;"></div>

## 源码

- [https://github.com/duke-git/lancet/blob/main/datastructure/stack/arraystack.go](https://github.com/duke-git/lancet/blob/main/datastructure/stack/arraystack.go)
- [https://github.com/duke-git/lancet/blob/main/datastructure/stack/linkedstack.go](https://github.com/duke-git/lancet/blob/main/datastructure/stack/linkedstack.go)


<div STYLE="page-break-after: always;"></div>

## 用法
```go
import (
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

### 1. ArrayStack（数组栈）

- [NewArrayStack](#NewArrayStack)
- [Push](#ArrayStack_Push)
- [Pop](#ArrayStack_Pop)
- [Peak](#ArrayStack_Peak)
- [Data](#ArrayStack_Data)
- [Size](#ArrayStack_Size)
- [IsEmpty](#ArrayStack_IsEmpty)
- [Clear](#ArrayStack_Clear)

### 2. LinkedStack（链表栈）

- [NewLinkedStack](#NewLinkedStack)
- [Push](#LinkedStack_Push)
- [Pop](#LinkedStack_Pop)
- [Peak](#LinkedStack_Peak)
- [Data](#LinkedStack_Data)
- [Size](#LinkedStack_Size)
- [IsEmpty](#LinkedStack_IsEmpty)
- [Clear](#LinkedStack_Clear)
- [Print](#LinkedStack_Print)

<div STYLE="page-break-after: always;"></div>

## 文档

### 1. ArrayStack
用切片实现栈结构

### <span id="NewArrayStack">NewArrayStack</span>
<p>返回ArrayStack指针实例</p>

<b>函数签名:</b>

```go
type ArrayStack[T any] struct {
	data   []T
	length int
}
func NewArrayStack[T any]() *ArrayStack[T]
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewArrayStack[int]()
    fmt.Println(sk)
}
```




### <span id="ArrayStack_Push">Push</span>
<p>将元素加入数组栈</p>

<b>函数签名:</b>

```go
func (s *ArrayStack[T]) Push(value T)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewArrayStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    fmt.Println(sk.Data()) //[]int{3, 2, 1}
}
```




### <span id="ArrayStack_Pop">Pop</span>
<p>删除栈顶元素并返回该元素指针</p>

<b>函数签名:</b>

```go
func (s *ArrayStack[T]) Pop() (*T, error)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewArrayStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    val, err := sk.Pop()
    fmt.Println(err) //nil
    fmt.Println(*val) //3

    fmt.Println(sk.Data()) //[]int{2, 1}
}
```




### <span id="ArrayStack_Peak">Peak</span>
<p>返回栈顶元素指针</p>

<b>函数签名:</b>

```go
func (s *ArrayStack[T]) Peak() (*T, error)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewArrayStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    val, err := sk.Peak()
    fmt.Println(err) //nil
    fmt.Println(*val) //3

    fmt.Println(sk.Data()) //[]int{3, 2, 1}
}
```




### <span id="ArrayStack_Data">Data</span>
<p>返回栈中所有元素组成的切片</p>

<b>函数签名:</b>

```go
func (s *ArrayStack[T]) Data() []T
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewArrayStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    fmt.Println(sk.Data()) //[]int{3, 2, 1}
}
```




### <span id="ArrayStack_Size">Size</span>
<p>返回栈中元素的数量</p>

<b>函数签名:</b>

```go
func (s *ArrayStack[T]) Size() int
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewArrayStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    fmt.Println(sk.Size()) //3
}
```




### <span id="ArrayStack_IsEmpty">IsEmpty</span>
<p>判断栈是否为空</p>

<b>函数签名:</b>

```go
func (s *ArrayStack[T]) IsEmpty() bool
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewArrayStack[int]()
    fmt.Println(sk.IsEmpty()) //true

    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    fmt.Println(sk.IsEmpty()) //false
}
```




### <span id="ArrayStack_Clear">Clear</span>
<p>清空栈元素，使栈为空</p>

<b>函数签名:</b>

```go
func (s *ArrayStack[T]) Clear()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewArrayStack[int]()

    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    sk.Clear()

    fmt.Println(sk.Data()) //[]int{}
}
```



### 2. LinkedStack
链表实现的栈结构。

### <span id="NewLinkedStack">NewLinkedStack</span>
<p>返回LinkedStack指针实例</p>

<b>函数签名:</b>

```go
type StackNode[T any] struct {
	Value T
	Next  *StackNode[T]
}
type LinkedStack[T any] struct {
	top    *datastructure.StackNode[T]
	length int
}
func NewLinkedStack[T any]() *LinkedStack[T]
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewLinkedStack[int]()
    fmt.Println(sk)
}
```




### <span id="LinkedStack_Push">Push</span>
<p>将元素加入链表栈</p>

<b>函数签名:</b>

```go
func (s *LinkedStack[T]) Push(value T)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewLinkedStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    fmt.Println(sk.Data()) //[]int{3, 2, 1}
}
```




### <span id="LinkedStack_Pop">Pop</span>
<p>删除栈顶元素并返回该元素指针</p>

<b>函数签名:</b>

```go
func (s *LinkedStack[T]) Pop() (*T, error)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewLinkedStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    val, err := sk.Pop()
    fmt.Println(err) //nil
    fmt.Println(*val) //3

    fmt.Println(sk.Data()) //[]int{2, 1}
}
```




### <span id="LinkedStack_Peak">Peak</span>
<p>返回栈顶元素指针</p>

<b>函数签名:</b>

```go
func (s *LinkedStack[T]) Peak() (*T, error)
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewLinkedStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    val, err := sk.Peak()
    fmt.Println(err) //nil
    fmt.Println(*val) //3

    fmt.Println(sk.Data()) //[]int{3, 2, 1}
}
```




### <span id="LinkedStack_Data">Data</span>
<p>返回栈中所有元素组成的切片</p>

<b>函数签名:</b>

```go
func (s *LinkedStack[T]) Data() []T
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewLinkedStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    fmt.Println(sk.Data()) //[]int{3, 2, 1}
}
```




### <span id="LinkedStack_Size">Size</span>
<p>返回栈中元素的数量</p>

<b>函数签名:</b>

```go
func (s *LinkedStack[T]) Size() int
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewLinkedStack[int]()
    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    fmt.Println(sk.Size()) //3
}
```




### <span id="LinkedStack_IsEmpty">IsEmpty</span>
<p>判断栈是否为空</p>

<b>函数签名:</b>

```go
func (s *LinkedStack[T]) IsEmpty() bool
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewLinkedStack[int]()
    fmt.Println(sk.IsEmpty()) //true

    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    fmt.Println(sk.IsEmpty()) //false
}
```




### <span id="LinkedStack_Clear">Clear</span>
<p>清空栈元素，使栈为空</p>

<b>函数签名:</b>

```go
func (s *LinkedStack[T]) Clear()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewLinkedStack[int]()

    sk.Push(1)
    sk.Push(2)
    sk.Push(3)

    sk.Clear()

    fmt.Println(sk.Data()) //[]int{}
}
```




### <span id="LinkedStack_Print">Print</span>
<p>打印链表栈结构</p>

<b>函数签名:</b>

```go
func (s *LinkedStack[T]) Print()
```
<b>示例:</b>

```go
package main

import (
    "fmt"
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)

func main() {
    sk := stack.NewLinkedStack[int]()

    sk.Push(1)
    sk.Push(2)
    sk.Push(3)


    sk.Print() //[ &{Value:3 Next:0xc000010260}, &{Value:2 Next:0xc000010250}, &{Value:1 Next:<nil>},  ]
}
```

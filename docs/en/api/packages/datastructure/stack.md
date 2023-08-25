# Stack
Stack is an abstract data type that serves as a collection of elements. Elements follow the LIFO principle. FIFO is last-in, first-out, meaning that the most recently produced items are recorded as sold first.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/datastructure/stack/arraystack.go](https://github.com/duke-git/lancet/blob/main/datastructure/stack/arraystack.go)
- [https://github.com/duke-git/lancet/blob/main/datastructure/stack/linkedstack.go](https://github.com/duke-git/lancet/blob/main/datastructure/stack/linkedstack.go)


<div STYLE="page-break-after: always;"></div>

## Usage
```go
import (
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

### 1. ArrayStack

- [NewArrayStack](#NewArrayStack)
- [Push](#ArrayStack_Push)
- [Pop](#ArrayStack_Pop)
- [Peak](#ArrayStack_Peak)
- [Data](#ArrayStack_Data)
- [Size](#ArrayStack_Size)
- [IsEmpty](#ArrayStack_IsEmpty)
- [Clear](#ArrayStack_Clear)

### 2. LinkedStack

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

## Documentation

### 1. ArrayStack
ArrayStack is a stack implemented by slice.

### <span id="NewArrayStack">NewArrayStack</span>
<p>Return a empty ArrayStack pointer</p>

<b>Signature:</b>

```go
type ArrayStack[T any] struct {
	data   []T
	length int
}
func NewArrayStack[T any]() *ArrayStack[T]
```
<b>Example:</b>

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
<p>Push element into array stack</p>

<b>Signature:</b>

```go
func (s *ArrayStack[T]) Push(value T)
```
<b>Example:</b>

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
<p>Delete the top element of stack then return it, if stack is empty, return nil and error</p>

<b>Signature:</b>

```go
func (s *ArrayStack[T]) Pop() (*T, error)
```
<b>Example:</b>

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
<p>Return the top element of array stack</p>

<b>Signature:</b>

```go
func (s *ArrayStack[T]) Peak() (*T, error)
```
<b>Example:</b>

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
<p>Return a slice of all data in array stack</p>

<b>Signature:</b>

```go
func (s *ArrayStack[T]) Data() []T
```
<b>Example:</b>

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
<p>Return number of elements in array stack</p>

<b>Signature:</b>

```go
func (s *ArrayStack[T]) Size() int
```
<b>Example:</b>

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
<p>Check if array stack is empty or not</p>

<b>Signature:</b>

```go
func (s *ArrayStack[T]) IsEmpty() bool
```
<b>Example:</b>

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
<p>Clear all elments in array stack</p>

<b>Signature:</b>

```go
func (s *ArrayStack[T]) Clear()
```
<b>Example:</b>

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
LinkedStack is a stack implemented by linked list.

### <span id="NewLinkedStack">NewLinkedStack</span>
<p>Return a empty LinkedStack pointer</p>

<b>Signature:</b>

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
<b>Example:</b>

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
<p>Push element into linked stack</p>

<b>Signature:</b>

```go
func (s *LinkedStack[T]) Push(value T)
```
<b>Example:</b>

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
<p>Delete the top element of stack then return it, if stack is empty, return nil and error</p>

<b>Signature:</b>

```go
func (s *LinkedStack[T]) Pop() (*T, error)
```
<b>Example:</b>

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
<p>Return the top element of linked stack</p>

<b>Signature:</b>

```go
func (s *LinkedStack[T]) Peak() (*T, error)
```
<b>Example:</b>

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
<p>Return a slice of all data in linked stack</p>

<b>Signature:</b>

```go
func (s *LinkedStack[T]) Data() []T
```
<b>Example:</b>

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
<p>Return number of elements in linked stack</p>

<b>Signature:</b>

```go
func (s *LinkedStack[T]) Size() int
```
<b>Example:</b>

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
<p>Check if linked stack is empty or not</p>

<b>Signature:</b>

```go
func (s *LinkedStack[T]) IsEmpty() bool
```
<b>Example:</b>

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
<p>Clear all elments in linked stack</p>

<b>Signature:</b>

```go
func (s *LinkedStack[T]) Clear()
```
<b>Example:</b>

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
<p>Print the structure of a linked stack</p>

<b>Signature:</b>

```go
func (s *LinkedStack[T]) Print()
```
<b>Example:</b>

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

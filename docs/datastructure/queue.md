# Queue
A queue is a kind of linear table. It only allows delete operations at the front of the table and insert operations at the rear of the table. This package includes ArrayQueue, LinkedQueue, CircularQueue, and PriorityQueue.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/datastructure/queue/arrayqueue.go](https://github.com/duke-git/lancet/blob/main/datastructure/queue/arrayqueue.go)
- [https://github.com/duke-git/lancet/blob/main/datastructure/queue/linkedqueue.go](https://github.com/duke-git/lancet/blob/main/datastructure/queue/linkedqueue.go)
- [https://github.com/duke-git/lancet/blob/main/datastructure/queue/circularqueue.go](https://github.com/duke-git/lancet/blob/main/datastructure/queue/circularqueue.go)
- [https://github.com/duke-git/lancet/blob/main/datastructure/queue/priorityqueue.go](https://github.com/duke-git/lancet/blob/main/datastructure/queue/priorityqueue.go)

<div STYLE="page-break-after: always;"></div>

## Usage
```go
import (
    stack "github.com/duke-git/lancet/v2/datastructure/stack"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

### 1. ArrayQueue
- [NewArrayQueue](#NewArrayQueue)
- [Enqueue](#ArrayQueue_Enqueue)
- [Dequeue](#ArrayQueue_Dequeue)
- [Front](#ArrayQueue_Front)
- [Back](#ArrayQueue_Back)
- [Front](#ArrayQueue_Size)
- [IsEmpty](#ArrayQueue_IsEmpty)
- [Clear](#ArrayQueue_Clear)
- [Contain](#ArrayQueue_Contain)



### 2. LinkedStack



<div STYLE="page-break-after: always;"></div>

## Documentation

### 1. ArrayQueue
Common queue implemented by slice.

### <span id="NewArrayQueue">NewArrayQueue</span>
<p>Return a ArrayQueue pointer with specific capacity </p>

<b>Signature:</b>

```go
func NewArrayQueue[T any](capacity int) *ArrayQueue[T]

type ArrayQueue[T any] struct {
	items    []T
	head     int
	tail     int
	capacity int
	size     int
}
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    fmt.Println(q.Data()) // []
}
```



### <span id="ArrayQueue_Data">Data</span>
<p>Get all queue data</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) Data() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    fmt.Println(q.Data()) // []
}
```




### <span id="ArrayQueue_Enqueue">Enqueue</span>
<p>Put element into queue, if queue is full, return false</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) Enqueue(item T) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Data()) // 1,2,3
}
```




### <span id="ArrayQueue_Dequeue">Dequeue</span>
<p>Remove head element of queue and return it</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) Dequeue() (T, bool)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Dequeue()) // 1
    fmt.Println(q.Data()) // 2,3
}
```




### <span id="ArrayQueue_Front">Front</span>
<p>Just get head element of queue</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) Front() T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Front()) // 1
    fmt.Println(q.Data()) // 1,2,3
}
```




### <span id="ArrayQueue_Back">Back</span>
<p>Just get tail element of queue</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) Back() T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Back()) // 3
    fmt.Println(q.Data()) // 1,2,3
}
```



### <span id="ArrayQueue_Size">Size</span>
<p>Get the number of elements in queue</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) Size() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Size()) // 3
}
```



### <span id="ArrayQueue_IsEmpty">IsEmpty</span>
<p>Check if queue is empty or not</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) IsEmpty() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    fmt.Println(q.IsEmpty()) // true

    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.IsEmpty()) // false
}
```




### <span id="ArrayQueue_Clear">Clear</span>
<p>Clean all data in queue</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) Clear()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Clear()

    fmt.Println(q.IsEmpty()) // true
}
```



### <span id="ArrayQueue_Contain">Contain</span>
<p>checks if the value is in queue or not</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) Contain(value T) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Contain(1)) // true
    fmt.Println(q.Contain(4)) // false
}
```



### 2. LinkedQueue
Common queue implemented by link.

### <span id="NewLinkedQueue">NewLinkedQueue</span>
<p>Return a LinkedQueue pointer with </p>

<b>Signature:</b>

```go
func NewLinkedQueue[T any]() *LinkedQueue[T]

type LinkedQueue[T any] struct {
	head   *datastructure.QueueNode[T]
	tail   *datastructure.QueueNode[T]
	length int
}
type QueueNode[T any] struct {
	Value T
	Next  *QueueNode[T]
}
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int]()
    fmt.Println(q.Data()) // []
}
```



### <span id="LinkedQueue_Data">Data</span>
<p>Get all queue data</p>

<b>Signature:</b>

```go
func (q *LinkedQueue[T]) Data() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int]()
    fmt.Println(q.Data()) // []
}
```




### <span id="LinkedQueue_Enqueue">Enqueue</span>
<p>Put element into queue</p>

<b>Signature:</b>

```go
func (q *LinkedQueue[T]) Enqueue(value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Data()) // 1,2,3
}
```




### <span id="LinkedQueue_Dequeue">Dequeue</span>
<p>Remove head element of queue and return it</p>

<b>Signature:</b>

```go
func (q *LinkedQueue[T]) Dequeue() (T, error)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Dequeue()) // 1
    fmt.Println(q.Data()) // 2,3
}
```




### <span id="LinkedQueue_Front">Front</span>
<p>Just get head element of queue</p>

<b>Signature:</b>

```go
func (q *LinkedQueue[T]) Front() (*T, error)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Front()) // 1
    fmt.Println(q.Data()) // 1,2,3
}
```




### <span id="LinkedQueue_Back">Back</span>
<p>Just get tail element of queue</p>

<b>Signature:</b>

```go
func (q *LinkedQueue[T]) Back() (*T, error) 
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Back()) // 3
    fmt.Println(q.Data()) // 1,2,3
}
```



### <span id="LinkedQueue_Size">Size</span>
<p>Get the number of elements in queue</p>

<b>Signature:</b>

```go
func (q *LinkedQueue[T]) Size() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Size()) // 3
}
```



### <span id="LinkedQueue_IsEmpty">IsEmpty</span>
<p>Check if queue is empty or not</p>

<b>Signature:</b>

```go
func (q *LinkedQueue[T]) IsEmpty() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int](5)
    fmt.Println(q.IsEmpty()) // true

    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.IsEmpty()) // false
}
```




### <span id="LinkedQueue_Clear">Clear</span>
<p>Clean all data in queue</p>

<b>Signature:</b>

```go
func (q *LinkedQueue[T]) Clear()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Clear()

    fmt.Println(q.IsEmpty()) // true
}
```



### <span id="LinkedQueue_Contain">Contain</span>
<p>checks if the value is in queue or not</p>

<b>Signature:</b>

```go
func (q *LinkedQueue[T]) Contain(value T) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewLinkedQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Contain(1)) // true
    fmt.Println(q.Contain(4)) // false
}
```


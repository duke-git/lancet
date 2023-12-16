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
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

### 1. ArrayQueue
- [NewArrayQueue](#NewArrayQueue)
- [Data](#ArrayQueue_Data)
- [Enqueue](#ArrayQueue_Enqueue)
- [Dequeue](#ArrayQueue_Dequeue)
- [Front](#ArrayQueue_Front)
- [Back](#ArrayQueue_Back)
- [Size](#ArrayQueue_Size)
- [IsEmpty](#ArrayQueue_IsEmpty)
- [IsFull](#ArrayQueue_IsFull)
- [Clear](#ArrayQueue_Clear)
- [Contain](#ArrayQueue_Contain)



### 2. LinkedQueue
- [NewLinkedQueue](#NewLinkedQueue)
- [Data](#LinkedQueue_Data)
- [Enqueue](#LinkedQueue_Enqueue)
- [Dequeue](#LinkedQueue_Dequeue)
- [Front](#LinkedQueue_Front)
- [Back](#LinkedQueue_Back)
- [Size](#LinkedQueue_Size)
- [IsEmpty](#LinkedQueue_IsEmpty)
- [Clear](#LinkedQueue_Clear)
- [Contain](#LinkedQueue_Contain)


### 3. CircularQueue
- [NewCircularQueue](#NewCircularQueue)
- [Data](#CircularQueue_Data)
- [Enqueue](#CircularQueue_Enqueue)
- [Dequeue](#CircularQueue_Dequeue)
- [Front](#CircularQueue_Front)
- [Back](#CircularQueue_Back)
- [Size](#CircularQueue_Size)
- [IsEmpty](#CircularQueue_IsEmpty)
- [IsFull](#CircularQueue_IsFull)
- [Clear](#CircularQueue_Clear)
- [Contain](#CircularQueue_Contain)



### 4. PriorityQueue
- [NewPriorityQueue](#NewPriorityQueue)
- [Data](#PriorityQueue_Data)
- [Enqueue](#PriorityQueue_Enqueue)
- [Dequeue](#PriorityQueue_Dequeue)
- [IsEmpty](#PriorityQueue_IsEmpty)
- [IsFull](#PriorityQueue_IsFull)
- [Size](#PriorityQueue_Size)


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
<p>Just get the head element of queue</p>

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
<p>Just get the tail element of queue</p>

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




### <span id="ArrayQueue_IsFull">IsFull</span>
<p>Check if queue is full or not</p>

<b>Signature:</b>

```go
func (q *ArrayQueue[T]) IsFull() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewArrayQueue[int](3)
    fmt.Println(q.IsFull()) // false

    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.IsFull()) // true
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
<p>Check if the value is in queue or not</p>

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
<p>Return a LinkedQueue pointer </p>

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
<p>Just get the head element of queue</p>

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
<p>Just get the tail element of queue</p>

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
<p>Check if the value is in queue or not</p>

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




### 3. CircularQueue
Circular queue implemented by slice.

### <span id="NewCircularQueue">NewCircularQueue</span>
<p>Return a CircularQueue pointer with specific capacity </p>

<b>Signature:</b>

```go
func NewCircularQueue[T any](capacity int) *CircularQueue[T]

type CircularQueue[T any] struct {
	data  []T
	front int
	rear  int
	capacity  int
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
    q := queue.NewCircularQueue[int](5)
    fmt.Println(q.Data()) // []
}
```



### <span id="CircularQueue_Data">Data</span>
<p>Get all queue data</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) Data() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](5)
    fmt.Println(q.Data()) // []
}
```




### <span id="CircularQueue_Enqueue">Enqueue</span>
<p>Put element into queue, if queue is full, return error</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) Enqueue(value T) error
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Data()) // 1,2,3
}
```




### <span id="CircularQueue_Dequeue">Dequeue</span>
<p>Remove head element of queue and return it</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) Dequeue() (*T, bool)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    val := q.Dequeue()
    fmt.Println(*val) // 1
    fmt.Println(q.Data()) // 2,3
}
```




### <span id="CircularQueue_Front">Front</span>
<p>Just get head element of queue</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) Front() T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Front()) // 1
    fmt.Println(q.Data()) // 1,2,3
}
```




### <span id="CircularQueue_Back">Back</span>
<p>Just get tail element of queue</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) Back() T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Back()) // 3
    fmt.Println(q.Data()) // 1,2,3
}
```



### <span id="CircularQueue_Size">Size</span>
<p>Get the number of elements in queue</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) Size() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Size()) // 3
}
```



### <span id="CircularQueue_IsEmpty">IsEmpty</span>
<p>Check if queue is empty or not</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) IsEmpty() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](5)
    fmt.Println(q.IsEmpty()) // true

    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.IsEmpty()) // false
}
```




### <span id="CircularQueue_IsFull">IsFull</span>
<p>Check if queue is full or not</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) IsFull() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](3)
    fmt.Println(q.IsFull()) // false

    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.IsFull()) // true
}
```



### <span id="CircularQueue_Clear">Clear</span>
<p>Clean all data in queue</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) Clear()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)
    q.Clear()

    fmt.Println(q.IsEmpty()) // true
}
```



### <span id="CircularQueue_Contain">Contain</span>
<p>Check if the value is in queue or not</p>

<b>Signature:</b>

```go
func (q *CircularQueue[T]) Contain(value T) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewCircularQueue[int](5)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    fmt.Println(q.Contain(1)) // true
    fmt.Println(q.Contain(4)) // false
}
```


### 4. PriorityQueue
Common queue implemented by slice.

### <span id="NewPriorityQueue">NewPriorityQueue</span>
<p>Return a PriorityQueue pointer with specific capacity, param `comparator` is used to compare values of type T in the queue. </p>

<b>Signature:</b>

```go
func NewPriorityQueue[T any](capacity int, comparator constraints.Comparator) *PriorityQueue[T]

type PriorityQueue[T any] struct {
	items      []T
	size       int
	comparator constraints.Comparator
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
    q := queue.NewPriorityQueue[int](3)
    fmt.Println(q.Data()) // []
}
```



### <span id="PriorityQueue_Data">Data</span>
<p>Get all queue data</p>

<b>Signature:</b>

```go
func (q *PriorityQueue[T]) Data() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

func main() {
    q := queue.NewPriorityQueue[int](3)
    fmt.Println(q.Data()) // []
}
```




### <span id="PriorityQueue_Enqueue">Enqueue</span>
<p>Put element into queue, if queue is full, return false</p>

<b>Signature:</b>

```go
func (q *PriorityQueue[T]) Enqueue(item T) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    comparator := &intComparator{}
    q := queue.NewPriorityQueue[int](10, comparator)
    for i := 1; i < 11; i++ {
		q.Enqueue(i)
	}

    fmt.Println(q.Data()) // 10, 9, 6, 7, 8, 2, 5, 1, 4, 3
}
```




### <span id="PriorityQueue_Dequeue">Dequeue</span>
<p>Remove head element of queue and return it</p>

<b>Signature:</b>

```go
func (q *PriorityQueue[T]) Dequeue() (T, bool)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)


type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    comparator := &intComparator{}
    q := queue.NewPriorityQueue[int](10, comparator)
    for i := 1; i < 11; i++ {
		q.Enqueue(i)
	}
    val, ok := pq.Dequeue()
    fmt.Println(val) // 10
    fmt.Println(q.Data()) // 9, 8, 6, 7, 3, 2, 5, 1, 4
}
```



### <span id="PriorityQueue_IsEmpty">IsEmpty</span>
<p>Check if queue is empty or not</p>

<b>Signature:</b>

```go
func (q *PriorityQueue[T]) IsEmpty() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    comparator := &intComparator{}
    q := queue.NewPriorityQueue[int](10, comparator)
    fmt.Println(q.IsEmpty()) // true

    for i := 1; i < 11; i++ {
		q.Enqueue(i)
	}
    fmt.Println(q.IsEmpty()) // false
}
```




### <span id="PriorityQueue_IsFull">IsFull</span>
<p>Check if queue is full or not</p>

<b>Signature:</b>

```go
func (q *PriorityQueue[T]) IsFull() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    comparator := &intComparator{}
    q := queue.NewPriorityQueue[int](10, comparator)
    fmt.Println(q.IsFull()) // false

    for i := 1; i < 11; i++ {
		q.Enqueue(i)
	}
    fmt.Println(q.IsFull()) // true
}
```




### <span id="PriorityQueue_Size">Size</span>
<p>Get nubmers of elements in queue</p>

<b>Signature:</b>

```go
func (q *PriorityQueue[T]) Size() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    queue "github.com/duke-git/lancet/v2/datastructure/queue"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    comparator := &intComparator{}
    q := queue.NewPriorityQueue[int](10, comparator)
    fmt.Println(q.IsFull()) // false

    for i := 1; i < 5; i++ {
		q.Enqueue(i)
	}
    fmt.Println(q.Size()) // 4
}
```



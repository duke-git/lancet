# Heap
Heap is a binary heap tree implemented by slice.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/datastructure/heap/maxheap.go](https://github.com/duke-git/lancet/blob/main/datastructure/heap/maxheap.go)


<div STYLE="page-break-after: always;"></div>

## Usage
```go
import (
    heap "github.com/duke-git/lancet/v2/datastructure/heap"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

- [MaxHeap](#MaxHeap)
- [Values](#Values)
- [Add](#Add)
- [Delete](#Delete)
- [Contain](#Contain)
- [ContainAll](#ContainAll)
- [Clone](#Clone)
- [Size](#Size)
- [Equal](#Equal)


<div STYLE="page-break-after: always;"></div>

## Documentation

### 1. MaxHeap
MaxHeap is a binary heap tree implemented by slice, The key of the root node is both greater than or equal to the key value of the left subtree and greater than or equal to the key value of the right subtree.

### <span id="NewMaxHeap">NewMaxHeap</span>
<p>Return a NewMaxHeap pointer instance.</p>

<b>Signature:</b>

```go
type MaxHeap[T any] struct {
	data       []T
	comparator lancetconstraints.Comparator
}
func NewMaxHeap[T any](comparator lancetconstraints.Comparator) *MaxHeap[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    heap "github.com/duke-git/lancet/v2/datastructure/heap"
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
    maxHeap := heap.NewMaxHeap[int](&intComparator{})
    fmt.Println(maxHeap)
}
```




### <span id="Push">Push</span>
<p>Push value into the heap</p>

<b>Signature:</b>

```go
func (h *MaxHeap[T]) Push(value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    heap "github.com/duke-git/lancet/v2/datastructure/heap"
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
    maxHeap := heap.NewMaxHeap[int](&intComparator{})
    values := []int{6, 5, 2, 4, 7, 10, 12, 1, 3, 8, 9, 11}

	for _, v := range values {
		maxHeap.Push(v)
	}

    fmt.Println(maxHeap.Data()) //[]int{12, 9, 11, 4, 8, 10, 7, 1, 3, 5, 6, 2}
}
```




### <span id="Pop">Pop</span>
<p>Pop return the largest value, and remove it from the heap if heap is empty, return zero value and fasle</p>

<b>Signature:</b>

```go
func (h *MaxHeap[T]) Pop() (T, bool)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    heap "github.com/duke-git/lancet/v2/datastructure/heap"
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
    maxHeap := heap.NewMaxHeap[int](&intComparator{})
    values := []int{6, 5, 2, 4, 7, 10, 12, 1, 3, 8, 9, 11}

	for _, v := range values {
		maxHeap.Push(v)
	}
    val, ok := maxHeap.Pop()

    fmt.Println(val) //12
    fmt.Println(ok) //true
}
```



### <span id="Peek">Peek</span>
<p>Return the largest element from the heap without removing it, if heap is empty, it returns zero value and false.</p>

<b>Signature:</b>

```go
func (h *MaxHeap[T]) Peek() (T, bool)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    heap "github.com/duke-git/lancet/v2/datastructure/heap"
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
    maxHeap := heap.NewMaxHeap[int](&intComparator{})
    values := []int{6, 5, 2, 4, 7, 10, 12, 1, 3, 8, 9, 11}

	for _, v := range values {
		maxHeap.Push(v)
	}
    val, ok := maxHeap.Peek()

    fmt.Println(val) //12
    fmt.Println(maxHeap.Size()) //12
}
```



### <span id="Data">Data</span>
<p>Return all element of the heap</p>

<b>Signature:</b>

```go
func (h *MaxHeap[T]) Data() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    heap "github.com/duke-git/lancet/v2/datastructure/heap"
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
    maxHeap := heap.NewMaxHeap[int](&intComparator{})
    values := []int{6, 5, 2, 4, 7, 10, 12, 1, 3, 8, 9, 11}

	for _, v := range values {
		maxHeap.Push(v)
	}

    fmt.Println(maxHeap.Data()) //[]int{12, 9, 11, 4, 8, 10, 7, 1, 3, 5, 6, 2}
}
```


### <span id="Size">Size</span>
<p>Return the number of elements in the heap</p>

<b>Signature:</b>

```go
func (h *MaxHeap[T]) Size() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    heap "github.com/duke-git/lancet/v2/datastructure/heap"
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
    maxHeap := heap.NewMaxHeap[int](&intComparator{})
    values := []int{6, 5, 2}

	for _, v := range values {
		maxHeap.Push(v)
	}

    fmt.Println(maxHeap.Size()) //3
}
```



### <span id="PrintStructure">PrintStructure</span>
<p>Print the tree structure of the heap</p>

<b>Signature:</b>

```go
func (h *MaxHeap[T]) PrintStructure()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    heap "github.com/duke-git/lancet/v2/datastructure/heap"
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
    maxHeap := heap.NewMaxHeap[int](&intComparator{})
    values := []int{6, 5, 2, 4, 7, 10, 12, 1, 3, 8, 9, 11}

	for _, v := range values {
		maxHeap.Push(v)
	}

    fmt.Println(maxHeap.PrintStructure())
//        12
//    9       11
//  4   8   10   7
// 1 3 5 6 2
}
```
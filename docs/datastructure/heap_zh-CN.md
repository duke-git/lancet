# Heap
堆，切片实现的二叉堆数据结构。

<div STYLE="page-break-after: always;"></div>

## 源码

- [https://github.com/duke-git/lancet/blob/main/datastructure/heap/maxheap.go](https://github.com/duke-git/lancet/blob/main/datastructure/heap/maxheap.go)


<div STYLE="page-break-after: always;"></div>

## 用法
```go
import (
    heap "github.com/duke-git/lancet/v2/datastructure/heap"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

- [MaxHeap](#MaxHeap)
- [Push](#Push)
- [Pop](#Pop)
- [Peek](#Peek)
- [Data](#Data)
- [Size](#Size)


<div STYLE="page-break-after: always;"></div>

## API文档

### 1. MaxHeap
MaxHeap是通过slice实现的二叉堆树，根节点的key既大于等于左子树的key值且大于等于右子树的key值。

### <span id="NewMaxHeap">NewMaxHeap</span>
<p>返回NewMaxHeap指针实例</p>

<b>函数签名:</b>

```go
type MaxHeap[T any] struct {
	data       []T
	comparator lancetconstraints.Comparator
}
func NewMaxHeap[T any](comparator lancetconstraints.Comparator) *MaxHeap[T]
```
<b>示例:</b>

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
<p>向堆中插入数据</p>

<b>函数签名:</b>

```go
func (h *MaxHeap[T]) Push(value T)
```
<b>示例:</b>

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
<p>返回堆中最大值并将其从堆中删除，如果堆为空，返回零值并返回false</p>

<b>函数签名:</b>

```go
func (h *MaxHeap[T]) Pop() (T, bool)
```
<b>示例:</b>

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
<p>返回堆中最大值，如果堆为空，返回零值并返回false</p>

<b>函数签名:</b>

```go
func (h *MaxHeap[T]) Peek() (T, bool)
```
<b>示例:</b>

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
<p>返回堆中全部元素的切片</p>

<b>函数签名:</b>

```go
func (h *MaxHeap[T]) Data() []T
```
<b>示例:</b>

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
<p>返回堆中元素的数量</p>

<b>函数签名:</b>

```go
func (h *MaxHeap[T]) Size() int
```
<b>示例:</b>

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
<p>打印堆的树形结构</p>

<b>函数签名:</b>

```go
func (h *MaxHeap[T]) PrintStructure()
```
<b>示例:</b>

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
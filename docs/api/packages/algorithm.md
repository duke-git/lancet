# Algorithm

algorithm 算法包实现一些基本算法，sort，search，lrucache。

<div STYLE="page-break-after: always;"></div>

## 源码

-   [https://github.com/duke-git/lancet/blob/main/algorithm/sort.go](https://github.com/duke-git/lancet/blob/main/algorithm/sort.go)
-   [https://github.com/duke-git/lancet/blob/main/algorithm/search.go](https://github.com/duke-git/lancet/blob/main/algorithm/search.go)
-   [https://github.com/duke-git/lancet/blob/main/algorithm/lru_cache.go](https://github.com/duke-git/lancet/blob/main/algorithm/lru_cache.go)

<div STYLE="page-break-after: always;"></div>

## 用法

```go
import (
    "github.com/duke-git/lancet/v2/algorithm"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [BubbleSort](#BubbleSort)
-   [InsertionSort](#InsertionSort)
-   [SelectionSort](#SelectionSort)
-   [ShellSort](#ShellSort)
-   [QuickSort](#QuickSort)
-   [HeapSort](#HeapSort)
-   [MergeSort](#MergeSort)
-   [CountSort](#CountSort)
-   [BinarySearch](#BinarySearch)
-   [BinaryIterativeSearch](#BinaryIterativeSearch)
-   [LinearSearch](#LinearSearch)
-   [LRUCache](#LRUCache)

<div STYLE="page-break-after: always;"></div>

<link rel="stylesheet" type="text/css" href="/styles/api_doc.css">

## 文档

### <span id="BubbleSort">BubbleSort</span>

<p>冒泡排序，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func BubbleSort[T any](slice []T, comparator lancetconstraints.Comparator)
```

<b>示例:<span class="run-container">[运行](https://go.dev/play/p/GNdv7Jg2Taj)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
    val1, _ := v1.(int)
    val2, _ := v2.(int)

    //ascending order
    if val1 < val2 {
        return -1
    } else if val1 > val2 {
        return 1
    }
    return 0
}

func main() {
    numbers := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}

    algorithm.BubbleSort(numbers, comparator)

    fmt.Println(numbers)

    // Output:
    // [1 2 3 4 5 6]
}
```

### <span id="InsertionSort">InsertionSort</span>

<p>插入排序，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func InsertionSort[T any](slice []T, comparator lancetconstraints.Comparator)
```

<b>示例:<span class="run-container">[运行](https://go.dev/play/p/G5LJiWgJJW6)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

type people struct {
    Name string
    Age  int
}

// PeopleAageComparator sort people slice by age field
type peopleAgeComparator struct{}

// Compare implements github.com/duke-git/lancet/lancetconstraints/constraints.go/Comparator
func (pc *peopleAgeComparator) Compare(v1 any, v2 any) int {
    p1, _ := v1.(people)
    p2, _ := v2.(people)

    //ascending order
    if p1.Age < p2.Age {
        return -1
    } else if p1.Age > p2.Age {
        return 1
    }

    return 0
}

func main() {
    peoples := []people{
        {Name: "a", Age: 20},
        {Name: "b", Age: 10},
        {Name: "c", Age: 17},
        {Name: "d", Age: 8},
        {Name: "e", Age: 28},
    }

    comparator := &peopleAgeComparator{}

    algorithm.InsertionSort(peoples, comparator)

    fmt.Println(peoples)

    // Output:
    // [{d 8} {b 10} {c 17} {a 20} {e 28}]
}
```

### <span id="SelectionSort">SelectionSort</span>

<p>选择排序，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func SelectionSort[T any](slice []T, comparator lancetconstraints.Comparator)
```

<b>示例:<span class="run-container">[运行](https://go.dev/play/p/oXovbkekayS)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
    val1, _ := v1.(int)
    val2, _ := v2.(int)

    //ascending order
    if val1 < val2 {
        return -1
    } else if val1 > val2 {
        return 1
    }
    return 0
}

func main() {
    numbers := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}

    algorithm.SelectionSort(numbers, comparator)

    fmt.Println(numbers)

    // Output:
    // [1 2 3 4 5 6]
}
```

### <span id="ShellSort">ShellSort</span>

<p>希尔排序，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func ShellSort[T any](slice []T, comparator lancetconstraints.Comparator)
```

<b>示例:<span class="run-container">[运行](https://go.dev/play/p/3ibkszpJEu3)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
    val1, _ := v1.(int)
    val2, _ := v2.(int)

    //ascending order
    if val1 < val2 {
        return -1
    } else if val1 > val2 {
        return 1
    }
    return 0
}

func main() {
    numbers := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}

    algorithm.ShellSort(numbers, comparator)

    fmt.Println(numbers)

    // Output:
    // [1 2 3 4 5 6]
}
```

### <span id="QuickSort">QuickSort</span>

<p>快速排序，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func QuickSort[T any](slice []T comparator lancetconstraints.Comparator)
```

<b>示例:<span class="run-container">[运行](https://go.dev/play/p/7Y7c1Elk3ax)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
    val1, _ := v1.(int)
    val2, _ := v2.(int)

    //ascending order
    if val1 < val2 {
        return -1
    } else if val1 > val2 {
        return 1
    }
    return 0
}

func main() {
    numbers := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}

    algorithm.QuickSort(numbers, comparator)

    fmt.Println(numbers)

    // Output:
    // [1 2 3 4 5 6]
}
```

### <span id="HeapSort">HeapSort</span>

<p>堆排序，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func HeapSort[T any](slice []T, comparator lancetconstraints.Comparator)
```

<b>示例:<span class="run-container">[运行](https://go.dev/play/p/u6Iwa1VZS_f)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
    val1, _ := v1.(int)
    val2, _ := v2.(int)

    //ascending order
    if val1 < val2 {
        return -1
    } else if val1 > val2 {
        return 1
    }
    return 0
}

func main() {
    numbers := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}

    algorithm.HeapSort(numbers, comparator)

    fmt.Println(numbers)

    // Output:
    // [1 2 3 4 5 6]
}
```

### <span id="MergeSort">MergeSort</span>

<p>归并排序，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func MergeSort[T any](slice []T, comparator lancetconstraints.Comparator)
```

<b>示例:<span class="run-container">[运行](https://go.dev/play/p/ydinn9YzUJn)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
    val1, _ := v1.(int)
    val2, _ := v2.(int)

    //ascending order
    if val1 < val2 {
        return -1
    } else if val1 > val2 {
        return 1
    }
    return 0
}

func main() {
    numbers := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}

    algorithm.MergeSort(numbers, comparator)

    fmt.Println(numbers)

    // Output:
    // [1 2 3 4 5 6]
}
```

### <span id="CountSort">CountSort</span>

<p>计数排序，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func CountSort[T any](slice []T, comparator lancetconstraints.Comparator) []T
```

<b>示例:<span class="run-container">[运行](https://go.dev/play/p/tB-Umgm0DrP)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)


type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
    val1, _ := v1.(int)
    val2, _ := v2.(int)

    //ascending order
    if val1 < val2 {
        return -1
    } else if val1 > val2 {
        return 1
    }
    return 0
}

func main() {
    numbers := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}

    sortedNums := algorithm.CountSort(numbers, comparator)

    fmt.Println(sortedNums)

    // Output:
    // [1 2 3 4 5 6]
}
```

### <span id="BinarySearch">BinarySearch</span>

<p>二分递归查找，返回元素索引，未找到元素返回-1，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func BinarySearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) int
```

<b>示例: <span class="run-container">[运行](https://go.dev/play/p/t6MeGiUSN47)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
    val1, _ := v1.(int)
    val2, _ := v2.(int)

    //ascending order
    if val1 < val2 {
        return -1
    } else if val1 > val2 {
        return 1
    }
    return 0
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
    comparator := &intComparator{}

    result1 := algorithm.BinarySearch(numbers, 5, 0, len(numbers)-1, comparator)
    result2 := algorithm.BinarySearch(numbers, 9, 0, len(numbers)-1, comparator)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 4
    // -1
}
```

### <span id="BinaryIterativeSearch">BinaryIterativeSearch</span>

<p>二分迭代查找，返回元素索引，未找到元素返回-1，参数comparator需要实现包lancetconstraints.Comparator。</p>

<b>函数签名:</b>

```go
func BinaryIterativeSearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) int
```

<b>示例: <span class="run-container">[运行](https://go.dev/play/p/Anozfr8ZLH3)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

type intComparator struct{}

func (c *intComparator) Compare(v1 any, v2 any) int {
    val1, _ := v1.(int)
    val2, _ := v2.(int)

    //ascending order
    if val1 < val2 {
        return -1
    } else if val1 > val2 {
        return 1
    }
    return 0
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
    comparator := &intComparator{}

    result1 := algorithm.BinaryIterativeSearch(numbers, 5, 0, len(numbers)-1, comparator)
    result2 := algorithm.BinaryIterativeSearch(numbers, 9, 0, len(numbers)-1, comparator)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 4
    // -1
}
```

### <span id="LinearSearch">LinearSearch</span>

<p>基于传入的相等函数线性查找元素，返回元素索引，未找到元素返回-1。</p>

<b>函数签名:</b>

```go
func LinearSearch[T any](slice []T, target T, equal func(a, b T) bool) int
```

<b>示例: <span class="run-container">[运行](https://go.dev/play/p/IsS7rgn5s3x)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
    numbers := []int{3, 4, 5, 3, 2, 1}

    equalFunc := func(a, b int) bool {
        return a == b
    }

    result1 := algorithm.LinearSearch(numbers, 3, equalFunc)
    result2 := algorithm.LinearSearch(numbers, 6, equalFunc)

    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // 0
    // -1
}
```

### <span id="LRUCache">LRUCache</span>

<p>lru算法实现缓存。</p>

<b>函数签名:</b>

```go
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V]
func (l *LRUCache[K, V]) Get(key K) (V, bool)
func (l *LRUCache[K, V]) Put(key K, value V)
func (l *LRUCache[K, V]) Delete(key K) bool
func (l *LRUCache[K, V]) Len() int
```

<b>示例:<span class="run-container">[运行](https://go.dev/play/p/-EZjgOURufP)</span></b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
    cache := algorithm.NewLRUCache[int, int](2)

    cache.Put(1, 1)
    cache.Put(2, 2)

    result1, ok1 := cache.Get(1)
    result2, ok2 := cache.Get(2)
    result3, ok3 := cache.Get(3)

    fmt.Println(result1, ok1)
    fmt.Println(result2, ok2)
    fmt.Println(result3, ok3)

    fmt.Println(cache.Len())

    ok := cache.Delete(2)
    fmt.Println(ok)


    // Output:
    // 1 true
    // 2 true
    // 0 false
    // 2
    // true
}
```

# Algorithm
algorithm算法包实现一些基本算法，sort，search，lrucache。

<div STYLE="page-break-after: always;"></div>

## 源码

- [https://github.com/duke-git/lancet/blob/main/algorithm/sort.go](https://github.com/duke-git/lancet/blob/main/algorithm/sort.go)
- [https://github.com/duke-git/lancet/blob/main/algorithm/search.go](https://github.com/duke-git/lancet/blob/main/algorithm/search.go)
- [https://github.com/duke-git/lancet/blob/main/algorithm/lru_cache.go](https://github.com/duke-git/lancet/blob/main/algorithm/lru_cache.go)

<div STYLE="page-break-after: always;"></div>

## 用法
```go
import (
    "github.com/duke-git/lancet/v2/algorithm"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

- [BubbleSort](#BubbleSort)
- [InsertionSort](#InsertionSort)
- [SelectionSort](#SelectionSort)
- [ShellSort](#ShellSort)
- [QuickSort](#QuickSort)
- [HeapSort](#HeapSort)
- [MergeSort](#MergeSort)

- [CountSort](#CountSort)
- [BinarySearch](#BinarySearch)
- [BinaryIterativeSearch](#BinaryIterativeSearch)
- [LinearSearch](#LinearSearch)
- [LRUCache](#LRUCache)


<div STYLE="page-break-after: always;"></div>

## 文档



### <span id="BubbleSort">BubbleSort</span>
<p>冒泡排序，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func BubbleSort[T any](slice []T, comparator lancetconstraints.Comparator)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    intSlice := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}
    algorithm.BubbleSort(intSlice, comparator)

    fmt.Println(intSlice) //[]int{1, 2, 3, 4, 5, 6}
}
```




### <span id="InsertionSort">InsertionSort</span>
<p>插入排序，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func InsertionSort[T any](slice []T, comparator lancetconstraints.Comparator)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

        //decending order
        // if p1.Age > p2.Age {
        // 	return -1
        // } else if p1.Age < p2.Age {
        // 	return 1
        // }
    }

    var peoples = []people{
        {Name: "a", Age: 20},
        {Name: "b", Age: 10},
        {Name: "c", Age: 17},
        {Name: "d", Age: 8},
        {Name: "e", Age: 28},
    }
    comparator := &peopleAgeComparator{}
    algorithm.InsertionSort(peoples, comparator)

    fmt.Println(intSlice) //[{d 8} {b 10} {c 17} {a 20} {e 28}]
}
```




### <span id="SelectionSort">SelectionSort</span>
<p>选择排序，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func SelectionSort[T any](slice []T, comparator lancetconstraints.Comparator)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    intSlice := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}
    algorithm.SelectionSort(intSlice, comparator)

    fmt.Println(intSlice) //[]int{1, 2, 3, 4, 5, 6}
}
```




### <span id="ShellSort">ShellSort</span>
<p>希尔排序，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func ShellSort[T any](slice []T, comparator lancetconstraints.Comparator)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    intSlice := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}
    algorithm.ShellSort(intSlice, comparator)

    fmt.Println(intSlice) //[]int{1, 2, 3, 4, 5, 6}
}
```




### <span id="QuickSort">QuickSort</span>
<p>快速排序，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func QuickSort[T any](slice []T, comparator lancetconstraints.Comparator) []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    intSlice := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}
    algorithm.QuickSort(intSlice, comparator)

    fmt.Println(intSlice) //[]int{1, 2, 3, 4, 5, 6}
}
```




### <span id="HeapSort">HeapSort</span>
<p>堆排序，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func HeapSort[T any](slice []T, comparator lancetconstraints.Comparator) []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    intSlice := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}
    algorithm.HeapSort(intSlice, comparator)

    fmt.Println(intSlice) //[]int{1, 2, 3, 4, 5, 6}
}
```




### <span id="MergeSort">MergeSort</span>
<p>归并排序，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func MergeSort[T any](slice []T, comparator lancetconstraints.Comparator)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    intSlice := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}
    algorithm.MergeSort(intSlice, comparator)

    fmt.Println(intSlice) //[]int{1, 2, 3, 4, 5, 6}
}
```



### <span id="CountSort">CountSort</span>
<p>计数排序，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func CountSort[T any](slice []T, comparator lancetconstraints.Comparator) []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    intSlice := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}
    sortedSlice := algorithm.CountSort(intSlice, comparator)

    fmt.Println(sortedSlice) //[]int{1, 2, 3, 4, 5, 6}
}
```




### <span id="BinarySearch">BinarySearch</span>
<p>二分递归查找，返回元素索引，未找到元素返回-1，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func BinarySearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    var sortedNumbers = []int{1, 2, 3, 4, 5, 6, 7, 8}
    comparator := &intComparator{}
    foundIndex := algorithm.BinarySearch(sortedNumbers, 5, 0, len(sortedNumbers)-1, comparator)
    fmt.Println(foundIndex) //4

    notFoundIndex := algorithm.BinarySearch(sortedNumbers, 9, 0, len(sortedNumbers)-1, comparator)
    fmt.Println(notFoundIndex) //-1
}
```



### <span id="BinaryIterativeSearch">BinaryIterativeSearch</span>
<p>二分迭代查找，返回元素索引，未找到元素返回-1，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func BinaryIterativeSearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator lancetconstraints.Comparator) int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    var sortedNumbers = []int{1, 2, 3, 4, 5, 6, 7, 8}
    comparator := &intComparator{}
    foundIndex := algorithm.BinaryIterativeSearch(sortedNumbers, 5, 0, len(sortedNumbers)-1, comparator)
    fmt.Println(foundIndex) //4

    notFoundIndex := algorithm.BinaryIterativeSearch(sortedNumbers, 9, 0, len(sortedNumbers)-1, comparator)
    fmt.Println(notFoundIndex) //-1
}
```




### <span id="LinearSearch">LinearSearch</span>
<p>线性查找，返回元素索引，未找到元素返回-1，参数comparator需要实现包lancetconstraints.Comparator</p>

<b>函数签名:</b>

```go
func LinearSearch[T any](slice []T, target T, comparator lancetconstraints.Comparator) int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    "github.com/duke-git/lancet/v2/algorithm"
)

func main() {
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

    intSlice := []int{2, 1, 5, 3, 6, 4}
    comparator := &intComparator{}
    foundIndex := algorithm.LinearSearch(intSlice, 5, comparator)
    fmt.Println(foundIndex) //2

    notFoundIndex := algorithm.LinearSearch(sortedNumbers, 0, comparator)
    fmt.Println(notFoundIndex) //-1
}
```




### <span id="LRUCache">LRUCache</span>
<p>lru实现缓存</p>

<b>函数签名:</b>

```go
func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V]
func (l *LRUCache[K, V]) Get(key K) (V, bool)
func (l *LRUCache[K, V]) Put(key K, value V)
```
<b>Example:</b>

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

    _, ok := cache.Get(0) // ok -> false

    v, ok := cache.Get(1) // v->1, ok->true

}
```
# Algorithm
Package algorithm implements some basic algorithm. eg. sort, search.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/algorithm/sorter.go](https://github.com/duke-git/lancet/blob/main/algorithm/sorter.go)
- [https://github.com/duke-git/lancet/blob/main/algorithm/search.go](https://github.com/duke-git/lancet/blob/main/algorithm/search.go)
- [https://github.com/duke-git/lancet/blob/main/algorithm/lru_cache.go](https://github.com/duke-git/lancet/blob/main/algorithm/lru_cache.go)

<div STYLE="page-break-after: always;"></div>

## Usage
```go
import (
    "github.com/duke-git/lancet/v2/algorithm"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
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

## Documentation



### <span id="BubbleSort">BubbleSort</span>
<p>Sort slice with bubble sort algorithm. Param comparator should implements lancetconstraints.Comparator.</p>

<b>Signature:</b>

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
<p>Sort slice with insertion sort algorithm. Param comparator should implements lancetconstraints.Comparator.</p>

<b>Signature:</b>

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

    fmt.Println(peoples) //[{d 8} {b 10} {c 17} {a 20} {e 28}]
}
```




### <span id="SelectionSort">SelectionSort</span>
<p>Sort slice with selection sort algorithm. Param comparator should implements lancetconstraints.Comparator.</p>

<b>Signature:</b>

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
<p>Sort slice with shell sort algorithm. Param comparator should implements lancetconstraints.Comparator.</p>

<b>Signature:</b>

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
<p>Sort slice with quick sort algorithm. Param comparator should implements lancetconstraints.Comparator.</p>

<b>Signature:</b>

```go
func QuickSort[T any](slice []T, lowIndex, highIndex int, comparator lancetconstraints.Comparator)
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
    algorithm.QuickSort(intSlice, 0, len(intSlice)-1, comparator)

    fmt.Println(intSlice) //[]int{1, 2, 3, 4, 5, 6}
}
```




### <span id="HeapSort">HeapSort</span>
<p>Sort slice with heap sort algorithm. Param comparator should implements lancetconstraints.Comparator.</p>

<b>Signature:</b>

```go
func HeapSort[T any](slice []T, comparator lancetconstraints.Comparator)
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
<p>Sort slice with merge sort algorithm. Param comparator should implements lancetconstraints.Comparator.</p>

<b>Signature:</b>

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
<p>Sort slice with count sort algorithm. Param comparator should implements lancetconstraints.Comparator.</p>

<b>Signature:</b>

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
<p>BinarySearch search for target within a sorted slice, recursive call itself. If a target is found, the index of the target is returned. Else the function return -1.</p>

<b>Signature:</b>

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
<p>BinaryIterativeSearch search for target within a sorted slice, recursive call itself. If a target is found, the index of the target is returned. Else the function return -1.</p>

<b>Signature:</b>

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
<p>LinearSearch Simple linear search algorithm that iterates over all elements of an slice. If a target is found, the index of the target is returned. Else the function return -1.</p>

<b>Signature:</b>

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
<p>LRUCache implements mem cache with lru.</p>

<b>Signature:</b>

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
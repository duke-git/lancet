# List
List is a linear table, implemented with slice.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/datastructure/list/list.go](https://github.com/duke-git/lancet/blob/main/datastructure/list/list.go)


<div STYLE="page-break-after: always;"></div>

## Usage
```go
import (
    list "github.com/duke-git/lancet/v2/datastructure/list"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

- [NewList](#NewList)
- [Contain](#Contain)
- [Data](#Data)
- [ValueOf](#ValueOf)
- [IndexOf](#IndexOf)
- [LastIndexOf](#LastIndexOf)
- [IndexOfFunc](#IndexOfFunc)
- [LastIndexOfFunc](#LastIndexOfFunc)
- [Push](#Push)
- [PopFirst](#PopFirst)
- [PopLast](#PopLast)
- [DeleteAt](#DeleteAt)
- [InsertAt](#InsertAt)
- [UpdateAt](#UpdateAt)
- [Equal](#Equal)
- [IsEmpty](#IsEmpty)
- [Clear](#Clear)
- [Clone](#Clone)
- [Merge](#Merge)
- [Size](#Size)
- [Cap](#Cap)
- [Swap](#Swap)
- [Reverse](#Reverse)
- [Unique](#Unique)
- [Union](#Union)
- [Intersection](#Intersection)
- [Difference](#Difference)
- [SymmetricDifference](#SymmetricDifference)
- [RetainAll](#RetainAll)
- [DeleteAll](#DeleteAll)
- [ForEach](#ForEach)
- [Iterator](#Iterator)
- [ListToMap](#ListToMap)
- [SubList](#SubList)
- [DeleteIf](#DeleteIf)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="NewList">NewList</span>
<p>List is a linear table, implemented with slice.
NewList function return a list pointer</p>

<b>Signature:</b>

```go
type List[T any] struct {
    data []T
}
func NewList[T any](data []T) *List[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})
    fmt.Println(li)
}
```



### <span id="Contain">Contain</span>
<p>Check if the value in the list or not</p>

<b>Signature:</b>

```go
func (l *List[T]) Contain(value T) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})

    fmt.Println(li.Contain(1)) //true
    fmt.Println(li.Contain(0)) //false
}
```




### <span id="Data">Data</span>
<p>Return slice of list data</p>

<b>Signature:</b>

```go
func (l *List[T]) Data() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})
    data := li.Data()

    fmt.Println(data) //[]int{1, 2, 3}
}
```




### <span id="ValueOf">ValueOf</span>
<p>Return the value pointer at index in list</p>

<b>Signature:</b>

```go
func (l *List[T]) ValueOf(index int) (*T, bool)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})
    v, ok := li.ValueOf(0)

    fmt.Println(*v) //1
    fmt.Println(ok) //true
}
```




### <span id="IndexOf">IndexOf</span>
<p>Returns the index of value in the list. if not found return -1</p>

<b>Signature:</b>

```go
func (l *List[T]) IndexOf(value T) int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})

    fmt.Println(li.IndexOf(1)) //0
    fmt.Println(li.IndexOf(0)) //-1
}
```

### <span id="LastIndexOf">LastIndexOf</span>
<p> Returns the index of the last occurrence of the value in this list if not found return -1</p>

<b>Signature:</b>

```go
func (l *List[T]) LastIndexOf(value T) int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3, 1})

    fmt.Println(li.LastIndexOf(1)) // 3
    fmt.Println(li.LastIndexOf(0)) //-1
}
```

### <span id="IndexOfFunc">IndexOfFunc</span>
<p> IndexOfFunc returns the first index satisfying f(v). if not found return -1</p>

<b>Signature:</b>

```go
func (l *List[T]) IndexOfFunc(f func(T) bool) int 
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})

    fmt.Println(li.IndexOfFunc(func(a int) bool { return a == 1 })) //0
    fmt.Println(li.IndexOfFunc(func(a int) bool { return a == 0 })) //-1
}
```

### <span id="LastIndexOfFunc">LastIndexOfFunc</span>
<p>LastIndexOfFunc returns the index of the last occurrence of the value in this list satisfying f(data[i]). if not found return -1</p>

<b>Signature:</b>

```go
func (l *List[T]) LastIndexOfFunc(f func(T) bool) int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3, 1})

    fmt.Println(li.LastIndexOfFunc(func(a int) bool { return a == 1 })) // 3
    fmt.Println(li.LastIndexOfFunc(func(a int) bool { return a == 0 })) //-1
}
```



### <span id="Push">Push</span>
<p>Append value to the list</p>

<b>Signature:</b>

```go
func (l *List[T]) Push(value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})
    li.Push(4)

    fmt.Println(li.Data()) //[]int{1, 2, 3, 4}
}
```




### <span id="PopFirst">PopFirst</span>
<p>Delete the first value of list and return it</p>

<b>Signature:</b>

```go
func (l *List[T]) PopFirst() (*T, bool)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})
    v, ok := li.PopFirst()

    fmt.Println(*v) //1
    fmt.Println(ok) //true
    fmt.Println(li.Data()) //2, 3
}
```





### <span id="PopLast">PopFirst</span>
<p>Delete the last value of list and return it</p>

<b>Signature:</b>

```go
func (l *List[T]) PopLast() (*T, bool)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})
    v, ok := li.PopLast()

    fmt.Println(*v) //3
    fmt.Println(ok) //true
    fmt.Println(li.Data()) //1, 2
}
```




### <span id="DeleteAt">DeleteAt</span>
<p>Delete the value of list at index, if index is not between 0 and length of list data, do nothing</p>

<b>Signature:</b>

```go
func (l *List[T]) DeleteAt(index int)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3, 4})

    li.DeleteAt(-1)
    fmt.Println(li.Data()) //1,2,3,4

    li.DeleteAt(4)
    fmt.Println(li.Data()) //1,2,3,4

    li.DeleteAt(0)
    fmt.Println(li.Data()) //2,3,4

    li.DeleteAt(2)
    fmt.Println(li.Data()) //2,3
}
```




### <span id="InsertAt">InsertAt</span>
<p>Insert value into list at index, if index is not between 0 and length of list data, do nothing</p>

<b>Signature:</b>

```go
func (l *List[T]) InsertAt(index int, value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})

    li.InsertAt(-1, 0)
    fmt.Println(li.Data()) //1,2,3

    li.InsertAt(4, 0)
    fmt.Println(li.Data()) //1,2,3

    li.InsertAt(3, 4)
    fmt.Println(li.Data()) //1,2,3,4

    // li.InsertAt(2, 4)
    // fmt.Println(li.Data()) //1,2,4,3
}
```



### <span id="UpdateAt">UpdateAt</span>
<p>Update value of list at index, index shoud between 0 and list size - 1</p>

<b>Signature:</b>

```go
func (l *List[T]) UpdateAt(index int, value T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})

    li.UpdateAt(-1, 0)
    fmt.Println(li.Data()) //1,2,3

    li.UpdateAt(2, 4)
    fmt.Println(li.Data()) //1,2,4

    li.UpdateAt(3, 5)
    fmt.Println(li.Data()) //1,2,4
}
```


### <span id="Equal">Equal</span>
<p>Compare a list to another list, use reflect.DeepEqual on every element</p>

<b>Signature:</b>

```go
func (l *List[T]) Equal(other *List[T]) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li1 := list.NewList([]int{1, 2, 3, 4})
    li2 := list.NewList([]int{1, 2, 3, 4})
    li3 := list.NewList([]int{1, 2, 3})

    fmt.Println(li1.Equal(li2)) //true
    fmt.Println(li1.Equal(li3)) //false
}
```



### <span id="IsEmpty">IsEmpty</span>
<p>Check if a list is empty or not</p>

<b>Signature:</b>

```go
func (l *List[T]) IsEmpty() bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li1 := list.NewList([]int{1, 2, 3})
    li2 := list.NewList([]int{})

    fmt.Println(li1.IsEmpty()) //false
    fmt.Println(li2.IsEmpty()) //true
}
```




### <span id="Clear">Clear</span>
<p>Clear the data of list</p>

<b>Signature:</b>

```go
func (l *List[T]) Clear()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})
    li.Clear()

    fmt.Println(li.Data()) // empty
}
```



### <span id="Clone">Clone</span>
<p>Return a copy of list</p>

<b>Signature:</b>

```go
func (l *List[T]) Clone() *List[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3})
    cloneList := li.Clone()

    fmt.Println(cloneList.Data()) // 1,2,3
}
```




### <span id="Merge">Merge</span>
<p>Merge two list, return new list, don't change original list</p>

<b>Signature:</b>

```go
func (l *List[T]) Merge(other *List[T]) *List[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li1 := list.NewList([]int{1, 2, 3, 4})
    li2 := list.NewList([]int{4, 5, 6})
    li3 := li1.Merge(li2)

    fmt.Println(li3.Data()) //1, 2, 3, 4, 4, 5, 6
}
```



### <span id="Size">Size</span>
<p>Return number of list data items</p>

<b>Signature:</b>

```go
func (l *List[T]) Size() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3, 4})

    fmt.Println(li.Size()) //4
}
```




### <span id="Cap">Cap</span>
<p>Cap return cap of the inner data</p>

<b>Signature:</b>

```go
func (l *List[T]) Cap() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    data := make([]int, 0, 100)
    
    li := list.NewList(data)

    fmt.Println(li.Cap()) // 100
}
```




### <span id="Swap">Swap</span>
<p>Swap the value at two index in list</p>

<b>Signature:</b>

```go
func (l *List[T]) Swap(i, j int)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3, 4})
    li.Swap(0, 3)

    fmt.Println(li.Data()) //4, 2, 3, 1
}
```




### <span id="Reverse">Reverse</span>
<p>Reverse the data item order of list</p>

<b>Signature:</b>

```go
func (l *List[T]) Reverse()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 3, 4})
    li.Reverse()

    fmt.Println(li.Data()) //4, 3, 2, 1
}
```




### <span id="Unique">Unique</span>
<p>Remove duplicate items in list</p>

<b>Signature:</b>

```go
func (l *List[T]) Unique()
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li := list.NewList([]int{1, 2, 2, 3, 4})
    li.Unique()

    fmt.Println(li.Data()) //1,2,3,4
}
```




### <span id="Union">Union</span>
<p>Creates a new list contain all elements in list l and other, remove duplicate element</p>

<b>Signature:</b>

```go
func (l *List[T]) Union(other *List[T]) *List[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li1 := list.NewList([]int{1, 2, 3, 4})
    li2 := list.NewList([]int{4, 5, 6})
    li3 := li1.Union(li2)

    fmt.Println(li3.Data()) //1,2,3,4,5,6
}
```




### <span id="Intersection">Intersection</span>
<p>Creates a new list whose element both be contained in list l and other</p>

<b>Signature:</b>

```go
func (l *List[T]) Intersection(other *List[T]) *List[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    li1 := list.NewList([]int{1, 2, 3, 4})
    li2 := list.NewList([]int{4, 5, 6})
    li3 := li1.Intersection(li2)

    fmt.Println(li3.Data()) //4
}
```



### <span id="Difference">Difference</span>
<p>Return a list whose element in the original list, not in the given list.</p>

<b>Signature:</b>

```go
func (l *List[T]) Difference(other *List[T]) *List[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    list1 := NewList([]int{1, 2, 3})
    list2 := NewList([]int{1, 2, 4})

    list3 := list1.Intersection(list2)

    fmt.Println(list3.Data()) //3
}
```


### <span id="SymmetricDifference">SymmetricDifference</span>
<p>Oppoiste operation of intersection function.</p>

<b>Signature:</b>

```go
func (l *List[T]) SymmetricDifference(other *List[T]) *List[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    list1 := NewList([]int{1, 2, 3})
    list2 := NewList([]int{1, 2, 4})

    list3 := list1.Intersection(list2)

    fmt.Println(list3.Data()) //3, 4
}
```


### <span id="RetainAll">RetainAll</span>
<p>Retains only the elements in this list that are contained in the given list.</p>

<b>Signature:</b>

```go
func (l *List[T]) RetainAll(list *List[T]) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    list := NewList([]int{1, 2, 3, 4})
    list1 := NewList([]int{1, 2, 3, 4})
    list2 := NewList([]int{1, 2, 3, 4})

    retain := NewList([]int{1, 2})
    retain1 := NewList([]int{2, 3})
    retain2 := NewList([]int{1, 2, 5})

    list.RetainAll(retain)
    list1.RetainAll(retain1)
    list2.RetainAll(retain2)

    fmt.Println(list.Data()) //1, 2
    fmt.Println(list1.Data()) //2, 3
    fmt.Println(list2.Data()) //1, 2
}
```


### <span id="DeleteAll">DeleteAll</span>
<p>Removes from this list all of its elements that are contained in the given list.</p>

<b>Signature:</b>

```go
func (l *List[T]) DeleteAll(list *List[T]) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    list := NewList([]int{1, 2, 3, 4})
    list1 := NewList([]int{1, 2, 3, 4})
    list2 := NewList([]int{1, 2, 3, 4})

    del := NewList([]int{1})
    del1 := NewList([]int{2, 3})
    del2 := NewList([]int{1, 2, 5})

    list.DeleteAll(del)
    list1.DeleteAll(del1)
    list2.DeleteAll(del2)

    fmt.Println(list.Data()) //2,3,4
    fmt.Println(list1.Data()) //1,4
    fmt.Println(list2.Data()) //3,4
}
```


### <span id="ForEach">ForEach</span>
<p>Performs the given action for each element of the list.</p>

<b>Signature:</b>

```go
func (l *List[T]) ForEach(consumer func(T))
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    list := NewList([]int{1, 2, 3, 4})

    result := make([]int, 0)
    list.ForEach(func(i int) {
        result = append(result, i)
    })

    fmt.Println(result.Data()) //1,2,3,4
}
```


### <span id="Iterator">Iterator</span>
<p>Returns an iterator over the elements in this list in proper sequence.</p>

<b>Signature:</b>

```go
func (l *List[T]) Iterator() iterator.Iterator[T]
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    list := NewList([]int{1, 2, 3, 4})

    iterator := list.Iterator()

    result := make([]int, 0)
    for iterator.HasNext() {
        item, _ := iterator.Next()
        result = append(result, item)
    }

    fmt.Println(result.Data()) //1,2,3,4
}
```


### <span id="ListToMap">ListToMap</span>
<p>Converts a list to a map based on iteratee function.</p>

<b>Signature:</b>

```go
func ListToMap[T any, K comparable, V any](list *List[T], iteratee func(T) (K, V)) map[K]V
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    list := NewList([]int{1, 2, 3, 4})

    result := ListToMap(list, func(n int) (int, bool) {
        return n, n > 1
    })

    fmt.Println(result) //map[int]bool{1: false, 2: true, 3: true, 4: true}
}
```

### <span id="SubList">SubList</span>
<p>SubList returns a sub list of the original list between the specified fromIndex, inclusive, and toIndex, exclusive.</p>

<b>Signature:</b>

```go
func (l *List[T]) SubList(fromIndex, toIndex int) *List[T] 
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    l := list.NewList([]int{1, 2, 3, 4, 5, 6})
   
    fmt.Println(l.SubList(2, 5)) // []int{3, 4, 5}
}
```




### <span id="DeleteIf">DeleteIf</span>
<p>DeleteIf delete all satisfying f(data[i]), returns count of removed elements</p>

<b>Signature:</b>

```go
func (l *List[T]) DeleteIf(f func(T) bool) int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    list "github.com/duke-git/lancet/v2/datastructure/list"
)

func main() {
    l := list.NewList([]int{1, 1, 1, 1, 2, 3, 1, 1, 4, 1, 1, 1, 1, 1, 1})

    fmt.Println(l.DeleteIf(func(a int) bool { return a == 1 })) // 12 
    fmt.Println(l.Data()) // []int{2, 3, 4}
}
```

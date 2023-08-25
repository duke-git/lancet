# List
List是线性表数据结构, 用go切片实现。

<div STYLE="page-break-after: always;"></div>

## 源码

- [https://github.com/duke-git/lancet/blob/main/datastructure/list/list.go](https://github.com/duke-git/lancet/blob/main/datastructure/list/list.go)


<div STYLE="page-break-after: always;"></div>

## 用法
```go
import (
    "github.com/duke-git/lancet/v2/datastructure"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

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

## 文档

### <span id="NewList">NewList</span>
<p>返回List指针实例</p>

<b>函数签名:</b>

```go
type List[T any] struct {
    data []T
}
func NewList[T any](data []T) *List[T]
```
<b>示例:</b>

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
<p>判断列表中是否包含特定值</p>

<b>函数签名:</b>

```go
func (l *List[T]) Contain(value T) bool
```
<b>示例:</b>

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
<p>返回List中所有数据（切片）</p>

<b>函数签名:</b>

```go
func (l *List[T]) Data() []T
```
<b>示例:</b>

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
<p>返回列表中索引处的值指针</p>

<b>函数签名:</b>

```go
func (l *List[T]) ValueOf(index int) (*T, bool)
```
<b>示例:</b>

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
<p>返回列表中值的索引，如果没有找到返回-1</p>

<b>函数签名:</b>

```go
func (l *List[T]) IndexOf(value T) int
```
<b>示例:</b>

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
<p>返回列表中最后一次出现的值的索引。如果未找到，则返回-1</p>

<b>函数签名:</b>

```go
func (l *List[T]) LastIndexOf(value T) int
```
<b>示例:</b>

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
<p>返回第一个符合函数条件的元素的索引。如果未找到，则返回-1</p>

<b>函数签名:</b>

```go
func (l *List[T]) IndexOfFunc(f func(T) bool) int 
```
<b>示例:</b>

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
<p>返回最后一个符合函数条件的元素的索引。如果未找到，则返回-1</p>

<b>函数签名:</b>

```go
func (l *List[T]) LastIndexOfFunc(f func(T) bool) int
```
<b>示例:</b>

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
<p>将值附加到列表末尾</p>

<b>函数签名:</b>

```go
func (l *List[T]) Push(value T)
```
<b>示例:</b>

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
<p>删除列表的第一个值并返回该值</p>

<b>函数签名:</b>

```go
func (l *List[T]) PopFirst() (*T, bool)
```
<b>示例:</b>

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
<p>删除列表的最后一个值并返回该值</p>

<b>函数签名:</b>

```go
func (l *List[T]) PopLast() (*T, bool)
```
<b>示例:</b>

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
<p>删除索引处列表的值，如果索引不在0和列表数据长度之间，则不执行任何操作</p>

<b>函数签名:</b>

```go
func (l *List[T]) DeleteAt(index int)
```
<b>示例:</b>

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
<p>在索引处插入值到列表中，如果索引不在 0 和列表数据长度之间，则不执行任何操作</p>

<b>函数签名:</b>

```go
func (l *List[T]) InsertAt(index int, value T)
```
<b>示例:</b>

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
<p>更新索引处列表的值，索引应该在0和列表数据长度-1之间</p>

<b>函数签名:</b>

```go
func (l *List[T]) UpdateAt(index int, value T)
```
<b>示例:</b>

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
<p>比较一个列表和另一个列表，在每个元素上使用 reflect.DeepEqual</p>

<b>函数签名:</b>

```go
func (l *List[T]) Equal(other *List[T]) bool
```
<b>示例:</b>

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
<p>判断列表是否为空</p>

<b>函数签名:</b>

```go
func (l *List[T]) IsEmpty() bool
```
<b>示例:</b>

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
<p>清空列表数据</p>

<b>函数签名:</b>

```go
func (l *List[T]) Clear()
```
<b>示例:</b>

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
<p>返回列表的一个拷贝</p>

<b>函数签名:</b>

```go
func (l *List[T]) Clone() *List[T]
```
<b>示例:</b>

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
<p>合并两个列表，返回新的列表</p>

<b>函数签名:</b>

```go
func (l *List[T]) Merge(other *List[T]) *List[T]
```
<b>示例:</b>

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
<p>返回列表数据项的数量</p>

<b>函数签名:</b>

```go
func (l *List[T]) Size() int
```
<b>示例:</b>

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
<p>返回列表数据容量</p>

<b>函数签名:</b>

```go
func (l *List[T]) Cap() int
```
<b>示例:</b>

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
<p>交换列表中两个索引位置的值</p>

<b>函数签名:</b>

```go
func (l *List[T]) Swap(i, j int)
```
<b>示例:</b>

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
<p>反转列表的数据项顺序</p>

<b>函数签名:</b>

```go
func (l *List[T]) Reverse()
```
<b>示例:</b>

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
<p>列表去除重复数据项</p>

<b>函数签名:</b>

```go
func (l *List[T]) Unique()
```
<b>示例:</b>

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
<p>两个列表取并集，去除重复数据项</p>

<b>函数签名:</b>

```go
func (l *List[T]) Union(other *List[T]) *List[T]
```
<b>示例:</b>

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
<p>两个列表取交集</p>

<b>函数签名:</b>

```go
func (l *List[T]) Intersection(other *List[T]) *List[T]
```
<b>示例:</b>

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
<p>差集运算。</p>

<b>函数签名:</b>

```go
func (l *List[T]) Difference(other *List[T]) *List[T]
```
<b>示例:</b>

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
<p>对称差集运算。</p>

<b>函数签名:</b>

```go
func (l *List[T]) SymmetricDifference(other *List[T]) *List[T]
```
<b>示例:</b>

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
<p>仅保留列表中包含在给定列表中的元素。</p>

<b>函数签名:</b>

```go
func (l *List[T]) RetainAll(list *List[T]) bool
```
<b>示例:</b>

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
<p>从列表中删除给定列表中包含的所有元素。</p>

<b>函数签名:</b>

```go
func (l *List[T]) DeleteAll(list *List[T]) bool
```
<b>示例:</b>

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
<p>对列表的每个元素执行给定的操作。</p>

<b>函数签名:</b>

```go
func (l *List[T]) ForEach(consumer func(T))
```
<b>示例:</b>

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
<p>按顺序返回列表中元素的迭代器。</p>

<b>函数签名:</b>

```go
func (l *List[T]) Iterator() iterator.Iterator[T]
```
<b>示例:</b>

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
<p>基于iteratee函数将列表转换为映射map。</p>

<b>函数签名:</b>

```go
func ListToMap[T any, K comparable, V any](list *List[T], iteratee func(T) (K, V)) map[K]V
```
<b>示例:</b>

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
<p>返回指定的fromIndex（包含）和toIndex（不包含）之间的原始列表的子列表。</p>

<b>函数签名:</b>

```go
func (l *List[T]) SubList(fromIndex, toIndex int) *List[T] 
```
<b>示例:</b>

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
<p>删除列表中所有符合函数（调用函数返回true)的元素，返回删除元素的数量</p>

<b>函数签名:</b>

```go
func (l *List[T]) DeleteIf(f func(T) bool) int
```
<b>示例:</b>

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
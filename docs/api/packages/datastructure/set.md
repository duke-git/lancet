# Set

集合数据结构，类似列表。Set中元素不重复。

<div STYLE="page-break-after: always;"></div>

## 源码

-   [https://github.com/duke-git/lancet/blob/main/datastructure/set/set.go](https://github.com/duke-git/lancet/blob/main/datastructure/set/set.go)

<div STYLE="page-break-after: always;"></div>

## 用法

```go
import (
    set "github.com/duke-git/lancet/v2/datastructure/set"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

-   [New](#New)
-   [FromSlice](#FromSlice)
-   [Values<sup>deprecated</sup>](#Values)
-   [Add](#Add)
-   [AddIfNotExist](#AddIfNotExist)
-   [AddIfNotExistBy](#AddIfNotExistBy)
-   [Delete](#Delete)
-   [Contain](#Contain)
-   [ContainAll](#ContainAll)
-   [Clone](#Clone)
-   [Size](#Size)
-   [Equal](#Equal)
-   [Iterate](#Iterate)
-   [IsEmpty](#IsEmpty)
-   [Union](#Union)
-   [Intersection](#Intersection)
-   [SymmetricDifference](#SymmetricDifference)
-   [Minus](#Minus)
-   [Pop](#Pop)
-   [ToSlice](#ToSlice)
-   [ToSortedSlice](#ToSortedSlice)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="New">New</span>

<p>返回Set结构体对象</p>

<b>函数签名:</b>

```go
type Set[T comparable] map[T]struct{}
func New[T comparable](items ...T) Set[T]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.New[int](1,2,2,3)
    fmt.Println(st.Values()) //1,2,3
}
```

### <span id="FromSlice">FromSlice</span>

<p>基于切片创建集合</p>

<b>函数签名:</b>

```go
func FromSlice[T comparable](items []T) Set[T]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.FromSlice([]int{1, 2, 2, 3})
    fmt.Println(st.Values()) //1,2,3
}
```

### <span id="Values">Values</span>

<p>获取集合中所有元素的切片。</p>

> ⚠️ 本函数已弃用，使用`ToSlice`代替。

<b>函数签名:</b>

```go
func (s Set[T]) Values() []T
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.New[int](1,2,2,3)
    fmt.Println(st.Values()) //1,2,3
}
```

### <span id="Add">Add</span>

<p>向集合中添加元素</p>

<b>函数签名:</b>

```go
func (s Set[T]) Add(items ...T)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.New[int]()
    st.Add(1, 2, 3)

    fmt.Println(st.Values()) //1,2,3
}
```

### <span id="AddIfNotExist">AddIfNotExist</span>

<p>如果集合中不存在元素，则添加该元素返回true, 如果集合中存在元素, 不做任何操作，返回false</p>

<b>函数签名:</b>

```go
func (s Set[T]) AddIfNotExist(item T) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.New[int]()
    st.Add(1, 2, 3)

    r1 := st.AddIfNotExist(1)
    r2 := st.AddIfNotExist(4)

    fmt.Println(r1) // false
    fmt.Println(r2) // true
    fmt.Println(st.Values()) // 1,2,3,4
}
```

### <span id="AddIfNotExistBy">AddIfNotExistBy</span>

<p>根据checker函数判断元素是否在集合中，如果集合中不存在元素且checker返回true，则添加该元素返回true, 否则不做任何操作，返回false</p>

<b>函数签名:</b>

```go
func (s Set[T]) AddIfNotExistBy(item T, checker func(element T) bool) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.New[int]()
    st.Add(1, 2)

    ok := st.AddIfNotExistBy(3, func(val int) bool {
        return val%2 != 0
    })
    fmt.Println(ok) // true


    notOk := st.AddIfNotExistBy(4, func(val int) bool {
        return val%2 != 0
    })
    fmt.Println(notOk) // false

    fmt.Println(st.Values()) // 1, 2, 3
}
```

### <span id="Delete">Delete</span>

<p>删除集合中元素</p>

<b>函数签名:</b>

```go
func (s Set[T]) Delete(items ...T)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.New[int]()
    st.Add(1, 2, 3)

    set.Delete(3)
    fmt.Println(st.Values()) //1,2
}
```

### <span id="Contain">Contain</span>

<p>判断集合是否包含某个值</p>

<b>函数签名:</b>

```go
func (s Set[T]) Contain(item T) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.New[int]()
    st.Add(1, 2, 3)

    fmt.Println(st.Contain(1)) //true
    fmt.Println(st.Contain(4)) //false
}
```

### <span id="ContainAll">ContainAll</span>

<p>判断集合是否包含另一个集合</p>

<b>函数签名:</b>

```go
func (s Set[T]) ContainAll(other Set[T]) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)
    set2 := set.New(1, 2)
    set3 := set.New(1, 2, 3, 4)

    fmt.Println(set1.ContainAll(set2)) //true
    fmt.Println(set1.ContainAll(set3)) //false
}
```

### <span id="Size">Size</span>

<p>获取集合中元素的个数</p>

<b>函数签名:</b>

```go
func (s Set[T]) Size() int
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)

    fmt.Println(set1.Size()) //3
}
```

### <span id="Clone">Clone</span>

<p>克隆一个集合</p>

<b>函数签名:</b>

```go
func (s Set[T]) Clone() Set[T]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)
    set2 := set1.Clone()

    fmt.Println(set1.Size() == set2.Size()) //true
    fmt.Println(set1.ContainAll(set2)) //true
}
```

### <span id="Equal">Equal</span>

<p>比较两个集合是否相等，包含相同元素为相等</p>

<b>函数签名:</b>

```go
func (s Set[T]) Equal(other Set[T]) bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)
    set2 := set.New(1, 2, 3)
    set3 := set.New(1, 2, 3, 4)

    fmt.Println(set1.Equal(set2)) //true
    fmt.Println(set1.Equal(set3)) //false
}
```

### <span id="Iterate">Iterate</span>

<p>迭代结合，在每个元素上调用函数</p>

<b>函数签名:</b>

```go
func (s Set[T]) Iterate(fn func(item T))
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)
    arr := []int{}
    set.Iterate(func(item int) {
        arr = append(arr, item)
    })

    fmt.Println(arr) //1,2,3
}
```

### <span id="EachWithBreak">EachWithBreak</span>

<p>遍历集合的元素并为每个元素调用iteratee函数，当iteratee函数返回false时，终止遍历。</p>

<b>函数签名:</b>

```go
func (s Set[T]) EachWithBreak(iteratee func(item T) bool)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    s := set.New(1, 2, 3, 4, 5)

    var sum int

    s.EachWithBreak(func(n int) bool {
        if n > 3 {
            return false
        }
        sum += n
        return true
    })

    fmt.Println(sum) //6
}
```

### <span id="IsEmpty">IsEmpty</span>

<p>判断集合是否为空</p>

<b>函数签名:</b>

```go
func (s Set[T]) IsEmpty() bool
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)
    set2 := set.New()

    fmt.Println(set1.IsEmpty()) //false
    fmt.Println(set2.IsEmpty()) //true
}
```

### <span id="Union">Union</span>

<p>求两个集合的并集</p>

<b>函数签名:</b>

```go
func (s Set[T]) Union(other Set[T]) Set[T]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)
    set2 := set.New(2, 3, 4, 5)
    set3 := set1.Union(set2)

    fmt.Println(set3.Values()) //1,2,3,4,5
}
```

### <span id="Intersection">Intersection</span>

<p>求两个集合的交集</p>

<b>函数签名:</b>

```go
func (s Set[T]) Intersection(other Set[T]) Set[T]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)
    set2 := set.New(2, 3, 4, 5)
    set3 := set1.Intersection(set2)

    fmt.Println(set3.Values()) //2,3
}
```

### <span id="SymmetricDifference">SymmetricDifference</span>

<p>返回一个集合，其中元素在第一个集合或第二个集合中，且不同时存在于两个集合中</p>

<b>函数签名:</b>

```go
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)
    set2 := set.New(2, 3, 4, 5)
    set3 := set1.SymmetricDifference(set2)

    fmt.Println(set3.Values()) //1,4,5
}
```

### <span id="Minus">Minus</span>

<p>创建一个集合，其元素在原始集中但不在比较集中</p>

<b>函数签名:</b>

```go
func (s Set[T]) Minus(comparedSet Set[T]) Set[T]
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.New(1, 2, 3)
    set2 := set.New(2, 3, 4, 5)
    set3 := set.New(2, 3)

    res1 := set1.Minus(set2)
    fmt.Println(res1.Values()) //1

    res2 := set2.Minus(set3)
    fmt.Println(res2.Values()) //4,5
}
```

### <span id="Pop">Pop</span>

<p>删除并返回集合中的顶部元素</p>

<b>函数签名:</b>

```go
func (s Set[T]) Pop() (v T, ok bool)
```

<b>示例:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    s := set.New[int]()
    s.Add(1)
    s.Add(2)
    s.Add(3)

    val, ok = s.Pop()

    fmt.Println(val) // 3
    fmt.Println(ok) // true
}
```

### <span id="ToSlice">ToSlice</span>

<p>以切片的形式返回集合中所有的元素（无序）</p>

<b>函数签名:</b>

```go
func (s Set[T]) ToSlice() (v T, ok bool)
```

<b>示例:</b>

```go
func main() {
    s := set.New(1, 2, 3, 4, 5)

    val := s.ToSlice()
    fmt.Println(val) // [2 3 4 5 1]
}
```

### <span id="ToSortedSlice">ToSortedSlice</span>

<p>以切片的形式返回集合中所有的元素（按给定的规则排序）</p>

<b>函数签名:</b>

```go
func (s Set[T]) ToSortedSlice() (v T, ok bool)
```

<b>示例:</b>

```go
func main() {
    s1 := set.New(1, 2, 3, 4, 5)
    type Person struct {
		Name string
		Age  int
	}
	s2 := FromSlice([]Person{{"Tom", 20}, {"Jerry", 18}, {"Spike", 25}})

    res1 := s1.ToSortedSlice(func(v1, v2 int) bool {
		return v1 < v2
	})
    
    res2 := s2.ToSortedSlice(func(v1, v2 Person) bool {
		return v1.Age < v2.Age
	})
    
    fmt.Println(res1) // [1 2 3 4 5]
    fmt.Println(res2) // [{Jerry 18} {Tom 20} {Spike 25}]
}
```

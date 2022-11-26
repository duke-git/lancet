# Set
Set集合数据结构，类似列表。Set中元素不重复。

<div STYLE="page-break-after: always;"></div>

## 源码

- [https://github.com/duke-git/lancet/blob/main/datastructure/set/set.go](https://github.com/duke-git/lancet/blob/main/datastructure/set/set.go)


<div STYLE="page-break-after: always;"></div>

## 用法
```go
import (
    set "github.com/duke-git/lancet/v2/datastructure/set"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

- [NewSet](#NewSet)
- [Values](#Values)
- [Add](#Add)
- [Delete](#Delete)
- [Contain](#Contain)
- [ContainAll](#ContainAll)
- [Clone](#Clone)
- [Size](#Size)
- [Equal](#Equal)
- [Iterate](#Iterate)
- [IsEmpty](#IsEmpty)
- [Union](#Union)
- [Intersection](#Intersection)
- [SymmetricDifference](#SymmetricDifference)
- [Minus](#Minus)



<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="NewSet">NewSet</span>
<p>返回Set结构体对象</p>

<b>函数签名:</b>

```go
type Set[T comparable] map[T]bool
func NewSet[T comparable](values ...T) Set[T]
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.NewSet[int](1,2,2,3)
    fmt.Println(st.Values()) //1,2,3
}
```




### <span id="Values">Values</span>
<p>获取集合中所有元素的切片</p>

<b>函数签名:</b>

```go
func (s Set[T]) Values() []T
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.NewSet[int](1,2,2,3)
    fmt.Println(st.Values()) //1,2,3
}
```




### <span id="Add">Add</span>
<p>向集合中添加元素</p>

<b>函数签名:</b>

```go
func (s Set[T]) Add(values ...T)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.NewSet[int]()
    st.Add(1, 2, 3)

    fmt.Println(st.Values()) //1,2,3
}
```



### <span id="Delete">Delete</span>
<p>删除集合中元素</p>

<b>函数签名:</b>

```go
func (s Set[T]) Delete(values ...T)
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.NewSet[int]()
    st.Add(1, 2, 3)

    set.Delete(3)
    fmt.Println(st.Values()) //1,2
}
```



### <span id="Contain">Contain</span>
<p>判断集合是否包含某个值</p>

<b>函数签名:</b>

```go
func (s Set[T]) Contain(value T) bool
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.NewSet[int]()
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
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
	set2 := set.NewSet(1, 2)
	set3 := set.NewSet(1, 2, 3, 4)

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
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)

    fmt.Println(set1.Size()) //3
}
```



### <span id="Clone">Clone</span>
<p>克隆一个集合</p>

<b>函数签名:</b>

```go
func (s Set[T]) Clone() Set[T]
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
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
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
    set2 := set.NewSet(1, 2, 3)
    set3 := set.NewSet(1, 2, 3, 4)

    fmt.Println(set1.Equal(set2)) //true
    fmt.Println(set1.Equal(set3)) //false
}
```



### <span id="Iterate">Iterate</span>
<p>迭代结合，在每个元素上调用函数</p>

<b>函数签名:</b>

```go
func (s Set[T]) Iterate(fn func(value T))
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
    arr := []int{}
    set.Iterate(func(value int) {
        arr = append(arr, value)
    })

    fmt.Println(arr) //1,2,3
}
```



### <span id="IsEmpty">IsEmpty</span>
<p>判断集合是否为空</p>

<b>函数签名:</b>

```go
func (s Set[T]) IsEmpty() bool
```
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
    set2 := set.NewSet()

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
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
    set2 := set.NewSet(2, 3, 4, 5)
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
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
    set2 := set.NewSet(2, 3, 4, 5)
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
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
	set2 := set.NewSet(2, 3, 4, 5)
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
<b>例子:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
	set2 := set.NewSet(2, 3, 4, 5)
	set3 := set.NewSet(2, 3)

    res1 := set1.Minus(set2)
    fmt.Println(res1.Values()) //1

    res2 := set2.Minus(set3)
    fmt.Println(res2.Values()) //4,5
}
```




# Set
Set is a data container, like list, but elements of set is not duplicate.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/datastructure/set/set.go](https://github.com/duke-git/lancet/blob/main/datastructure/set/set.go)


<div STYLE="page-break-after: always;"></div>

## Usage
```go
import (
    set "github.com/duke-git/lancet/v2/datastructure/set"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

- [NewSet](#NewSet)
- [Values](#Values)
- [Add](#Add)
- [AddIfNotExist](#AddIfNotExist)
- [AddIfNotExistBy](#AddIfNotExistBy)
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

## Documentation

### <span id="NewSet">NewSet</span>
<p>Make a Set instance</p>

<b>Signature:</b>

```go
type Set[T comparable] map[T]bool
func NewSet[T comparable](items ...T) Set[T]
```
<b>Example:</b>

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
<p>Return slice of all set data</p>

<b>Signature:</b>

```go
func (s Set[T]) Values() []T
```
<b>Example:</b>

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
<p>Add items to set</p>

<b>Signature:</b>

```go
func (s Set[T]) Add(items ...T)
```
<b>Example:</b>

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


### <span id="AddIfNotExist">AddIfNotExist</span>
<p>AddIfNotExist checks if item exists in the set, it adds the item to set and returns true if it does not exist in the set, or else it does nothing and returns false.</p>

<b>Signature:</b>

```go
func (s Set[T]) AddIfNotExist(item T) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.NewSet[int]()
    st.Add(1, 2, 3)

    r1 := st.AddIfNotExist(1)
	r2 := st.AddIfNotExist(4)

    fmt.Println(r1) // false
    fmt.Println(r2) // true
    fmt.Println(st.Values()) // 1,2,3,4
}
```


### <span id="AddIfNotExistBy">AddIfNotExistBy</span>
<p>AddIfNotExistBy checks if item exists in the set and pass the `checker` function it adds the item to set and returns true if it does not exists in the set and function `checker` returns true, or else it does nothing and returns false.</p>

<b>Signature:</b>

```go
func (s Set[T]) AddIfNotExistBy(item T, checker func(element T) bool) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    st := set.NewSet[int]()
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
<p>Delete item in set</p>

<b>Signature:</b>

```go
func (s Set[T]) Delete(items ...T)
```
<b>Example:</b>

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
<p>Check if item is in set or not</p>

<b>Signature:</b>

```go
func (s Set[T]) Contain(item T) bool
```
<b>Example:</b>

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
<p>Checks if set contains another set</p>

<b>Signature:</b>

```go
func (s Set[T]) ContainAll(other Set[T]) bool
```
<b>Example:</b>

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
<p>Get the number of elements in set</p>

<b>Signature:</b>

```go
func (s Set[T]) Size() int
```
<b>Example:</b>

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
<p>Make a copy of set</p>

<b>Signature:</b>

```go
func (s Set[T]) Clone() Set[T]
```
<b>Example:</b>

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
<p>Check if two sets has same elements or not</p>

<b>Signature:</b>

```go
func (s Set[T]) Equal(other Set[T]) bool
```
<b>Example:</b>

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
<p>Call function by every element of set</p>

<b>Signature:</b>

```go
func (s Set[T]) Iterate(fn func(item T))
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    set "github.com/duke-git/lancet/v2/datastructure/set"
)

func main() {
    set1 := set.NewSet(1, 2, 3)
    arr := []int{}
    set.Iterate(func(item int) {
        arr = append(arr, item)
    })

    fmt.Println(arr) //1,2,3
}
```



### <span id="IsEmpty">IsEmpty</span>
<p>Check if the set is empty or not</p>

<b>Signature:</b>

```go
func (s Set[T]) IsEmpty() bool
```
<b>Example:</b>

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
<p>Create a new set contain all element of set s and other</p>

<b>Signature:</b>

```go
func (s Set[T]) Union(other Set[T]) Set[T]
```
<b>Example:</b>

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
<p>Create a new set whose element both be contained in set s and other</p>

<b>Signature:</b>

```go
func (s Set[T]) Intersection(other Set[T]) Set[T]
```
<b>Example:</b>

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
<p>Create a new set whose element is in set1 or set2, but not in both set1 and set2</p>

<b>Signature:</b>

```go
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T]
```
<b>Example:</b>

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
<p>Create an set of whose element in origin set but not in compared set</p>

<b>Signature:</b>

```go
func (s Set[T]) Minus(comparedSet Set[T]) Set[T]
```
<b>Example:</b>

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




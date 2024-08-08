# Set

Set is a data container, like list, but elements of set is not duplicate.

<div STYLE="page-break-after: always;"></div>

## Source

-   [https://github.com/duke-git/lancet/blob/main/datastructure/set/set.go](https://github.com/duke-git/lancet/blob/main/datastructure/set/set.go)

<div STYLE="page-break-after: always;"></div>

## Usage

```go
import (
    set "github.com/duke-git/lancet/v2/datastructure/set"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

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
-   [EachWithBreak](#EachWithBreak)
-   [IsEmpty](#IsEmpty)
-   [Union](#Union)
-   [Intersection](#Intersection)
-   [SymmetricDifference](#SymmetricDifference)
-   [Minus](#Minus)
-   [Pop](#Pop)
-   [ToSlice](#ToSlice)
-   [ToSortedSlice](#ToSortedSlice)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="New">New</span>

<p>Create a set instance</p>

<b>Signature:</b>

```go
type Set[T comparable] map[T]struct{}
func New[T comparable](items ...T) Set[T]
```

<b>Example:</b>

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

<p>Create a set from slice</p>

<b>Signature:</b>

```go
func FromSlice[T comparable](items []T) Set[T]
```

<b>Example:</b>

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

<p>Return slice of all set data.</p>

> ⚠️ This function is deprecated. use `ToSlice` instead.

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
    st := set.New[int](1,2,2,3)
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
    st := set.New[int]()
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
    st := set.New[int]()
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
    st := set.New[int]()
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
    set1 := set.New(1, 2, 3)
    set2 := set.New(1, 2)
    set3 := set.New(1, 2, 3, 4)

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
    set1 := set.New(1, 2, 3)

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
    set1 := set.New(1, 2, 3)
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
    set1 := set.New(1, 2, 3)
    set2 := set.New(1, 2, 3)
    set3 := set.New(1, 2, 3, 4)

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
    set1 := set.New(1, 2, 3)
    arr := []int{}
    set.Iterate(func(item int) {
        arr = append(arr, item)
    })

    fmt.Println(arr) //1,2,3
}
```

### <span id="EachWithBreak">EachWithBreak</span>

<p>Iterates over elements of a set and invokes function for each element, when iteratee return false, will break the for each loop.</p>

<b>Signature:</b>

```go
func (s Set[T]) EachWithBreak(iteratee func(item T) bool)
```

<b>Example:</b>

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
    set1 := set.New(1, 2, 3)
    set2 := set.New()

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
    set1 := set.New(1, 2, 3)
    set2 := set.New(2, 3, 4, 5)
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
    set1 := set.New(1, 2, 3)
    set2 := set.New(2, 3, 4, 5)
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
    set1 := set.New(1, 2, 3)
    set2 := set.New(2, 3, 4, 5)
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

<p>Delete the top element of set then return it, if set is empty, return nil-value of T and false.</p>

<b>Signature:</b>

```go
func (s Set[T]) Pop() (v T, ok bool)
```

<b>Example:</b>

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

<p>returns a slice containing all values of the set.</p>

<b>Signature:</b>

```go
func (s Set[T]) ToSlice() (v T, ok bool)
```

<b>Example:</b>

```go
func main() {
    s := set.New(1, 2, 3, 4, 5)

    val := s.ToSlice()
    fmt.Println(val) // [2 3 4 5 1]
}
```

### <span id="ToSortedSlice">ToSortedSlice</span>

<p>returns a sorted slice containing all values of the set</p>

<b>Signature:</b>

```go
func (s Set[T]) ToSortedSlice() (v T, ok bool)
```

<b>Example:</b>

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

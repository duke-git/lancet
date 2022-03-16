# Slice
Package slice implements some functions to manipulate slice.

<div STYLE="page-break-after: always;"></div>

## Source:

- [https://github.com/duke-git/lancet/blob/main/slice/slice.go](https://github.com/duke-git/lancet/blob/main/slice/slice.go)


<div STYLE="page-break-after: always;"></div>

## Usage:
```go
import (
    "github.com/duke-git/lancet/slice"
)
```

<div STYLE="page-break-after: always;"></div>

## Index
- [Contain](#Contain)
- [ContainSubSlice](#ContainSubSlice)
- [Chunk](#Chunk)
- [Compact](#Compact)
- [Concat](#Concat)
- [Count](#Count)
- [Difference](#Difference)
- [DifferenceBy](#DifferenceBy)
- [DifferenceWith](#DifferenceWith)
- [DeleteAt](#DeleteAt)
- [Drop](#Drop)
- [Every](#Every)
- [Filter](#Filter)
- [Find](#Find)
- [FindLast](#FindLast)
- [FlattenDeep](#FlattenDeep)
- [ForEach](#ForEach)
  
- [GroupBy](#GroupBy)
- [IntSlice](#IntSlice)
- [InterfaceSlice](#InterfaceSlice)
- [Intersection](#Intersection)
- [InsertAt](#InsertAt)
- [Map](#Map)
- [Reverse](#Reverse)
- [Reduce](#Reduce)
- [Shuffle](#Shuffle)
- [SortByField](#SortByField)
- [Some](#Some)
- [StringSlice](#StringSlice)
- [Unique](#Unique)
- [Union](#Union)
- [UpdateAt](#UpdateAt)
- [Without](#Without)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="Contain">Contain</span>
<p>Check if the value is in the slice or not.</p>

<b>Signature:</b>

```go
func Contain[T any](slice []T, value T) bool
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res := slice.Contain([]string{"a", "b", "c"}, "a")
	fmt.Println(res) //true
}
```


### <span id="ContainSubSlice">ContainSubSlice</span>
<p>Check if the slice contain subslice or not.</p>

<b>Signature:</b>

```go
func ContainSubSlice[T any](slice, subslice []T) bool
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res := slice.ContainSubSlice([]string{"a", "b", "c"}, []string{"a", "b"})
	fmt.Println(res) //true
}
```




### <span id="Chunk">Chunk</span>
<p>Creates an slice of elements split into groups the length of `size`.</p>

<b>Signature:</b>

```go
func Chunk[T any](slice []T, size int) [][]T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
	res := slice.Chunk(InterfaceSlice(arr), 3)
	fmt.Println(res) //[][]interface{}{{"a", "b", "c"}, {"d", "e"}}
}
```



### <span id="Compact">Compact</span>
<p>Creates an slice with all falsey values removed. The values false, nil, 0, and "" are falsey.</p>

<b>Signature:</b>

```go
func Compact[T any](slice []T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res := slice.Compact([]int{0, 1, 2, 3})
	fmt.Println(res) //[]int{1, 2, 3}
}
```


### <span id="Concat">Concat</span>
<p>Creates a new slice concatenating slice with any additional slices and/or values.</p>

<b>Signature:</b>

```go
func Concat[T any](slice []T, values ...[]T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res1 := slice.Concat([]int{1, 2, 3}, 4, 5)
	fmt.Println(res1) //[]int{1, 2, 3, 4, 5}

	res2 := slice.Concat([]int{1, 2, 3}, []int{4, 5})
	fmt.Println(res2) //[]int{1, 2, 3, 4, 5}
}
```



### <span id="Count">Count</span>
<p>Count iterates over elements of slice, returns a count of all matched elements.</p>

<b>Signature:</b>

```go
func Count[T any](slice []T, predicate func(index int, t T) bool) int
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	evenFunc := func(i, num int) bool {
		return (num % 2) == 0
	}

	res := slice.Count(nums, evenFunc)
	fmt.Println(res) //3
}
```




### <span id="Difference">Difference</span>
<p>Creates an slice of whose element not included in the other given slice.</p>

<b>Signature:</b>

```go
func Difference[T comparable](slice, comparedSlice []T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6}

	res := slice.Difference(s1, s2)
	fmt.Println(res) //[]int{1, 2, 3}
}
```




### <span id="DifferenceBy">DifferenceBy</span>
<p>DifferenceBy accepts iteratee func which is invoked for each element of slice and values to generate the criterion by which they're compared.</p>

<b>Signature:</b>

```go
func DifferenceBy[T any](slice []T, comparedSlice []T, iteratee func(index int, t T) T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6}
	addOne := func(i int, v int) int {
		return v + 1
	}

	res := slice.DifferenceBy(s1, s2, addOne)
	fmt.Println(res) //[]int{1, 2}
}
```


### <span id="DifferenceWith">DifferenceWith</span>
<p>DifferenceWith accepts comparator which is invoked to compare elements of slice to values. The order and references of result values are determined by the first slice.</p>

<b>Signature:</b>

```go
func DifferenceWith[T any](slice []T, comparedSlice []T, comparator func(value, otherValue T) bool) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6, 7, 8}
	isDouble := func(v1, v2 int) bool {
		return v2 == 2*v1
	}

	res := slice.DifferenceWith(s1, s2, isDouble)
	fmt.Println(res) //[]int{1, 5}
}
```

### <span id="DeleteAt">DeleteAt</span>
<p>Delete the element of slice from start index to end index - 1.</p>

<b>Signature:</b>

```go
func DeleteAt[T any](slice []T, start int, end ...int)
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res1 := slice.DeleteAt([]string{"a", "b", "c", "d", "e"}, 3)
	fmt.Println(res1) //[]string{"a", "b", "c", "e"}

	res2 := slice.DeleteAt([]string{"a", "b", "c", "d", "e"}, 0, 2)
	fmt.Println(res2) //[]string{"c", "d", "e"}

}
```




### <span id="Drop">Drop</span>
<p>Creates a slice with `n` elements dropped from the beginning when n > 0, or `n` elements dropped from the ending when n < 0.</p>

<b>Signature:</b>

```go
func Drop[T any](slice []T, n int) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res1 := slice.Drop([]int{}, 0)
	fmt.Println(res1) //[]int{}

	res2 := slice.Drop([]int{1, 2, 3, 4, 5}, 1)
	fmt.Println(res2) //[]int{2, 3, 4, 5}

	res3 := slice.Drop([]int{1, 2, 3, 4, 5}, -1)
	fmt.Println(res3) //[]int{1, 2, 3, 4}
}
```




### <span id="Every">Every</span>
<p>Return true if all of the values in the slice pass the predicate function.</p>

<b>Signature:</b>

```go
func Every[T any](slice []T, predicate func(index int, t T) bool) bool
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	res := slice.Every(nums, isEven)
	fmt.Println(res) //false
}
```




### <span id="Filter">Filter</span>
<p>Return all elements which match the function.</p>

<b>Signature:</b>

```go
func Filter[T any](slice []T, predicate func(index int, t T) bool) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	res := slice.Filter(nums, isEven)
	fmt.Println(res) //[]int{2, 4}
}
```



### <span id="Find">Find</span>
<p>Iterates over elements of slice, returning the first one that passes a truth test on function.</p>

<b>Signature:</b>

```go
func Find[T any](slice []T, predicate func(index int, t T) bool) (*T, bool)
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	res, ok := slice.Find(nums, even)
	fmt.Println(res) //2
	fmt.Println(ok) //true
}
```




### <span id="FindLast">FindLast</span>
<p>iterates over elements of slice from end to begin, returning the last one that passes a truth test on function.</p>

<b>Signature:</b>

```go
func FindLast[T any](slice []T, predicate func(index int, t T) bool) (*T, bool)
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	res, ok := slice.FindLast(nums, even)
	fmt.Println(res) //4
	fmt.Println(ok) //true
}
```



### <span id="FlattenDeep">FlattenDeep</span>
<p>flattens slice recursive.</p>

<b>Signature:</b>

```go
func FlattenDeep(slice interface{}) interface{}
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	arr := [][][]string{{{"a", "b"}}, {{"c", "d"}}}
	res := slice.FlattenDeep(arr)
	fmt.Println(res) //[]string{"a", "b", "c", "d"}
}
```





### <span id="ForEach">ForEach</span>
<p>Iterates over elements of slice and invokes function for each element.</p>

<b>Signature:</b>

```go
func ForEach[T any](slice []T, iteratee func(index int, t T))
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var numbersAddTwo []int
	slice.ForEach(numbers, func(index int, value int) {
		numbersAddTwo = append(numbersAddTwo, value+2)
	})
	fmt.Println(numbersAddTwo) //[]int{3, 4, 5, 6, 7}
}
```




### <span id="GroupBy">GroupBy</span>
<p>Iterates over elements of the slice, each element will be group by criteria, returns two slices.</p>

<b>Signature:</b>

```go
func GroupBy[T any](slice []T, groupFn func(index int, t T) bool) ([]T, []T)
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	evenFunc := func(i, num int) bool {
		return (num % 2) == 0
	}
	even, odd := slice.GroupBy(nums, evenFunc)

	fmt.Println(even) //[]int{2, 4, 6}
	fmt.Println(odd) //]int{1, 3, 5}
}
```




### <span id="IntSlice">IntSlice</span>
<p>Convert interface slice to int slice.</p>

<b>Signature:</b>

```go
func IntSlice(slice interface{}) []int
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	var nums = []interface{}{1, 2, 3}
	res := slice.IntSlice(nums)
	fmt.Println(res) //[]int{1, 2, 3}
}
```




### <span id="InterfaceSlice">InterfaceSlice</span>
<p>Convert value to interface slice.</p>

<b>Signature:</b>

```go
func InterfaceSlice(slice interface{}) []interface{}
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	var nums = []int{}{1, 2, 3}
	res := slice.InterfaceSlice(nums)
	fmt.Println(res) //[]interface{}{1, 2, 3}
}
```




### <span id="Intersection">Intersection</span>
<p>Creates a slice of unique values that included by all slices.</p>

<b>Signature:</b>

```go
func Intersection[T any](slices ...[]T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	s1 := []int{1, 2, 2, 3}
	s2 := []int{1, 2, 3, 4}
	res := slice.Intersection(s1, s2),

	fmt.Println(res) //[]int{1, 2, 3}
}
```




### <span id="InsertAt">InsertAt</span>
<p>insert the element into slice at index.</p>

<b>Signature:</b>

```go
func InsertAt[T any](slice []T, index int, value interface{}) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	s := []string{"a", "b", "c"}
	
	res1, _ := slice.InsertAt(s, 0, "1")
	fmt.Println(res1) //[]string{"1", "a", "b", "c"}

	res2, _ := slice.InsertAt(s, 3, []string{"1", "2", "3"})
	fmt.Println(res2) //[]string{"a", "b", "c", "1", "2", "3"}
}
```




### <span id="Map">Map</span>
<p>Creates an slice of values by running each element in slice thru function.</p>

<b>Signature:</b>

```go
func Map[T any, U any](slice []T, iteratee func(index int, t T) U) []U
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 4}
	multiplyTwo := func(i, num int) int {
		return num * 2
	}
	res := slice.Map(nums, multiplyTwo)
	fmt.Println(res) //[]int{2, 4, 6, 8}
}
```




### <span id="Reverse">Reverse</span>
<p>Reverse the elements order in slice.</p>

<b>Signature:</b>

```go
func Reverse[T any](slice []T)
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 4}
	slice.Reverse(nums)
	fmt.Println(res) //[]int{4, 3, 2, 1}
}
```



### <span id="Reduce">Reduce</span>
<p>Reduce slice.</p>

<b>Signature:</b>

```go
func Reduce[T any](slice []T, iteratee func(index int, t1, t2 T) T, initial T) T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 4}
	reduceFunc := func(i, v1, v2 int) int {
		return v1 + v2
	}
	res := slice.Reduce(nums, reduceFunc, 0)
	fmt.Println(res) //10
}
```




### <span id="Shuffle">Shuffle</span>
<p>Creates an slice of shuffled values.</p>

<b>Signature:</b>

```go
func Shuffle[T any](slice []T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := slice.Shuffle(nums)
	fmt.Println(res) //3,1,5,4,2
}
```



### <span id="SortByField">SortByField</span>
<p>Sort struct slice by field. Slice element should be struct, field type should be int, uint, string, or bool. Default sort type is ascending (asc), if descending order, set sortType to desc</p>

<b>Signature:</b>

```go
func SortByField(slice interface{}, field string, sortType ...string) error
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	type student struct {
		name string
		age  int
	}
	students := []student{
		{"a", 10},
		{"b", 15},
		{"c", 5},
		{"d", 6},
	}
	err := slice.SortByField(students, "age", "desc")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(students) 
	// []students{
	// 	{"b", 15},
	// 	{"a", 10},
	// 	{"d", 6},
	// 	{"c", 5},
	// }
}
```



### <span id="Some">Some</span>
<p>Return true if any of the values in the list pass the predicate function.</p>

<b>Signature:</b>

```go
func Some[T any](slice []T, predicate func(index int, t T) bool) bool
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	res := slice.Some(nums, isEven)
	fmt.Println(res) //true
}
```



### <span id="StringSlice">StringSlice</span>
<p>Convert interface slice to string slice.</p>

<b>Signature:</b>

```go
func StringSlice(slice interface{}) []string
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	var s = []interface{}{"a", "b", "c"}
	res := slice.StringSlice(s)
	fmt.Println(res) //[]string{"a", "b", "c"}
}
```




### <span id="Unique">Unique</span>
<p>Remove duplicate elements in slice.</p>

<b>Signature:</b>

```go
func Unique[T any](slice []T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res := slice.Unique([]int{1, 2, 2, 3})
	fmt.Println(res) //[]int{1, 2, 3}
}
```



### <span id="Union">Unique</span>
<p>Creates a slice of unique values, in order, from all given slices. using == for equality comparisons.</p>

<b>Signature:</b>

```go
func Union[T any](slices ...[]T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	s1 := []int{1, 3, 4, 6}
	s2 := []int{1, 2, 5, 6}
	res := slice.Union(s1, s2)

	fmt.Println(res) //[]int{1, 3, 4, 6, 2, 5}
}
```



### <span id="UpdateAt">UpdateAt</span>
<p>Update the slice element at index. if param index < 0 or index >= len(slice), will return error. </p>

<b>Signature:</b>

```go
func UpdateAt[T any](slice []T, index int, value T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	s := []string{"a", "b", "c"}
	
	res1, _ := slice.UpdateAt(s, 0, "1")
	fmt.Println(res1) //[]string{"1", "b", "c"}
}
```




### <span id="Without">Without</span>
<p>Creates a slice excluding all given values. </p>

<b>Signature:</b>

```go
func Without[T any](slice []T, values ...T) []T
```
<b>Example:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res := slice.Without([]int{1, 2, 3, 4, 5}, 1, 2)
	fmt.Println(res) //[]int{3, 4, 5}
}
```












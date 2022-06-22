# Slice
slice包包含操作切片的方法集合。

<div STYLE="page-break-after: always;"></div>

## 源码:

[https://github.com/duke-git/lancet/blob/v1/slice/slice.go](https://github.com/duke-git/lancet/blob/v1/slice/slice.go)


<div STYLE="page-break-after: always;"></div>

## 用法:
```go
import (
    "github.com/duke-git/lancet/slice"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录
- [Contain](#Contain)
- [ContainSubSlice](#ContainSubSlice)
- [Chunk](#Chunk)
- [Compact](#Compact)
- [Concat](#Concat)
- [Count](#Count)
- [Difference](#Difference)
- [DifferenceBy](#DifferenceBy)
- [DeleteByIndex](#DeleteByIndex)
- [Drop](#Drop)
- [Equal](#Equal)
- [EqualWith](#EqualWith)
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
- [InsertByIndex](#InsertByIndex)
- [IndexOf](#IndexOf)
- [LastIndexOf](#LastIndexOf)
- [Map](#Map)
- [ReverseSlice](#ReverseSlice)
- [Reduce](#Reduce)
- [Shuffle](#Shuffle)
- [SortByField](#SortByField)
- [Some](#Some)
- [StringSlice](#StringSlice)
- [Unique](#Unique)
- [Union](#Union)
- [UpdateByIndex](#UpdateByIndex)
- [Without](#Without)

<div STYLE="page-break-after: always;"></div>

## 文档

### <span id="Contain">Contain</span>
<p>判断slice是否包含value</p>

<b>函数签名:</b>

```go
func Contain(iterableType interface{}, value interface{}) bool
```
<b>例子:</b>

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
<p>判断slice是否包含subslice</p>

<b>函数签名:</b>

```go
func ContainSubSlice(slice interface{}, subslice interface{}) bool
```
<b>例子:</b>

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
<p>按照size参数均分slice</p>

<b>函数签名:</b>

```go
func Chunk(slice []interface{}, size int) [][]interface{}
```
<b>例子:</b>

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
<p>去除slice中的假值（false values are false, nil, 0, ""）</p>

<b>函数签名:</b>

```go
func Compact(slice interface{}) interface{}
```
<b>例子:</b>

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
<p>连接values到slice中，values类型可以是切片或多个值</p>

<b>函数签名:</b>

```go
func Concat(slice interface{}, values ...interface{}) interface{}
```
<b>例子:</b>

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
<p>遍历切片，对每个元素执行函数function. 返回符合函数返回值为true的元素的个数，函数签名必须是func(index int, value interface{}) bool</p>

<b>函数签名:</b>

```go
func Count(slice, function interface{}) int
```
<b>例子:</b>

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
<p>创建一个切片，其元素不包含在另一个给定切片中</p>

<b>函数签名:</b>

```go
func Difference(slice1, slice2 interface{}) interface{}
```
<b>例子:</b>

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
<p>在slice和comparedSlice中的每个元素调用iteratee函数，并比较它们的返回值，如果不想等返回在slice中对应的值</p>

<b>函数签名:</b>

```go
func DifferenceBy(slice interface{}, comparedSlice interface{}, iterateeFn interface{}) interface{}
```
<b>例子:</b>

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




### <span id="DeleteByIndex">DeleteByIndex</span>
<p>删除切片中从开始索引到结束索引-1的元素</p>

<b>函数签名:</b>

```go
func DeleteByIndex(slice interface{}, start int, end ...int) (interface{}, error)
```
<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res1 := slice.DeleteByIndex([]string{"a", "b", "c", "d", "e"}, 3)
	fmt.Println(res1) //[]string{"a", "b", "c", "e"}

	res2 := slice.DeleteByIndex([]string{"a", "b", "c", "d", "e"}, 0, 2)
	fmt.Println(res2) //[]string{"c", "d", "e"}

}
```




### <span id="Drop">Drop</span>
<p>创建一个切片，当 n > 0 时从开头删除 n 个元素，或者当 n < 0 时从结尾删除 n 个元素</p>

<b>函数签名:</b>

```go
func Drop(slice interface{}, n int) interface{}
```
<b>例子:</b>

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



### <span id="Equal">Equal</span>
<p>检查两个切片是否相等，相等条件：切片长度相同，元素顺序和值都相同</p>

<b>函数签名:</b>

```go
func Equal(slice1, slice2 interface{}) bool
```
<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{3, 2, 1}

	res1 := slice.Equal(slice1, slice2)
	res2 := slice.Equal(slice1, slice3)

	fmt.Println(res1) //true
	fmt.Println(res2) //false
}
```



### <span id="EqualWith">EqualWith</span>
<p>检查两个切片是否相等，相等条件：对两个切片的元素调用比较函数comparator，返回true。 comparator函数签名: func(a interface{}, b interface{}) bool</p>

<b>函数签名:</b>

```go
func EqualWith(slice1, slice2 interface{}, comparator interface{}) bool
```
<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{2, 4, 6}

	isDouble := func(a, b int) bool {
		return b == a*2
	}

	res := slice.EqualWith(slice1, slice2, isDouble)

	fmt.Println(res) //true
}
```



### <span id="Every">Every</span>
<p>如果切片中的所有值都通过谓词函数，则返回true。 函数签名应该是func(index int, value interface{}) bool</p>

<b>函数签名:</b>

```go
func Every(slice, function interface{}) bool
```
<b>例子:</b>

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
<p>返回与函数匹配的所有元素。 函数签名应该是 func(index int, value interface{}) bool</p>

<b>函数签名:</b>

```go
func Filter(slice, function interface{}) interface{}
```
<b>例子:</b>

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
<p>遍历slice的元素，返回第一个通过function真值测试的元素。函数签名应该是 func(index int, value interface{}) bool</p>

<b>函数签名:</b>

```go
func Find(slice, function interface{}) (interface{}, bool)
```
<b>例子:</b>

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
<p>从头到尾遍历 slice 的元素，返回最后一个通过函数真值测试的元素。 函数签名应该是 func(index int, value interface{}) bool。</p>

<b>函数签名:</b>

```go
func FindLast(slice, function interface{}) (interface{}, bool)
```
<b>例子:</b>

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

<b>函数签名:</b>

```go
func FlattenDeep(slice interface{}) interface{}
```
<b>例子:</b>

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
<p>遍历slice的元素并为每个元素调用函数，函数签名应该是func(index int, value interface{})</p>

<b>函数签名:</b>

```go
func ForEach(slice, function interface{})
```
<b>例子:</b>

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
<p>迭代切片的元素，每个元素将按条件分组，返回两个切片。 函数签名应该是func(index int, value interface{}) bool</p>

<b>函数签名:</b>

```go
func GroupBy(slice, function interface{}) (interface{}, interface{})
```
<b>例子:</b>

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
<p>将接口切片转换为int切片</p>

<b>函数签名:</b>

```go
func IntSlice(slice interface{}) []int
```
<b>例子:</b>

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
<p>将值转换为接口切片</p>

<b>函数签名:</b>

```go
func InterfaceSlice(slice interface{}) []interface{}
```
<b>例子:</b>

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
<p>多个切片的交集</p>

<b>函数签名:</b>

```go
func Intersection(slices ...interface{}) interface{}
```
<b>例子:</b>

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



### <span id="IndexOf">IndexOf</span>
<p>返回在切片中找到值的第一个匹配项的索引，如果找不到值，则返回-1</p>

<b>函数签名:</b>

```go
func IndexOf(slice, value interface{}) int
```
<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	arr := []string{"a", "a", "b", "c"}
	res1 := slice.IndexOf(arr, "a")
	fmt.Println(res1) //0

	res2 := slice.IndexOf(arr, "d")
	fmt.Println(res2) //-1
}
```



### <span id="LastIndexOf">LastIndexOf</span>
<p>返回在切片中找到最后一个值的索引，如果找不到该值，则返回-1</p>

<b>函数签名:</b>

```go
func LastIndexOf(slice, value interface{}) int
```
<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	arr := []string{"a", "a", "b", "c"}
	res1 := slice.LastIndexOf(arr, "a")
	fmt.Println(res1) //1

	res2 := slice.LastIndexOf(arr, "d")
	fmt.Println(res2) //-1
}
```



### <span id="InsertByIndex">InsertByIndex</span>
<p>将元素插入到索引处的切片中</p>

<b>函数签名:</b>

```go
func InsertByIndex(slice interface{}, index int, value interface{}) (interface{}, error)
```
<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	s := []string{"a", "b", "c"}
	
	res1, _ := slice.InsertByIndex(s, 0, "1")
	fmt.Println(res1) //[]string{"1", "a", "b", "c"}

	res2, _ := slice.InsertByIndex(s, 3, []string{"1", "2", "3"})
	fmt.Println(res2) //[]string{"a", "b", "c", "1", "2", "3"}
}
```




### <span id="Map">Map</span>
<p>通过运行函数slice中的每个元素来创建一个值切片，函数签名应该是func(index int, value interface{}) interface{}。</p>

<b>函数签名:</b>

```go
func Map(slice, function interface{}) interface{}
```
<b>例子:</b>

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




### <span id="ReverseSlice">ReverseSlice</span>
<p>反转切片中的元素顺序</p>

<b>函数签名:</b>

```go
func ReverseSlice(slice interface{})
```
<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	nums := []int{1, 2, 3, 4}
	slice.ReverseSlice(nums)
	fmt.Println(res) //[]int{4, 3, 2, 1}
}
```



### <span id="Reduce">Reduce</span>
<p>将slice中的元素运行函数，返回运行结果。函数签名应该是func(index int, value1, value2 interface{}) interface{}。</p>

<b>函数签名:</b>

```go
func Reduce(slice, function, zero interface{}) interface{}
```
<b>例子:</b>

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
<p>随机打乱切片中的元素顺序</p>

<b>函数签名:</b>

```go
func Shuffle(slice interface{}) interface{}
```
<b>例子:</b>

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
<p>按字段对结构切片进行排序。slice元素应为struct，字段类型应为int、uint、string或bool。 默认排序类型是升序（asc），如果是降序，设置 sortType 为 desc</p>

<b>函数签名:</b>

```go
func SortByField(slice interface{}, field string, sortType ...string) error
```
<b>例子:</b>

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
<p>如果列表中的任何值通过谓词函数，则返回true，函数签名应该是func(index int, value interface{}) bool .</p>

<b>函数签名:</b>

```go
func Some(slice, function interface{}) bool
```
<b>例子:</b>

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
<p>将接口切片转换为字符串切片</p>

<b>函数签名:</b>

```go
func StringSlice(slice interface{}) []string
```
<b>例子:</b>

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
<p>删除切片中的重复元素</p>

<b>函数签名:</b>

```go
func Unique(slice interface{}) interface{}
```
<b>例子:</b>

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



### <span id="UniqueBy">UniqueBy</span>
<p>对切片的每个项目调用iteratee函数，然后删除重复的</p>

<b>函数签名:</b>

```go
func UniqueBy(slice, iteratee interface{}) interface{}
```
<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	res := slice.UniqueBy([]int{1, 2, 3, 4, 5, 6}, func(val int) int {
		return val % 4
	})
	fmt.Println(res) //[]int{1, 2, 3, 0}
}
```


### <span id="Union">Union</span>
<p>从所有给定的切片按顺序创建一个唯一值切片。 使用 == 进行相等比较。</p>

<b>函数签名:</b>

```go
func Union(slices ...interface{}) interface{}
```
<b>例子:</b>

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



### <span id="UpdateByIndex">UpdateByIndex</span>
<p>更新索引处的切片元素。 如果 param index < 0 或 index >= len(slice)，将返回错误</p>

<b>函数签名:</b>

```go
func UpdateByIndex(slice interface{}, index int, value interface{}) (interface{}, error)
```
<b>例子:</b>

```go
import (
	"fmt"
	"github.com/duke-git/lancet/slice"
)

func main() {
	s := []string{"a", "b", "c"}
	
	res1, _ := slice.UpdateByIndex(s, 0, "1")
	fmt.Println(res1) //[]string{"1", "b", "c"}
}
```




### <span id="Without">Without</span>
<p>创建一个不包括所有给定值的切片</p>

<b>函数签名:</b>

```go
func Without(slice interface{}, values ...interface{}) interface{}
```
<b>例子:</b>

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












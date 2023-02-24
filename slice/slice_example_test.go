package slice

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func ExampleContain() {
	result1 := Contain([]string{"a", "b", "c"}, "a")
	result2 := Contain([]int{1, 2, 3}, 4)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleContainBy() {
	type foo struct {
		A string
		B int
	}

	array1 := []foo{{A: "1", B: 1}, {A: "2", B: 2}}
	result1 := ContainBy(array1, func(f foo) bool { return f.A == "1" && f.B == 1 })
	result2 := ContainBy(array1, func(f foo) bool { return f.A == "2" && f.B == 1 })

	array2 := []string{"a", "b", "c"}
	result3 := ContainBy(array2, func(t string) bool { return t == "a" })
	result4 := ContainBy(array2, func(t string) bool { return t == "d" })

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleContainSubSlice() {
	result1 := ContainSubSlice([]string{"a", "b", "c"}, []string{"a", "b"})
	result2 := ContainSubSlice([]string{"a", "b", "c"}, []string{"a", "d"})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleChunk() {
	arr := []string{"a", "b", "c", "d", "e"}

	result1 := Chunk(arr, 1)
	result2 := Chunk(arr, 2)
	result3 := Chunk(arr, 3)
	result4 := Chunk(arr, 4)
	result5 := Chunk(arr, 5)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// [[a] [b] [c] [d] [e]]
	// [[a b] [c d] [e]]
	// [[a b c] [d e]]
	// [[a b c d] [e]]
	// [[a b c d e]]
}

func ExampleCompact() {
	result1 := Compact([]int{0})
	result2 := Compact([]int{0, 1, 2, 3})
	result3 := Compact([]string{"", "a", "b", "0"})
	result4 := Compact([]bool{false, true, true})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// []
	// [1 2 3]
	// [a b 0]
	// [true true]
}

func ExampleConcat() {
	result1 := Concat([]int{1, 2}, []int{3, 4})
	result2 := Concat([]string{"a", "b"}, []string{"c"}, []string{"d"})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// [1 2 3 4]
	// [a b c d]
}

func ExampleDifference() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6}

	result := Difference(slice1, slice2)

	fmt.Println(result)

	// Output:
	// [1 2 3]
}

func ExampleDifferenceBy() {
	slice1 := []int{1, 2, 3, 4, 5} //after add one: 2 3 4 5 6
	slice2 := []int{3, 4, 5}       //after add one: 4 5 6

	addOne := func(i int, v int) int {
		return v + 1
	}

	result := DifferenceBy(slice1, slice2, addOne)

	fmt.Println(result)

	// Output:
	// [1 2]
}

func ExampleDifferenceWith() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4, 5, 6, 7, 8}

	isDouble := func(v1, v2 int) bool {
		return v2 == 2*v1
	}

	result := DifferenceWith(slice1, slice2, isDouble)

	fmt.Println(result)

	// Output:
	// [1 5]
}

func ExampleEqual() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{1, 3, 2}

	result1 := Equal(slice1, slice2)
	result2 := Equal(slice1, slice3)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleEqualWith() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{2, 4, 6}

	isDouble := func(a, b int) bool {
		return b == a*2
	}

	result := EqualWith(slice1, slice2, isDouble)

	fmt.Println(result)

	// Output:
	// true
}

func ExampleEvery() {
	nums := []int{1, 2, 3, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result := Every(nums, isEven)

	fmt.Println(result)

	// Output:
	// false
}

func ExampleNone() {
	nums := []int{1, 3, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result := None(nums, isEven)

	fmt.Println(result)

	// Output:
	// true
}

func ExampleSome() {
	nums := []int{1, 2, 3, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result := Some(nums, isEven)

	fmt.Println(result)

	// Output:
	// true
}

func ExampleFilter() {
	nums := []int{1, 2, 3, 4, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result := Filter(nums, isEven)

	fmt.Println(result)

	// Output:
	// [2 4]
}

func ExampleCount() {
	nums := []int{1, 2, 3, 3, 4}

	result1 := Count(nums, 1)
	result2 := Count(nums, 3)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 1
	// 2
}

func ExampleCountBy() {
	nums := []int{1, 2, 3, 4, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result := CountBy(nums, isEven)

	fmt.Println(result)

	// Output:
	// 2
}

func ExampleGroupBy() {
	nums := []int{1, 2, 3, 4, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	even, odd := GroupBy(nums, isEven)

	fmt.Println(even)
	fmt.Println(odd)

	// Output:
	// [2 4]
	// [1 3 5]
}

func ExampleGroupWith() {
	nums := []float64{6.1, 4.2, 6.3}

	floor := func(num float64) float64 {
		return math.Floor(num)
	}

	result := GroupWith(nums, floor) //map[float64][]float64

	fmt.Println(result)

	// Output:
	// map[4:[4.2] 6:[6.1 6.3]]
}

func ExampleFind() {
	nums := []int{1, 2, 3, 4, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result, ok := Find(nums, isEven)

	fmt.Println(*result)
	fmt.Println(ok)

	// Output:
	// 2
	// true
}

func ExampleFindLast() {
	nums := []int{1, 2, 3, 4, 5}

	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	result, ok := FindLast(nums, isEven)

	fmt.Println(*result)
	fmt.Println(ok)

	// Output:
	// 4
	// true
}

func ExampleFlatten() {
	arrs := [][][]string{{{"a", "b"}}, {{"c", "d"}}}

	result := Flatten(arrs)

	fmt.Println(result)

	// Output:
	// [[a b] [c d]]
}

func ExampleFlattenDeep() {
	arrs := [][][]string{{{"a", "b"}}, {{"c", "d"}}}

	result := FlattenDeep(arrs)

	fmt.Println(result)

	// Output:
	// [a b c d]
}

func ExampleForEach() {
	nums := []int{1, 2, 3}

	var result []int
	addOne := func(_ int, v int) {
		result = append(result, v+1)
	}

	ForEach(nums, addOne)

	fmt.Println(result)

	// Output:
	// [2 3 4]
}

func ExampleMap() {
	nums := []int{1, 2, 3}

	addOne := func(_ int, v int) int {
		return v + 1
	}

	result := Map(nums, addOne)

	fmt.Println(result)

	// Output:
	// [2 3 4]
}

func ExampleFilterMap() {
	nums := []int{1, 2, 3, 4, 5}

	getEvenNumStr := func(i, num int) (string, bool) {
		if num%2 == 0 {
			return strconv.FormatInt(int64(num), 10), true
		}
		return "", false
	}

	result := FilterMap(nums, getEvenNumStr)

	fmt.Printf("%#v", result)

	// Output:
	// []string{"2", "4"}
}

func ExampleFlatMap() {
	nums := []int{1, 2, 3, 4}

	result := FlatMap(nums, func(i int, num int) []string {
		s := "hi-" + strconv.FormatInt(int64(num), 10)
		return []string{s}
	})

	fmt.Printf("%#v", result)

	// Output:
	// []string{"hi-1", "hi-2", "hi-3", "hi-4"}
}

func ExampleReduce() {
	nums := []int{1, 2, 3}

	sum := func(_ int, v1, v2 int) int {
		return v1 + v2
	}

	result := Reduce(nums, sum, 0)

	fmt.Println(result)

	// Output:
	// 6
}

func ExampleReplace() {
	strs := []string{"a", "b", "c", "a"}

	result1 := Replace(strs, "a", "x", 0)
	result2 := Replace(strs, "a", "x", 1)
	result3 := Replace(strs, "a", "x", 2)
	result4 := Replace(strs, "a", "x", 3)
	result5 := Replace(strs, "a", "x", -1)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// [a b c a]
	// [x b c a]
	// [x b c x]
	// [x b c x]
	// [x b c x]
}

func ExampleReplaceAll() {
	result := ReplaceAll([]string{"a", "b", "c", "a"}, "a", "x")

	fmt.Println(result)

	// Output:
	// [x b c x]
}

func ExampleRepeat() {
	result := Repeat("a", 3)

	fmt.Println(result)

	// Output:
	// [a a a]
}

func ExampleInterfaceSlice() {
	strs := []string{"a", "b", "c"}

	result := InterfaceSlice(strs) //[]interface{}{"a", "b", "c"}
	fmt.Println(result)

	// Output:
	// [a b c]
}

func ExampleStringSlice() {
	strs := []interface{}{"a", "b", "c"}

	result := StringSlice(strs) //[]string{"a", "b", "c"}
	fmt.Println(result)

	// Output:
	// [a b c]
}

func ExampleIntSlice() {
	nums := []interface{}{1, 2, 3}

	result := IntSlice(nums) //[]int{1, 2, 3}
	fmt.Println(result)

	// Output:
	// [1 2 3]
}

func ExampleDeleteAt() {
	result1 := DeleteAt([]string{"a", "b", "c"}, -1)
	result2 := DeleteAt([]string{"a", "b", "c"}, 0)
	result3 := DeleteAt([]string{"a", "b", "c"}, 0, 2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// [a b c]
	// [b c]
	// [c]
}

func ExampleDrop() {
	result1 := Drop([]string{"a", "b", "c"}, 0)
	result2 := Drop([]string{"a", "b", "c"}, 1)
	result3 := Drop([]string{"a", "b", "c"}, -1)
	result4 := Drop([]string{"a", "b", "c"}, 4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// [a b c]
	// [b c]
	// [a b c]
	// []
}

func ExampleDropRight() {
	result1 := DropRight([]string{"a", "b", "c"}, 0)
	result2 := DropRight([]string{"a", "b", "c"}, 1)
	result3 := DropRight([]string{"a", "b", "c"}, -1)
	result4 := DropRight([]string{"a", "b", "c"}, 4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// [a b c]
	// [a b]
	// [a b c]
	// []
}

func ExampleDropWhile() {
	numbers := []int{1, 2, 3, 4, 5}

	result1 := DropWhile(numbers, func(n int) bool {
		return n != 2
	})
	result2 := DropWhile(numbers, func(n int) bool {
		return true
	})
	result3 := DropWhile(numbers, func(n int) bool {
		return n == 0
	})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// [2 3 4 5]
	// []
	// [1 2 3 4 5]
}

func ExampleDropRightWhile() {
	numbers := []int{1, 2, 3, 4, 5}

	result1 := DropRightWhile(numbers, func(n int) bool {
		return n != 2
	})
	result2 := DropRightWhile(numbers, func(n int) bool {
		return true
	})
	result3 := DropRightWhile(numbers, func(n int) bool {
		return n == 0
	})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// [1 2]
	// []
	// [1 2 3 4 5]
}

func ExampleInsertAt() {
	result1 := InsertAt([]string{"a", "b", "c"}, 0, "1")
	result2 := InsertAt([]string{"a", "b", "c"}, 1, "1")
	result3 := InsertAt([]string{"a", "b", "c"}, 2, "1")
	result4 := InsertAt([]string{"a", "b", "c"}, 3, "1")
	result5 := InsertAt([]string{"a", "b", "c"}, 0, []string{"1", "2", "3"})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// [1 a b c]
	// [a 1 b c]
	// [a b 1 c]
	// [a b c 1]
	// [1 2 3 a b c]
}

func ExampleUpdateAt() {
	result1 := UpdateAt([]string{"a", "b", "c"}, -1, "1")
	result2 := UpdateAt([]string{"a", "b", "c"}, 0, "1")
	result3 := UpdateAt([]string{"a", "b", "c"}, 1, "1")
	result4 := UpdateAt([]string{"a", "b", "c"}, 2, "1")
	result5 := UpdateAt([]string{"a", "b", "c"}, 3, "1")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// [a b c]
	// [1 b c]
	// [a 1 c]
	// [a b 1]
	// [a b c]
}

func ExampleUnique() {
	result := Unique([]string{"a", "a", "b"})
	fmt.Println(result)

	// Output:
	// [a b]
}

func ExampleUniqueBy() {
	nums := []int{1, 2, 3, 4, 5, 6}
	result := UniqueBy(nums, func(val int) int {
		return val % 3
	})

	fmt.Println(result)

	// Output:
	// [1 2 0]
}

func ExampleUnion() {
	nums1 := []int{1, 3, 4, 6}
	nums2 := []int{1, 2, 5, 6}

	result := Union(nums1, nums2)

	fmt.Println(result)

	// Output:
	// [1 3 4 6 2 5]
}

func ExampleUnionBy() {
	nums := []int{1, 2, 3, 4}

	divideTwo := func(n int) int {
		return n / 2
	}
	result := UnionBy(divideTwo, nums)

	fmt.Println(result)

	// Output:
	// [1 2 4]
}

func ExampleMerge() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{3, 4}

	result := Merge(nums1, nums2)

	fmt.Println(result)

	// Output:
	// [1 2 3 3 4]
}

func ExampleIntersection() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 3, 4}

	result := Intersection(nums1, nums2)

	fmt.Println(result)

	// Output:
	// [2 3]
}

func ExampleSymmetricDifference() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{1, 2, 4}

	result := SymmetricDifference(nums1, nums2)

	fmt.Println(result)

	// Output:
	// [3 4]
}

func ExampleReverse() {
	strs := []string{"a", "b", "c", "d"}

	Reverse(strs)

	fmt.Println(strs)

	// Output:
	// [d c b a]
}

func ExampleIsAscending() {

	result1 := IsAscending([]int{1, 2, 3, 4, 5})
	result2 := IsAscending([]int{5, 4, 3, 2, 1})
	result3 := IsAscending([]int{2, 1, 3, 4, 5})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// false
	// false
}

func ExampleIsDescending() {

	result1 := IsDescending([]int{5, 4, 3, 2, 1})
	result2 := IsDescending([]int{1, 2, 3, 4, 5})
	result3 := IsDescending([]int{2, 1, 3, 4, 5})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// false
	// false
}

func ExampleIsSorted() {

	result1 := IsSorted([]int{1, 2, 3, 4, 5})
	result2 := IsSorted([]int{5, 4, 3, 2, 1})
	result3 := IsSorted([]int{2, 1, 3, 4, 5})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// true
	// false
}

func ExampleIsSortedByKey() {
	result1 := IsSortedByKey([]string{"a", "ab", "abc"}, func(s string) int {
		return len(s)
	})
	result2 := IsSortedByKey([]string{"abc", "ab", "a"}, func(s string) int {
		return len(s)
	})
	result3 := IsSortedByKey([]string{"abc", "a", "ab"}, func(s string) int {
		return len(s)
	})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// true
	// true
	// false
}

func ExampleSort() {
	nums := []int{1, 4, 3, 2, 5}

	Sort(nums)

	fmt.Println(nums)

	// Output:
	// [1 2 3 4 5]
}

func ExampleSortBy() {
	type User struct {
		Name string
		Age  uint
	}

	users := []User{
		{Name: "a", Age: 21},
		{Name: "b", Age: 15},
		{Name: "c", Age: 100}}

	SortBy(users, func(a, b User) bool {
		return a.Age < b.Age
	})

	fmt.Println(users)

	// Output:
	// [{b 15} {a 21} {c 100}]
}

func ExampleSortByField() {
	type User struct {
		Name string
		Age  uint
	}

	users := []User{
		{Name: "a", Age: 21},
		{Name: "b", Age: 15},
		{Name: "c", Age: 100}}

	err := SortByField(users, "Age", "desc")
	if err != nil {
		return
	}

	fmt.Println(users)

	// Output:
	// [{c 100} {a 21} {b 15}]
}

func ExampleWithout() {
	result := Without([]int{1, 2, 3, 4}, 1, 2)

	fmt.Println(result)

	// Output:
	// [3 4]
}

func ExampleIndexOf() {
	strs := []string{"a", "a", "b", "c"}

	result1 := IndexOf(strs, "a")
	result2 := IndexOf(strs, "d")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 0
	// -1
}

func ExampleLastIndexOf() {
	strs := []string{"a", "a", "b", "c"}

	result1 := LastIndexOf(strs, "a")
	result2 := LastIndexOf(strs, "d")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// 1
	// -1
}

func ExampleToSlice() {
	result := ToSlice("a", "b", "c")

	fmt.Println(result)

	// Output:
	// [a b c]
}

func ExampleToSlicePointer() {
	str1 := "a"
	str2 := "b"

	result := ToSlicePointer(str1, str2)

	expect := []*string{&str1, &str2}

	isEqual := reflect.DeepEqual(result, expect)

	fmt.Println(isEqual)

	// Output:
	// true
}

func ExampleAppendIfAbsent() {
	result1 := AppendIfAbsent([]string{"a", "b"}, "b")
	result2 := AppendIfAbsent([]string{"a", "b"}, "c")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// [a b]
	// [a b c]
}

func ExampleKeyBy() {
	result := KeyBy([]string{"a", "ab", "abc"}, func(str string) int {
		return len(str)
	})

	fmt.Println(result)

	// Output:
	// map[1:a 2:ab 3:abc]
}

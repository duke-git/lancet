package slice

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestContain(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContain")

	assert.Equal(true, Contain([]string{"a", "b", "c"}, "a"))
	assert.Equal(false, Contain([]string{"a", "b", "c"}, "d"))
	assert.Equal(true, Contain([]string{""}, ""))
	assert.Equal(false, Contain([]string{}, ""))

	assert.Equal(true, Contain([]int{1, 2, 3}, 1))
}

func TestContainBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContainBy")

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

	assert.Equal(true, result1)
	assert.Equal(false, result2)
	assert.Equal(true, result3)
	assert.Equal(false, result4)
}

func TestContainSubSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestContainSubSlice")
	assert.Equal(true, ContainSubSlice([]string{"a", "a", "b", "c"}, []string{"a", "a"}))
	assert.Equal(false, ContainSubSlice([]string{"a", "a", "b", "c"}, []string{"a", "d"}))
	assert.Equal(true, ContainSubSlice([]int{1, 2, 3}, []int{1, 2}))
	assert.Equal(false, ContainSubSlice([]int{1, 2, 3}, []int{0, 1}))
}

func TestChunk(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestChunk")

	arr := []string{"a", "b", "c", "d", "e"}

	assert.Equal([][]string{}, Chunk(arr, -1))

	assert.Equal([][]string{}, Chunk(arr, 0))

	r1 := [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}}
	assert.Equal(r1, Chunk(arr, 1))

	r2 := [][]string{{"a", "b"}, {"c", "d"}, {"e"}}
	assert.Equal(r2, Chunk(arr, 2))

	r3 := [][]string{{"a", "b", "c"}, {"d", "e"}}
	assert.Equal(r3, Chunk(arr, 3))

	r4 := [][]string{{"a", "b", "c", "d"}, {"e"}}
	assert.Equal(r4, Chunk(arr, 4))

	r5 := [][]string{{"a", "b", "c", "d", "e"}}
	assert.Equal(r5, Chunk(arr, 5))

	r6 := [][]string{{"a", "b", "c", "d", "e"}}
	assert.Equal(r6, Chunk(arr, 6))
}

func TestCompact(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TesCompact")

	assert.Equal([]int{}, Compact([]int{0}))
	assert.Equal([]int{1, 2, 3}, Compact([]int{0, 1, 2, 3}))
	assert.Equal([]string{}, Compact([]string{""}))
	assert.Equal([]string{" "}, Compact([]string{" "}))
	assert.Equal([]string{"a", "b", "0"}, Compact([]string{"", "a", "b", "0"}))
	assert.Equal([]bool{true, true}, Compact([]bool{false, true, true}))
}

func TestConcat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "Concat")

	assert.Equal([]int{1, 2, 3, 4, 5}, Concat([]int{1, 2, 3}, []int{4, 5}))
	assert.Equal([]int{1, 2, 3, 4, 5}, Concat([]int{1, 2, 3}, []int{4}, []int{5}))
}

func BenchmarkConcat(b *testing.B) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := []int{7, 8, 9}

	for i := 0; i < b.N; i++ {
		result := Concat(slice1, slice2, slice3)
		if len(result) == 0 {
			b.Fatal("unexpected empty result")
		}
	}
}

func TestEqual(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEqual")

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slice3 := []int{3, 2, 1}

	assert.Equal(true, Equal(slice1, slice2))
	assert.Equal(false, Equal(slice1, slice3))
}

// go test -fuzz=Fuzz -fuzztime=10s .
func FuzzEqual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b []byte) {
		Equal(a, b)
	})
}

func TestEqualWith(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEqualWith")

	slice1 := []int{1, 2, 3}
	slice2 := []int{2, 4, 6}

	isDouble := func(a, b int) bool {
		return b == a*2
	}

	assert.Equal(true, EqualWith(slice1, slice2, isDouble))
}

func TestEvery(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEvery")

	nums := []int{1, 2, 3, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	assert.Equal(false, Every(nums, isEven))
}

func TestNone(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestNone")

	nums := []int{1, 2, 3, 5}
	check := func(i, num int) bool {
		return num%2 == 1
	}

	assert.Equal(false, None(nums, check))
}

func TestSome(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSome")

	nums := []int{1, 2, 3, 5}
	hasEven := func(i, num int) bool {
		return num%2 == 0
	}

	assert.Equal(true, Some(nums, hasEven))
}

func TestFilter(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFilter")

	nums := []int{1, 2, 3, 4, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	assert.Equal([]int{2, 4}, Filter(nums, isEven))

	type student struct {
		name string
		age  int
	}
	students := []student{
		{"a", 10},
		{"b", 11},
		{"c", 12},
		{"d", 13},
		{"e", 14},
	}
	studentsOfAageGreat12 := []student{
		{"d", 13},
		{"e", 14},
	}
	filterFunc := func(i int, s student) bool {
		return s.age > 12
	}

	assert.Equal(studentsOfAageGreat12, Filter(students, filterFunc))
}

func TestGroupBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGroupBy")

	nums := []int{1, 2, 3, 4, 5, 6}
	evenFunc := func(i, num int) bool {
		return (num % 2) == 0
	}
	even, odd := GroupBy(nums, evenFunc)

	assert.Equal([]int{2, 4, 6}, even)
	assert.Equal([]int{1, 3, 5}, odd)
}

func TestGroupWith(t *testing.T) {
	t.Parallel()

	nums := []float64{6.1, 4.2, 6.3}
	floor := func(num float64) float64 {
		return math.Floor(num)
	}
	expected := map[float64][]float64{
		4: {4.2},
		6: {6.1, 6.3},
	}
	actual := GroupWith(nums, floor)
	assert := internal.NewAssert(t, "TestGroupWith")
	assert.Equal(expected, actual)
}

func TestCount(t *testing.T) {
	t.Parallel()

	numbers := []int{1, 2, 3, 3, 5, 6}

	assert := internal.NewAssert(t, "TestCountBy")

	assert.Equal(1, Count(numbers, 1))
	assert.Equal(2, Count(numbers, 3))
}

func TestCountBy(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5, 6}
	evenFunc := func(i, num int) bool {
		return (num % 2) == 0
	}

	assert := internal.NewAssert(t, "TestCountBy")
	assert.Equal(3, CountBy(nums, evenFunc))
}

func TestFind(t *testing.T) {
	t.Parallel()

	nums := []int{1, 2, 3, 4, 5}
	even := func(i, num int) bool {
		return num%2 == 0
	}
	result, ok := Find(nums, even)

	assert := internal.NewAssert(t, "TestFind")
	assert.Equal(true, ok)
	assert.Equal(2, *result)
}

func TestFindBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFindBy")

	nums := []int{1, 2, 3, 4, 5}

	result, ok := FindBy(nums, func(i, num int) bool {
		return num%2 == 0
	})
	assert.Equal(true, ok)
	assert.Equal(2, result)

	result, ok = FindBy(nums, func(_ int, v int) bool {
		return v == 6
	})
	assert.Equal(false, ok)
	assert.Equal(0, result)
}

func TestFindLastBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFindLastBy")

	nums := []int{1, 2, 3, 4, 5}
	even := func(i, num int) bool {
		return num%2 == 0
	}
	result, ok := FindLastBy(nums, even)
	if !ok {
		t.Fatal("found nothing")
	}
	assert.Equal(4, result)

	result, ok = FindLastBy(nums, func(_ int, v int) bool {
		return v == 6
	})

	assert.Equal(result == 0 && ok == false, true)
}

func TestFindLast(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFindLast")

	nums := []int{1, 2, 3, 4, 5}
	even := func(i, num int) bool {
		return num%2 == 0
	}
	result, ok := FindLast(nums, even)
	if !ok {
		t.Fatal("found nothing")
	}

	assert.Equal(4, *result)
}

func TestFindFoundNothing(t *testing.T) {
	t.Parallel()

	nums := []int{1, 1, 1, 1, 1, 1}
	findFunc := func(i, num int) bool {
		return num > 1
	}
	_, ok := Find(nums, findFunc)

	assert := internal.NewAssert(t, "TestFindFoundNothing")
	assert.Equal(false, ok)
}

func TestFlatten(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFlatten")

	input := [][][]string{{{"a", "b"}}, {{"c", "d"}}}
	expected := [][]string{{"a", "b"}, {"c", "d"}}

	assert.Equal(expected, Flatten(input))
}

func TestFlattenDeep(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFlattenDeep")

	input := [][][]string{{{"a", "b"}}, {{"c", "d"}}}
	expected := []string{"a", "b", "c", "d"}

	assert.Equal(expected, FlattenDeep(input))
}

func TestForEach(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestForEach")

	numbers := []int{1, 2, 3, 4, 5}

	var result []int
	addTwo := func(index int, value int) {
		result = append(result, value+2)
	}

	ForEach(numbers, addTwo)

	assert.Equal([]int{3, 4, 5, 6, 7}, result)
}

func TestForEachWithBreak(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestForEach")

	numbers := []int{1, 2, 3, 4, 5}

	var sum int

	ForEachWithBreak(numbers, func(_, n int) bool {
		if n > 3 {
			return false
		}
		sum += n
		return true
	})

	assert.Equal(6, sum)
}

func TestMap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMap")

	nums := []int{1, 2, 3, 4}
	multiplyTwo := func(i, num int) int {
		return num * 2
	}

	assert.Equal([]int{2, 4, 6, 8}, Map(nums, multiplyTwo))

	type student struct {
		name string
		age  int
	}
	students := []student{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}
	studentsOfAdd10Aage := []student{
		{"a", 11},
		{"b", 12},
		{"c", 13},
	}
	mapFunc := func(i int, s student) student {
		s.age += 10
		return s
	}

	assert.Equal(studentsOfAdd10Aage, Map(students, mapFunc))
}

func TestFilterMap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFilterMap")

	nums := []int{1, 2, 3, 4, 5}

	getEvenNumStr := func(i, num int) (string, bool) {
		if num%2 == 0 {
			return strconv.FormatInt(int64(num), 10), true
		}
		return "", false
	}

	result := FilterMap(nums, getEvenNumStr)

	assert.Equal([]string{"2", "4"}, result)
}

func TestFlatMap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFlatMap")

	nums := []int{1, 2, 3, 4}

	result := FlatMap(nums, func(i int, num int) []string {
		s := "hi-" + strconv.FormatInt(int64(num), 10)
		return []string{s}
	})

	assert.Equal([]string{"hi-1", "hi-2", "hi-3", "hi-4"}, result)
}

func TestReduce(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestReduce")

	cases := [][]int{
		{},
		{1},
		{1, 2, 3, 4}}

	expected := []int{0, 1, 10}

	f := func(i, v1, v2 int) int {
		return v1 + v2
	}

	for i := 0; i < len(cases); i++ {
		actual := Reduce(cases[i], f, 0)
		assert.Equal(expected[i], actual)
	}
}

func TestReduceBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestReduceBy")

	result1 := ReduceBy([]int{1, 2, 3, 4}, 0, func(_ int, item int, agg int) int {
		return agg + item
	})

	result2 := ReduceBy([]int{1, 2, 3, 4}, "", func(_ int, item int, agg string) string {
		return agg + fmt.Sprintf("%v", item)
	})

	assert.Equal(10, result1)
	assert.Equal("1234", result2)

}

func TestReduceRight(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestReduceRight")

	result := ReduceRight([]int{1, 2, 3, 4}, "", func(_ int, item int, agg string) string {
		return agg + fmt.Sprintf("%v", item)
	})

	assert.Equal("4321", result)
}

func TestIntSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIntSlice")

	var nums []any
	nums = append(nums, 1, 2, 3)

	assert.Equal([]int{1, 2, 3}, IntSlice(nums))
}

func TestStringSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStringSlice")

	var strs []any
	strs = append(strs, "a", "b", "c")

	assert.Equal([]string{"a", "b", "c"}, StringSlice(strs))
}

func TestInterfaceSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestInterfaceSlice")

	strs := []string{"a", "b", "c"}
	expect := []any{"a", "b", "c"}

	assert.Equal(expect, InterfaceSlice(strs))
}

func TestDeleteAt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDeleteAt")
	arr := []int{1, 2, 3, 4, 5}

	assert.Equal([]int{2, 3, 4, 5}, DeleteAt(arr, 0))
	assert.Equal([]int{1, 2, 3, 4}, DeleteAt(arr, 4))

	assert.Equal([]int{1, 2, 3, 4}, DeleteAt(arr, 5))
	assert.Equal([]int{1, 2, 3, 4}, DeleteAt(arr, 6))

	assert.Equal([]int{1, 2, 3, 4, 5}, arr)
}

func TestDeleteRange(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDeleteRange")

	arr := []int{1, 2, 3, 4, 5}

	assert.Equal([]int{1, 2, 3, 4, 5}, DeleteRange(arr, 0, 0))
	assert.Equal([]int{2, 3, 4, 5}, DeleteRange(arr, 0, 1))
	assert.Equal([]int{4, 5}, DeleteRange(arr, 0, 3))
}

func TestDrop(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDrop")

	assert.Equal([]int{}, Drop([]int{}, 0))
	assert.Equal([]int{}, Drop([]int{}, 1))
	assert.Equal([]int{}, Drop([]int{}, -1))

	assert.Equal([]int{1, 2, 3, 4, 5}, Drop([]int{1, 2, 3, 4, 5}, 0))
	assert.Equal([]int{2, 3, 4, 5}, Drop([]int{1, 2, 3, 4, 5}, 1))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, 5))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, 6))

	assert.Equal([]int{1, 2, 3, 4, 5}, Drop([]int{1, 2, 3, 4, 5}, -1))
}

func TestDropRight(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDropRight")

	assert.Equal([]int{}, DropRight([]int{}, 0))
	assert.Equal([]int{}, DropRight([]int{}, 1))
	assert.Equal([]int{}, DropRight([]int{}, -1))

	assert.Equal([]int{1, 2, 3, 4, 5}, DropRight([]int{1, 2, 3, 4, 5}, 0))
	assert.Equal([]int{1, 2, 3, 4}, DropRight([]int{1, 2, 3, 4, 5}, 1))
	assert.Equal([]int{}, DropRight([]int{1, 2, 3, 4, 5}, 5))
	assert.Equal([]int{}, DropRight([]int{1, 2, 3, 4, 5}, 6))

	assert.Equal([]int{1, 2, 3, 4, 5}, DropRight([]int{1, 2, 3, 4, 5}, -1))
}

func TestDropWhile(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDropWhile")

	numbers := []int{1, 2, 3, 4, 5}

	r1 := DropWhile(numbers, func(n int) bool {
		return n != 2
	})
	assert.Equal([]int{2, 3, 4, 5}, r1)

	r2 := DropWhile(numbers, func(n int) bool {
		return true
	})
	assert.Equal([]int{}, r2)

	r3 := DropWhile(numbers, func(n int) bool {
		return n == 0
	})
	assert.Equal([]int{1, 2, 3, 4, 5}, r3)
}

func TestDropRightWhile(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDropRightWhile")

	numbers := []int{1, 2, 3, 4, 5}

	r1 := DropRightWhile(numbers, func(n int) bool {
		return n != 2
	})
	assert.Equal([]int{1, 2}, r1)

	r2 := DropRightWhile(numbers, func(n int) bool {
		return true
	})
	assert.Equal([]int{}, r2)

	r3 := DropRightWhile(numbers, func(n int) bool {
		return n == 0
	})
	assert.Equal([]int{1, 2, 3, 4, 5}, r3)
}

func TestInsertAt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestInsertAt")

	strs := []string{"a", "b", "c"}
	assert.Equal([]string{"a", "b", "c"}, InsertAt(strs, -1, "1"))
	assert.Equal([]string{"a", "b", "c"}, InsertAt(strs, 4, "1"))
	assert.Equal([]string{"1", "a", "b", "c"}, InsertAt(strs, 0, "1"))
	assert.Equal([]string{"a", "1", "b", "c"}, InsertAt(strs, 1, "1"))
	assert.Equal([]string{"a", "b", "1", "c"}, InsertAt(strs, 2, "1"))
	assert.Equal([]string{"a", "b", "c", "1"}, InsertAt(strs, 3, "1"))
	assert.Equal([]string{"1", "2", "3", "a", "b", "c"}, InsertAt(strs, 0, []string{"1", "2", "3"}))
	assert.Equal([]string{"a", "b", "c", "1", "2", "3"}, InsertAt(strs, 3, []string{"1", "2", "3"}))
}

func TestUpdateAt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUpdateAt")

	assert.Equal([]string{"a", "b", "c"}, UpdateAt([]string{"a", "b", "c"}, -1, "1"))
	assert.Equal([]string{"1", "b", "c"}, UpdateAt([]string{"a", "b", "c"}, 0, "1"))
	assert.Equal([]string{"a", "b", "2"}, UpdateAt([]string{"a", "b", "c"}, 2, "2"))
	assert.Equal([]string{"a", "b", "c"}, UpdateAt([]string{"a", "b", "c"}, 3, "2"))
}

func TestUnique(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUnique")

	assert.Equal([]int{1, 2, 3}, Unique([]int{1, 2, 2, 3}))
	assert.Equal([]string{"a", "b", "c"}, Unique([]string{"a", "a", "b", "c"}))
}

func TestUniqueBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUniqueBy")

	actual := UniqueBy([]int{1, 2, 3, 4, 5, 6}, func(val int) int {
		return val % 4
	})

	assert.Equal([]int{1, 2, 3, 4}, actual)
}

func TestUniqueByField(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUniqueByField")

	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	users := []User{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "b"},
		{ID: 1, Name: "c"},
	}

	uniqueUsers, err := UniqueByField(users, "ID")
	if err != nil {
		t.Error(err)
	}

	assert.Equal([]User{
		{ID: 1, Name: "a"},
		{ID: 2, Name: "b"},
	}, uniqueUsers)
}

func TestUniqueByComparator(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUniqueByComparator")

	t.Run("equal comparison", func(t *testing.T) {
		nums := []int{1, 2, 3, 1, 2, 4, 5, 6, 4, 7}
		comparator := func(item int, other int) bool {
			return item == other
		}
		result := UniqueByComparator(nums, comparator)

		assert.Equal([]int{1, 2, 3, 4, 5, 6, 7}, result)
	})

	t.Run("unique struct slice by field", func(t *testing.T) {
		type student struct {
			Name string
			Age  int
		}

		students := []student{
			{Name: "a", Age: 11},
			{Name: "b", Age: 12},
			{Name: "a", Age: 13},
			{Name: "c", Age: 14},
		}

		comparator := func(item, other student) bool { return item.Name == other.Name }

		result := UniqueByComparator(students, comparator)

		assert.Equal([]student{
			{Name: "a", Age: 11},
			{Name: "b", Age: 12},
			{Name: "c", Age: 14},
		}, result)
	})

	t.Run("case-insensitive string comparison", func(t *testing.T) {
		stringSlice := []string{"apple", "banana", "Apple", "cherry", "Banana", "date"}
		caseInsensitiveComparator := func(item, other string) bool {
			return strings.ToLower(item) == strings.ToLower(other)
		}

		result := UniqueByComparator(stringSlice, caseInsensitiveComparator)

		assert.Equal([]string{"apple", "banana", "cherry", "date"}, result)
	})

}

func TestUnion(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUnion")

	s1 := []int{1, 3, 4, 6}
	s2 := []int{1, 2, 5, 6}
	s3 := []int{0, 4, 5, 7}

	assert.Equal([]int{1, 3, 4, 6, 2, 5, 0, 7}, Union(s1, s2, s3))
	assert.Equal([]int{1, 3, 4, 6, 2, 5}, Union(s1, s2))
	assert.Equal([]int{1, 3, 4, 6}, Union(s1))
}

func TestUnionBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUnionBy")

	testFunc := func(i int) int {
		return i / 2
	}

	result := UnionBy(testFunc, []int{0, 1, 2, 3, 4, 5}, []int{0, 2, 10})
	assert.Equal(result, []int{0, 2, 4, 10})
}

func TestMerge(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMerge")

	s1 := []int{1, 2, 3, 4}
	s2 := []int{2, 3, 4, 5}
	s3 := []int{4, 5, 6}

	assert.Equal([]int{1, 2, 3, 4, 2, 3, 4, 5, 4, 5, 6}, Merge(s1, s2, s3))
	assert.Equal([]int{1, 2, 3, 4, 2, 3, 4, 5}, Merge(s1, s2))
	assert.Equal([]int{2, 3, 4, 5, 4, 5, 6}, Merge(s2, s3))
}

func TestIntersection(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIntersection")

	s1 := []int{1, 2, 2, 3}
	s2 := []int{1, 2, 3, 4}
	s3 := []int{0, 2, 3, 5, 6}
	s4 := []int{0, 5, 6}

	expected := [][]int{
		{2, 3},
		{1, 2, 3},
		{1, 2, 3},
		{},
	}
	result := []any{
		Intersection(s1, s2, s3),
		Intersection(s1, s2),
		Intersection(s1),
		Intersection(s1, s4),
	}

	for i := 0; i < len(result); i++ {
		assert.Equal(expected[i], result[i])
	}
}

func TestSymmetricDifference(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSymmetricDifference")

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s3 := []int{1, 2, 3, 5}

	assert.Equal([]int{1, 2, 3}, SymmetricDifference(s1))
	assert.Equal([]int{3, 4}, SymmetricDifference(s1, s2))
	assert.Equal([]int{3, 4, 5}, SymmetricDifference(s1, s2, s3))
}

func TestReverse(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestReverse")

	s1 := []int{1, 2, 3, 4, 5}
	Reverse(s1)
	assert.Equal([]int{5, 4, 3, 2, 1}, s1)

	s2 := []string{"a", "b", "c", "d", "e"}
	Reverse(s2)
	assert.Equal([]string{"e", "d", "c", "b", "a"}, s2)
}

func TestDifference(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDifference")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6}

	assert.Equal([]int{1, 2, 3}, Difference(s1, s2))
}

func TestDifferenceWith(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDifferenceWith")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6, 7, 8}
	isDouble := func(v1, v2 int) bool {
		return v2 == 2*v1
	}

	assert.Equal([]int{1, 5}, DifferenceWith(s1, s2, isDouble))
}

func TestDifferenceBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDifferenceBy")

	s1 := []int{1, 2, 3, 4, 5} // after add one: 2 3 4 5 6
	s2 := []int{3, 4, 5}       // after add one: 4 5 6
	addOne := func(i int, v int) int {
		return v + 1
	}

	assert.Equal([]int{1, 2}, DifferenceBy(s1, s2, addOne))
}

func TestIsAscending(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsAscending")

	assert.Equal(true, IsAscending([]int{1, 2, 3, 4, 5}))
	assert.Equal(false, IsAscending([]int{5, 4, 3, 2, 1}))
	assert.Equal(false, IsAscending([]int{2, 1, 3, 4, 5}))
}

func TestIsDescending(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsDescending")

	assert.Equal(true, IsDescending([]int{5, 4, 3, 2, 1}))
	assert.Equal(false, IsDescending([]int{1, 2, 3, 4, 5}))
	assert.Equal(false, IsDescending([]int{2, 1, 3, 4, 5}))
}

func TestIsSorted(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsSorted")

	assert.Equal(true, IsSorted([]int{5, 4, 3, 2, 1}))
	assert.Equal(true, IsSorted([]int{1, 2, 3, 4, 5}))
	assert.Equal(false, IsSorted([]int{2, 1, 3, 4, 5}))
}

func TestIsSortedByKey(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIsSortedByKey")

	assert.Equal(true, IsSortedByKey([]string{"a", "ab", "abc"}, func(s string) int {
		return len(s)
	}))

	assert.Equal(true, IsSortedByKey([]string{"abc", "ab", "a"}, func(s string) int {
		return len(s)
	}))

	assert.Equal(false, IsSortedByKey([]string{"abc", "a", "ab"}, func(s string) int {
		return len(s)
	}))
}

func TestSort(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSort")

	numbers := []int{1, 4, 3, 2, 5}

	Sort(numbers)
	assert.Equal([]int{1, 2, 3, 4, 5}, numbers)

	Sort(numbers, "desc")
	assert.Equal([]int{5, 4, 3, 2, 1}, numbers)

	strings := []string{"a", "d", "c", "b", "e"}

	Sort(strings)
	assert.Equal([]string{"a", "b", "c", "d", "e"}, strings)

	Sort(strings, "desc")
	assert.Equal([]string{"e", "d", "c", "b", "a"}, strings)
}

func TestSortBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSortBy")

	numbers := []int{1, 4, 3, 2, 5}

	SortBy(numbers, func(a, b int) bool {
		return a < b
	})
	assert.Equal([]int{1, 2, 3, 4, 5}, numbers)

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

	assert.EqualValues(15, users[0].Age)
	assert.EqualValues(21, users[1].Age)
	assert.EqualValues(100, users[2].Age)
}

func TestSortByFielDesc(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSortByFielDesc")

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
	studentsOfSortByAge := []student{
		{"b", 15},
		{"a", 10},
		{"d", 6},
		{"c", 5},
	}

	err := SortByField(students, "age", "desc")
	assert.IsNil(err)

	assert.Equal(studentsOfSortByAge, students)
}

func TestSortByFieldAsc(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSortByFieldAsc")

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
	studentsOfSortByAge := []student{
		{"c", 5},
		{"d", 6},
		{"a", 10},
		{"b", 15},
	}

	err := SortByField(students, "age")
	assert.IsNil(err)

	assert.Equal(studentsOfSortByAge, students)
}

func TestWithout(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestWithout")
	assert.Equal([]int{3, 4, 5}, Without([]int{1, 2, 3, 4, 5}, 1, 2))
	assert.Equal([]int{1, 2, 3, 4, 5}, Without([]int{1, 2, 3, 4, 5}))
}

func TestShuffle(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestShuffle")

	numbers := []int{1, 2, 3, 4, 5}
	result := Shuffle(numbers)

	assert.Equal(5, len(result))
}

func TestIndexOf(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIndexOf")

	arr := []string{"a", "a", "b", "c"}
	key := fmt.Sprintf("%p", arr)
	assert.Equal(0, IndexOf(arr, "a"))
	assert.Equal(-1, IndexOf(arr, "d"))
	assert.Equal(2, memoryHashCounter[key])

	arr1 := []int{1, 2, 3, 4, 5}
	key1 := fmt.Sprintf("%p", arr1)
	assert.Equal(3, IndexOf(arr1, 4))
	assert.Equal(-1, IndexOf(arr1, 6))
	assert.Equal(2, memoryHashCounter[key1])

	arr2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	key2 := fmt.Sprintf("%p", arr2)
	assert.Equal(2, IndexOf(arr2, 3.3))
	assert.Equal(3, IndexOf(arr2, 4.4))
	assert.Equal(-1, IndexOf(arr2, 6.6))
	assert.Equal(3, memoryHashCounter[key2])

	for i := 0; i < 6; i++ {
		a := []string{"a", "b", "c"}
		IndexOf(a, "a")
		IndexOf(a, "b")
	}
	minArr := []string{"c", "b", "a"}
	minKey := fmt.Sprintf("%p", minArr)
	assert.Equal(0, IndexOf(minArr, "c"))

	arr3 := []string{"q", "w", "e"}
	key3 := fmt.Sprintf("%p", arr3)
	assert.Equal(1, IndexOf(arr3, "w"))
	assert.Equal(-1, IndexOf(arr3, "r"))
	assert.Equal(2, memoryHashCounter[key3])
	assert.Equal(0, memoryHashCounter[minKey])
}

func TestLastIndexOf(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestLastIndexOf")

	arr := []string{"a", "b", "b", "c"}
	assert.Equal(0, LastIndexOf(arr, "a"))
	assert.Equal(2, LastIndexOf(arr, "b"))
	assert.Equal(-1, LastIndexOf(arr, "d"))
}

func TestToSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToSlice")

	str1 := "a"
	str2 := "b"
	assert.Equal([]string{"a"}, ToSlice(str1))
	assert.Equal([]string{"a", "b"}, ToSlice(str1, str2))
}

func TestToSlicePointer(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToSlicePointer")

	str1 := "a"
	str2 := "b"
	assert.Equal([]*string{&str1}, ToSlicePointer(str1))
	assert.Equal([]*string{&str1, &str2}, ToSlicePointer(str1, str2))
}

func TestAppendIfAbsent(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestAppendIfAbsent")

	str1 := []string{"a", "b"}
	assert.Equal([]string{"a", "b"}, AppendIfAbsent(str1, "a"))
	assert.Equal([]string{"a", "b", "c"}, AppendIfAbsent(str1, "c"))
}

func TestReplace(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestReplace")

	strs := []string{"a", "b", "a", "c", "d", "a"}

	assert.Equal([]string{"a", "b", "a", "c", "d", "a"}, Replace(strs, "a", "x", 0))
	assert.Equal([]string{"x", "b", "a", "c", "d", "a"}, Replace(strs, "a", "x", 1))
	assert.Equal([]string{"x", "b", "x", "c", "d", "a"}, Replace(strs, "a", "x", 2))
	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, Replace(strs, "a", "x", 3))
	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, Replace(strs, "a", "x", 4))

	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, Replace(strs, "a", "x", -1))
	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, Replace(strs, "a", "x", -2))

	assert.Equal([]string{"a", "b", "a", "c", "d", "a"}, Replace(strs, "x", "y", 1))
	assert.Equal([]string{"a", "b", "a", "c", "d", "a"}, Replace(strs, "x", "y", -1))
}

func TestReplaceAll(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestReplaceAll")

	strs := []string{"a", "b", "a", "c", "d", "a"}

	assert.Equal([]string{"x", "b", "x", "c", "d", "x"}, ReplaceAll(strs, "a", "x"))
	assert.Equal([]string{"a", "b", "a", "c", "d", "a"}, ReplaceAll(strs, "e", "x"))
}

func TestKeyBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestKeyBy")

	result := KeyBy([]string{"a", "ab", "abc"}, func(str string) int {
		return len(str)
	})

	assert.Equal(map[int]string{1: "a", 2: "ab", 3: "abc"}, result)
}

func TestRepeat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRepeat")

	assert.Equal([]string{}, Repeat("a", 0))
	assert.Equal([]string{"a", "a", "a", "a", "a", "a"}, Repeat("a", 6))
}

func TestJoin(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestJoin")

	nums := []int{1, 2, 3, 4, 5}

	result1 := Join(nums, ",")
	result2 := Join(nums, "-")

	assert.Equal("1,2,3,4,5", result1)
	assert.Equal("1-2-3-4-5", result2)
}

func TestPartition(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestPartition")

	assert.Equal([][]int{{1, 2, 3, 4, 5}}, Partition([]int{1, 2, 3, 4, 5}))
	assert.Equal([][]int{{2, 4}, {1, 3, 5}}, Partition([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 }))
	assert.Equal([][]int{{1, 2}, {3, 4}, {5}}, Partition([]int{1, 2, 3, 4, 5}, func(n int) bool { return n == 1 || n == 2 }, func(n int) bool { return n == 2 || n == 3 || n == 4 }))
}

func TestRandom(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRandom")

	var arr []int
	var val, idx int

	_, idx = Random(arr)
	assert.Equal(-1, idx)

	arr = []int{}
	_, idx = Random(arr)
	assert.Equal(-1, idx)

	arr = []int{1}
	val, idx = Random(arr)
	assert.Equal(0, idx)
	assert.Equal(arr[idx], val)

	arr = []int{1, 2, 3}
	val, idx = Random(arr)
	assert.Greater(len(arr), idx)
	assert.Equal(arr[idx], val)
}

func TestSetToDefaultIf(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "SetToDefaultIf")

	// Subtest for strings
	t.Run("strings", func(t *testing.T) {
		strs := []string{"a", "b", "a", "c", "d", "a"}
		actualStrs, count := SetToDefaultIf(strs, func(s string) bool { return "a" == s })
		assert.Equal([]string{"", "b", "", "c", "d", ""}, actualStrs)
		assert.Equal(3, count)
	})

	// Subtest for integers
	t.Run("integers", func(t *testing.T) {
		ints := []int{1, 2, 3, 2, 4, 2}
		actualInts, count := SetToDefaultIf(ints, func(i int) bool { return i == 2 })
		assert.Equal([]int{1, 0, 3, 0, 4, 0}, actualInts)
		assert.Equal(3, count)
	})

	// Subtest for floating-point numbers
	t.Run("floats", func(t *testing.T) {
		floats := []float64{1.1, 2.2, 3.3, 2.2, 4.4, 2.2}
		actualFloats, count := SetToDefaultIf(floats, func(f float64) bool { return f == 2.2 })
		assert.Equal([]float64{1.1, 0, 3.3, 0, 4.4, 0}, actualFloats)
		assert.Equal(3, count)
	})

	// Subtest for booleans
	t.Run("booleans", func(t *testing.T) {
		bools := []bool{true, false, true, true, false}
		actualBools, count := SetToDefaultIf(bools, func(b bool) bool { return b })
		assert.Equal([]bool{false, false, false, false, false}, actualBools)
		assert.Equal(3, count)
	})

	// Subtest for a custom type
	type customType struct {
		field string
	}
	t.Run("customType", func(t *testing.T) {
		customs := []customType{{"a"}, {"b"}, {"a"}, {"c"}}
		actualCustoms, count := SetToDefaultIf(customs, func(c customType) bool { return c.field == "a" })
		expected := []customType{{""}, {"b"}, {""}, {"c"}}
		assert.Equal(expected, actualCustoms)
		assert.Equal(2, count)
	})

	// Subtest for slice of integers
	t.Run("sliceOfInts", func(t *testing.T) {
		sliceOfInts := [][]int{{1, 2}, {3, 4}, {5, 6}, {1, 2}}
		actualSlice, count := SetToDefaultIf(sliceOfInts, func(s []int) bool { return reflect.DeepEqual(s, []int{1, 2}) })
		expected := [][]int{nil, {3, 4}, {5, 6}, nil}
		assert.Equal(expected, actualSlice)
		assert.Equal(2, count)
	})

	// Subtest for maps (simple use case)
	t.Run("mapOfStringToInts", func(t *testing.T) {
		maps := []map[string]int{{"a": 1}, {"b": 2}, {"a": 1}, {"c": 3}}
		actualMaps, count := SetToDefaultIf(maps, func(m map[string]int) bool { _, ok := m["a"]; return ok })
		expected := []map[string]int{nil, {"b": 2}, nil, {"c": 3}}
		assert.Equal(expected, actualMaps)
		assert.Equal(2, count)
	})

	// Subtest for pointers to integers
	t.Run("pointersToInts", func(t *testing.T) {
		one, two, three := 1, 2, 3
		pointers := []*int{&one, &two, &one, &three}
		actualPointers, count := SetToDefaultIf(pointers, func(p *int) bool { return p != nil && *p == 1 })
		expected := []*int{nil, &two, nil, &three}
		assert.Equal(expected, actualPointers)
		assert.Equal(2, count)
	})

	// Subtest for channels
	t.Run("channels", func(t *testing.T) {
		ch1, ch2 := make(chan int), make(chan int)
		channels := []chan int{ch1, ch2, ch1}
		actualChannels, count := SetToDefaultIf(channels, func(ch chan int) bool { return ch == ch1 })
		expected := []chan int{nil, ch2, nil}
		assert.Equal(expected, actualChannels)
		assert.Equal(2, count)
	})

	// Subtest for interfaces
	t.Run("interfaces", func(t *testing.T) {
		var i1, i2 interface{} = "hello", 42
		interfaces := []interface{}{i1, i2, i1}
		actualInterfaces, count := SetToDefaultIf(interfaces, func(i interface{}) bool { _, ok := i.(string); return ok })
		expected := []interface{}{nil, 42, nil}
		assert.Equal(expected, actualInterfaces)
		assert.Equal(2, count)
	})

	// Subtest for complex structs
	t.Run("complexStructs", func(t *testing.T) {
		type ComplexStruct struct {
			Name  string
			Value int
			Data  []byte
		}
		cs1, cs2 := ComplexStruct{Name: "Test", Value: 1, Data: []byte{1, 2, 3}}, ComplexStruct{Name: "Another", Value: 2, Data: []byte{4, 5, 6}}
		complexStructs := []ComplexStruct{cs1, cs2, cs1}
		actualComplexStructs, count := SetToDefaultIf(complexStructs, func(cs ComplexStruct) bool { return cs.Name == "Test" })
		expected := []ComplexStruct{{}, cs2, {}}
		assert.Equal(expected, actualComplexStructs)
		assert.Equal(2, count)
	})

	// Subtest for uints
	t.Run("uints", func(t *testing.T) {
		uints := []uint{1, 2, 3, 2, 4, 2}
		actualUints, count := SetToDefaultIf(uints, func(u uint) bool { return u == 2 })
		assert.Equal([]uint{1, 0, 3, 0, 4, 0}, actualUints)
		assert.Equal(3, count)
	})

	// Subtest for float32
	t.Run("float32s", func(t *testing.T) {
		floats := []float32{1.1, 2.2, 3.3, 2.2, 4.4, 2.2}
		actualFloats, count := SetToDefaultIf(floats, func(f float32) bool { return f == 2.2 })
		assert.Equal([]float32{1.1, 0, 3.3, 0, 4.4, 0}, actualFloats)
		assert.Equal(3, count)
	})

	// Subtest for []byte
	t.Run("byteSlices", func(t *testing.T) {
		bytes := [][]byte{{'a', 'b'}, {'c', 'd'}, {'a', 'b'}}
		actualBytes, count := SetToDefaultIf(bytes, func(b []byte) bool { return string(b) == "ab" })
		expected := [][]byte{nil, {'c', 'd'}, nil}
		assert.Equal(expected, actualBytes)
		assert.Equal(2, count)
	})
}

func TestBreak(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBreak")

	// Test with integers
	nums := []int{1, 2, 3, 4, 5}
	even := func(n int) bool { return n%2 == 0 }

	resultEven, resultAfterFirstEven := Break(nums, even)
	assert.Equal([]int{1}, resultEven)
	assert.Equal([]int{2, 3, 4, 5}, resultAfterFirstEven)

	// Test with strings
	strings := []string{"apple", "banana", "cherry", "date", "elderberry"}
	startsWithA := func(s string) bool { return s[0] == 'a' }

	resultStartsWithA, resultAfterFirstStartsWithA := Break(strings, startsWithA)
	assert.Equal([]string{}, resultStartsWithA)
	assert.Equal([]string{"apple", "banana", "cherry", "date", "elderberry"}, resultAfterFirstStartsWithA)

	// Test with empty slice
	emptySlice := []int{}
	resultEmpty, _ := Break(emptySlice, even)
	assert.Equal([]int{}, resultEmpty)

	// Test with all elements satisfying the predicate
	allEven := []int{2, 4, 6, 8, 10}
	emptyResult, resultAllEven := Break(allEven, even)
	assert.Equal([]int{2, 4, 6, 8, 10}, resultAllEven)
	assert.Equal([]int{}, emptyResult)

	// Test with no elements satisfying the predicate
	allOdd := []int{1, 3, 5, 7, 9}
	resultAllOdd, emptyResult := Break(allOdd, even)
	assert.Equal([]int{1, 3, 5, 7, 9}, resultAllOdd)
	assert.Equal([]int{}, emptyResult)
}

func TestRightPaddingAndLeftPadding(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "RightPaddingAndLeftPadding")

	// Test with integers
	nums := []int{1, 2, 3, 4, 5}

	padded := LeftPadding(RightPadding(nums, 0, 3), 0, 3)
	assert.Equal([]int{0, 0, 0, 1, 2, 3, 4, 5, 0, 0, 0}, padded)
}

func TestUniqueByConcurrent(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUniqueByConcurrent")

	nums := []int{1, 2, 3, 1, 2, 4, 5, 6, 4, 7}
	comparator := func(item int, other int) bool { return item == other }

	result := UniqueByConcurrent(nums, comparator, 4)

	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7}, result)
}

func TestMapConcurrent(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMapConcurrent")

	t.Run("empty slice", func(t *testing.T) {
		actual := MapConcurrent([]int{}, func(_, n int) int { return n * n }, 4)
		assert.Equal([]int{}, actual)
	})

	t.Run("single thread", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6}
		expected := []int{1, 4, 9, 16, 25, 36}
		actual := MapConcurrent(nums, func(_, n int) int { return n * n }, 1)
		assert.Equal(expected, actual)
	})

	t.Run("multiple threads", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6}
		expected := []int{1, 4, 9, 16, 25, 36}
		actual := MapConcurrent(nums, func(_, n int) int { return n * n }, 4)
		assert.Equal(expected, actual)
	})

}

func TestFilterConcurrent(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFilterConcurrent")

	t.Run("empty slice", func(t *testing.T) {
		actual := FilterConcurrent([]int{}, func(_, n int) bool { return n != 0 }, 4)
		assert.Equal([]int{}, actual)
	})

	t.Run("single thread", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6}
		expected := []int{4, 5, 6}
		actual := FilterConcurrent(nums, func(_, n int) bool { return n > 3 }, 1)
		assert.Equal(expected, actual)
	})

	t.Run("multiple threads", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6}
		expected := []int{4, 5, 6}
		actual := FilterConcurrent(nums, func(_, n int) bool { return n > 3 }, 4)
		assert.Equal(expected, actual)
	})

}

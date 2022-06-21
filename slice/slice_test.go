package slice

import (
	"math"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestContain(t *testing.T) {
	assert := internal.NewAssert(t, "TestContain")

	assert.Equal(true, Contain([]string{"a", "b", "c"}, "a"))
	assert.Equal(false, Contain([]string{"a", "b", "c"}, "d"))
	assert.Equal(true, Contain([]string{""}, ""))
	assert.Equal(false, Contain([]string{}, ""))

	assert.Equal(true, Contain([]int{1, 2, 3}, 1))
}

func TestContainSubSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestContainSubSlice")
	assert.Equal(true, ContainSubSlice([]string{"a", "a", "b", "c"}, []string{"a", "a"}))
	assert.Equal(false, ContainSubSlice([]string{"a", "a", "b", "c"}, []string{"a", "d"}))
	assert.Equal(true, ContainSubSlice([]int{1, 2, 3}, []int{1, 2}))
	assert.Equal(false, ContainSubSlice([]int{1, 2, 3}, []int{0, 1}))
}

func TestChunk(t *testing.T) {
	assert := internal.NewAssert(t, "TestChunk")

	arr := []string{"a", "b", "c", "d", "e"}

	r1 := [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}}
	assert.Equal(r1, Chunk(arr, 1))

	r2 := [][]string{{"a", "b"}, {"c", "d"}, {"e"}}
	assert.Equal(r2, Chunk(arr, 2))

	r3 := [][]string{{"a", "b", "c"}, {"d", "e"}}
	assert.Equal(r3, Chunk(arr, 3))

	r4 := [][]string{{"a", "b", "c", "d"}, {"e"}}
	assert.Equal(r4, Chunk(arr, 4))

	r5 := [][]string{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}}
	assert.Equal(r5, Chunk(arr, 5))
}

func TestCompact(t *testing.T) {
	assert := internal.NewAssert(t, "TesCompact")

	assert.Equal([]int{}, Compact([]int{0}))
	assert.Equal([]int{1, 2, 3}, Compact([]int{0, 1, 2, 3}))
	assert.Equal([]string{}, Compact([]string{""}))
	assert.Equal([]string{" "}, Compact([]string{" "}))
	assert.Equal([]string{"a", "b", "0"}, Compact([]string{"", "a", "b", "0"}))
	assert.Equal([]bool{true, true}, Compact([]bool{false, true, true}))
}

func TestConcat(t *testing.T) {
	assert := internal.NewAssert(t, "Concat")

	assert.Equal([]int{1, 2, 3, 4, 5}, Concat([]int{1, 2, 3}, []int{4, 5}))
	assert.Equal([]int{1, 2, 3, 4, 5}, Concat([]int{1, 2, 3}, []int{4}, []int{5}))
}

func TestEqual(t *testing.T) {
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
	assert := internal.NewAssert(t, "TestEqualWith")

	slice1 := []int{1, 2, 3}
	slice2 := []int{2, 4, 6}

	isDouble := func(a, b int) bool {
		return b == a*2
	}

	assert.Equal(true, EqualWith(slice1, slice2, isDouble))
}

func TestEvery(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	assert := internal.NewAssert(t, "TestEvery")
	assert.Equal(false, Every(nums, isEven))
}

func TestNone(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	check := func(i, num int) bool {
		return num%2 == 1
	}

	assert := internal.NewAssert(t, "TestNone")
	assert.Equal(false, None(nums, check))
}

func TestSome(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	hasEven := func(i, num int) bool {
		return num%2 == 0
	}

	assert := internal.NewAssert(t, "TestSome")
	assert.Equal(true, Some(nums, hasEven))
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}

	assert := internal.NewAssert(t, "TestFilter")
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
	nums := []int{1, 2, 3, 4, 5, 6}
	evenFunc := func(i, num int) bool {
		return (num % 2) == 0
	}
	expectedEven := []int{2, 4, 6}
	expectedOdd := []int{1, 3, 5}
	even, odd := GroupBy(nums, evenFunc)

	assert := internal.NewAssert(t, "TestGroupBy")
	assert.Equal(expectedEven, even)
	assert.Equal(expectedOdd, odd)
}

func TestGroupWith(t *testing.T) {
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
	nums := []int{1, 2, 3, 4, 5, 6}
	evenFunc := func(i, num int) bool {
		return (num % 2) == 0
	}

	assert := internal.NewAssert(t, "TestCount")
	assert.Equal(3, Count(nums, evenFunc))
}

func TestFind(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	even := func(i, num int) bool {
		return num%2 == 0
	}
	res, ok := Find(nums, even)
	if !ok {
		t.Fatal("found nothing")
	}

	assert := internal.NewAssert(t, "TestFind")
	assert.Equal(2, *res)
}

func TestFindLast(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	even := func(i, num int) bool {
		return num%2 == 0
	}
	res, ok := FindLast(nums, even)
	if !ok {
		t.Fatal("found nothing")
	}

	assert := internal.NewAssert(t, "TestFindLast")
	assert.Equal(4, *res)
}

func TestFindFoundNothing(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1, 1}
	findFunc := func(i, num int) bool {
		return num > 1
	}
	_, ok := Find(nums, findFunc)
	// if ok {
	// 	t.Fatal("found something")
	// }
	assert := internal.NewAssert(t, "TestFindFoundNothing")
	assert.Equal(false, ok)
}

func TestFlattenDeep(t *testing.T) {
	input := [][][]string{{{"a", "b"}}, {{"c", "d"}}}
	expected := []string{"a", "b", "c", "d"}

	assert := internal.NewAssert(t, "TestFlattenDeep")
	assert.Equal(expected, FlattenDeep(input))
}

func TestForEach(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{3, 4, 5, 6, 7}

	var numbersAddTwo []int
	addTwo := func(index int, value int) {
		numbersAddTwo = append(numbersAddTwo, value+2)
	}

	ForEach(numbers, addTwo)

	assert := internal.NewAssert(t, "TestForEach")
	assert.Equal(expected, numbersAddTwo)
}

func TestMap(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	multiplyTwo := func(i, num int) int {
		return num * 2
	}

	assert := internal.NewAssert(t, "TestMap")
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

func TestReduce(t *testing.T) {
	cases := [][]int{
		{},
		{1},
		{1, 2, 3, 4}}

	expected := []int{0, 1, 10}

	f := func(i, v1, v2 int) int {
		return v1 + v2
	}

	assert := internal.NewAssert(t, "TestReduce")

	for i := 0; i < len(cases); i++ {
		actual := Reduce(cases[i], f, 0)
		assert.Equal(expected[i], actual)
	}
}

func TestIntSlice(t *testing.T) {
	var nums []any
	nums = append(nums, 1, 2, 3)

	assert := internal.NewAssert(t, "TestIntSlice")
	assert.Equal([]int{1, 2, 3}, IntSlice(nums))
}

func TestStringSlice(t *testing.T) {
	var strs []any
	strs = append(strs, "a", "b", "c")

	assert := internal.NewAssert(t, "TestStringSlice")
	assert.Equal([]string{"a", "b", "c"}, StringSlice(strs))
}

func TestInterfaceSlice(t *testing.T) {
	strs := []string{"a", "b", "c"}
	expect := []any{"a", "b", "c"}

	assert := internal.NewAssert(t, "TestInterfaceSlice")
	assert.Equal(expect, InterfaceSlice(strs))
}

func TestDeleteAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteAt")

	assert.Equal([]string{"a", "b", "c"}, DeleteAt([]string{"a", "b", "c"}, -1))
	assert.Equal([]string{"a", "b", "c"}, DeleteAt([]string{"a", "b", "c"}, 3))
	assert.Equal([]string{"b", "c"}, DeleteAt([]string{"a", "b", "c"}, 0))
	assert.Equal([]string{"a", "c"}, DeleteAt([]string{"a", "b", "c"}, 1))
	assert.Equal([]string{"a", "b"}, DeleteAt([]string{"a", "b", "c"}, 2))

	assert.Equal([]string{"b", "c"}, DeleteAt([]string{"a", "b", "c"}, 0, 1))
	assert.Equal([]string{"c"}, DeleteAt([]string{"a", "b", "c"}, 0, 2))
	assert.Equal([]string{}, DeleteAt([]string{"a", "b", "c"}, 0, 3))
	assert.Equal([]string{}, DeleteAt([]string{"a", "b", "c"}, 0, 4))
	assert.Equal([]string{"a"}, DeleteAt([]string{"a", "b", "c"}, 1, 3))
	assert.Equal([]string{"a"}, DeleteAt([]string{"a", "b", "c"}, 1, 4))
}

func TestDrop(t *testing.T) {
	assert := internal.NewAssert(t, "TestInterfaceSlice")

	assert.Equal([]int{}, Drop([]int{}, 0))
	assert.Equal([]int{}, Drop([]int{}, 1))
	assert.Equal([]int{}, Drop([]int{}, -1))

	assert.Equal([]int{1, 2, 3, 4, 5}, Drop([]int{1, 2, 3, 4, 5}, 0))
	assert.Equal([]int{2, 3, 4, 5}, Drop([]int{1, 2, 3, 4, 5}, 1))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, 5))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, 6))

	assert.Equal([]int{1, 2, 3, 4}, Drop([]int{1, 2, 3, 4, 5}, -1))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, -6))
	assert.Equal([]int{}, Drop([]int{1, 2, 3, 4, 5}, -6))
}

func TestInsertAt(t *testing.T) {
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
	t.Log(strs)
}

func TestUpdateAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateAt")

	assert.Equal([]string{"a", "b", "c"}, UpdateAt([]string{"a", "b", "c"}, -1, "1"))
	assert.Equal([]string{"1", "b", "c"}, UpdateAt([]string{"a", "b", "c"}, 0, "1"))
	assert.Equal([]string{"a", "b", "2"}, UpdateAt([]string{"a", "b", "c"}, 2, "2"))
	assert.Equal([]string{"a", "b", "c"}, UpdateAt([]string{"a", "b", "c"}, 3, "2"))
}

func TestUnique(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnique")

	assert.Equal([]int{1, 2, 3}, Unique([]int{1, 2, 2, 3}))
	assert.Equal([]string{"a", "b", "c"}, Unique([]string{"a", "a", "b", "c"}))
}

func TestUniqueBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestUniqueBy")

	actual := UniqueBy([]int{1, 2, 3, 4, 5, 6}, func(val int) int {
		return val % 4
	})
	assert.Equal([]int{1, 2, 3, 0}, actual)
}

func TestUnion(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnion")

	s1 := []int{1, 3, 4, 6}
	s2 := []int{1, 2, 5, 6}
	s3 := []int{0, 4, 5, 7}

	assert.Equal([]int{1, 3, 4, 6, 2, 5, 0, 7}, Union(s1, s2, s3))
	assert.Equal([]int{1, 3, 4, 6, 2, 5}, Union(s1, s2))
	assert.Equal([]int{1, 3, 4, 6}, Union(s1))
}

func TestIntersection(t *testing.T) {
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
	res := []any{
		Intersection(s1, s2, s3),
		Intersection(s1, s2),
		Intersection(s1),
		Intersection(s1, s4),
	}

	assert := internal.NewAssert(t, "TestIntersection")

	for i := 0; i < len(res); i++ {
		assert.Equal(expected[i], res[i])
	}
}

func TestSymmetricDifference(t *testing.T) {
	assert := internal.NewAssert(t, "TestSymmetricDifference")

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s3 := []int{1, 2, 3, 5}

	assert.Equal([]int{1, 2, 3}, SymmetricDifference(s1))
	assert.Equal([]int{3, 4}, SymmetricDifference(s1, s2))
	assert.Equal([]int{3, 4, 5}, SymmetricDifference(s1, s2, s3))
}

func TestReverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestReverse")

	s1 := []int{1, 2, 3, 4, 5}
	Reverse(s1)
	assert.Equal([]int{5, 4, 3, 2, 1}, s1)

	s2 := []string{"a", "b", "c", "d", "e"}
	Reverse(s2)
	assert.Equal([]string{"e", "d", "c", "b", "a"}, s2)
}

func TestDifference(t *testing.T) {
	assert := internal.NewAssert(t, "TestDifference")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6}
	assert.Equal([]int{1, 2, 3}, Difference(s1, s2))
}

func TestDifferenceWith(t *testing.T) {
	assert := internal.NewAssert(t, "TestDifferenceWith")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6, 7, 8}
	isDouble := func(v1, v2 int) bool {
		return v2 == 2*v1
	}
	assert.Equal([]int{1, 5}, DifferenceWith(s1, s2, isDouble))
}

func TestDifferenceBy(t *testing.T) {
	assert := internal.NewAssert(t, "TestDifferenceBy")

	s1 := []int{1, 2, 3, 4, 5} //after add one: 2 3 4 5 6
	s2 := []int{3, 4, 5}       //after add one: 4 5 6
	addOne := func(i int, v int) int {
		return v + 1
	}
	assert.Equal([]int{1, 2}, DifferenceBy(s1, s2, addOne))
}

func TestSortByFielDesc(t *testing.T) {
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

	assert.Equal(students, studentsOfSortByAge)
}

func TestSortByFieldAsc(t *testing.T) {
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

	assert.Equal(students, studentsOfSortByAge)
}

func TestWithout(t *testing.T) {
	assert := internal.NewAssert(t, "TestWithout")
	assert.Equal([]int{3, 4, 5}, Without([]int{1, 2, 3, 4, 5}, 1, 2))
	assert.Equal([]int{1, 2, 3, 4, 5}, Without([]int{1, 2, 3, 4, 5}))
}

func TestShuffle(t *testing.T) {
	assert := internal.NewAssert(t, "TestShuffle")

	s := []int{1, 2, 3, 4, 5}
	res := Shuffle(s)
	t.Log("Shuffle result: ", res)

	assert.Equal(5, len(res))
}

func TestIndexOf(t *testing.T) {
	assert := internal.NewAssert(t, "TestIndexOf")

	arr := []string{"a", "a", "b", "c"}
	assert.Equal(0, IndexOf(arr, "a"))
	assert.Equal(-1, IndexOf(arr, "d"))
}

func TestLastIndexOf(t *testing.T) {
	assert := internal.NewAssert(t, "TestLastIndexOf")

	arr := []string{"a", "a", "b", "c"}
	assert.Equal(1, LastIndexOf(arr, "a"))
	assert.Equal(-1, LastIndexOf(arr, "d"))
}

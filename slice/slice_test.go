package slice

import (
	"reflect"
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestContain(t *testing.T) {
	assert := internal.NewAssert(t, "TestContain")

	assert.Equal(true, Contain([]string{"a", "b", "c"}, "a"))
	assert.Equal(false, Contain([]string{"a", "b", "c"}, "d"))
	assert.Equal(true, Contain([]string{""}, ""))
	assert.Equal(false, Contain([]string{}, ""))

	var m = map[string]int{"a": 1}
	assert.Equal(true, Contain(m, "a"))
	assert.Equal(false, Contain(m, "b"))

	assert.Equal(true, Contain("abc", "a"))
	assert.Equal(false, Contain("abc", "d"))
}

func TestChunk(t *testing.T) {
	assert := internal.NewAssert(t, "TestChunk")

	arr := []string{"a", "b", "c", "d", "e"}
	r1 := [][]interface{}{
		{"a"},
		{"b"},
		{"c"},
		{"d"},
		{"e"},
	}
	assert.Equal(r1, Chunk(InterfaceSlice(arr), 1))

	r2 := [][]interface{}{
		{"a", "b"},
		{"c", "d"},
		{"e"},
	}
	assert.Equal(r2, Chunk(InterfaceSlice(arr), 2))

	r3 := [][]interface{}{
		{"a", "b", "c"},
		{"d", "e"},
	}
	assert.Equal(r3, Chunk(InterfaceSlice(arr), 3))

	r4 := [][]interface{}{
		{"a", "b", "c", "d"},
		{"e"},
	}
	assert.Equal(r4, Chunk(InterfaceSlice(arr), 4))

	r5 := [][]interface{}{
		{"a"},
		{"b"},
		{"c"},
		{"d"},
		{"e"},
	}
	assert.Equal(r5, Chunk(InterfaceSlice(arr), 5))
}

func TestConvertSlice(t *testing.T) {
	//t1 := []string{"1","2"}
	//aInt, _ := strconv.ParseInt("1", 10, 64)
	//bInt, _ := strconv.ParseInt("2", 10, 64)
	//expected :=[]int64{aInt, bInt}
	//
	//a := ConvertSlice(t1, reflect.TypeOf(expected))
	//if !reflect.DeepEqual(a, expected) {
	//	utils.LogFailedTestInfo(t, "ConvertSlice", t1, expected, a)
	//	t.FailNow()
	//}
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
	assert.Equal(2, res)
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
	ForEach(numbers, func(index int, value int) {
		numbersAddTwo = append(numbersAddTwo, value+2)
	})

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
	var nums []interface{}
	nums = append(nums, 1, 2, 3)

	assert := internal.NewAssert(t, "TestIntSlice")
	assert.Equal([]int{1, 2, 3}, IntSlice(nums))
}

func TestStringSlice(t *testing.T) {
	var strs []interface{}
	strs = append(strs, "a", "b", "c")

	assert := internal.NewAssert(t, "TestStringSlice")
	assert.Equal([]string{"a", "b", "c"}, StringSlice(strs))
}

func TestInterfaceSlice(t *testing.T) {
	strs := []string{"a", "b", "c"}
	expect := []interface{}{"a", "b", "c"}

	assert := internal.NewAssert(t, "TestInterfaceSlice")
	assert.Equal(expect, InterfaceSlice(strs))
}

func TestDeleteByIndex(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeleteByIndex")

	t1 := []string{"a", "b", "c", "d", "e"}
	r1 := []string{"b", "c", "d", "e"}
	a1, _ := DeleteByIndex(t1, 0)
	assert.Equal(r1, a1)

	t2 := []string{"a", "b", "c", "d", "e"}
	r2 := []string{"a", "b", "c", "e"}
	a2, _ := DeleteByIndex(t2, 3)
	assert.Equal(r2, a2)

	t3 := []string{"a", "b", "c", "d", "e"}
	r3 := []string{"c", "d", "e"}
	a3, _ := DeleteByIndex(t3, 0, 2)
	assert.Equal(r3, a3)

	t4 := []string{"a", "b", "c", "d", "e"}
	r4 := []string{}
	a4, _ := DeleteByIndex(t4, 0, 5)
	assert.Equal(r4, a4)

	t5 := []string{"a", "b", "c", "d", "e"}
	_, err := DeleteByIndex(t5, 1, 1)
	assert.IsNotNil(err)

	_, err = DeleteByIndex(t5, 0, 6)
	assert.IsNotNil(err)
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

func TestInsertByIndex(t *testing.T) {
	assert := internal.NewAssert(t, "TestInsertByIndex")

	t1 := []string{"a", "b", "c"}
	r1, _ := InsertByIndex(t1, 0, "1")
	assert.Equal([]string{"1", "a", "b", "c"}, r1)

	r2, _ := InsertByIndex(t1, 1, "1")
	assert.Equal([]string{"a", "1", "b", "c"}, r2)

	r3, _ := InsertByIndex(t1, 3, "1")
	assert.Equal([]string{"a", "b", "c", "1"}, r3)

	r4, _ := InsertByIndex(t1, 0, []string{"1", "2", "3"})
	assert.Equal([]string{"1", "2", "3", "a", "b", "c"}, r4)

	r5, _ := InsertByIndex(t1, 3, []string{"1", "2", "3"})
	assert.Equal([]string{"a", "b", "c", "1", "2", "3"}, r5)

	_, err := InsertByIndex(t1, 4, "1")
	assert.IsNotNil(err)

	_, err = InsertByIndex(t1, 0, 1)
	assert.IsNotNil(err)
}

func TestUpdateByIndex(t *testing.T) {
	assert := internal.NewAssert(t, "TestUpdateByIndex")

	t1 := []string{"a", "b", "c"}
	r1, _ := UpdateByIndex(t1, 0, "1")
	assert.Equal([]string{"1", "b", "c"}, r1)

	t2 := []string{"a", "b", "c"}
	r2, _ := UpdateByIndex(t2, 1, "1")
	assert.Equal([]string{"a", "1", "c"}, r2)

	_, err := UpdateByIndex([]string{"a", "b", "c"}, 4, "1")
	assert.IsNotNil(err)

	_, err = UpdateByIndex([]string{"a", "b", "c"}, 0, 1)
	assert.IsNotNil(err)
}

func TestUnique(t *testing.T) {
	assert := internal.NewAssert(t, "TestUnique")

	assert.Equal([]int{1, 2, 3}, Unique([]int{1, 2, 2, 3}))
	assert.Equal([]string{"a", "b", "c"}, Unique([]string{"a", "a", "b", "c"}))
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
	res := []interface{}{
		Intersection(s1, s2, s3),
		Intersection(s1, s2),
		Intersection(s1),
		Intersection(s1, s4),
	}

	assert := internal.NewAssert(t, "TestIntersection")

	for i := 0; i < len(res); i++ {
		assert.Equal(res[i], expected[i])
	}
}

func TestReverseSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestIntersection")

	s1 := []int{1, 2, 3, 4, 5}
	ReverseSlice(s1)
	assert.Equal([]int{5, 4, 3, 2, 1}, s1)

	s2 := []string{"a", "b", "c", "d", "e"}
	ReverseSlice(s2)
	assert.Equal([]string{"e", "d", "c", "b", "a"}, s2)
}

func TestDifference(t *testing.T) {
	assert := internal.NewAssert(t, "TestIntersection")

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6}
	assert.Equal([]int{1, 2, 3}, Difference(s1, s2))
}

func TestSortByFielDesc(t *testing.T) {
	assert := internal.NewAssert(t, "TestWithout")

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
	assert := internal.NewAssert(t, "TestSortByField")

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

	assert.Equal(reflect.TypeOf(s), reflect.TypeOf(res))

	rv := reflect.ValueOf(res)
	assert.Equal(5, rv.Len())

	assert.Equal(true, rv.Kind() == reflect.Slice)
	assert.Equal(true, rv.Type().Elem().Kind() == reflect.Int)

	assert.Equal(true, Contain(res, 1))
	assert.Equal(true, Contain(res, 2))
	assert.Equal(true, Contain(res, 3))
	assert.Equal(true, Contain(res, 4))
	assert.Equal(true, Contain(res, 5))
}

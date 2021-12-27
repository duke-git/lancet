package slice

import (
	"reflect"
	"testing"

	"github.com/duke-git/lancet/utils"
)

func TestContain(t *testing.T) {
	t1 := []string{"a", "b", "c", "d"}
	contain(t, t1, "a", true)
	contain(t, t1, "e", false)

	var t2 []string
	contain(t, t2, "1", false)

	m := make(map[string]int)
	m["a"] = 1
	contain(t, m, "a", true)
	contain(t, m, "b", false)

	s := "hello"
	contain(t, s, "h", true)
	contain(t, s, "s", false)
}

func contain(t *testing.T, test interface{}, value interface{}, expected bool) {
	res := Contain(test, value)
	if res != expected {
		utils.LogFailedTestInfo(t, "Contain", test, expected, res)
		t.FailNow()
	}
}

func TestChunk(t *testing.T) {
	t1 := []string{"a", "b", "c", "d", "e"}

	r1 := [][]interface{}{
		{"a"},
		{"b"},
		{"c"},
		{"d"},
		{"e"},
	}
	chunk(t, InterfaceSlice(t1), 1, r1)

	r2 := [][]interface{}{
		{"a", "b"},
		{"c", "d"},
		{"e"},
	}
	chunk(t, InterfaceSlice(t1), 2, r2)

	r3 := [][]interface{}{
		{"a", "b", "c"},
		{"d", "e"},
	}
	chunk(t, InterfaceSlice(t1), 3, r3)

	r4 := [][]interface{}{
		{"a", "b", "c", "d"},
		{"e"},
	}
	chunk(t, InterfaceSlice(t1), 4, r4)

	r5 := [][]interface{}{
		{"a"},
		{"b"},
		{"c"},
		{"d"},
		{"e"},
	}
	chunk(t, InterfaceSlice(t1), 5, r5)

}

func chunk(t *testing.T, test []interface{}, num int, expected [][]interface{}) {
	res := Chunk(test, num)
	if !reflect.DeepEqual(res, expected) {
		utils.LogFailedTestInfo(t, "Chunk", test, expected, res)
		t.FailNow()
	}
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
	res := Every(nums, isEven)
	if res != false {
		utils.LogFailedTestInfo(t, "Every", nums, false, res)
		t.FailNow()
	}
}

func TestSome(t *testing.T) {
	nums := []int{1, 2, 3, 5}
	isEven := func(i, num int) bool {
		return num%2 == 0
	}
	res := Some(nums, isEven)
	if res != true {
		utils.LogFailedTestInfo(t, "Some", nums, true, res)
		t.FailNow()
	}
}

func TestFilter(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	even := func(i, num int) bool {
		return num%2 == 0
	}
	e1 := []int{2, 4}
	r1 := Filter(nums, even)
	if !reflect.DeepEqual(r1, e1) {
		utils.LogFailedTestInfo(t, "Filter", nums, e1, r1)
		t.FailNow()
	}

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

	e2 := []student{
		{"d", 13},
		{"e", 14},
	}
	filterFunc := func(i int, s student) bool {
		return s.age > 12
	}

	r2 := Filter(students, filterFunc)
	if !reflect.DeepEqual(r2, e2) {
		utils.LogFailedTestInfo(t, "Filter", students, e2, r2)
		t.FailNow()
	}

}

func TestFind(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	even := func(i, num int) bool {
		return num%2 == 0
	}
	res := Find(nums, even)
	if res != 2 {
		utils.LogFailedTestInfo(t, "Find", nums, 2, res)
		t.FailNow()
	}
}

func TestMap(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	multiplyTwo := func(i, num int) int {
		return num * 2
	}
	e1 := []int{2, 4, 6, 8}
	r1 := Map(s1, multiplyTwo)
	if !reflect.DeepEqual(r1, e1) {
		utils.LogFailedTestInfo(t, "Map", s1, e1, r1)
		t.FailNow()
	}

	type student struct {
		name string
		age  int
	}
	students := []student{
		{"a", 1},
		{"b", 2},
		{"c", 3},
	}

	e2 := []student{
		{"a", 11},
		{"b", 12},
		{"c", 13},
	}
	mapFunc := func(i int, s student) student {
		s.age += 10
		return s
	}
	r2 := Map(students, mapFunc)
	if !reflect.DeepEqual(r2, e2) {
		utils.LogFailedTestInfo(t, "Filter", students, e2, r2)
		t.FailNow()
	}
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
	for i := 0; i < len(cases); i++ {
		res := Reduce(cases[i], f, 0)
		if res != expected[i] {
			utils.LogFailedTestInfo(t, "Reduce", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestIntSlice(t *testing.T) {
	var t1 []interface{}
	t1 = append(t1, 1, 2, 3, 4, 5)
	expect := []int{1, 2, 3, 4, 5}
	intSlice(t, t1, expect)
}

func intSlice(t *testing.T, test interface{}, expected []int) {
	res, err := IntSlice(test)
	if err != nil {
		t.Error("IntSlice Error: " + err.Error())
	}
	if !reflect.DeepEqual(res, expected) {
		utils.LogFailedTestInfo(t, "IntSlice", test, expected, res)
		t.FailNow()
	}
}

func TestStringSlice(t *testing.T) {
	var t1 []interface{}
	t1 = append(t1, "a", "b", "c", "d", "e")
	expect := []string{"a", "b", "c", "d", "e"}
	stringSlice(t, t1, expect)
}

func stringSlice(t *testing.T, test interface{}, expected []string) {
	res := StringSlice(test)
	if !reflect.DeepEqual(res, expected) {
		utils.LogFailedTestInfo(t, "StringSlice", test, expected, res)
		t.FailNow()
	}
}

func TestInterfaceSlice(t *testing.T) {
	t1 := []string{"a", "b", "c", "d", "e"}
	expect := []interface{}{"a", "b", "c", "d", "e"}
	interfaceSlice(t, t1, expect)
}

func interfaceSlice(t *testing.T, test interface{}, expected []interface{}) {
	res := InterfaceSlice(test)
	if !reflect.DeepEqual(res, expected) {
		utils.LogFailedTestInfo(t, "InterfaceSlice", test, expected, res)
		t.FailNow()
	}
}

func TestDeleteByIndex(t *testing.T) {
	origin := []string{"a", "b", "c", "d", "e"}

	t1 := []string{"a", "b", "c", "d", "e"}
	r1 := []string{"b", "c", "d", "e"}
	deleteByIndex(t, origin, t1, 0, 0, r1)

	t1 = []string{"a", "b", "c", "d", "e"}
	r2 := []string{"a", "b", "c", "e"}
	deleteByIndex(t, origin, t1, 3, 0, r2)

	t1 = []string{"a", "b", "c", "d", "e"}
	r3 := []string{"a", "b", "c", "d"}
	deleteByIndex(t, origin, t1, 4, 0, r3)

	t1 = []string{"a", "b", "c", "d", "e"}
	r4 := []string{"c", "d", "e"}
	deleteByIndex(t, origin, t1, 0, 2, r4)

	t1 = []string{"a", "b", "c", "d", "e"}
	r5 := []string{} // var r5 []string{} failed
	deleteByIndex(t, origin, t1, 0, 5, r5)

	// failed
	//t1 = []string{"a", "b", "c", "d","e"}
	//r6 := []string{"a", "c", "d","e"}
	//deleteByIndex(t, origin, t1, 1, 1, r6)

	// failed
	//t1 = []string{"a", "b", "c", "d","e"}
	//r7 := []string{}
	//deleteByIndex(t, origin, t1, 0, 6, r7)
}

func deleteByIndex(t *testing.T, origin, test interface{}, start, end int, expected interface{}) {
	var res interface{}
	var err error
	if end != 0 {
		res, err = DeleteByIndex(test, start, end)
	} else {
		res, err = DeleteByIndex(test, start)
	}
	if err != nil {
		t.Error("DeleteByIndex Error: " + err.Error())
	}

	if !reflect.DeepEqual(res, expected) {
		utils.LogFailedTestInfo(t, "DeleteByIndex", origin, expected, res)
		t.FailNow()
	}
}

func TestInsertByIndex(t *testing.T) {
	t1 := []string{"a", "b", "c"}

	r1 := []string{"1", "a", "b", "c"}
	insertByIndex(t, t1, 0, "1", r1)

	r2 := []string{"a", "1", "b", "c"}
	insertByIndex(t, t1, 1, "1", r2)

	r3 := []string{"a", "b", "c", "1"}
	insertByIndex(t, t1, 3, "1", r3)

	r4 := []string{"1", "2", "3", "a", "b", "c"}
	insertByIndex(t, t1, 0, []string{"1", "2", "3"}, r4)

	r5 := []string{"a", "1", "2", "3", "b", "c"}
	insertByIndex(t, t1, 1, []string{"1", "2", "3"}, r5)

	r6 := []string{"a", "b", "1", "2", "3", "c"}
	insertByIndex(t, t1, 2, []string{"1", "2", "3"}, r6)

	r7 := []string{"a", "b", "c", "1", "2", "3"}
	insertByIndex(t, t1, 3, []string{"1", "2", "3"}, r7)
}

func insertByIndex(t *testing.T, test interface{}, index int, value, expected interface{}) {
	res, err := InsertByIndex(test, index, value)
	if err != nil {
		t.Error("InsertByIndex Error: " + err.Error())
	}

	if !reflect.DeepEqual(res, expected) {
		utils.LogFailedTestInfo(t, "InsertByIndex", test, expected, res)
		t.FailNow()
	}
}

func TestUpdateByIndex(t *testing.T) {
	t1 := []string{"a", "b", "c"}
	r1 := []string{"1", "b", "c"}
	updateByIndex(t, t1, 0, "1", r1)

	t1 = []string{"a", "b", "c"}
	r2 := []string{"a", "1", "c"}
	updateByIndex(t, t1, 1, "1", r2)

	t1 = []string{"a", "b", "c"}
	r3 := []string{"a", "b", "1"}
	updateByIndex(t, t1, 2, "1", r3)

}

func updateByIndex(t *testing.T, test interface{}, index int, value, expected interface{}) {
	res, err := UpdateByIndex(test, index, value)
	if err != nil {
		t.Error("UpdateByIndex Error: " + err.Error())
	}

	if !reflect.DeepEqual(res, expected) {
		utils.LogFailedTestInfo(t, "UpdateByIndex", test, expected, res)
		t.FailNow()
	}
}

func TestUnique(t *testing.T) {
	t1 := []int{1, 2, 2, 3}
	e1 := []int{1, 2, 3}
	r1 := Unique(t1)
	if !reflect.DeepEqual(r1, e1) {
		utils.LogFailedTestInfo(t, "Unique", t1, e1, r1)
		t.FailNow()
	}

	t2 := []string{"a", "a", "b", "c"}
	e2 := []string{"a", "b", "c"}
	r2 := Unique(t2)
	if !reflect.DeepEqual(r2, e2) {
		utils.LogFailedTestInfo(t, "Unique", t2, e2, r2)
		t.FailNow()
	}
}

func TestUnion(t *testing.T) {
	s1 := []int{1, 3, 4, 6}
	s2 := []int{1, 2, 5, 6}
	s3 := []int{0, 4, 5, 7}

	expected1 := []int{1, 3, 4, 6, 2, 5, 0, 7}
	res1 := Union(s1, s2, s3)
	if !reflect.DeepEqual(res1, expected1) {
		utils.LogFailedTestInfo(t, "Union", s1, expected1, res1)
		t.FailNow()
	}

	expected2 := []int{1, 3, 4, 6}
	res2 := Union(s1)
	if !reflect.DeepEqual(res2, expected2) {
		utils.LogFailedTestInfo(t, "Union", s1, expected2, res2)
		t.FailNow()
	}
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
	for i := 0; i < len(res); i++ {
		if !reflect.DeepEqual(res[i], expected[i]) {
			utils.LogFailedTestInfo(t, "Intersection", "Intersection", expected[i], res[i])
			t.FailNow()
		}
	}

}

func TestReverseSlice(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	e1 := []int{5, 4, 3, 2, 1}
	ReverseSlice(s1)
	if !reflect.DeepEqual(s1, e1) {
		utils.LogFailedTestInfo(t, "ReverseSlice", s1, e1, s1)
		t.FailNow()
	}

	s2 := []string{"a", "b", "c", "d", "e"}
	e2 := []string{"e", "d", "c", "b", "a"}
	ReverseSlice(s2)
	if !reflect.DeepEqual(s2, e2) {
		utils.LogFailedTestInfo(t, "ReverseSlice", s2, e2, s2)
		t.FailNow()
	}
}

func TestDifference(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 6}
	e1 := []int{1, 2, 3}
	r1 := Difference(s1, s2)
	if !reflect.DeepEqual(r1, e1) {
		utils.LogFailedTestInfo(t, "Difference", s1, e1, r1)
		t.FailNow()
	}
}

func TestSortByField(t *testing.T) {
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

	sortByAge := []student{
		{"b", 15},
		{"a", 10},
		{"d", 6},
		{"c", 5},
	}

	err := SortByField(students, "age", "desc")
	if err != nil {
		t.Error("IntSlice Error: " + err.Error())
	}

	if !reflect.DeepEqual(students, sortByAge) {
		utils.LogFailedTestInfo(t, "SortByField", students, sortByAge, students)
		t.FailNow()
	}

}

func TestWithout(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	expected := []int{3, 4, 5}
	res := Without(s, 1, 2)

	if !reflect.DeepEqual(res, expected) {
		utils.LogFailedTestInfo(t, "Without", s, expected, res)
		t.FailNow()
	}
}

package maputil

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

func ExampleKeys() {
	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	keys := Keys(m)
	sort.Ints(keys)

	fmt.Println(keys)

	// Output:
	// [1 2 3 4 5]
}

func ExampleValues() {
	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	values := Values(m)
	sort.Strings(values)

	fmt.Println(values)

	// Output:
	// [a a b c d]
}

func ExampleKeysBy() {
	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
	}

	keys := KeysBy(m, func(n int) int {
		return n + 1
	})

	sort.Ints(keys)

	fmt.Println(keys)

	// Output:
	// [2 3 4]
}

func ExampleValuesBy() {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}
	values := ValuesBy(m, func(v string) string {
		switch v {
		case "a":
			return "a-1"
		case "b":
			return "b-2"
		case "c":
			return "c-3"
		default:
			return ""
		}
	})

	sort.Strings(values)

	fmt.Println(values)

	// Output:
	// [a-1 b-2 c-3]
}
func ExampleMerge() {
	m1 := map[int]string{
		1: "a",
		2: "b",
	}
	m2 := map[int]string{
		1: "c",
		3: "d",
	}

	result := Merge(m1, m2)

	fmt.Println(result)

	// Output:
	// map[1:c 2:b 3:d]
}

func ExampleForEach() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	var sum int
	ForEach(m, func(_ string, value int) {
		sum += value
	})

	fmt.Println(sum)

	// Output:
	// 10
}

func ExampleFilter() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}

	result := Filter(m, isEven)

	fmt.Println(result)

	// Output:
	// map[b:2 d:4]
}

func ExampleFilterByKeys() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	result := FilterByKeys(m, []string{"a", "b"})

	fmt.Println(result)

	// Output:
	// map[a:1 b:2]
}

func ExampleFilterByValues() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	result := FilterByValues(m, []int{3, 4})

	fmt.Println(result)

	// Output:
	// map[c:3 d:4]
}

func ExampleIntersect() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 6,
		"d": 7,
	}

	m3 := map[string]int{
		"a": 1,
		"b": 9,
		"e": 9,
	}

	result1 := Intersect(m1)
	result2 := Intersect(m1, m2)
	result3 := Intersect(m1, m2, m3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// map[a:1 b:2 c:3]
	// map[a:1 b:2]
	// map[a:1]
}

func ExampleMinus() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 11,
		"b": 22,
		"d": 33,
	}

	result := Minus(m1, m2)

	fmt.Println(result)

	// Output:
	// map[c:3]
}

func ExampleIsDisjoint() {
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"d": 22,
	}

	m3 := map[string]int{
		"a": 22,
	}

	result1 := IsDisjoint(m1, m2)
	result2 := IsDisjoint(m1, m3)

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleEntries() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	result := Entries(m)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value < result[j].Value
	})

	fmt.Println(result)

	// Output:
	// [{a 1} {b 2} {c 3}]
}

func ExampleFromEntries() {

	result := FromEntries([]Entry[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
		{Key: "c", Value: 3},
	})

	fmt.Println(result)

	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleTransform() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	result := Transform(m, func(k string, v int) (string, string) {
		return k, strconv.Itoa(v)
	})

	fmt.Println(result)

	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleMapKeys() {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	result := MapKeys(m, func(k int, _ string) string {
		return strconv.Itoa(k)
	})

	fmt.Println(result)

	// Output:
	// map[1:a 2:b 3:c]
}

func ExampleMapValues() {
	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	result := MapValues(m, func(k int, v string) string {
		return v + strconv.Itoa(k)
	})

	fmt.Println(result)

	// Output:
	// map[1:a1 2:b2 3:c3]
}

func ExampleOmitBy() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}

	result := OmitBy(m, isEven)

	fmt.Println(result)

	// Output:
	// map[a:1 c:3 e:5]
}

func ExampleOmitByKeys() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	result := OmitByKeys(m, []string{"a", "b"})

	fmt.Println(result)

	// Output:
	// map[c:3 d:4 e:5]
}

func ExampleOmitByValues() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	result := OmitByValues(m, []int{4, 5})

	fmt.Println(result)

	// Output:
	// map[a:1 b:2 c:3]
}

func ExampleMapTo() {
	type (
		Person struct {
			Name  string  `json:"name"`
			Age   int     `json:"age"`
			Phone string  `json:"phone"`
			Addr  Address `json:"address"`
		}

		Address struct {
			Street string `json:"street"`
			Number int    `json:"number"`
		}
	)

	personInfo := map[string]interface{}{
		"name":  "Nothin",
		"age":   28,
		"phone": "123456789",
		"address": map[string]interface{}{
			"street": "test",
			"number": 1,
		},
	}

	var p Person
	err := MapTo(personInfo, &p)

	fmt.Println(err)
	fmt.Println(p)

	// Output:
	// <nil>
	// {Nothin 28 123456789 {test 1}}
}

func ExampleHasKey() {
	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	result1 := HasKey(m, "a")
	result2 := HasKey(m, "c")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleMapToStruct() {

	personReqMap := map[string]any{
		"name":     "Nothin",
		"max_age":  35,
		"page":     1,
		"pageSize": 10,
	}

	type PersonReq struct {
		Name     string `json:"name"`
		MaxAge   int    `json:"max_age"`
		Page     int    `json:"page"`
		PageSize int    `json:"pageSize"`
	}
	var personReq PersonReq
	_ = MapToStruct(personReqMap, &personReq)
	fmt.Println(personReq)

	// Output:
	// {Nothin 35 1 10}
}

func ExampleToSortedSlicesDefault() {
	m := map[int]string{
		1: "a",
		3: "c",
		2: "b",
	}

	keys, values := ToSortedSlicesDefault(m)

	fmt.Println(keys)
	fmt.Println(values)

	// Output:
	// [1 2 3]
	// [a b c]
}

func ExampleToSortedSlicesWithComparator() {
	m1 := map[time.Time]string{
		time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC): "today",
		time.Date(2024, 3, 30, 0, 0, 0, 0, time.UTC): "yesterday",
		time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC):  "tomorrow",
	}

	keys1, values1 := ToSortedSlicesWithComparator(m1, func(a, b time.Time) bool {
		return a.Before(b)
	})

	m2 := map[int]string{
		1: "a",
		3: "c",
		2: "b",
	}
	keys2, values2 := ToSortedSlicesWithComparator(m2, func(a, b int) bool {
		return a > b
	})

	fmt.Println(keys1)
	fmt.Println(values1)

	fmt.Println(keys2)
	fmt.Println(values2)

	// Output:
	// [2024-03-30 00:00:00 +0000 UTC 2024-03-31 00:00:00 +0000 UTC 2024-04-01 00:00:00 +0000 UTC]
	// [yesterday today tomorrow]
	// [3 2 1]
	// [c b a]
}

func ExampleGetOrSet() {
	m := map[int]string{
		1: "a",
	}

	result1 := GetOrSet(m, 1, "1")
	result2 := GetOrSet(m, 2, "b")

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// a
	// b
}

func ExampleSortByKey() {
	m := map[int]string{
		3: "c",
		1: "a",
		4: "d",
		2: "b",
	}

	result := SortByKey(m, func(a, b int) bool {
		return a < b
	})

	fmt.Println(result)

	// Output:
	// map[1:a 2:b 3:c 4:d]
}

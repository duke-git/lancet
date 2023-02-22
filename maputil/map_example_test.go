package maputil

import (
	"fmt"
	"sort"
	"strconv"
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

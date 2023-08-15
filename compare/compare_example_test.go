package compare

import (
	"fmt"
	"time"
)

func ExampleEqual() {
	result1 := Equal(1, 1)
	result2 := Equal("1", "1")
	result3 := Equal([]int{1, 2, 3}, []int{1, 2, 3})
	result4 := Equal(map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"})

	result5 := Equal(1, "1")
	result6 := Equal(1, int64(1))
	result7 := Equal([]int{1, 2}, []int{1, 2, 3})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)

	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
	// false
}

func ExampleEqualValue() {
	result1 := EqualValue(1, 1)
	result2 := EqualValue(int(1), int64(1))
	result3 := EqualValue(1, "1")
	result4 := EqualValue(1, "2")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// true
	// true
	// true
	// false
}

func ExampleLessThan() {
	result1 := LessThan(1, 2)
	result2 := LessThan(1.1, 2.2)
	result3 := LessThan("a", "b")

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	result4 := LessThan(time1, time2)

	result5 := LessThan(2, 1)
	result6 := LessThan(1, int64(2))

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
}

func ExampleGreaterThan() {
	result1 := GreaterThan(2, 1)
	result2 := GreaterThan(2.2, 1.1)
	result3 := GreaterThan("b", "a")

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	result4 := GreaterThan(time2, time1)

	result5 := GreaterThan(1, 2)
	result6 := GreaterThan(int64(2), 1)
	result7 := GreaterThan("b", "c")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)

	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
	// false
}

func ExampleLessOrEqual() {
	result1 := LessOrEqual(1, 1)
	result2 := LessOrEqual(1.1, 2.2)
	result3 := LessOrEqual("a", "b")

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	result4 := LessOrEqual(time1, time2)

	result5 := LessOrEqual(2, 1)
	result6 := LessOrEqual(1, int64(2))

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
}

func ExampleGreaterOrEqual() {
	result1 := GreaterOrEqual(1, 1)
	result2 := GreaterOrEqual(2.2, 1.1)
	result3 := GreaterOrEqual("b", "b")

	time1 := time.Now()
	time2 := time1.Add(time.Second)
	result4 := GreaterOrEqual(time2, time1)

	result5 := GreaterOrEqual(1, 2)
	result6 := GreaterOrEqual(int64(2), 1)
	result7 := GreaterOrEqual("b", "c")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)

	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
	// false
}

func ExampleInDelta() {
	result1 := InDelta(1, 1, 0)
	result2 := InDelta(1, 2, 0)

	result3 := InDelta(2.0/3.0, 0.66667, 0.001)
	result4 := InDelta(2.0/3.0, 0.0, 0.001)

	result5 := InDelta(float64(74.96)-float64(20.48), 54.48, 0)
	result6 := InDelta(float64(74.96)-float64(20.48), 54.48, 1e-14)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

	// Output:
	// true
	// false
	// true
	// false
	// false
	// true
}

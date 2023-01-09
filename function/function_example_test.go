package function

import (
	"fmt"
	"strings"
	"time"
)

func ExampleAfter() {
	fn := After(2, func() {
		fmt.Println("test")
	})

	fn()
	fn()

	// Output:
	// test
}

func ExampleBefore() {
	fn := Before(2, func() {
		fmt.Println("test")
	})

	fn()
	fn()
	fn()
	fn()

	// Output:
	// test
	// test
}

func ExampleCurryFn_New() {
	add := func(a, b int) int {
		return a + b
	}

	var addCurry CurryFn[int] = func(values ...int) int {
		return add(values[0], values[1])
	}
	add1 := addCurry.New(1)

	result := add1(2)

	fmt.Println(result)

	// Output:
	// 3
}

func ExampleCompose() {
	toUpper := func(strs ...string) string {
		return strings.ToUpper(strs[0])
	}
	toLower := func(strs ...string) string {
		return strings.ToLower(strs[0])
	}
	transform := Compose(toUpper, toLower)

	result := transform("aBCde")

	fmt.Println(result)

	// Output:
	// ABCDE
}

func ExampleDelay() {
	var print = func(s string) {
		fmt.Println(s)
	}

	Delay(2*time.Second, print, "hello")

	// Output:
	// hello
}

func ExampleDebounced() {
	count := 0
	add := func() {
		count++
	}

	debouncedAdd := Debounced(add, 50*time.Microsecond)

	debouncedAdd()
	debouncedAdd()
	debouncedAdd()
	debouncedAdd()

	time.Sleep(100 * time.Millisecond)

	fmt.Println(count)

	debouncedAdd()

	time.Sleep(100 * time.Millisecond)

	fmt.Println(count)

	// Output:
	// 1
	// 2
}

func ExampleSchedule() {
	count := 0

	increase := func() {
		count++
	}

	stop := Schedule(1*time.Second, increase)

	time.Sleep(3 * time.Second)
	close(stop)

	fmt.Println(count)

	// Output:
	// 3
}

func ExamplePipeline() {
	addOne := func(x int) int {
		return x + 1
	}
	double := func(x int) int {
		return 2 * x
	}
	square := func(x int) int {
		return x * x
	}

	fn := Pipeline(addOne, double, square)

	result := fn(2)

	fmt.Println(result)

	// Output:
	// 36
}

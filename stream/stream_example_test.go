package stream

import (
	"fmt"
)

func ExampleOf() {
	s := Of(1, 2, 3)

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3]
}

func ExampleFromSlice() {
	s := FromSlice([]int{1, 2, 3})

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3]
}

func ExampleFromChannel() {
	ch := make(chan int)
	go func() {
		for i := 1; i < 4; i++ {
			ch <- i
		}
		close(ch)
	}()

	s := FromChannel(ch)

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3]
}

func ExampleFromRange() {
	s := FromRange(1, 5, 1)

	data := s.ToSlice()
	fmt.Println(data)

	// Output:
	// [1 2 3 4 5]
}

func ExampleGenerate() {
	n := 0
	max := 4

	generator := func() func() (int, bool) {
		return func() (int, bool) {
			n++
			return n, n < max
		}
	}

	s := Generate(generator)

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3]
}

func ExampleConcat() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{4, 5, 6})

	s := Concat(s1, s2)

	data := s.ToSlice()

	fmt.Println(data)

	// Output:
	// [1 2 3 4 5 6]
}

func ExampleStream_Distinct() {
	original := FromSlice([]int{1, 2, 2, 3, 3, 3})
	distinct := original.Distinct()

	data1 := original.ToSlice()
	data2 := distinct.ToSlice()

	fmt.Println(data1)
	fmt.Println(data2)

	// Output:
	// [1 2 2 3 3 3]
	// [1 2 3]
}

func ExampleStream_Filter() {
	original := FromSlice([]int{1, 2, 3, 4, 5})

	isEven := func(n int) bool {
		return n%2 == 0
	}

	even := original.Filter(isEven)

	fmt.Println(even.ToSlice())

	// Output:
	// [2 4]
}

func ExampleStream_Map() {
	original := FromSlice([]int{1, 2, 3})

	addOne := func(n int) int {
		return n + 1
	}

	increament := original.Map(addOne)

	fmt.Println(increament.ToSlice())

	// Output:
	// [2 3 4]
}

func ExampleStream_Peek() {
	original := FromSlice([]int{1, 2, 3})

	data := []string{}
	peekStream := original.Peek(func(n int) {
		data = append(data, fmt.Sprint("value", n))
	})

	fmt.Println(original.ToSlice())
	fmt.Println(peekStream.ToSlice())
	fmt.Println(data)

	// Output:
	// [1 2 3]
	// [1 2 3]
	// [value1 value2 value3]
}

func ExampleStream_Skip() {
	original := FromSlice([]int{1, 2, 3, 4})

	s1 := original.Skip(-1)
	s2 := original.Skip(0)
	s3 := original.Skip(1)
	s4 := original.Skip(5)

	fmt.Println(s1.ToSlice())
	fmt.Println(s2.ToSlice())
	fmt.Println(s3.ToSlice())
	fmt.Println(s4.ToSlice())

	// Output:
	// [1 2 3 4]
	// [1 2 3 4]
	// [2 3 4]
	// []
}

func ExampleStream_Limit() {
	original := FromSlice([]int{1, 2, 3, 4})

	s1 := original.Limit(-1)
	s2 := original.Limit(0)
	s3 := original.Limit(1)
	s4 := original.Limit(5)

	fmt.Println(s1.ToSlice())
	fmt.Println(s2.ToSlice())
	fmt.Println(s3.ToSlice())
	fmt.Println(s4.ToSlice())

	// Output:
	// []
	// []
	// [1]
	// [1 2 3 4]
}

func ExampleStream_AllMatch() {
	original := FromSlice([]int{1, 2, 3})

	result1 := original.AllMatch(func(item int) bool {
		return item > 0
	})

	result2 := original.AllMatch(func(item int) bool {
		return item > 1
	})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleStream_AnyMatch() {
	original := FromSlice([]int{1, 2, 3})

	result1 := original.AnyMatch(func(item int) bool {
		return item > 1
	})

	result2 := original.AnyMatch(func(item int) bool {
		return item > 3
	})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleStream_NoneMatch() {
	original := FromSlice([]int{1, 2, 3})

	result1 := original.NoneMatch(func(item int) bool {
		return item > 3
	})

	result2 := original.NoneMatch(func(item int) bool {
		return item > 1
	})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// true
	// false
}

func ExampleStream_ForEach() {
	original := FromSlice([]int{1, 2, 3})

	result := 0
	original.ForEach(func(item int) {
		result += item
	})

	fmt.Println(result)

	// Output:
	// 6
}

func ExampleStream_Reduce() {
	original := FromSlice([]int{1, 2, 3})

	result := original.Reduce(0, func(a, b int) int {
		return a + b
	})

	fmt.Println(result)

	// Output:
	// 6
}

func ExampleStream_FindFirst() {
	original := FromSlice([]int{1, 2, 3})

	result, ok := original.FindFirst()

	fmt.Println(result)
	fmt.Println(ok)

	// Output:
	// 1
	// true
}

func ExampleStream_FindLast() {
	original := FromSlice([]int{3, 2, 1})

	result, ok := original.FindLast()

	fmt.Println(result)
	fmt.Println(ok)

	// Output:
	// 1
	// true
}

func ExampleStream_Reverse() {
	original := FromSlice([]int{1, 2, 3})

	reverse := original.Reverse()

	fmt.Println(reverse.ToSlice())

	// Output:
	// [3 2 1]
}

func ExampleStream_Range() {
	original := FromSlice([]int{1, 2, 3})

	s1 := original.Range(0, 0)
	s2 := original.Range(0, 1)
	s3 := original.Range(0, 3)
	s4 := original.Range(1, 2)

	fmt.Println(s1.ToSlice())
	fmt.Println(s2.ToSlice())
	fmt.Println(s3.ToSlice())
	fmt.Println(s4.ToSlice())

	// Output:
	// []
	// [1]
	// [1 2 3]
	// [2]
}

func ExampleStream_Sorted() {
	original := FromSlice([]int{4, 2, 1, 3})

	sorted := original.Sorted(func(a, b int) bool { return a < b })

	fmt.Println(original.ToSlice())
	fmt.Println(sorted.ToSlice())

	// Output:
	// [4 2 1 3]
	// [1 2 3 4]
}

func ExampleStream_Max() {
	original := FromSlice([]int{4, 2, 1, 3})

	max, ok := original.Max(func(a, b int) bool { return a > b })

	fmt.Println(max)
	fmt.Println(ok)

	// Output:
	// 4
	// true
}

func ExampleStream_Min() {
	original := FromSlice([]int{4, 2, 1, 3})

	min, ok := original.Min(func(a, b int) bool { return a < b })

	fmt.Println(min)
	fmt.Println(ok)

	// Output:
	// 1
	// true
}

func ExampleStream_Count() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{})

	fmt.Println(s1.Count())
	fmt.Println(s2.Count())

	// Output:
	// 3
	// 0
}

func ExampleStream_IndexOf() {
	s := FromSlice([]int{1, 2, 3, 2})

	result1 := s.IndexOf(0, func(a, b int) bool { return a == b })
	result2 := s.IndexOf(2, func(a, b int) bool { return a == b })

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// -1
	// 1
}

func ExampleStream_LastIndexOf() {
	s := FromSlice([]int{1, 2, 3, 2})

	result1 := s.LastIndexOf(0, func(a, b int) bool { return a == b })
	result2 := s.LastIndexOf(2, func(a, b int) bool { return a == b })

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// -1
	// 3
}

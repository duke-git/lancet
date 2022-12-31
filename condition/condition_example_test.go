package condition

import "fmt"

func ExampleBool() {
	// bool
	result1 := Bool(false)
	result2 := Bool(true)
	fmt.Println(result1)
	fmt.Println(result2)

	// integer
	result3 := Bool(0)
	result4 := Bool(1)
	fmt.Println(result3)
	fmt.Println(result4)

	// string
	result5 := Bool("")
	result6 := Bool(" ")
	fmt.Println(result5)
	fmt.Println(result6)

	// slice
	var nums = []int{}
	result7 := Bool(nums)
	nums = append(nums, 1, 2)
	result8 := Bool(nums)
	fmt.Println(result7)
	fmt.Println(result8)

	// Output:
	// false
	// true
	// false
	// true
	// false
	// true
	// false
	// true
}

func ExampleAnd() {
	result1 := And(0, 1)
	result2 := And(0, "")
	result3 := And(0, "0")
	result4 := And(1, "0")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// false
	// false
	// false
	// true
}

func ExampleOr() {
	result1 := Or(0, "")
	result2 := Or(0, 1)
	result3 := Or(0, "0")
	result4 := Or(1, "0")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// false
	// true
	// true
	// true
}

func ExampleXor() {
	result1 := Xor(0, 0)
	result2 := Xor(1, 1)
	result3 := Xor(0, 1)
	result4 := Xor(1, 0)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// false
	// false
	// true
	// true
}

func ExampleNor() {
	result1 := Nor(0, 0)
	result2 := Nor(1, 1)
	result3 := Nor(0, 1)
	result4 := Nor(1, 0)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// true
	// false
	// false
	// false
}

func ExampleXnor() {
	result1 := Xnor(0, 0)
	result2 := Xnor(1, 1)
	result3 := Xnor(0, 1)
	result4 := Xnor(1, 0)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// true
	// true
	// false
	// false
}

func ExampleNand() {
	result1 := Nand(0, 0)
	result2 := Nand(1, 0)
	result3 := Nand(0, 1)
	result4 := Nand(1, 1)

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

func ExampleTernaryOperator() {
	conditionTrue := 2 > 1
	result1 := TernaryOperator(conditionTrue, 0, 1)
	fmt.Println(result1)

	conditionFalse := 2 > 3
	result2 := TernaryOperator(conditionFalse, 0, 1)
	fmt.Println(result2)

	// Output:
	// 0
	// 1
}

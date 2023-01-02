package formatter

import "fmt"

func ExampleComma() {
	result1 := Comma("123", "")
	result2 := Comma("12345", "$")
	result3 := Comma(1234567, "￥")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 123
	// $12,345
	// ￥1,234,567
}

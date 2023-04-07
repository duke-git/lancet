package formatter

import (
	"bytes"
	"fmt"
)

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

func ExamplePretty() {
	result1, _ := Pretty([]string{"a", "b", "c"})
	result2, _ := Pretty(map[string]int{"a": 1})

	fmt.Println(result1)
	fmt.Println(result2)

	// Output:
	// [
	//     "a",
	//     "b",
	//     "c"
	// ]
	// {
	//     "a": 1
	// }
}

func ExamplePrettyToWriter() {
	type User struct {
		Name string `json:"name"`
		Aage uint   `json:"age"`
	}
	user := User{Name: "King", Aage: 10000}

	buf := &bytes.Buffer{}
	err := PrettyToWriter(user, buf)

	fmt.Println(buf)
	fmt.Println(err)

	// Output:
	// {
	//     "name": "King",
	//     "age": 10000
	// }
	//
	// <nil>
}

func ExampleDecimalBytes() {
	result1 := DecimalBytes(1000)
	result2 := DecimalBytes(1024)
	result3 := DecimalBytes(1234567)
	result4 := DecimalBytes(1234567, 3)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 1KB
	// 1.024KB
	// 1.2346MB
	// 1.235MB
}

func ExampleBinaryBytes() {
	result1 := BinaryBytes(1024)
	result2 := BinaryBytes(1024 * 1024)
	result3 := BinaryBytes(1234567)
	result4 := BinaryBytes(1234567, 2)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 1KiB
	// 1MiB
	// 1.1774MiB
	// 1.18MiB
}

func ExampleParseDecimalBytes() {
	result1, _ := ParseDecimalBytes("12")
	result2, _ := ParseDecimalBytes("12k")
	result3, _ := ParseDecimalBytes("12 Kb")
	result4, _ := ParseDecimalBytes("12.2 kb")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 12
	// 12000
	// 12000
	// 12200
}

func ExampleParseBinaryBytes() {
	result1, _ := ParseBinaryBytes("12")
	result2, _ := ParseBinaryBytes("12ki")
	result3, _ := ParseBinaryBytes("12 KiB")
	result4, _ := ParseBinaryBytes("12.2 kib")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)

	// Output:
	// 12
	// 12288
	// 12288
	// 12492
}

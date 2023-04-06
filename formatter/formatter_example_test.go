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

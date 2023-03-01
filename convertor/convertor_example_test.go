package convertor

import (
	"fmt"
	"strconv"
)

func ExampleToBool() {
	cases := []string{"1", "true", "True", "false", "False", "0", "123", "0.0", "abc"}

	for i := 0; i < len(cases); i++ {
		result, _ := ToBool(cases[i])
		fmt.Println(result)
	}

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
}

func ExampleToBytes() {
	result1, _ := ToBytes(1)
	result2, _ := ToBytes("abc")
	result3, _ := ToBytes(true)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// [0 0 0 0 0 0 0 1]
	// [97 98 99]
	// [116 114 117 101]
}

func ExampleToChar() {
	result1 := ToChar("")
	result2 := ToChar("abc")
	result3 := ToChar("1 2#3")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// []
	// [a b c]
	// [1   2 # 3]
}

func ExampleToChannel() {
	ch := ToChannel([]int{1, 2, 3})
	result1 := <-ch
	result2 := <-ch
	result3 := <-ch

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 1
	// 2
	// 3
}

func ExampleToString() {
	result1 := ToString("")
	result2 := ToString(nil)
	result3 := ToString(0)
	result4 := ToString(1.23)
	result5 := ToString(true)
	result6 := ToString(false)
	result7 := ToString([]int{1, 2, 3})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)

	// Output:
	//
	//
	// 0
	// 1.23
	// true
	// false
	// [1,2,3]
}

func ExampleToJson() {

	aMap := map[string]int{"a": 1, "b": 2, "c": 3}
	result1, err := ToJson(aMap)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(result1)

	// Output:
	// {"a":1,"b":2,"c":3}
}

func ExampleToFloat() {
	result1, _ := ToFloat("")
	result2, _ := ToFloat("abc")
	result3, _ := ToFloat("-1")
	result4, _ := ToFloat("-.11")
	result5, _ := ToFloat("1.23e3")
	result6, _ := ToFloat(true)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

	// Output:
	// 0
	// 0
	// -1
	// -0.11
	// 1230
	// 0
}

func ExampleToInt() {
	result1, _ := ToInt("123")
	result2, _ := ToInt("-123")
	result3, _ := ToInt(float64(12.3))
	result4, _ := ToInt("abc")
	result5, _ := ToInt(true)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// 123
	// -123
	// 12
	// 0
	// 0
}

func ExampleToPointer() {
	result := ToPointer(123)
	fmt.Println(*result)

	// Output:
	// 123
}

func ExampleToMap() {
	type Message struct {
		name string
		code int
	}
	messages := []Message{
		{name: "Hello", code: 100},
		{name: "Hi", code: 101},
	}
	result := ToMap(messages, func(msg Message) (int, string) {
		return msg.code, msg.name
	})

	fmt.Println(result)

	// Output:
	// map[100:Hello 101:Hi]
}

func ExampleStructToMap() {
	type People struct {
		Name string `json:"name"`
		age  int
	}
	p := People{
		"test",
		100,
	}
	pm, _ := StructToMap(p)

	fmt.Println(pm)

	// Output:
	// map[name:test]
}

func ExampleMapToSlice() {
	aMap := map[string]int{"a": 1, "b": 2, "c": 3}
	result := MapToSlice(aMap, func(key string, value int) string {
		return key + ":" + strconv.Itoa(value)
	})

	fmt.Println(result) //[]string{"a:1", "c:3", "b:2"} (random order)
}

func ExampleColorHexToRGB() {
	colorHex := "#003366"
	r, g, b := ColorHexToRGB(colorHex)

	fmt.Println(r, g, b)

	// Output:
	// 0 51 102
}

func ExampleColorRGBToHex() {
	r := 0
	g := 51
	b := 102
	colorHex := ColorRGBToHex(r, g, b)

	fmt.Println(colorHex)

	// Output:
	// #003366
}

func ExampleEncodeByte() {
	byteData, _ := EncodeByte("abc")
	fmt.Println(byteData)

	// Output:
	// [6 12 0 3 97 98 99]
}

func ExampleDecodeByte() {
	var obj string
	byteData := []byte{6, 12, 0, 3, 97, 98, 99}
	err := DecodeByte(byteData, &obj)
	if err != nil {
		return
	}

	fmt.Println(obj)

	// Output:
	// abc
}

func ExampleDeepClone() {
	type Struct struct {
		Str   string
		Int   int
		Float float64
		Bool  bool
		Nil   interface{}
		// unexported string
	}

	cases := []interface{}{
		true,
		1,
		0.1,
		map[string]int{
			"a": 1,
			"b": 2,
		},
		&Struct{
			Str:   "test",
			Int:   1,
			Float: 0.1,
			Bool:  true,
			Nil:   nil,
			// unexported: "can't be cloned",
		},
	}

	for _, item := range cases {
		cloned := DeepClone(item)

		isPointerEqual := &cloned == &item
		fmt.Println(cloned, isPointerEqual)
	}

	// Output:
	// true false
	// 1 false
	// 0.1 false
	// map[a:1 b:2] false
	// &{test 1 0.1 true <nil>} false
}

func ExampleCopyProperties() {
	type Address struct {
		Country string
		ZipCode string
	}

	type User struct {
		Name   string
		Age    int
		Role   string
		Addr   Address
		Hobbys []string
		salary int
	}

	type Employee struct {
		Name   string
		Age    int
		Role   string
		Addr   Address
		Hobbys []string
		salary int
	}

	user := User{Name: "user001", Age: 10, Role: "Admin", Addr: Address{Country: "CN", ZipCode: "001"}, Hobbys: []string{"a", "b"}, salary: 1000}

	employee1 := Employee{}
	err := CopyProperties(&employee1, &user)
	if err != nil {
		return
	}

	employee2 := Employee{Name: "employee001", Age: 20, Role: "User",
		Addr: Address{Country: "UK", ZipCode: "002"}, salary: 500}

	err = CopyProperties(&employee2, &user)
	if err != nil {
		return
	}

	fmt.Println(employee1)
	fmt.Println(employee2)

	// Output:
	// {user001 10 Admin {CN 001} [a b] 0}
	// {user001 10 Admin {CN 001} [a b] 500}
}

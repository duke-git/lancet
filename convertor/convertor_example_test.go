package convertor

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unicode/utf8"

	"github.com/duke-git/lancet/v2/validator"
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
	type Disk struct {
		Name    string  `json:"name"`
		Total   string  `json:"total"`
		Used    string  `json:"used"`
		Percent float64 `json:"percent"`
	}

	type DiskVO struct {
		Name    string  `json:"name"`
		Total   string  `json:"total"`
		Used    string  `json:"used"`
		Percent float64 `json:"percent"`
	}

	type Indicator struct {
		Id      string    `json:"id"`
		Ip      string    `json:"ip"`
		UpTime  string    `json:"upTime"`
		LoadAvg string    `json:"loadAvg"`
		Cpu     int       `json:"cpu"`
		Disk    []Disk    `json:"disk"`
		Stop    chan bool `json:"-"`
	}

	type IndicatorVO struct {
		Id      string   `json:"id"`
		Ip      string   `json:"ip"`
		UpTime  string   `json:"upTime"`
		LoadAvg string   `json:"loadAvg"`
		Cpu     int64    `json:"cpu"`
		Disk    []DiskVO `json:"disk"`
	}

	indicator := &Indicator{Id: "001", Ip: "127.0.0.1", Cpu: 1, Disk: []Disk{
		{Name: "disk-001", Total: "100", Used: "1", Percent: 10},
		{Name: "disk-002", Total: "200", Used: "1", Percent: 20},
		{Name: "disk-003", Total: "300", Used: "1", Percent: 30},
	}}

	indicatorVO := IndicatorVO{}

	CopyProperties(&indicatorVO, indicator)

	fmt.Println(indicatorVO.Id)
	fmt.Println(indicatorVO.Ip)
	fmt.Println(len(indicatorVO.Disk))

	// Output:
	// 001
	// 127.0.0.1
	// 3
}

func ExampleToInterface() {
	val := reflect.ValueOf("abc")
	iVal, ok := ToInterface(val)

	fmt.Printf("%T\n", iVal)
	fmt.Printf("%v\n", iVal)
	fmt.Println(ok)

	// Output:
	// string
	// abc
	// true
}

func ExampleUtf8ToGbk() {
	utf8Data := []byte("hello")
	gbkData, _ := Utf8ToGbk(utf8Data)

	fmt.Println(utf8.Valid(utf8Data))
	fmt.Println(validator.IsGBK(gbkData))

	// Output:
	// true
	// true
}

func ExampleGbkToUtf8() {
	gbkData, _ := Utf8ToGbk([]byte("hello"))
	utf8Data, _ := GbkToUtf8(gbkData)

	fmt.Println(utf8.Valid(utf8Data))
	fmt.Println(string(utf8Data))

	// Output:
	// true
	// hello
}

func ExampleToStdBase64() {
	// if you want to see the result, please use 'base64.StdEncoding.DecodeString()' to decode the result

	afterEncode := ToStdBase64(nil)
	fmt.Println(afterEncode)

	stringVal := "hello"
	afterEncode = ToStdBase64(stringVal)
	fmt.Println(afterEncode)

	byteSliceVal := []byte("hello")
	afterEncode = ToStdBase64(byteSliceVal)
	fmt.Println(afterEncode)

	intVal := 123
	afterEncode = ToStdBase64(intVal)
	fmt.Println(afterEncode)

	mapVal := map[string]any{"a": "hi", "b": 2, "c": struct {
		A string
		B int
	}{"hello", 3}}
	afterEncode = ToStdBase64(mapVal)
	fmt.Println(afterEncode)

	floatVal := 123.456
	afterEncode = ToStdBase64(floatVal)
	fmt.Println(afterEncode)

	boolVal := true
	afterEncode = ToStdBase64(boolVal)
	fmt.Println(afterEncode)

	errVal := errors.New("err")
	afterEncode = ToStdBase64(errVal)
	fmt.Println(afterEncode)

	// Output:
	//
	// aGVsbG8=
	// aGVsbG8=
	// MTIz
	// eyJhIjoiaGkiLCJiIjoyLCJjIjp7IkEiOiJoZWxsbyIsIkIiOjN9fQ==
	// MTIzLjQ1Ng==
	// dHJ1ZQ==
	// ZXJy
}

func ExampleToUrlBase64() {
	// if you want to see the result, please use 'base64.URLEncoding.DecodeString()' to decode the result

	stringVal := "hello"
	afterEncode := ToUrlBase64(stringVal)
	fmt.Println(afterEncode)

	byteSliceVal := []byte("hello")
	afterEncode = ToUrlBase64(byteSliceVal)
	fmt.Println(afterEncode)

	intVal := 123
	afterEncode = ToUrlBase64(intVal)
	fmt.Println(afterEncode)

	mapVal := map[string]any{"a": "hi", "b": 2, "c": struct {
		A string
		B int
	}{"hello", 3}}
	afterEncode = ToUrlBase64(mapVal)
	fmt.Println(afterEncode)

	floatVal := 123.456
	afterEncode = ToUrlBase64(floatVal)
	fmt.Println(afterEncode)

	boolVal := true
	afterEncode = ToUrlBase64(boolVal)
	fmt.Println(afterEncode)

	errVal := errors.New("err")
	afterEncode = ToUrlBase64(errVal)
	fmt.Println(afterEncode)

	// Output:
	// aGVsbG8=
	// aGVsbG8=
	// MTIz
	// eyJhIjoiaGkiLCJiIjoyLCJjIjp7IkEiOiJoZWxsbyIsIkIiOjN9fQ==
	// MTIzLjQ1Ng==
	// dHJ1ZQ==
	// ZXJy
}

func ExampleToRawStdBase64() {
	// if you want to see the result, please use 'base64.RawStdEncoding.DecodeString()' to decode the result
	stringVal := "hello"
	afterEncode := ToRawStdBase64(stringVal)
	fmt.Println(afterEncode)

	byteSliceVal := []byte("hello")
	afterEncode = ToRawStdBase64(byteSliceVal)
	fmt.Println(afterEncode)

	intVal := 123
	afterEncode = ToRawStdBase64(intVal)
	fmt.Println(afterEncode)

	mapVal := map[string]any{"a": "hi", "b": 2, "c": struct {
		A string
		B int
	}{"hello", 3}}
	afterEncode = ToRawStdBase64(mapVal)
	fmt.Println(afterEncode)

	floatVal := 123.456
	afterEncode = ToRawStdBase64(floatVal)
	fmt.Println(afterEncode)

	boolVal := true
	afterEncode = ToRawStdBase64(boolVal)
	fmt.Println(afterEncode)

	errVal := errors.New("err")
	afterEncode = ToRawStdBase64(errVal)
	fmt.Println(afterEncode)

	// Output:
	// aGVsbG8
	// aGVsbG8
	// MTIz
	// eyJhIjoiaGkiLCJiIjoyLCJjIjp7IkEiOiJoZWxsbyIsIkIiOjN9fQ
	// MTIzLjQ1Ng
	// dHJ1ZQ
	// ZXJy
}

func ExampleToRawUrlBase64() {
	// if you want to see the result, please use 'base64.RawURLEncoding.DecodeString()' to decode the result

	stringVal := "hello"
	afterEncode := ToRawUrlBase64(stringVal)
	fmt.Println(afterEncode)

	byteSliceVal := []byte("hello")
	afterEncode = ToRawUrlBase64(byteSliceVal)
	fmt.Println(afterEncode)

	intVal := 123
	afterEncode = ToRawUrlBase64(intVal)
	fmt.Println(afterEncode)

	mapVal := map[string]any{"a": "hi", "b": 2, "c": struct {
		A string
		B int
	}{"hello", 3}}
	afterEncode = ToRawUrlBase64(mapVal)
	fmt.Println(afterEncode)

	floatVal := 123.456
	afterEncode = ToRawUrlBase64(floatVal)
	fmt.Println(afterEncode)

	boolVal := true
	afterEncode = ToRawUrlBase64(boolVal)
	fmt.Println(afterEncode)

	errVal := errors.New("err")
	afterEncode = ToRawUrlBase64(errVal)
	fmt.Println(afterEncode)

	// Output:
	// aGVsbG8
	// aGVsbG8
	// MTIz
	// eyJhIjoiaGkiLCJiIjoyLCJjIjp7IkEiOiJoZWxsbyIsIkIiOjN9fQ
	// MTIzLjQ1Ng
	// dHJ1ZQ
	// ZXJy
}

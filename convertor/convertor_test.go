package convertor

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strconv"
	"testing"
	"unicode/utf8"
	"unsafe"

	"github.com/duke-git/lancet/v2/internal"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/duke-git/lancet/v2/validator"
)

func TestToChar(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToChar")

	cases := []string{"", "abc", "1 2#3"}
	expected := [][]string{
		{""},
		{"a", "b", "c"},
		{"1", " ", "2", "#", "3"},
	}
	for i := 0; i < len(cases); i++ {
		assert.Equal(expected[i], ToChar(cases[i]))
	}
}

func TestToChannel(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToChannel")

	ch := ToChannel([]int{1, 2, 3})
	assert.Equal(1, <-ch)
	assert.Equal(2, <-ch)
	assert.Equal(3, <-ch)

	_, ok := <-ch
	assert.Equal(false, ok)
}

func TestToBool(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToBool")

	cases := []string{"1", "true", "True", "false", "False", "0", "123", "0.0", "abc"}
	expected := []bool{true, true, true, false, false, false, false, false, false}

	for i := 0; i < len(cases); i++ {
		actual, _ := ToBool(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestToBytes(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToBytes")

	cases := []any{
		0,
		false,
		"1",
	}
	expected := [][]byte{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{102, 97, 108, 115, 101},
		{49},
	}
	for i := 0; i < len(cases); i++ {
		actual, _ := ToBytes(cases[i])
		assert.Equal(expected[i], actual)
	}

	bytesData, err := ToBytes("abc")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	assert.Equal("abc", ToString(bytesData))
}

func TestToInt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToInt")

	cases := []any{"123", "-123", 123,
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float32(12.3), float64(12.3),
		"abc", false, "111111111111111111111111111111111111111"}

	expected := []int64{123, -123, 123, 123, 123, 123, 123, 123, 12, 12, 0, 0, 0}

	for i := 0; i < len(cases); i++ {
		actual, _ := ToInt(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestToFloat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToFloat")

	cases := []any{
		"", "-1", "-.11", "1.23e3", ".123e10", "abc",
		int(0), int8(1), int16(-1), int32(123), int64(123),
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float64(12.3), float32(12.3),
	}
	expected := []float64{0, -1, -0.11, 1230, 0.123e10, 0,
		0, 1, -1, 123, 123, 123, 123, 123, 123, 123, 12.3, 12.300000190734863}

	for i := 0; i < len(cases); i++ {
		actual, _ := ToFloat(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestToString(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToString")

	aMap := make(map[string]int)
	aMap["a"] = 1
	aMap["b"] = 2
	aMap["c"] = 3

	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "TestStruct"}

	cases := []any{
		"", nil,
		int(0), int8(1), int16(-1), int32(123), int64(123),
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float64(12.3), float32(12.3),
		true, false,
		[]int{1, 2, 3}, aMap, aStruct, []byte{104, 101, 108, 108, 111}}

	expected := []string{
		"", "",
		"0", "1", "-1",
		"123", "123", "123", "123", "123", "123", "123",
		"12.3", "12.3",
		"true", "false",
		"[1,2,3]", "{\"a\":1,\"b\":2,\"c\":3}", "{\"Name\":\"TestStruct\"}", "hello",
	}

	for i := 0; i < len(cases); i++ {
		actual := ToString(cases[i])
		assert.Equal(expected[i], actual)
	}
}
func TestToJson(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToJson")

	var aMap = map[string]int{"a": 1, "b": 2, "c": 3}
	mapJsonStr, _ := ToJson(aMap)
	assert.Equal("{\"a\":1,\"b\":2,\"c\":3}", mapJsonStr)

	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "TestStruct"}
	structJsonStr, _ := ToJson(aStruct)
	assert.Equal("{\"Name\":\"TestStruct\"}", structJsonStr)
}

func TestToMap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToMap")

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
	expected := map[int]string{100: "Hello", 101: "Hi"}

	assert.Equal(expected, result)
}

func TestStructToMap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStructToMap")

	t.Run("StructToMap", func(_ *testing.T) {
		type People struct {
			Name string `json:"name"`
			age  int
		}
		p := People{
			"test",
			100,
		}
		pm, _ := StructToMap(p)
		var expected = map[string]any{"name": "test"}
		assert.Equal(expected, pm)
	})

	t.Run("StructToMapWithJsonAttr", func(_ *testing.T) {
		type People struct {
			Name  string `json:"name,omitempty"` // json tag with attribute
			Phone string `json:"phone"`          // json tag without attribute
			Sex   string `json:"-"`              // ignore
			age   int    // no tag
		}
		p := People{
			Phone: "1111",
			Sex:   "male",
			age:   100,
		}
		pm, _ := StructToMap(p)
		var expected = map[string]any{"phone": "1111"}
		assert.Equal(expected, pm)
	})
}

func TestMapToSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMapToSlice")

	aMap := map[string]int{"a": 1, "b": 2, "c": 3}
	result := MapToSlice(aMap, func(key string, value int) string {
		return key + ":" + strconv.Itoa(value)
	})

	assert.Equal(3, len(result))
	assert.Equal(true, slice.Contain(result, "a:1"))
	assert.Equal(true, slice.Contain(result, "b:2"))
	assert.Equal(true, slice.Contain(result, "c:3"))
}

func TestColorHexToRGB(t *testing.T) {
	t.Parallel()

	colorHex := "#003366"
	r, g, b := ColorHexToRGB(colorHex)
	colorRGB := fmt.Sprintf("%d,%d,%d", r, g, b)
	expected := "0,51,102"

	assert := internal.NewAssert(t, "TestColorHexToRGB")
	assert.Equal(expected, colorRGB)
}

func TestColorRGBToHex(t *testing.T) {
	t.Parallel()

	r := 0
	g := 51
	b := 102
	colorHex := ColorRGBToHex(r, g, b)
	expected := "#003366"

	assert := internal.NewAssert(t, "TestColorRGBToHex")
	assert.Equal(expected, colorHex)
}

func TestToPointer(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToPointer")
	result := ToPointer(123)

	assert.Equal(*result, 123)
}

func TestEncodeByte(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEncodeByte")

	byteData, _ := EncodeByte("abc")
	expected := []byte{6, 12, 0, 3, 97, 98, 99}

	assert.Equal(expected, byteData)
}

func TestDecodeByte(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestDecodeByte")

	var obj string
	byteData := []byte{6, 12, 0, 3, 97, 98, 99}
	err := DecodeByte(byteData, &obj)
	assert.IsNil(err)
	assert.Equal("abc", obj)
}

func TestDeepClone(t *testing.T) {
	t.Parallel()

	// assert := internal.NewAssert(t, "TestDeepClone")

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
		[]interface{}{1, &Struct{Str: "test"}, Struct{Str: "test2"}},
	}

	for i, item := range cases {
		cloned := DeepClone(item)

		if &cloned == &item {
			t.Fatalf("[TestDeepClone case #%d failed]: equal pointer", i)
		}

		if !reflect.DeepEqual(item, cloned) {
			t.Fatalf("[TestDeepClone case #%d failed] unequal objects", i)
		}
	}
}

func TestCopyProperties(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestCopyProperties")

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

	err := CopyProperties(&indicatorVO, indicator)

	assert.IsNil(err)
	assert.Equal("001", indicatorVO.Id)
	assert.Equal("127.0.0.1", indicatorVO.Ip)
	assert.Equal(3, len(indicatorVO.Disk))
}

func TestToInterface(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToInterface")

	cases := []reflect.Value{
		reflect.ValueOf("abc"),
		reflect.ValueOf(int(0)), reflect.ValueOf(int8(1)), reflect.ValueOf(int16(-1)), reflect.ValueOf(int32(123)), reflect.ValueOf(int64(123)),
		reflect.ValueOf(uint(123)), reflect.ValueOf(uint8(123)), reflect.ValueOf(uint16(123)), reflect.ValueOf(uint32(123)), reflect.ValueOf(uint64(123)),
		reflect.ValueOf(float64(12.3)), reflect.ValueOf(float32(12.3)),
		reflect.ValueOf(true), reflect.ValueOf(false),
	}

	expected := []interface{}{
		"abc",
		0, int8(1), int16(-1), int32(123), int64(123),
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float64(12.3), float32(12.3),
		true, false,
	}

	for i := 0; i < len(cases); i++ {
		actual, _ := ToInterface(cases[i])
		assert.Equal(expected[i], actual)
	}

	nilVal, ok := ToInterface(reflect.ValueOf(nil))
	assert.EqualValues(nil, nilVal)
	assert.Equal(false, ok)
}

func TestUtf8ToGbk(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestUtf8ToGbk")

	utf8Data := []byte("hello")
	gbkData, err := Utf8ToGbk(utf8Data)

	assert.Equal(true, utf8.Valid(utf8Data))
	assert.Equal(true, validator.IsGBK(gbkData))
	assert.IsNil(err)
}

func TestGbkToUtf8(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestGbkToUtf8")

	gbkData, err := Utf8ToGbk([]byte("hello"))
	utf8Data, err := GbkToUtf8(gbkData)

	assert.IsNil(err)
	assert.Equal(true, utf8.Valid(utf8Data))
	assert.Equal("hello", string(utf8Data))
}

func TestToStdBase64(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestToStdBase64")

	r1 := ToStdBase64("abc")
	d1, _ := base64.StdEncoding.DecodeString(r1)
	assert.Equal("abc", string(d1))

	r2 := ToStdBase64([]byte("abc"))
	d2, _ := base64.StdEncoding.DecodeString(r2)
	assert.Equal("abc", string(d2))

	r3 := ToStdBase64(123)
	d3, _ := base64.StdEncoding.DecodeString(r3)
	assert.Equal("123", string(d3))

	r4 := ToStdBase64(11.11)
	d4, _ := base64.StdEncoding.DecodeString(r4)
	assert.Equal("11.11", string(d4))

	r5 := ToStdBase64(map[string]any{"name": "duke", "quantity": 1})
	d5, _ := base64.StdEncoding.DecodeString(r5)
	assert.Equal("{\"name\":\"duke\",\"quantity\":1}", string(d5))

	r6 := ToStdBase64([]int64{7, 5, 9, 4, 23})
	d6, _ := base64.StdEncoding.DecodeString(r6)
	assert.Equal("[7,5,9,4,23]", string(d6))

	r7 := ToStdBase64([]string{"7", "5", "9", "4", "23"})
	d7, _ := base64.StdEncoding.DecodeString(r7)
	assert.Equal("[\"7\",\"5\",\"9\",\"4\",\"23\"]", string(d7))

	r8 := ToStdBase64(nil)
	d8, _ := base64.StdEncoding.DecodeString(r8)
	assert.Equal("", string(d8))

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	r9 := ToStdBase64(ch)
	d9, _ := base64.StdEncoding.DecodeString(r9)
	assert.Equal("", string(d9))

	r10 := ToStdBase64(io.EOF)
	d10, _ := base64.StdEncoding.DecodeString(r10)
	assert.Equal("EOF", string(d10))

	r11 := ToStdBase64(errors.New("test"))
	d11, _ := base64.StdEncoding.DecodeString(r11)
	assert.Equal("test", string(d11))

	typedNil := (*int)(nil)
	r12 := ToStdBase64(typedNil)
	d12, _ := base64.StdEncoding.DecodeString(r12)
	assert.Equal("", string(d12))

	type nilInterface interface {
	}
	var nI nilInterface = nil
	d13, _ := base64.StdEncoding.DecodeString(ToStdBase64(nI))
	assert.Equal("", string(d13))

	var p unsafe.Pointer
	d14, _ := base64.StdEncoding.DecodeString(ToStdBase64(p))
	assert.Equal("", string(d14))
}

func TestToUrlBase64(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestToUrlBase64")

	r1 := ToUrlBase64("abc")
	d1, _ := base64.URLEncoding.DecodeString(r1)
	assert.Equal("abc", string(d1))

	r2 := ToUrlBase64([]byte("abc"))
	d2, _ := base64.URLEncoding.DecodeString(r2)
	assert.Equal("abc", string(d2))

	r3 := ToUrlBase64(123)
	d3, _ := base64.URLEncoding.DecodeString(r3)
	assert.Equal("123", string(d3))

	r4 := ToUrlBase64(11.11)
	d4, _ := base64.URLEncoding.DecodeString(r4)
	assert.Equal("11.11", string(d4))

	r5 := ToUrlBase64(map[string]any{"name": "duke", "quantity": 1})
	d5, _ := base64.URLEncoding.DecodeString(r5)
	assert.Equal("{\"name\":\"duke\",\"quantity\":1}", string(d5))

	r6 := ToUrlBase64([]int64{7, 5, 9, 4, 23})
	d6, _ := base64.URLEncoding.DecodeString(r6)
	assert.Equal("[7,5,9,4,23]", string(d6))

	r7 := ToUrlBase64([]string{"7", "5", "9", "4", "23"})
	d7, _ := base64.URLEncoding.DecodeString(r7)
	assert.Equal("[\"7\",\"5\",\"9\",\"4\",\"23\"]", string(d7))

	r8 := ToUrlBase64(nil)
	d8, _ := base64.URLEncoding.DecodeString(r8)
	assert.Equal("", string(d8))

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	r9 := ToUrlBase64(ch)
	d9, _ := base64.URLEncoding.DecodeString(r9)
	assert.Equal("", string(d9))

	r10 := ToUrlBase64(io.EOF)
	d10, _ := base64.URLEncoding.DecodeString(r10)
	assert.Equal("EOF", string(d10))

	r11 := ToUrlBase64(errors.New("test"))
	d11, _ := base64.URLEncoding.DecodeString(r11)
	assert.Equal("test", string(d11))

	typedNil := (*int)(nil)
	r12 := ToUrlBase64(typedNil)
	d12, _ := base64.URLEncoding.DecodeString(r12)
	assert.Equal("", string(d12))

	type nilInterface interface {
	}
	var nI nilInterface = nil
	d13, _ := base64.URLEncoding.DecodeString(ToUrlBase64(nI))
	assert.Equal("", string(d13))

	var p unsafe.Pointer
	d14, _ := base64.URLEncoding.DecodeString(ToUrlBase64(p))
	assert.Equal("", string(d14))

	r15 := ToUrlBase64("4+3/4?=")
	d15, _ := base64.URLEncoding.DecodeString(r15)
	assert.Equal("4+3/4?=", string(d15))
}

func TestToRawStdBase64(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestToRawStdBase64")

	r1 := ToRawStdBase64("abc")
	d1, _ := base64.RawStdEncoding.DecodeString(r1)
	assert.Equal("abc", string(d1))

	r2 := ToRawStdBase64([]byte("abc"))
	d2, _ := base64.RawStdEncoding.DecodeString(r2)
	assert.Equal("abc", string(d2))

	r3 := ToRawStdBase64(123)
	d3, _ := base64.RawStdEncoding.DecodeString(r3)
	assert.Equal("123", string(d3))

	r4 := ToRawStdBase64(11.11)
	d4, _ := base64.RawStdEncoding.DecodeString(r4)
	assert.Equal("11.11", string(d4))

	r5 := ToRawStdBase64(map[string]any{"name": "duke", "quantity": 1})
	d5, _ := base64.RawStdEncoding.DecodeString(r5)
	assert.Equal("{\"name\":\"duke\",\"quantity\":1}", string(d5))

	r6 := ToRawStdBase64([]int64{7, 5, 9, 4, 23})
	d6, _ := base64.RawStdEncoding.DecodeString(r6)
	assert.Equal("[7,5,9,4,23]", string(d6))

	r7 := ToRawStdBase64([]string{"7", "5", "9", "4", "23"})
	d7, _ := base64.RawStdEncoding.DecodeString(r7)
	assert.Equal("[\"7\",\"5\",\"9\",\"4\",\"23\"]", string(d7))

	r8 := ToRawStdBase64(nil)
	d8, _ := base64.RawStdEncoding.DecodeString(r8)
	assert.Equal("", string(d8))

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	r9 := ToRawStdBase64(ch)
	d9, _ := base64.RawStdEncoding.DecodeString(r9)
	assert.Equal("", string(d9))

	r10 := ToRawStdBase64(io.EOF)
	d10, _ := base64.RawStdEncoding.DecodeString(r10)
	assert.Equal("EOF", string(d10))

	r11 := ToRawStdBase64(errors.New("test"))
	d11, _ := base64.RawStdEncoding.DecodeString(r11)
	assert.Equal("test", string(d11))

	typedNil := (*int)(nil)
	r12 := ToRawStdBase64(typedNil)
	d12, _ := base64.RawStdEncoding.DecodeString(r12)
	assert.Equal("", string(d12))

	type nilInterface interface {
	}
	var nI nilInterface = nil
	d13, _ := base64.RawStdEncoding.DecodeString(ToRawStdBase64(nI))
	assert.Equal("", string(d13))

	var p unsafe.Pointer
	d14, _ := base64.RawStdEncoding.DecodeString(ToRawStdBase64(p))
	assert.Equal("", string(d14))
}

func TestToRawUrlBase64(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestToRawUrlBase64")

	r1 := ToRawUrlBase64("abc")
	d1, _ := base64.RawURLEncoding.DecodeString(r1)
	assert.Equal("abc", string(d1))

	r2 := ToRawUrlBase64([]byte("abc"))
	d2, _ := base64.RawURLEncoding.DecodeString(r2)
	assert.Equal("abc", string(d2))

	r3 := ToRawUrlBase64(123)
	d3, _ := base64.RawURLEncoding.DecodeString(r3)
	assert.Equal("123", string(d3))

	r4 := ToRawUrlBase64(11.11)
	d4, _ := base64.RawURLEncoding.DecodeString(r4)
	assert.Equal("11.11", string(d4))

	r5 := ToRawUrlBase64(map[string]any{"name": "duke", "quantity": 1})
	d5, _ := base64.RawURLEncoding.DecodeString(r5)
	assert.Equal("{\"name\":\"duke\",\"quantity\":1}", string(d5))

	r6 := ToRawUrlBase64([]int64{7, 5, 9, 4, 23})
	d6, _ := base64.RawURLEncoding.DecodeString(r6)
	assert.Equal("[7,5,9,4,23]", string(d6))

	r7 := ToRawUrlBase64([]string{"7", "5", "9", "4", "23"})
	d7, _ := base64.RawURLEncoding.DecodeString(r7)
	assert.Equal("[\"7\",\"5\",\"9\",\"4\",\"23\"]", string(d7))

	r8 := ToRawUrlBase64(nil)
	d8, _ := base64.RawURLEncoding.DecodeString(r8)
	assert.Equal("", string(d8))

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	r9 := ToRawUrlBase64(ch)
	d9, _ := base64.RawURLEncoding.DecodeString(r9)
	assert.Equal("", string(d9))

	r10 := ToRawUrlBase64(io.EOF)
	d10, _ := base64.RawURLEncoding.DecodeString(r10)
	assert.Equal("EOF", string(d10))

	r11 := ToRawUrlBase64(errors.New("test"))
	d11, _ := base64.RawURLEncoding.DecodeString(r11)
	assert.Equal("test", string(d11))

	typedNil := (*int)(nil)
	r12 := ToRawUrlBase64(typedNil)
	d12, _ := base64.RawURLEncoding.DecodeString(r12)
	assert.Equal("", string(d12))

	type nilInterface interface {
	}
	var nI nilInterface = nil
	d13, _ := base64.RawURLEncoding.DecodeString(ToRawUrlBase64(nI))
	assert.Equal("", string(d13))

	var p unsafe.Pointer
	d14, _ := base64.RawURLEncoding.DecodeString(ToRawUrlBase64(p))
	assert.Equal("", string(d14))

	r15 := ToRawUrlBase64("4+3/4?=")
	d15, _ := base64.RawURLEncoding.DecodeString(r15)
	assert.Equal("4+3/4?=", string(d15))
}

func TestToBigInt(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestToBigInt")

	tests := []struct {
		name   string
		input  any
		want   *big.Int
		hasErr bool
	}{
		{
			name:  "int",
			input: 42,
			want:  big.NewInt(42),
		},
		{
			name:  "int8",
			input: int8(127),
			want:  big.NewInt(127),
		},
		{
			name:  "int16",
			input: int16(32000),
			want:  big.NewInt(32000),
		},
		{
			name:  "int32",
			input: int32(123456),
			want:  big.NewInt(123456),
		},
		{
			name:  "int64",
			input: int64(987654321),
			want:  big.NewInt(987654321),
		},
		{
			name:  "uint",
			input: uint(987654321),
			want:  big.NewInt(987654321),
		},
		{
			name:  "uint8",
			input: uint8(255),
			want:  big.NewInt(255),
		},
		{
			name:  "uint16",
			input: uint16(65535),
			want:  big.NewInt(65535),
		},
		{
			name:  "uint32",
			input: uint32(4294967295),
			want:  big.NewInt(4294967295),
		},
		{
			name:  "uint64",
			input: uint64(18446744073709551615),
			want:  new(big.Int).SetUint64(18446744073709551615),
		},
		{
			name:   "unsupported type",
			input:  3.14, // Unsupported type
			hasErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBigInt(tt.input)
			if (err != nil) != tt.hasErr {
				t.Errorf("ToBigInt() error = %v, hasErr %v", err, tt.hasErr)
				return
			}

			assert.Equal(tt.want, got)
		})
	}
}

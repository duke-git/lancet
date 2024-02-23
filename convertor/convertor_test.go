package convertor

import (
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"

	"github.com/duke-git/lancet/internal"
	"github.com/duke-git/lancet/validator"
)

func TestToChar(t *testing.T) {
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

func TestToBool(t *testing.T) {
	assert := internal.NewAssert(t, "TestToBool")

	cases := []string{"1", "true", "True", "false", "False", "0", "123", "0.0", "abc"}
	expected := []bool{true, true, true, false, false, false, false, false, false}

	for i := 0; i < len(cases); i++ {
		actual, _ := ToBool(cases[i])
		assert.Equal(expected[i], actual)
	}
}

func TestToBytes(t *testing.T) {
	assert := internal.NewAssert(t, "TestToBytes")

	cases := []interface{}{
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
	assert := internal.NewAssert(t, "TestToInt")

	cases := []interface{}{"123", "-123", 123,
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
	assert := internal.NewAssert(t, "TestToFloat")

	cases := []interface{}{
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
	assert := internal.NewAssert(t, "TestToString")

	aMap := make(map[string]int)
	aMap["a"] = 1
	aMap["b"] = 2
	aMap["c"] = 3

	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "TestStruct"}

	cases := []interface{}{
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
		"[1,2,3]", "{\"a\":1,\"b\":2,\"c\":3}", "{\"Name\":\"TestStruct\"}", "hello"}

	for i := 0; i < len(cases); i++ {
		actual := ToString(cases[i])
		assert.Equal(expected[i], actual)
	}
}
func TestToJson(t *testing.T) {
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

func TestStructToMap(t *testing.T) {
	assert := internal.NewAssert(t, "TestStructToMap")

	type People struct {
		Name string `json:"name"`
		age  int
	}
	p := People{
		"test",
		100,
	}
	pm, _ := StructToMap(p)
	var expected = map[string]interface{}{"name": "test"}
	assert.Equal(expected, pm)
}

func TestColorHexToRGB(t *testing.T) {
	colorHex := "#003366"
	r, g, b := ColorHexToRGB(colorHex)
	colorRGB := fmt.Sprintf("%d,%d,%d", r, g, b)
	expected := "0,51,102"

	assert := internal.NewAssert(t, "TestColorHexToRGB")
	assert.Equal(expected, colorRGB)
}

func TestColorRGBToHex(t *testing.T) {
	r := 0
	g := 51
	b := 102
	colorHex := ColorRGBToHex(r, g, b)
	expected := "#003366"

	assert := internal.NewAssert(t, "TestColorRGBToHex")
	assert.Equal(expected, colorHex)
}

func TestToChannel(t *testing.T) {
	assert := internal.NewAssert(t, "TestToChannel")

	ch := ToChannel([]interface{}{1, 2, 3})
	val1, _ := <-ch
	assert.Equal(1, val1)

	val2, _ := <-ch
	assert.Equal(2, val2)

	val3, _ := <-ch
	assert.Equal(3, val3)

	_, ok := <-ch
	assert.Equal(false, ok)
}

func TestEncodeByte(t *testing.T) {
	assert := internal.NewAssert(t, "TestEncodeByte")

	byteData, _ := EncodeByte("abc")
	expected := []byte{6, 12, 0, 3, 97, 98, 99}

	assert.Equal(expected, byteData)
}

func TestDecodeByte(t *testing.T) {
	assert := internal.NewAssert(t, "TestDecodeByte")

	var obj string
	byteData := []byte{6, 12, 0, 3, 97, 98, 99}
	DecodeByte(byteData, &obj)
	assert.Equal("abc", obj)
}

func TestDeepClone(t *testing.T) {
	// assert := internal.NewAssert(t, "TestDeepClone")

	type Struct struct {
		Str        string
		Int        int
		Float      float64
		Bool       bool
		Nil        interface{}
		unexported string
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

	for i, item := range cases {
		cloned := DeepClone(item)

		t.Log(cloned)
		if &cloned == &item {
			t.Fatalf("[TestDeepClone case #%d failed]: equal pointer", i)
		}

		if !reflect.DeepEqual(item, cloned) {
			t.Fatalf("[TestDeepClone case #%d failed] unequal objects", i)
		}
	}
}

func TestCopyProperties(t *testing.T) {
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
	assert := internal.NewAssert(t, "TestUtf8ToGbk")

	utf8Data := []byte("hello")
	gbkData, err := Utf8ToGbk(utf8Data)

	assert.Equal(true, utf8.Valid(utf8Data))
	assert.Equal(true, validator.IsGBK(gbkData))
	assert.IsNil(err)
}

func TestGbkToUtf8(t *testing.T) {
	assert := internal.NewAssert(t, "TestGbkToUtf8")

	gbkData, err := Utf8ToGbk([]byte("hello"))
	utf8Data, err := GbkToUtf8(gbkData)

	assert.IsNil(err)
	assert.Equal(true, utf8.Valid(utf8Data))
	assert.Equal("hello", string(utf8Data))
}

func TestMapToStruct(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMapToStruct")

	type Address struct {
		Street string `json:"street"`
		Number int    `json:"number"`
	}

	type Person struct {
		Name  string   `json:"name"`
		Age   int      `json:"age"`
		Phone string   `json:"phone"`
		Addr  *Address `json:"address"`
	}

	m := map[string]interface{}{
		"name":  "Nothin",
		"age":   28,
		"phone": "123456789",
		"address": map[string]interface{}{
			"street": "test",
			"number": 1,
		},
	}

	var p Person
	err := MapToStruct(m, &p)
	assert.IsNil(err)
	assert.Equal(m["name"], p.Name)
	assert.Equal(m["age"], p.Age)
	assert.Equal(m["phone"], p.Phone)
	assert.Equal("test", p.Addr.Street)
	assert.Equal(1, p.Addr.Number)
}

package convertor

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/duke-git/lancet/internal"
)

func TestToChar(t *testing.T) {
	cases := []string{"", "abc", "1 2#3"}
	expected := [][]string{
		{""},
		{"a", "b", "c"},
		{"1", " ", "2", "#", "3"},
	}
	for i := 0; i < len(cases); i++ {
		res := ToChar(cases[i])
		if !reflect.DeepEqual(res, expected[i]) {
			internal.LogFailedTestInfo(t, "ToChar", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestToBool(t *testing.T) {
	cases := []string{"true", "True", "false", "False", "0", "1", "123"}
	expected := []bool{true, true, false, false, false, true, false}

	for i := 0; i < len(cases); i++ {
		res, _ := ToBool(cases[i])
		if res != expected[i] {
			internal.LogFailedTestInfo(t, "ToBool", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestToBytes(t *testing.T) {
	cases := []interface{}{
		0,
		false,
		"1",
	}
	expected := [][]byte{
		{3, 4, 0, 0},
		{3, 2, 0, 0},
		{4, 12, 0, 1, 49},
	}
	for i := 0; i < len(cases); i++ {
		res, _ := ToBytes(cases[i])
		fmt.Println(res)
		if !reflect.DeepEqual(res, expected[i]) {
			internal.LogFailedTestInfo(t, "ToBytes", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestToInt(t *testing.T) {
	cases := []interface{}{"123", "-123", 123,
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float32(12.3), float64(12.3),
		"abc", false, "111111111111111111111111111111111111111"}

	expected := []int64{123, -123, 123, 123, 123, 123, 123, 123, 12, 12, 0, 0, 0}

	for i := 0; i < len(cases); i++ {
		res, _ := ToInt(cases[i])
		if res != expected[i] {
			internal.LogFailedTestInfo(t, "ToInt", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestToFloat(t *testing.T) {
	cases := []interface{}{
		"", "-1", "-.11", "1.23e3", ".123e10", "abc",
		int(0), int8(1), int16(-1), int32(123), int64(123),
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float64(12.3), float32(12.3),
	}
	expected := []float64{0, -1, -0.11, 1230, 0.123e10, 0,
		0, 1, -1, 123, 123, 123, 123, 123, 123, 123, 12.3, 12.300000190734863}

	for i := 0; i < len(cases); i++ {
		res, _ := ToFloat(cases[i])
		if res != expected[i] {
			internal.LogFailedTestInfo(t, "ToFloat", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}

func TestToString(t *testing.T) {
	// map
	aMap := make(map[string]int)
	aMap["a"] = 1
	aMap["b"] = 2
	aMap["c"] = 3

	// struct
	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "TestStruct"}

	cases := []interface{}{
		int(0), int8(1), int16(-1), int32(123), int64(123),
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float64(12.3), float32(12.3),
		true, false,
		[]int{1, 2, 3}, aMap, aStruct, []byte{104, 101, 108, 108, 111}}

	expected := []string{"0", "1", "-1", "123", "123", "123", "123", "123",
		"123", "123", "12.3", "12.300000190734863", "true", "false",
		"[1,2,3]", "{\"a\":1,\"b\":2,\"c\":3}", "{\"Name\":\"TestStruct\"}", "hello"}

	for i := 0; i < len(cases); i++ {
		res := ToString(cases[i])
		if res != expected[i] {
			internal.LogFailedTestInfo(t, "ToString", cases[i], expected[i], res)
			t.FailNow()
		}
	}
}
func TestToJson(t *testing.T) {
	// map
	aMap := make(map[string]int)
	aMap["a"] = 1
	aMap["b"] = 2
	aMap["c"] = 3

	mapJson := "{\"a\":1,\"b\":2,\"c\":3}"
	r1, _ := ToJson(aMap)
	if r1 != mapJson {
		internal.LogFailedTestInfo(t, "ToJson", aMap, mapJson, r1)
		t.FailNow()
	}

	// struct
	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "TestStruct"}
	structJson := "{\"Name\":\"TestStruct\"}"
	r2, _ := ToJson(aStruct)
	if r2 != structJson {
		internal.LogFailedTestInfo(t, "ToJson", aMap, mapJson, r1)
		t.FailNow()
	}
}

func TestStructToMap(t *testing.T) {
	type People struct {
		Name string `json:"name"`
		age  int
	}

	p1 := People{
		"test",
		100,
	}

	pm1, _ := StructToMap(p1)
	m1 := make(map[string]interface{})
	m1["name"] = "test"
	//exp1["100"] = 100

	if !reflect.DeepEqual(pm1, m1) {
		internal.LogFailedTestInfo(t, "StructToMap", p1, m1, pm1)
		t.FailNow()
	}

	p2 := People{
		"test",
		100,
	}

	pm2, _ := StructToMap(p1)
	m2 := make(map[string]interface{})
	m2["name"] = "test"
	m2["100"] = 100

	if reflect.DeepEqual(pm2, m2) {
		internal.LogFailedTestInfo(t, "StructToMap", p2, m2, pm2)
		t.FailNow()
	}
}

func TestColorHexToRGB(t *testing.T) {
	colorHex := "#003366"
	r, g, b := ColorHexToRGB(colorHex)
	colorRGB := fmt.Sprintf("%d,%d,%d", r, g, b)
	expected := "0,51,102"

	if colorRGB != expected {
		internal.LogFailedTestInfo(t, "ColorHexToRGB", colorHex, expected, colorRGB)
		t.FailNow()
	}
}

func TestColorRGBToHex(t *testing.T) {
	r := 0
	g := 51
	b := 102
	colorRGB := fmt.Sprintf("%d,%d,%d", r, g, b)
	colorHex := ColorRGBToHex(r, g, b)
	expected := "#003366"

	if colorHex != expected {
		internal.LogFailedTestInfo(t, "ColorHexToRGB", colorRGB, expected, colorHex)
		t.FailNow()
	}
}

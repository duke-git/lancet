// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package convertor implements some functions to convert data.
package convertor

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// ToBool convert string to a boolean
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// ToBool convert interface to bytes
func ToBytes(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToChar convert string to char slice
func ToChar(s string) []string {
	c := make([]string, 0)
	if len(s) == 0 {
		c = append(c, "")
	}
	for _, v := range s {
		c = append(c, string(v))
	}
	return c
}

// ToString convert value to string
func ToString(value interface{}) string {
	var res string
	if value == nil {
		return res
	}
	switch v := value.(type) {
	case float64:
		res = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		res = strconv.FormatFloat(float64(v), 'f', -1, 64)
	case int:
		res = strconv.Itoa(v)
	case uint:
		res = strconv.Itoa(int(v))
	case int8:
		res = strconv.Itoa(int(v))
	case uint8:
		res = strconv.Itoa(int(v))
	case int16:
		res = strconv.Itoa(int(v))
	case uint16:
		res = strconv.Itoa(int(v))
	case int32:
		res = strconv.Itoa(int(v))
	case uint32:
		res = strconv.Itoa(int(v))
	case int64:
		res = strconv.FormatInt(v, 10)
	case uint64:
		res = strconv.FormatUint(v, 10)
	case string:
		res = value.(string)
	case []byte:
		res = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		res = string(newValue)
	}
	return res
}

// ToJson convert value to a valid json string
func ToJson(value interface{}) (string, error) {
	res, err := json.Marshal(value)
	if err != nil {
		res = []byte("")
	}

	return string(res), err
}

// ToFloat convert value to a float64, if input is not a float return 0.0 and error
func ToFloat(value interface{}) (float64, error) {
	v := reflect.ValueOf(value)

	res := 0.0
	err := fmt.Errorf("ToInt: unvalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = float64(v.Int())
		return res, nil
	case uint, uint8, uint16, uint32, uint64:
		res = float64(v.Uint())
		return res, nil
	case float32, float64:
		res = v.Float()
		return res, nil
	case string:
		res, err = strconv.ParseFloat(v.String(), 64)
		if err != nil {
			res = 0.0
		}
		return res, err
	default:
		return res, err
	}
}

// ToInt convert value to a int64, if input is not a numeric format return 0 and error
func ToInt(value interface{}) (int64, error) {
	v := reflect.ValueOf(value)

	var res int64
	err := fmt.Errorf("ToInt: invalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = v.Int()
		return res, nil
	case uint, uint8, uint16, uint32, uint64:
		res = int64(v.Uint())
		return res, nil
	case float32, float64:
		res = int64(v.Float())
		return res, nil
	case string:
		res, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			res = 0
		}
		return res, err
	default:
		return res, err
	}
}

// StructToMap convert struct to map, only convert exported struct field
// map key is specified same as struct field tag `json` value
func StructToMap(value interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", value)
	}

	res := make(map[string]interface{})

	fieldNum := t.NumField()
	pattern := `^[A-Z]`
	regex := regexp.MustCompile(pattern)
	for i := 0; i < fieldNum; i++ {
		name := t.Field(i).Name
		tag := t.Field(i).Tag.Get("json")
		if regex.MatchString(name) && tag != "" {
			//res[name] = v.Field(i).Interface()
			res[tag] = v.Field(i).Interface()
		}
	}

	return res, nil
}

// ColorHexToRGB convert hex color to rgb color
func ColorHexToRGB(colorHex string) (red, green, blue int) {
	colorHex = strings.TrimPrefix(colorHex, "#")
	color64, err := strconv.ParseInt(colorHex, 16, 32)
	if err != nil {
		return
	}
	color := int(color64)
	return color >> 16, (color & 0x00FF00) >> 8, color & 0x0000FF
}

// ColorRGBToHex convert rgb color to hex color
func ColorRGBToHex(red, green, blue int) string {
	r := strconv.FormatInt(int64(red), 16)
	g := strconv.FormatInt(int64(green), 16)
	b := strconv.FormatInt(int64(blue), 16)

	if len(r) == 1 {
		r = "0" + r
	}
	if len(g) == 1 {
		g = "0" + g
	}
	if len(b) == 1 {
		b = "0" + b
	}

	return "#" + r + g + b
}

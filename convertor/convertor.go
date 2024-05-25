// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package convertor implements some functions to convert data.
package convertor

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ToBool convert string to a boolean
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// ToBytes convert interface to bytes
func ToBytes(value interface{}) ([]byte, error) {
	v := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		number := v.Int()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case uint, uint8, uint16, uint32, uint64:
		number := v.Uint()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err := binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes(), err
	case float32:
		number := float32(v.Float())
		bits := math.Float32bits(number)
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, bits)
		return bytes, nil
	case float64:
		number := v.Float()
		bits := math.Float64bits(number)
		bytes := make([]byte, 8)
		binary.BigEndian.PutUint64(bytes, bits)
		return bytes, nil
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool()), nil
	case string:
		return []byte(v.String()), nil
	case []byte:
		return v.Bytes(), nil
	default:
		newValue, err := json.Marshal(value)
		return newValue, err
	}
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
	if value == nil {
		return ""
	}

	switch val := value.(type) {
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case string:
		return val
	case []byte:
		return string(val)
	default:
		b, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return string(b)
	}
}

// ToJson convert value to a valid json string
func ToJson(value interface{}) (string, error) {
	res, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(res), nil
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
		v = v.Elem()
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

// ToChannel convert a array of elements to a read-only channels
func ToChannel(array []interface{}) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		for _, item := range array {
			ch <- item
		}
		close(ch)
	}()

	return ch
}

// EncodeByte encode data to byte
func EncodeByte(data interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// DecodeByte decode byte data to target object
func DecodeByte(data []byte, target interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(target)
}

// DeepClone creates a deep copy of passed item.
// can't clone unexported field of struct
func DeepClone(src interface{}) interface{} {
	c := cloner{
		ptrs: map[reflect.Type]map[uintptr]reflect.Value{},
	}
	result := c.clone(reflect.ValueOf(src))
	if result.Kind() == reflect.Invalid {
		return nil
	}

	return result.Interface()
}

// CopyProperties copies each field from the source into the destination. It recursively copies struct pointers and interfaces that contain struct pointers.
// use json.Marshal/Unmarshal, so json tag should be set for fields of dst and src struct.
func CopyProperties(dst, src interface{}) error {
	dstType, srcType := reflect.TypeOf(dst), reflect.TypeOf(src)

	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("CopyProperties: parameter dst should be struct pointer")
	}

	if srcType.Kind() == reflect.Ptr {
		srcType = srcType.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		return errors.New("CopyProperties: parameter src should be a struct or struct pointer")
	}

	bytes, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("CopyProperties: unable to marshal src: %s", err)
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return fmt.Errorf("CopyProperties: unable to unmarshal into dst: %s", err)
	}

	return nil
}

// ToInterface converts reflect value to its interface type.
func ToInterface(v reflect.Value) (value interface{}, ok bool) {
	if v.IsValid() && v.CanInterface() {
		return v.Interface(), true
	}
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint(), true
	case reflect.Float32, reflect.Float64:
		return v.Float(), true
	case reflect.Complex64, reflect.Complex128:
		return v.Complex(), true
	case reflect.String:
		return v.String(), true
	case reflect.Ptr:
		return ToInterface(v.Elem())
	case reflect.Interface:
		return ToInterface(v.Elem())
	default:
		return nil, false
	}
}

// Utf8ToGbk convert utf8 encoding data to GBK encoding data.
func Utf8ToGbk(bs []byte) ([]byte, error) {
	r := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewEncoder())
	b, err := io.ReadAll(r)
	return b, err
}

// GbkToUtf8 convert GBK encoding data to utf8 encoding data.
func GbkToUtf8(bs []byte) ([]byte, error) {
	r := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewDecoder())
	b, err := io.ReadAll(r)
	return b, err
}

// MapToStruct converts map to struct
func MapToStruct(m map[string]interface{}, structObj interface{}) error {
	for k, v := range m {
		err := setStructField(structObj, k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

// ToStdBase64 convert data to standard base64 encoding.
func ToStdBase64(value interface{}) string {
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		return ""
	}
	switch value.(type) {
	case []byte:
		return base64.StdEncoding.EncodeToString(value.([]byte))
	case string:
		return base64.StdEncoding.EncodeToString([]byte(value.(string)))
	case error:
		return base64.StdEncoding.EncodeToString([]byte(value.(error).Error()))
	default:
		marshal, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		return base64.StdEncoding.EncodeToString(marshal)
	}
}

// ToUrlBase64 convert data to URL base64 encoding.
func ToUrlBase64(value interface{}) string {
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		return ""
	}
	switch value.(type) {
	case []byte:
		return base64.URLEncoding.EncodeToString(value.([]byte))
	case string:
		return base64.URLEncoding.EncodeToString([]byte(value.(string)))
	case error:
		return base64.URLEncoding.EncodeToString([]byte(value.(error).Error()))
	default:
		marshal, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		return base64.URLEncoding.EncodeToString(marshal)
	}
}

// ToRawStdBase64 convert data to raw standard base64 encoding.
func ToRawStdBase64(value interface{}) string {
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		return ""
	}
	switch value.(type) {
	case []byte:
		return base64.RawStdEncoding.EncodeToString(value.([]byte))
	case string:
		return base64.RawStdEncoding.EncodeToString([]byte(value.(string)))
	case error:
		return base64.RawStdEncoding.EncodeToString([]byte(value.(error).Error()))
	default:
		marshal, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		return base64.RawStdEncoding.EncodeToString(marshal)
	}
}

// ToRawUrlBase64 convert data to raw URL base64 encoding.
func ToRawUrlBase64(value interface{}) string {
	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		return ""
	}
	switch value.(type) {
	case []byte:
		return base64.RawURLEncoding.EncodeToString(value.([]byte))
	case string:
		return base64.RawURLEncoding.EncodeToString([]byte(value.(string)))
	case error:
		return base64.RawURLEncoding.EncodeToString([]byte(value.(error).Error()))
	default:
		marshal, err := json.Marshal(value)
		if err != nil {
			return ""
		}
		return base64.RawURLEncoding.EncodeToString(marshal)
	}
}

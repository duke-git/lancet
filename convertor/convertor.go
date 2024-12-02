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
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/duke-git/lancet/v2/structs"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// ToBool convert string to boolean.
// Play: https://go.dev/play/p/ARht2WnGdIN
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// ToBytes convert value to byte slice.
// Play: https://go.dev/play/p/fAMXYFDvOvr
func ToBytes(value any) ([]byte, error) {
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

// ToChar convert string to char slice.
// Play: https://go.dev/play/p/JJ1SvbFkVdM
func ToChar(s string) []string {
	c := make([]string, 0, len(s))
	if len(s) == 0 {
		c = append(c, "")
	}
	for _, v := range s {
		c = append(c, string(v))
	}
	return c
}

// ToChannel convert a slice of elements to a read-only channel.
// Play: https://go.dev/play/p/hOx_oYZbAnL
func ToChannel[T any](array []T) <-chan T {
	ch := make(chan T)

	go func() {
		for _, item := range array {
			ch <- item
		}
		close(ch)
	}()

	return ch
}

// ToString convert value to string
// for number, string, []byte, will convert to string
// for other type (slice, map, array, struct) will call json.Marshal.
// Play: https://go.dev/play/p/nF1zOOslpQq
func ToString(value any) string {
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

		// todo: maybe we should't supprt other type conversion
		// v := reflect.ValueOf(value)
		// log.Panicf("Unsupported data type: %s ", v.String())
		// return ""
	}
}

// ToJson convert value to a json string.
// Play: https://go.dev/play/p/2rLIkMmXWvR
func ToJson(value any) (string, error) {
	result, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// ToFloat convert value to float64, if input is not a float return 0.0 and error.
// Play: https://go.dev/play/p/4YTmPCibqHJ
func ToFloat(value any) (float64, error) {
	v := reflect.ValueOf(value)

	result := 0.0
	err := fmt.Errorf("ToInt: unvalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = float64(v.Int())
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = float64(v.Uint())
		return result, nil
	case float32, float64:
		result = v.Float()
		return result, nil
	case string:
		result, err = strconv.ParseFloat(v.String(), 64)
		if err != nil {
			result = 0.0
		}
		return result, err
	default:
		return result, err
	}
}

// ToInt convert value to int64 value, if input is not numerical, return 0 and error.
// Play: https://go.dev/play/p/9_h9vIt-QZ_b
func ToInt(value any) (int64, error) {
	v := reflect.ValueOf(value)

	var result int64
	err := fmt.Errorf("ToInt: invalid value type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = v.Int()
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = int64(v.Uint())
		return result, nil
	case float32, float64:
		result = int64(v.Float())
		return result, nil
	case string:
		result, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			result = 0
		}
		return result, err
	default:
		return result, err
	}
}

// ToPointer returns a pointer to passed value.
// Play: https://go.dev/play/p/ASf_etHNlw1
func ToPointer[T any](value T) *T {
	return &value
}

// ToMap convert a slice of structs to a map based on iteratee function.
// Play: https://go.dev/play/p/tVFy7E-t24l
func ToMap[T any, K comparable, V any](array []T, iteratee func(T) (K, V)) map[K]V {
	result := make(map[K]V, len(array))
	for _, item := range array {
		k, v := iteratee(item)
		result[k] = v
	}

	return result
}

// StructToMap convert struct to map, only convert exported struct field
// map key is specified same as struct field tag `json` value.
// Play: https://go.dev/play/p/KYGYJqNUBOI
func StructToMap(value any) (map[string]any, error) {
	return structs.ToMap(value)
}

// MapToSlice convert map to slice based on iteratee function.
// Play: https://go.dev/play/p/dmX4Ix5V6Wl
func MapToSlice[T any, K comparable, V any](aMap map[K]V, iteratee func(K, V) T) []T {
	result := make([]T, 0, len(aMap))

	for k, v := range aMap {
		result = append(result, iteratee(k, v))
	}

	return result
}

// ColorHexToRGB convert hex color to rgb color.
// Play: https://go.dev/play/p/o7_ft-JCJBV
func ColorHexToRGB(colorHex string) (red, green, blue int) {
	colorHex = strings.TrimPrefix(colorHex, "#")
	color64, err := strconv.ParseInt(colorHex, 16, 32)
	if err != nil {
		return
	}
	color := int(color64)
	return color >> 16, (color & 0x00FF00) >> 8, color & 0x0000FF
}

// ColorRGBToHex convert rgb color to hex color.
// Play: https://go.dev/play/p/nzKS2Ro87J1
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

// EncodeByte encode data to byte slice.
// Play: https://go.dev/play/p/DVmM1G5JfuP
func EncodeByte(data any) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// DecodeByte decode byte slice data to target object.
// Play: https://go.dev/play/p/zI6xsmuQRbn
func DecodeByte(data []byte, target any) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(target)
}

// DeepClone creates a deep copy of passed item.
// can't clone unexported field of struct
// Play: https://go.dev/play/p/j4DP5dquxnk
func DeepClone[T any](src T) T {
	c := cloner{
		ptrs: map[reflect.Type]map[uintptr]reflect.Value{},
	}
	result := c.clone(reflect.ValueOf(src))
	if result.Kind() == reflect.Invalid {
		var zeroValue T
		return zeroValue
	}

	return result.Interface().(T)
}

// CopyProperties copies each field from the source into the destination. It recursively copies struct pointers and interfaces that contain struct pointers.
// use json.Marshal/Unmarshal, so json tag should be set for fields of dst and src struct.
// Play: https://go.dev/play/p/oZujoB5Sgg5
func CopyProperties[T, U any](dst T, src U) error {
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
// Play: https://go.dev/play/p/syqw0-WG7Xd
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
// Play: https://go.dev/play/p/9FlIaFLArIL
func Utf8ToGbk(bs []byte) ([]byte, error) {
	r := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewEncoder())
	b, err := io.ReadAll(r)
	return b, err
}

// GbkToUtf8 convert GBK encoding data to utf8 encoding data.
// Play: https://go.dev/play/p/OphmHCN_9u8
func GbkToUtf8(bs []byte) ([]byte, error) {
	r := transform.NewReader(bytes.NewReader(bs), simplifiedchinese.GBK.NewDecoder())
	b, err := io.ReadAll(r)
	return b, err
}

// ToStdBase64 convert data to standard base64 encoding.
// Play: https://go.dev/play/p/_fLJqJD3NMo
func ToStdBase64(value any) string {
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
// Play: https://go.dev/play/p/C_d0GlvEeUR
func ToUrlBase64(value any) string {
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
// Play: https://go.dev/play/p/wSAr3sfkDcv
func ToRawStdBase64(value any) string {
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
// Play: https://go.dev/play/p/HwdDPFcza1O
func ToRawUrlBase64(value any) string {
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

// ToBigInt converts an integer of any supported type (int, int64, uint64, etc.) to *big.Int
// Play: todo
func ToBigInt[T any](v T) (*big.Int, error) {
	result := new(big.Int)

	switch v := any(v).(type) {
	case int:
		result.SetInt64(int64(v)) // Convert to int64 for big.Int
	case int8:
		result.SetInt64(int64(v))
	case int16:
		result.SetInt64(int64(v))
	case int32:
		result.SetInt64(int64(v))
	case int64:
		result.SetInt64(v)
	case uint:
		result.SetUint64(uint64(v)) // Convert to uint64 for big.Int
	case uint8:
		result.SetUint64(uint64(v))
	case uint16:
		result.SetUint64(uint64(v))
	case uint32:
		result.SetUint64(uint64(v))
	case uint64:
		result.SetUint64(v)
	default:
		return nil, fmt.Errorf("unsupported type: %T", v)
	}

	return result, nil
}

// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package slice implements some functions to manipulate slice.
package slice

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"unsafe"
)

// Contain check if the value is in the slice or not
func Contain(slice interface{}, value interface{}) bool {
	v := reflect.ValueOf(slice)
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == value {
				return true
			}
		}

	case reflect.Map:
		if v.MapIndex(reflect.ValueOf(value)).IsValid() {
			return true
		}
	case reflect.String:
		s := fmt.Sprintf("%v", slice)
		ss := fmt.Sprintf("%v", value)
		return strings.Contains(s, ss)
	}

	return false
}

// Chunk creates an slice of elements split into groups the length of `size`.
func Chunk(slice []interface{}, size int) [][]interface{} {
	var res [][]interface{}

	if len(slice) == 0 || size <= 0 {
		return res
	}

	length := len(slice)
	if size == 1 || size >= length {
		for _, v := range slice {
			var tmp []interface{}
			tmp = append(tmp, v)
			res = append(res, tmp)
		}
		return res
	}

	// divide slice equally
	divideNum := length/size + 1
	for i := 0; i < divideNum; i++ {
		if i == divideNum-1 {
			if len(slice[i*size:]) > 0 {
				res = append(res, slice[i*size:])
			}
		} else {
			res = append(res, slice[i*size:(i+1)*size])
		}
	}

	return res
}

// Difference creates an slice of whose element not included in the other given slice
func Difference(slice1, slice2 interface{}) interface{} {
	v := sliceValue(slice1)

	var indexes []int
	for i := 0; i < v.Len(); i++ {
		vi := v.Index(i).Interface()
		if !Contain(slice2, vi) {
			indexes = append(indexes, i)
		}
	}

	res := reflect.MakeSlice(v.Type(), len(indexes), len(indexes))
	for i := range indexes {
		res.Index(i).Set(v.Index(indexes[i]))
	}
	return res.Interface()
}

// Every return true if all of the values in the slice pass the predicate function.
// The function signature should be func(index int, value interface{}) bool .
func Every(slice, function interface{}) bool {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("Filter function must be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	var indexes []int
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			indexes = append(indexes, i)
		}
	}

	return len(indexes) == sv.Len()
}

// Some return true if any of the values in the list pass the predicate function.
// The function signature should be func(index int, value interface{}) bool .
func Some(slice, function interface{}) bool {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("Filter function must be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	var indexes []int
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			indexes = append(indexes, i)
		}
	}

	return len(indexes) > 0
}

// Filter iterates over elements of slice, returning an slice of all elements `signature` returns truthy for.
// The function signature should be func(index int, value interface{}) bool .
func Filter(slice, function interface{}) interface{} {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("Filter function must be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	var indexes []int
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			indexes = append(indexes, i)
		}
	}

	res := reflect.MakeSlice(sv.Type(), len(indexes), len(indexes))
	for i := range indexes {
		res.Index(i).Set(sv.Index(indexes[i]))
	}
	return res.Interface()
}

// Find iterates over elements of slice, returning the first one that passes a truth test on function.
// The function signature should be func(index int, value interface{}) bool .
func Find(slice, function interface{}) interface{} {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("Filter function must be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	var index int
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			index = i
			break
		}
	}
	return sv.Index(index).Interface()
}

// Map creates an slice of values by running each element of `slice` thru `function`.
// The function signature should be func(index int, value interface{}) interface{}.
func Map(slice, function interface{}) interface{} {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, nil) {
		panic("Map function must be of type func(int, " + elemType.String() + ")" + elemType.String())
	}

	res := reflect.MakeSlice(sv.Type(), sv.Len(), sv.Len())
	for i := 0; i < sv.Len(); i++ {
		res.Index(i).Set(fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0])
	}
	return res.Interface()
}

// Reduce creates an slice of values by running each element of `slice` thru `function`.
// The function signature should be func(index int, value1, value2 interface{}) interface{} .
func Reduce(slice, function, zero interface{}) interface{} {
	sv := sliceValue(slice)
	elementType := sv.Type().Elem()

	len := sv.Len()
	if len == 0 {
		return zero
	} else if len == 1 {
		return sv.Index(0).Interface()
	}

	fn := functionValue(function)
	if checkSliceCallbackFuncSignature(fn, elementType, elementType, elementType) {
		t := elementType.String()
		panic("Reduce function must be of type func(int, " + t + ", " + t + ")" + t)
	}

	var params [3]reflect.Value
	params[0] = reflect.ValueOf(0)
	params[1] = sv.Index(0)
	params[2] = sv.Index(1)

	res := fn.Call(params[:])[0]

	for i := 2; i < len; i++ {
		params[0] = reflect.ValueOf(i)
		params[1] = res
		params[2] = sv.Index(i)
		res = fn.Call(params[:])[0]
	}

	return res.Interface()
}

// InterfaceSlice convert param to slice of interface.
func InterfaceSlice(slice interface{}) []interface{} {
	sv := sliceValue(slice)
	if sv.IsNil() {
		return nil
	}

	res := make([]interface{}, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		res[i] = sv.Index(i).Interface()
	}

	return res
}

// StringSlice convert param to slice of string.
func StringSlice(slice interface{}) []string {
	var res []string

	v := sliceValue(slice)
	for i := 0; i < v.Len(); i++ {
		res = append(res, fmt.Sprint(v.Index(i)))
	}

	return res
}

// IntSlice convert param to slice of int.
func IntSlice(slice interface{}) ([]int, error) {
	var res []int

	sv := sliceValue(slice)
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		switch v := v.(type) {
		case int:
			res = append(res, v)
		default:
			return nil, errors.New("InvalidSliceElementType")
		}
	}

	return res, nil
}

// ConvertSlice convert original slice to new data type element of slice.
func ConvertSlice(originalSlice interface{}, newSliceType reflect.Type) interface{} {
	sv := sliceValue(originalSlice)
	if newSliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid newSliceType(non-slice type of type %T)", newSliceType))
	}

	newSlice := reflect.New(newSliceType)

	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))

	var newElemSize = int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())

	hdr.Cap = sv.Cap() * newElemSize
	hdr.Len = sv.Len() * newElemSize
	hdr.Data = sv.Pointer()

	return newSlice.Elem().Interface()
}

// DeleteByIndex delete the element of slice from start index to end index - 1.
// Delete i: s = append(s[:i], s[i+1:]...)
// Delete i to j: s = append(s[:i], s[j:]...)
func DeleteByIndex(slice interface{}, start int, end ...int) (interface{}, error) {
	v := sliceValue(slice)
	i := start
	if v.Len() == 0 || i < 0 || i > v.Len() {
		return nil, errors.New("InvalidStartIndex")
	}
	if len(end) > 0 {
		j := end[0]
		if j <= i || j > v.Len() {
			return nil, errors.New("InvalidEndIndex")
		}
		v = reflect.AppendSlice(v.Slice(0, i), v.Slice(j, v.Len()))
	} else {
		v = reflect.AppendSlice(v.Slice(0, i), v.Slice(i+1, v.Len()))
	}

	return v.Interface(), nil
}

// InsertByIndex insert the element into slice at index.
// Insert value: s = append(s[:i], append([]T{x}, s[i:]...)...)
// Insert slice: a = append(a[:i], append(b, a[i:]...)...)
func InsertByIndex(slice interface{}, index int, value interface{}) (interface{}, error) {
	v := sliceValue(slice)

	if index < 0 || index > v.Len() {
		return slice, errors.New("InvalidSliceIndex")
	}

	// value is slice
	vv := reflect.ValueOf(value)
	if vv.Kind() == reflect.Slice {
		if reflect.TypeOf(slice).Elem() != reflect.TypeOf(value).Elem() {
			return slice, errors.New("InvalidValueType")
		}
		v = reflect.AppendSlice(v.Slice(0, index), reflect.AppendSlice(vv.Slice(0, vv.Len()), v.Slice(index, v.Len())))
		return v.Interface(), nil
	}

	// value is not slice
	if reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		return slice, errors.New("InvalidValueType")
	}
	if index == v.Len() {
		return reflect.Append(v, reflect.ValueOf(value)).Interface(), nil
	}

	v = reflect.AppendSlice(v.Slice(0, index+1), v.Slice(index, v.Len()))
	v.Index(index).Set(reflect.ValueOf(value))

	return v.Interface(), nil
}

// UpdateByIndex update the slice element at index.
func UpdateByIndex(slice interface{}, index int, value interface{}) (interface{}, error) {
	v := sliceValue(slice)

	if index < 0 || index >= v.Len() {
		return slice, errors.New("InvalidSliceIndex")
	}
	if reflect.TypeOf(slice).Elem() != reflect.TypeOf(value) {
		return slice, errors.New("InvalidValueType")
	}

	v.Index(index).Set(reflect.ValueOf(value))

	return v.Interface(), nil
}

// Unique remove duplicate elements in slice.
func Unique(slice interface{}) interface{} {
	sv := sliceValue(slice)
	if sv.Len() == 0 {
		return slice
	}

	var res []interface{}
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		flag := true
		for j := range res {
			if v == res[j] {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, v)
		}
	}

	return res

	// if use map filter, the result slice element order is random, not same as origin slice
	//mp := make(map[interface{}]bool)
	//for i := 0; i < sv.Len(); i++ {
	//	v := sv.Index(i).Interface()
	//	mp[v] = true
	//}
	//
	//var res []interface{}
	//for k := range mp {
	//	res = append(res, mp[k])
	//}
	//return res

}

// ReverseSlice return slice of element order is reversed to the given slice
func ReverseSlice(slice interface{}) {
	v := sliceValue(slice)
	swp := reflect.Swapper(v.Interface())
	for i, j := 0, v.Len()-1; i < j; i, j = i+1, j-1 {
		swp(i, j)
	}
}

// SortByField return sorted slice by field
// Slice element should be struct, field type should be int, uint, string, or bool
// default sortType is ascending (asc), if descending order, set sortType to desc
func SortByField(slice interface{}, field string, sortType ...string) error {
	v := sliceValue(slice)
	t := v.Type().Elem()

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", slice)
	}

	// Find the field.
	sf, ok := t.FieldByName(field)
	if !ok {
		return fmt.Errorf("field name %s not found", field)
	}

	// Create a less function based on the field's kind.
	var less func(a, b reflect.Value) bool
	switch sf.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		less = func(a, b reflect.Value) bool { return a.Int() < b.Int() }
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		less = func(a, b reflect.Value) bool { return a.Uint() < b.Uint() }
	case reflect.Float32, reflect.Float64:
		less = func(a, b reflect.Value) bool { return a.Float() < b.Float() }
	case reflect.String:
		less = func(a, b reflect.Value) bool { return a.String() < b.String() }
	case reflect.Bool:
		less = func(a, b reflect.Value) bool { return !a.Bool() && b.Bool() }
	default:
		return fmt.Errorf("field type %s not supported", sf.Type)
	}

	sort.Slice(slice, func(i, j int) bool {
		a := v.Index(i)
		b := v.Index(j)
		if t.Kind() == reflect.Ptr {
			a = a.Elem()
			b = b.Elem()
		}
		a = a.FieldByIndex(sf.Index)
		b = b.FieldByIndex(sf.Index)
		return less(a, b)
	})

	if sortType[0] == "desc" {
		ReverseSlice(slice)
	}
	return nil
}

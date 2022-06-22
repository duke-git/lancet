// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package slice implements some functions to manipulate slice.
package slice

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strings"
)

// Contain check if the value is in the iterable type or not
func Contain(iterableType interface{}, value interface{}) bool {
	v := reflect.ValueOf(iterableType)

	switch kind := reflect.TypeOf(iterableType).Kind(); kind {
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
		s := iterableType.(string)
		ss, ok := value.(string)
		if !ok {
			panic("kind mismatch")
		}

		return strings.Contains(s, ss)
	default:
		panic(fmt.Sprintf("kind %s is not support", iterableType))
	}

	return false
}

// ContainSubSlice check if the slice contain subslice or not
func ContainSubSlice(slice interface{}, subslice interface{}) bool {
	super := sliceValue(slice)
	sub := sliceValue(subslice)

	if super.Type().Elem().Kind() != sub.Type().Elem().Kind() {
		return false
	}

	unique := make(map[interface{}]bool)
	for i := 0; i < super.Len(); i++ {
		v := super.Index(i).Interface()
		unique[v] = true
	}
	for i := 0; i < sub.Len(); i++ {
		v := sub.Index(i).Interface()
		if !unique[v] {
			return false
		}
	}

	return true
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

// Compact creates an slice with all falsey values removed. The values false, nil, 0, and "" are falsey
func Compact(slice interface{}) interface{} {
	sv := sliceValue(slice)

	var indexes []int
	for i := 0; i < sv.Len(); i++ {
		item := sv.Index(i).Interface()
		if item != nil && item != false && item != "" && item != 0 {
			indexes = append(indexes, i)
		}
	}
	res := reflect.MakeSlice(sv.Type(), len(indexes), len(indexes))
	for i := range indexes {
		res.Index(i).Set(sv.Index(indexes[i]))
	}

	return res.Interface()
}

// Concat creates a new slice concatenating slice with any additional slices and/or values.
func Concat(slice interface{}, values ...interface{}) interface{} {
	sv := sliceValue(slice)
	size := sv.Len()

	res := reflect.MakeSlice(sv.Type(), size, size)
	for i := 0; i < size; i++ {
		res.Index(i).Set(sv.Index(i))
	}

	for _, v := range values {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			vv := reflect.ValueOf(v)
			for i := 0; i < vv.Len(); i++ {
				res = reflect.Append(res, vv.Index(i))
			}
		} else {
			res = reflect.Append(res, reflect.ValueOf(v))
		}
	}

	return res.Interface()
}

// Difference creates an slice of whose element in slice1, not in slice2
func Difference(slice1, slice2 interface{}) interface{} {
	sv := sliceValue(slice1)

	var indexes []int
	for i := 0; i < sv.Len(); i++ {
		item := sv.Index(i).Interface()
		if !Contain(slice2, item) {
			indexes = append(indexes, i)
		}
	}

	res := reflect.MakeSlice(sv.Type(), len(indexes), len(indexes))
	for i := range indexes {
		res.Index(i).Set(sv.Index(indexes[i]))
	}
	return res.Interface()
}

// DifferenceBy it accepts iteratee which is invoked for each element of slice
// and values to generate the criterion by which they're compared.
// like lodash.js differenceBy: https://lodash.com/docs/4.17.15#differenceBy,
// the iterateeFn function signature should be func(index int, value interface{}) interface{}.
func DifferenceBy(slice interface{}, comparedSlice interface{}, iterateeFn interface{}) interface{} {
	sv := sliceValue(slice)
	smv := sliceValue(comparedSlice)
	fn := functionValue(iterateeFn)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, nil) {
		panic("function param should be of type func(" + elemType.String() + ")" + elemType.String())
	}

	slice1 := reflect.MakeSlice(sv.Type(), sv.Len(), sv.Len())
	for i := 0; i < sv.Len(); i++ {
		slice1.Index(i).Set(fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0])
	}

	slice2 := reflect.MakeSlice(smv.Type(), smv.Len(), smv.Len())
	for i := 0; i < smv.Len(); i++ {
		slice2.Index(i).Set(fn.Call([]reflect.Value{reflect.ValueOf(i), smv.Index(i)})[0])
	}

	sliceAfterMap := slice1.Interface()
	comparedSliceAfterMap := slice2.Interface()

	res := reflect.MakeSlice(sv.Type(), 0, 0)
	sm := sliceValue(sliceAfterMap)
	for i := 0; i < sm.Len(); i++ {
		item := sm.Index(i).Interface()
		if !Contain(comparedSliceAfterMap, item) {
			res = reflect.Append(res, sv.Index(i))
		}
	}

	return res.Interface()
}

// Equal checks if two slices are equal: the same length and all elements' order and value are equal
func Equal(slice1, slice2 interface{}) bool {
	sv1 := sliceValue(slice1)
	sv2 := sliceValue(slice2)

	if sv1.Type().Elem() != sv2.Type().Elem() {
		return false
	}

	if sv1.Len() != sv2.Len() {
		return false
	}

	for i := 0; i < sv1.Len(); i++ {
		if sv1.Index(i).Interface() != sv2.Index(i).Interface() {
			return false
		}
	}

	return true
}

// EqualWith checks if two slices are equal with comparator func
func EqualWith(slice1, slice2 interface{}, comparator interface{}) bool {
	sv1 := sliceValue(slice1)
	sv2 := sliceValue(slice2)

	fn := functionValue(comparator)
	// elemType1 := sv1.Type().Elem()
	// elemType2 := sv2.Type().Elem()
	// todo:  check fn signature: func(a elemType1.Kind(), b elemType2.Kind()) bool

	for i := 0; i < sv1.Len(); i++ {
		flag := fn.Call([]reflect.Value{sv1.Index(i), sv2.Index(i)})[0]
		if !flag.Bool() {
			return false
		}
	}

	return true
}

// Every return true if all of the values in the slice pass the predicate function.
// The function signature should be func(index int, value interface{}) bool .
func Every(slice, function interface{}) bool {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	var currentLength int
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			currentLength++
		}
	}

	return currentLength == sv.Len()
}

// None return true if all the values in the slice mismatch the criteria
// The function signature should be func(index int, value interface{}) bool .
func None(slice, function interface{}) bool {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	var currentLength int
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if !flag.Bool() {
			currentLength++
		}
	}

	return currentLength == sv.Len()
}

// Some return true if any of the values in the list pass the predicate function.
// The function signature should be func(index int, value interface{}) bool .
func Some(slice, function interface{}) bool {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	has := false
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			has = true
		}
	}

	return has
}

// Filter iterates over elements of slice, returning an slice of all elements `signature` returns truthy for.
// The function signature should be func(index int, value interface{}) bool .
func Filter(slice, function interface{}) interface{} {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	res := reflect.MakeSlice(sv.Type(), 0, 0)
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			res = reflect.Append(res, sv.Index(i))
		}
	}

	return res.Interface()
}

// Count iterates over elements of slice, returns a count of all matched elements
// The function signature should be func(index int, value interface{}) bool .
func Count(slice, function interface{}) int {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	var counter int
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			counter++
		}
	}

	return counter
}

// GroupBy iterate over elements of the slice, each element will be group by criteria, returns two slices
// The function signature should be func(index int, value interface{}) bool .
func GroupBy(slice, function interface{}) (interface{}, interface{}) {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	groupB := reflect.MakeSlice(sv.Type(), 0, 0)
	groupA := reflect.MakeSlice(sv.Type(), 0, 0)

	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			groupA = reflect.Append(groupA, sv.Index(i))
		} else {
			groupB = reflect.Append(groupB, sv.Index(i))
		}
	}

	return groupA.Interface(), groupB.Interface()
}

// Find iterates over elements of slice, returning the first one that passes a truth test on function.
// The function signature should be func(index int, value interface{}) bool .
func Find(slice, function interface{}) (interface{}, bool) {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	index := -1
	for i := 0; i < sv.Len(); i++ {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			index = i
			break
		}
	}

	if index == -1 {
		var none interface{}
		return none, false
	}

	return sv.Index(index).Interface(), true
}

// FindLast iterates over elements of slice from end to begin, returning the first one that passes a truth test on function.
// The function signature should be func(index int, value interface{}) bool .
func FindLast(slice, function interface{}) (interface{}, bool) {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, reflect.ValueOf(true).Type()) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + reflect.ValueOf(true).Type().String())
	}

	index := -1
	for i := sv.Len() - 1; i >= 0; i-- {
		flag := fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})[0]
		if flag.Bool() {
			index = i
			break
		}
	}

	if index == -1 {
		var none interface{}
		return none, false
	}

	return sv.Index(index).Interface(), true
}

// FlattenDeep flattens slice recursive
func FlattenDeep(slice interface{}) interface{} {
	sv := sliceValue(slice)
	st := sliceElemType(sv.Type())
	tmp := reflect.MakeSlice(reflect.SliceOf(st), 0, 0)
	res := flattenRecursive(sv, tmp)
	return res.Interface()
}

func flattenRecursive(value reflect.Value, result reflect.Value) reflect.Value {
	for i := 0; i < value.Len(); i++ {
		item := value.Index(i)
		kind := item.Kind()

		if kind == reflect.Slice {
			result = flattenRecursive(item, result)
		} else {
			result = reflect.Append(result, item)
		}
	}

	return result
}

// ForEach iterates over elements of slice and invokes function for each element
// The function signature should be func(index int, value interface{}).
func ForEach(slice, function interface{}) {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, nil) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + elemType.String())
	}

	for i := 0; i < sv.Len(); i++ {
		fn.Call([]reflect.Value{reflect.ValueOf(i), sv.Index(i)})
	}
}

// Map creates an slice of values by running each element of `slice` thru `function`.
// The function signature should be func(index int, value interface{}) interface{}.
func Map(slice, function interface{}) interface{} {
	sv := sliceValue(slice)
	fn := functionValue(function)

	elemType := sv.Type().Elem()
	if checkSliceCallbackFuncSignature(fn, elemType, nil) {
		panic("function param should be of type func(int, " + elemType.String() + ")" + elemType.String())
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
		panic("function param should be of type func(int, " + t + ", " + t + ")" + t)
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
	v := sliceValue(slice)

	out := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		v, ok := v.Index(i).Interface().(string)
		if !ok {
			panic("invalid element type")
		}
		out[i] = v
	}

	return out
}

// IntSlice convert param to slice of int.
func IntSlice(slice interface{}) []int {
	sv := sliceValue(slice)

	out := make([]int, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		v, ok := sv.Index(i).Interface().(int)
		if !ok {
			panic("invalid element type")
		}
		out[i] = v
	}

	return out
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

// Drop creates a slice with `n` elements dropped from the beginning when n > 0, or `n` elements dropped from the ending when n < 0
func Drop(slice interface{}, n int) interface{} {
	sv := sliceValue(slice)

	if n == 0 {
		return slice
	}

	svLen := sv.Len()

	if math.Abs(float64(n)) >= float64(svLen) {
		return reflect.MakeSlice(sv.Type(), 0, 0).Interface()
	}

	if n > 0 {
		res := reflect.MakeSlice(sv.Type(), svLen-n, svLen-n)
		for i := 0; i < res.Len(); i++ {
			res.Index(i).Set(sv.Index(i + n))
		}

		return res.Interface()
	}

	res := reflect.MakeSlice(sv.Type(), svLen+n, svLen+n)
	for i := 0; i < res.Len(); i++ {
		res.Index(i).Set(sv.Index(i))
	}

	return res.Interface()
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

	var temp []interface{}

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		skip := true
		for j := range temp {
			if v == temp[j] {
				skip = false
				break
			}
		}
		if skip {
			temp = append(temp, v)
		}
	}

	res := reflect.MakeSlice(sv.Type(), len(temp), len(temp))
	for i := 0; i < len(temp); i++ {
		res.Index(i).Set(reflect.ValueOf(temp[i]))
	}
	return res.Interface()

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

// UniqueBy call iteratee func with every item of slice, then remove duplicated.
// The iteratee function signature should be func(value interface{}) interface{}.
func UniqueBy(slice, iteratee interface{}) interface{} {
	sv := sliceValue(slice)
	fn := functionValue(iteratee)

	// elemType := sv.Type().Elem()
	// if !checkCallbackFuncSignature2(fn, elemType, elemType) {
	// 	panic("iteratee function signature should be of type func(" + elemType.String() + ")" + elemType.String())
	// }

	if sv.Len() == 0 {
		return slice
	}

	res := reflect.MakeSlice(sv.Type(), sv.Len(), sv.Len())

	for i := 0; i < sv.Len(); i++ {
		val := fn.Call([]reflect.Value{sv.Index(i)})[0]
		res.Index(i).Set(val)
	}

	return Unique(res.Interface())
}

// Union creates a slice of unique values, in order, from all given slices. using == for equality comparisons.
func Union(slices ...interface{}) interface{} {
	if len(slices) == 0 {
		return nil
	}
	// append all slices, then unique it
	var allSlices []interface{}
	len := 0
	for i := range slices {
		sv := sliceValue(slices[i])
		len += sv.Len()
		for j := 0; j < sv.Len(); j++ {
			v := sv.Index(j).Interface()
			allSlices = append(allSlices, v)
		}
	}

	sv := sliceValue(slices[0])
	res := reflect.MakeSlice(sv.Type(), len, len)
	for i := 0; i < len; i++ {
		res.Index(i).Set(reflect.ValueOf(allSlices[i]))
	}

	return Unique(res.Interface())
}

// Intersection creates a slice of unique values that included by all slices.
func Intersection(slices ...interface{}) interface{} {
	if len(slices) == 0 {
		return nil
	}

	reduceFunc := func(index int, slice1, slice2 interface{}) interface{} {
		set := make([]interface{}, 0)
		hash := make(map[interface{}]bool)

		sv1 := reflect.ValueOf(slice1)
		for i := 0; i < sv1.Len(); i++ {
			v := sv1.Index(i).Interface()
			hash[v] = true
		}

		sv2 := reflect.ValueOf(slice2)
		for i := 0; i < sv2.Len(); i++ {
			el := sv2.Index(i).Interface()
			if _, found := hash[el]; found {
				set = append(set, el)
			}
		}
		res := reflect.MakeSlice(sv1.Type(), len(set), len(set))
		for i := 0; i < len(set); i++ {
			res.Index(i).Set(reflect.ValueOf(set[i]))
		}
		return res.Interface()
	}

	res := Reduce(slices, reduceFunc, nil)
	return Unique(res)
}

// ReverseSlice return slice of element order is reversed to the given slice
func ReverseSlice(slice interface{}) {
	sv := sliceValue(slice)
	swp := reflect.Swapper(sv.Interface())
	for i, j := 0, sv.Len()-1; i < j; i, j = i+1, j-1 {
		swp(i, j)
	}
}

// Shuffle creates an slice of shuffled values
func Shuffle(slice interface{}) interface{} {
	sv := sliceValue(slice)
	length := sv.Len()

	res := reflect.MakeSlice(sv.Type(), length, length)
	for i, v := range rand.Perm(length) {
		res.Index(i).Set(sv.Index(v))
	}

	return res.Interface()
}

// SortByField return sorted slice by field
// Slice element should be struct, field type should be int, uint, string, or bool
// default sortType is ascending (asc), if descending order, set sortType to desc
func SortByField(slice interface{}, field string, sortType ...string) error {
	sv := sliceValue(slice)
	t := sv.Type().Elem()

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
	var compare func(a, b reflect.Value) bool
	switch sf.Type.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Int() > b.Int() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.Int() < b.Int() }
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Uint() > b.Uint() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.Uint() < b.Uint() }
		}
	case reflect.Float32, reflect.Float64:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Float() > b.Float() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.Float() < b.Float() }
		}
	case reflect.String:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.String() > b.String() }
		} else {
			compare = func(a, b reflect.Value) bool { return a.String() < b.String() }
		}
	case reflect.Bool:
		if len(sortType) > 0 && sortType[0] == "desc" {
			compare = func(a, b reflect.Value) bool { return a.Bool() && !b.Bool() }
		} else {
			compare = func(a, b reflect.Value) bool { return !a.Bool() && b.Bool() }
		}
	default:
		return fmt.Errorf("field type %s not supported", sf.Type)
	}

	sort.Slice(slice, func(i, j int) bool {
		a := sv.Index(i)
		b := sv.Index(j)
		if t.Kind() == reflect.Ptr {
			a = a.Elem()
			b = b.Elem()
		}
		a = a.FieldByIndex(sf.Index)
		b = b.FieldByIndex(sf.Index)
		return compare(a, b)
	})

	return nil
}

// Without creates a slice excluding all given values
func Without(slice interface{}, values ...interface{}) interface{} {
	sv := sliceValue(slice)
	if sv.Len() == 0 {
		return slice
	}

	var indexes []int
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		if !Contain(values, v) {
			indexes = append(indexes, i)
		}
	}

	res := reflect.MakeSlice(sv.Type(), len(indexes), len(indexes))
	for i := range indexes {
		res.Index(i).Set(sv.Index(indexes[i]))
	}

	return res.Interface()
}

// IndexOf returns the index at which the first occurrence of a value is found in a slice or return -1
// if the value cannot be found.
func IndexOf(slice, value interface{}) int {
	sv := sliceValue(slice)

	for i := 0; i < sv.Len(); i++ {
		if reflect.DeepEqual(sv.Index(i).Interface(), value) {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index at which the last occurrence of a value is found in a slice or return -1
// if the value cannot be found.
func LastIndexOf(slice, value interface{}) int {
	sv := sliceValue(slice)

	for i := sv.Len() - 1; i > 0; i-- {
		if reflect.DeepEqual(sv.Index(i).Interface(), value) {
			return i
		}
	}

	return -1
}

// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package slice implements some functions to manipulate slice.
package slice

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
)

// Contain check if the value is in the slice or not
func Contain[T any](slice []T, value T) bool {
	for _, v := range slice {
		if reflect.DeepEqual(v, value) {
			return true
		}
	}

	return false
}

// ContainSubSlice check if the slice contain subslice or not
func ContainSubSlice[T any](slice, subslice []T) bool {
	for _, v := range subslice {
		if !Contain(slice, v) {
			return false
		}
	}

	return true
}

// Chunk creates an slice of elements split into groups the length of size.
func Chunk[T any](slice []T, size int) [][]T {
	var res [][]T

	if len(slice) == 0 || size <= 0 {
		return res
	}

	length := len(slice)
	if size == 1 || size >= length {
		for _, v := range slice {
			var tmp []T
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
func Compact[T any](slice []T) []T {
	res := make([]T, 0, 0)
	for _, v := range slice {
		if !reflect.DeepEqual(v, nil) &&
			!reflect.DeepEqual(v, false) &&
			!reflect.DeepEqual(v, "") &&
			!reflect.DeepEqual(v, 0) {
			res = append(res, v)
		}
	}

	return res
}

// Concat creates a new slice concatenating slice with any additional slices and/or values.
func Concat[T any](slice []T, values ...[]T) []T {
	res := append([]T{}, slice...)

	for _, v := range values {
		res = append(res, v...)
	}

	return res
}

// Difference creates an slice of whose element in slice but not in comparedSlice
func Difference[T comparable](slice, comparedSlice []T) []T {
	var res []T
	for _, v := range slice {
		if !Contain(comparedSlice, v) {
			res = append(res, v)
		}
	}

	return res
}

// DifferenceBy it accepts iteratee which is invoked for each element of slice
// and values to generate the criterion by which they're compared.
// like lodash.js differenceBy: https://lodash.com/docs/4.17.15#differenceBy,
func DifferenceBy[T any](slice []T, comparedSlice []T, iteratee func(index int, item T) T) []T {
	orginSliceAfterMap := Map(slice, iteratee)
	comparedSliceAfterMap := Map(comparedSlice, iteratee)

	res := make([]T, 0, 0)
	for i, v := range orginSliceAfterMap {
		if !Contain(comparedSliceAfterMap, v) {
			res = append(res, slice[i])
		}
	}

	return res
}

//DifferenceWith accepts comparator which is invoked to compare elements of slice to values. The order and references of result values are determined by the first slice. The comparator is invoked with two arguments: (arrVal, othVal).
func DifferenceWith[T any](slice []T, comparedSlice []T, comparator func(value, otherValue T) bool) []T {
	res := make([]T, 0, 0)

	getIndex := func(arr []T, item T, comparison func(v1, v2 T) bool) int {
		index := -1
		for i, v := range arr {
			if comparison(item, v) {
				index = i
				break
			}
		}
		return index
	}

	for i, v := range slice {
		index := getIndex(comparedSlice, v, comparator)
		if index == -1 {
			res = append(res, slice[i])
		}
	}

	return res
}

// Equal checks if two slices are equal: the same length and all elements' order and value are equal
func Equal[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

// EqualWith checks if two slices are equal with comparator func
func EqualWith[T, U any](slice1 []T, slice2 []U, comparator func(T, U) bool) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v1 := range slice1 {
		v2 := slice2[i]
		if !comparator(v1, v2) {
			return false
		}
	}

	return true
}

// Every return true if all of the values in the slice pass the predicate function.
func Every[T any](slice []T, predicate func(index int, item T) bool) bool {
	if predicate == nil {
		panic("predicate func is missing")
	}

	var currentLength int
	for i, v := range slice {
		if predicate(i, v) {
			currentLength++
		}
	}

	return currentLength == len(slice)
}

// None return true if all the values in the slice mismatch the criteria
func None[T any](slice []T, predicate func(index int, item T) bool) bool {
	if predicate == nil {
		panic("predicate func is missing")
	}

	var currentLength int
	for i, v := range slice {
		if !predicate(i, v) {
			currentLength++
		}
	}

	return currentLength == len(slice)
}

// Some return true if any of the values in the list pass the predicate function.
func Some[T any](slice []T, predicate func(index int, item T) bool) bool {
	if predicate == nil {
		panic("predicate func is missing")
	}

	for i, v := range slice {
		if predicate(i, v) {
			return true
		}
	}
	return false
}

// Filter iterates over elements of slice, returning an slice of all elements pass the predicate function
func Filter[T any](slice []T, predicate func(index int, item T) bool) []T {
	if predicate == nil {
		panic("predicate func is missing")
	}

	res := make([]T, 0, 0)
	for i, v := range slice {
		if predicate(i, v) {
			res = append(res, v)
		}
	}
	return res
}

// Count iterates over elements of slice, returns a count of all matched elements
func Count[T any](slice []T, predicate func(index int, item T) bool) int {
	if predicate == nil {
		panic("predicate func is missing")
	}

	if len(slice) == 0 {
		return 0
	}

	var count int
	for i, v := range slice {
		if predicate(i, v) {
			count++
		}
	}

	return count
}

// GroupBy iterate over elements of the slice, each element will be group by criteria, returns two slices
func GroupBy[T any](slice []T, groupFn func(index int, item T) bool) ([]T, []T) {
	if groupFn == nil {
		panic("groupFn func is missing")
	}

	if len(slice) == 0 {
		return make([]T, 0), make([]T, 0)
	}

	groupB := make([]T, 0)
	groupA := make([]T, 0)

	for i, v := range slice {
		ok := groupFn(i, v)
		if ok {
			groupA = append(groupA, v)
		} else {
			groupB = append(groupB, v)
		}
	}

	return groupA, groupB
}

// GroupWith return a map composed of keys generated from the results of running each element of slice thru iteratee.
func GroupWith[T any, U comparable](slice []T, iteratee func(T) U) map[U][]T {
	if iteratee == nil {
		panic("iteratee func is missing")
	}

	res := make(map[U][]T)

	for _, v := range slice {
		key := iteratee(v)
		if _, ok := res[key]; !ok {
			res[key] = []T{}
		}
		res[key] = append(res[key], v)
	}

	return res
}

// Find iterates over elements of slice, returning the first one that passes a truth test on predicate function.
// If return T is nil then no items matched the predicate func
func Find[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	if predicate == nil {
		panic("predicate func is missing")
	}

	if len(slice) == 0 {
		return nil, false
	}

	index := -1
	for i, v := range slice {
		if predicate(i, v) {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, false
	}

	return &slice[index], true
}

// FindLast iterates over elements of slice from end to begin, returning the first one that passes a truth test on predicate function.
// If return T is nil then no items matched the predicate func
func FindLast[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	if predicate == nil {
		panic("predicate func is missing")
	}

	if len(slice) == 0 {
		return nil, false
	}

	index := -1
	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(i, slice[i]) {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, false
	}

	return &slice[index], true
}

// FlattenDeep flattens slice recursive
func FlattenDeep(slice any) any {
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
func ForEach[T any](slice []T, iteratee func(index int, item T)) {
	if iteratee == nil {
		panic("iteratee func is missing")
	}

	for i, v := range slice {
		iteratee(i, v)
	}
}

// Map creates an slice of values by running each element of slice thru iteratee function.
func Map[T any, U any](slice []T, iteratee func(index int, item T) U) []U {
	if iteratee == nil {
		panic("iteratee func is missing")
	}

	res := make([]U, len(slice), cap(slice))
	for i, v := range slice {
		res[i] = iteratee(i, v)
	}

	return res
}

// Reduce creates an slice of values by running each element of slice thru iteratee function.
func Reduce[T any](slice []T, iteratee func(index int, item1, item2 T) T, initial T) T {
	if iteratee == nil {
		panic("iteratee func is missing")
	}

	if len(slice) == 0 {
		return initial
	}

	res := iteratee(0, initial, slice[0])

	tmp := make([]T, 2, 2)
	for i := 1; i < len(slice); i++ {
		tmp[0] = res
		tmp[1] = slice[i]
		res = iteratee(i, tmp[0], tmp[1])
	}

	return res
}

// InterfaceSlice convert param to slice of interface.
func InterfaceSlice(slice any) []any {
	sv := sliceValue(slice)
	if sv.IsNil() {
		return nil
	}

	res := make([]any, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		res[i] = sv.Index(i).Interface()
	}

	return res
}

// StringSlice convert param to slice of string.
func StringSlice(slice any) []string {
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
func IntSlice(slice any) []int {
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

// DeleteAt delete the element of slice from start index to end index - 1.
func DeleteAt[T any](slice []T, start int, end ...int) []T {
	size := len(slice)

	if start < 0 || start >= size {
		return slice
	}

	if len(end) > 0 {
		end := end[0]
		if end <= start {
			return slice
		}
		if end > size {
			end = size
		}

		slice = append(slice[:start], slice[end:]...)
		return slice
	}

	if start == size-1 {
		slice = append(slice[:start])
	} else {
		slice = append(slice[:start], slice[start+1:]...)
	}

	return slice
}

// Drop creates a slice with `n` elements dropped from the beginning when n > 0, or `n` elements dropped from the ending when n < 0
func Drop[T any](slice []T, n int) []T {
	size := len(slice)

	if size == 0 || n == 0 {
		return slice
	}

	if math.Abs(float64(n)) >= float64(size) {
		return []T{}
	}

	if n < 0 {
		return slice[0 : size+n]
	}

	return slice[n:size]
}

// InsertAt insert the value or other slice into slice at index.
func InsertAt[T any](slice []T, index int, value any) []T {
	size := len(slice)

	if index < 0 || index > size {
		return slice
	}

	// value is T
	if v, ok := value.(T); ok {
		slice = append(slice[:index], append([]T{v}, slice[index:]...)...)
		return slice
	}

	// value is []T
	if v, ok := value.([]T); ok {
		slice = append(slice[:index], append(v, slice[index:]...)...)
		return slice
	}

	return slice
}

// UpdateAt update the slice element at index.
func UpdateAt[T any](slice []T, index int, value T) []T {
	size := len(slice)

	if index < 0 || index >= size {
		return slice
	}
	slice = append(slice[:index], append([]T{value}, slice[index+1:]...)...)

	return slice
}

// Unique remove duplicate elements in slice.
func Unique[T any](slice []T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	// here no use map filter. if use it, the result slice element order is random, not same as origin slice
	var res []T
	for i := 0; i < len(slice); i++ {
		v := slice[i]
		skip := true
		for j := range res {
			if reflect.DeepEqual(v, res[j]) {
				skip = false
				break
			}
		}
		if skip {
			res = append(res, v)
		}
	}

	return res
}

// UniqueBy call iteratee func with every item of slice, then remove duplicated.
func UniqueBy[T any](slice []T, iteratee func(item T) T) []T {
	if len(slice) == 0 {
		return []T{}
	}

	var res []T
	for _, v := range slice {
		val := iteratee(v)
		res = append(res, val)
	}

	return Unique(res)
}

// Union creates a slice of unique values, in order, from all given slices. using == for equality comparisons.
func Union[T any](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}

	// append all slices, then unique it
	var allElements []T

	for _, slice := range slices {
		for _, v := range slice {
			allElements = append(allElements, v)
		}
	}

	return Unique(allElements)
}

// Intersection creates a slice of unique values that included by all slices.
func Intersection[T any](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return Unique(slices[0])
	}

	var res []T

	reducer := func(s1, s2 []T) []T {
		s := make([]T, 0, 0)
		for _, v := range s1 {
			if Contain(s2, v) {
				s = append(s, v)
			}
		}
		return s
	}

	res = reducer(slices[0], slices[1])

	reduceSlice := make([][]T, 2, 2)
	for i := 2; i < len(slices); i++ {
		reduceSlice[0] = res
		reduceSlice[1] = slices[i]
		res = reducer(reduceSlice[0], reduceSlice[1])
	}

	return Unique(res)
}

// SymmetricDifference oppoiste operation of intersection function
func SymmetricDifference[T any](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return Unique(slices[0])
	}

	res := make([]T, 0)

	intersectSlice := Intersection(slices...)

	for i := 0; i < len(slices); i++ {
		slice := slices[i]
		for _, v := range slice {
			if !Contain(intersectSlice, v) {
				res = append(res, v)
			}
		}

	}

	return Unique(res)
}

// Reverse return slice of element order is reversed to the given slice
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle creates an slice of shuffled values
func Shuffle[T any](slice []T) []T {

	res := make([]T, len(slice))
	for i, v := range rand.Perm(len(slice)) {
		res[i] = slice[v]
	}

	return res
}

// SortByField return sorted slice by field
// Slice element should be struct, field type should be int, uint, string, or bool
// default sortType is ascending (asc), if descending order, set sortType to desc
func SortByField(slice any, field string, sortType ...string) error {
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
func Without[T any](slice []T, values ...T) []T {
	if len(values) == 0 || len(slice) == 0 {
		return slice
	}

	out := make([]T, 0, len(slice))
	for _, v := range slice {
		if !Contain(values, v) {
			out = append(out, v)
		}
	}

	return out
}

// IndexOf returns the index at which the first occurrence of a value is found in a slice or return -1
// if the value cannot be found.
func IndexOf[T any](slice []T, value T) int {
	for i, v := range slice {
		if reflect.DeepEqual(v, value) {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the index at which the last occurrence of a value is found in a slice or return -1
// if the value cannot be found.
func LastIndexOf[T any](slice []T, value T) int {
	for i := len(slice) - 1; i > 0; i-- {
		if reflect.DeepEqual(value, slice[i]) {
			return i
		}
	}

	return -1
}

// ToSlicePointer returns a pointer to the slices of a variable parameter transformation
func ToSlicePointer[T any](value ...T) []*T {
	out := make([]*T, len(value))
	for i := range value {
		out[i] = &value[i]
	}
	return out
}

// ToSlice returns a slices of a variable parameter transformation
func ToSlice[T any](value ...T) []T {
	out := make([]T, len(value))
	for i := range value {
		out[i] = value[i]
	}
	return out
}

// Copyright 2021 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package slice implements some functions to manipulate slice.
package slice

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"golang.org/x/exp/constraints"
)

// Create a static variable to store the hash table.
// This variable has the same lifetime as the entire program and can be shared by functions that are called more than once.
var (
	memoryHashMap     = make(map[string]map[any]int)
	memoryHashCounter = make(map[string]int)
	muForMemoryHash   sync.RWMutex
)

// Contain check if the target value is in the slice or not.
// Play: https://go.dev/play/p/_454yEHcNjf
func Contain[T comparable](slice []T, target T) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}

	return false
}

// ContainBy returns true if predicate function return true.
// Play: https://go.dev/play/p/49tkHfX4GNc
func ContainBy[T any](slice []T, predicate func(item T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}

	return false
}

// ContainSubSlice check if the slice contain a given subslice or not.
// Play: https://go.dev/play/p/bcuQ3UT6Sev
func ContainSubSlice[T comparable](slice, subSlice []T) bool {
	if len(subSlice) == 0 {
		return true
	}
	if len(slice) == 0 {
		return false
	}

	elementMap := make(map[T]struct{}, len(slice))
	for _, item := range slice {
		elementMap[item] = struct{}{}
	}

	for _, item := range subSlice {
		if _, ok := elementMap[item]; !ok {
			return false
		}
	}
	return true
}

// Chunk creates a slice of elements split into groups the length of size.
// Play: https://go.dev/play/p/b4Pou5j2L_C
func Chunk[T any](slice []T, size int) [][]T {
	result := [][]T{}

	if len(slice) == 0 || size <= 0 {
		return result
	}

	for _, item := range slice {
		l := len(result)
		if l == 0 || len(result[l-1]) == size {
			result = append(result, []T{})
			l++
		}

		result[l-1] = append(result[l-1], item)
	}

	return result
}

// Compact creates a slice with all falsey values removed. The values false, nil, 0, and "" are falsey.
// Play: https://go.dev/play/p/pO5AnxEr3TK
func Compact[T comparable](slice []T) []T {
	var zero T

	result := make([]T, 0, len(slice))

	for _, v := range slice {
		if v != zero {
			result = append(result, v)
		}
	}
	return result[:len(result):len(result)]
}

// Concat creates a new slice concatenating slice with any additional slices.
// Play: https://go.dev/play/p/gPt-q7zr5mk
func Concat[T any](slices ...[]T) []T {
	totalLen := 0
	for _, v := range slices {
		totalLen += len(v)
		if totalLen < 0 {
			panic("len out of range")
		}
	}
	result := make([]T, 0, totalLen)

	for _, v := range slices {
		result = append(result, v...)
	}

	return result
}

// Difference creates a slice of whose element in slice but not in comparedSlice.
// Play: https://go.dev/play/p/VXvadzLzhDa
func Difference[T comparable](slice, comparedSlice []T) []T {
	result := []T{}

	for _, v := range slice {
		if !Contain(comparedSlice, v) {
			result = append(result, v)
		}
	}

	return result
}

// DifferenceBy it accepts iteratee which is invoked for each element of slice
// and values to generate the criterion by which they're compared.
// like lodash.js differenceBy: https://lodash.com/docs/4.17.15#differenceBy.
// Play: https://go.dev/play/p/DiivgwM5OnC
func DifferenceBy[T comparable](slice []T, comparedSlice []T, iteratee func(index int, item T) T) []T {
	orginSliceAfterMap := Map(slice, iteratee)
	comparedSliceAfterMap := Map(comparedSlice, iteratee)

	result := make([]T, 0)
	for i, v := range orginSliceAfterMap {
		if !Contain(comparedSliceAfterMap, v) {
			result = append(result, slice[i])
		}
	}

	return result
}

// DifferenceWith accepts comparator which is invoked to compare elements of slice to values.
// The order and references of result values are determined by the first slice.
// The comparator is invoked with two arguments: (arrVal, othVal).
// Play: https://go.dev/play/p/v2U2deugKuV
func DifferenceWith[T any](slice []T, comparedSlice []T, comparator func(item1, item2 T) bool) []T {
	result := make([]T, 0)

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
			result = append(result, slice[i])
		}
	}

	return result
}

// Equal checks if two slices are equal: the same length and all elements' order and value are equal.
// Play: https://go.dev/play/p/WcRQJ37ifPa
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

// EqualWith checks if two slices are equal with comparator func.
// Play: https://go.dev/play/p/b9iygtgsHI1
func EqualWith[T, U any](slice1 []T, slice2 []U, comparator func(T, U) bool) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, v := range slice1 {
		if !comparator(v, slice2[i]) {
			return false
		}
	}

	return true
}

// EqualUnordered checks if two slices are equal: the same length and all elements' value are equal (unordered).
// Play: todo
func EqualUnordered[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	seen := make(map[T]int)
	for _, v := range slice1 {
		seen[v]++
	}

	for _, v := range slice2 {
		if seen[v] == 0 {
			return false
		}
		seen[v]--
	}

	return true
}

// Every return true if all of the values in the slice pass the predicate function.
// Play: https://go.dev/play/p/R8U6Sl-j8cD
func Every[T any](slice []T, predicate func(index int, item T) bool) bool {
	for i, v := range slice {
		if !predicate(i, v) {
			return false
		}
	}

	return true
}

// None return true if all the values in the slice mismatch the criteria.
// Play: https://go.dev/play/p/KimdalUlC-T
func None[T any](slice []T, predicate func(index int, item T) bool) bool {
	l := 0
	for i, v := range slice {
		if !predicate(i, v) {
			l++
		}
	}

	return l == len(slice)
}

// Some return true if any of the values in the list pass the predicate function.
// Play: https://go.dev/play/p/4pO9Xf9NDGS
func Some[T any](slice []T, predicate func(index int, item T) bool) bool {
	for i, v := range slice {
		if predicate(i, v) {
			return true
		}
	}

	return false
}

// Filter iterates over elements of slice, returning an slice of all elements pass the predicate function.
// Play: https://go.dev/play/p/SdPna-7qK4T
func Filter[T any](slice []T, predicate func(index int, item T) bool) []T {
	result := make([]T, 0)

	for i, v := range slice {
		if predicate(i, v) {
			result = append(result, v)
		}
	}

	return result
}

// Count returns the number of occurrences of the given item in the slice.
// Play: https://go.dev/play/p/Mj4oiEnQvRJ
func Count[T comparable](slice []T, item T) int {
	count := 0

	for _, v := range slice {
		if item == v {
			count++
		}
	}

	return count
}

// CountBy iterates over elements of slice with predicate function, returns the number of all matched elements.
// Play: https://go.dev/play/p/tHOccTMDZCC
func CountBy[T any](slice []T, predicate func(index int, item T) bool) int {
	count := 0

	for i, v := range slice {
		if predicate(i, v) {
			count++
		}
	}

	return count
}

// GroupBy iterate over elements of the slice, each element will be group by criteria, returns two slices.
// Play: https://go.dev/play/p/QVkPxzPR0iA
func GroupBy[T any](slice []T, groupFn func(index int, item T) bool) ([]T, []T) {
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

// GroupWith return a map composed of keys generated from the resultults of running each element of slice thru iteratee.
// Play: https://go.dev/play/p/ApCvMNTLO8a
func GroupWith[T any, U comparable](slice []T, iteratee func(item T) U) map[U][]T {
	result := make(map[U][]T)

	for _, v := range slice {
		key := iteratee(v)
		if _, ok := result[key]; !ok {
			result[key] = []T{}
		}
		result[key] = append(result[key], v)
	}

	return result
}

// Find iterates over elements of slice, returning the first one that passes a truth test on predicate function.
// If return T is nil then no items matched the predicate func.
// Play: https://go.dev/play/p/CBKeBoHVLgq
// Deprecated
func Find[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	v, ok := FindBy(slice, predicate)
	return &v, ok
}

// FindLast iterates over elements of slice from end to begin,
// return the first one that passes a truth test on predicate function.
// If return T is nil then no items matched the predicate func.
// Play: https://go.dev/play/p/FFDPV_j7URd
// Deprecated
func FindLast[T any](slice []T, predicate func(index int, item T) bool) (*T, bool) {
	v, ok := FindLastBy(slice, predicate)
	return &v, ok
}

// FindBy iterates over elements of slice, returning the first one that passes a truth test on predicate function.
// If return T is nil or zero value then no items matched the predicate func.
// In contrast to Find or FindLast, its return value no longer requires dereferencing
// Play: https://go.dev/play/p/n1lysBYl-GB
func FindBy[T any](slice []T, predicate func(index int, item T) bool) (v T, ok bool) {
	index := -1

	for i, v := range slice {
		if predicate(i, v) {
			index = i
			break
		}
	}

	if index == -1 {
		return v, false
	}

	return slice[index], true
}

// FindLastBy iterates over elements of slice, returning the last one that passes a truth test on predicate function.
// If return T is nil or zero value then no items matched the predicate func.
// In contrast to Find or FindLast, its return value no longer requires dereferencing
// Play: https://go.dev/play/p/8iqomzyCl_s
func FindLastBy[T any](slice []T, predicate func(index int, item T) bool) (v T, ok bool) {
	index := -1

	for i := len(slice) - 1; i >= 0; i-- {
		if predicate(i, slice[i]) {
			index = i
			break
		}
	}

	if index == -1 {
		return v, false
	}

	return slice[index], true
}

// Flatten flattens slice with one level.
// Play: https://go.dev/play/p/hYa3cBEevtm
func Flatten(slice any) any {
	sv := sliceValue(slice)

	var result reflect.Value
	if sv.Type().Elem().Kind() == reflect.Interface {
		result = reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, sv.Len())
	} else if sv.Type().Elem().Kind() == reflect.Slice {
		result = reflect.MakeSlice(sv.Type().Elem(), 0, sv.Len())
	} else {
		return result
	}

	for i := 0; i < sv.Len(); i++ {
		item := reflect.ValueOf(sv.Index(i).Interface())
		if item.Kind() == reflect.Slice {
			for j := 0; j < item.Len(); j++ {
				result = reflect.Append(result, item.Index(j))
			}
		} else {
			result = reflect.Append(result, item)
		}
	}

	return result.Interface()
}

// FlattenDeep flattens slice recursive.
// Play: https://go.dev/play/p/yjYNHPyCFaF
func FlattenDeep(slice any) any {
	sv := sliceValue(slice)
	st := sliceElemType(sv.Type())

	tmp := reflect.MakeSlice(reflect.SliceOf(st), 0, 0)

	result := flattenRecursive(sv, tmp)

	return result.Interface()
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

// ForEach iterates over elements of slice and invokes function for each element.
// Play: https://go.dev/play/p/DrPaa4YsHRF
func ForEach[T any](slice []T, iteratee func(index int, item T)) {
	for i := 0; i < len(slice); i++ {
		iteratee(i, slice[i])
	}
}

// ForEachWithBreak iterates over elements of slice and invokes function for each element,
// when iteratee return false, will break the for each loop.
// Play: https://go.dev/play/p/qScs39f3D9W
func ForEachWithBreak[T any](slice []T, iteratee func(index int, item T) bool) {
	for i := 0; i < len(slice); i++ {
		if !iteratee(i, slice[i]) {
			break
		}
	}
}

// Map creates an slice of values by running each element of slice thru iteratee function.
// Play: https://go.dev/play/p/biaTefqPquw
func Map[T any, U any](slice []T, iteratee func(index int, item T) U) []U {
	result := make([]U, len(slice), cap(slice))

	for i := 0; i < len(slice); i++ {
		result[i] = iteratee(i, slice[i])
	}

	return result
}

// FilterMap returns a slice which apply both filtering and mapping to the given slice.
// iteratee callback function should returntwo values:
// 1, mapping result.
// 2, whether the result element should be included or not
// Play: https://go.dev/play/p/J94SZ_9MiIe
func FilterMap[T any, U any](slice []T, iteratee func(index int, item T) (U, bool)) []U {
	result := []U{}

	for i, v := range slice {
		if a, ok := iteratee(i, v); ok {
			result = append(result, a)
		}
	}

	return result
}

// FlatMap manipulates a slice and transforms and flattens it to a slice of another type.
// Play: https://go.dev/play/p/_QARWlWs1N_F
func FlatMap[T any, U any](slice []T, iteratee func(index int, item T) []U) []U {
	result := make([]U, 0, len(slice))

	for i, v := range slice {
		result = append(result, iteratee(i, v)...)
	}

	return result
}

// Reduce creates an slice of values by running each element of slice thru iteratee function.
// Play: https://go.dev/play/p/_RfXJJWIsIm
func Reduce[T any](slice []T, iteratee func(index int, item1, item2 T) T, initial T) T {
	accumulator := initial

	for i, v := range slice {
		accumulator = iteratee(i, v, accumulator)
	}

	return accumulator
}

// ReduceBy produces a value from slice by accumulating the result of each element as passed through the reducer function.
// Play: https://go.dev/play/p/YKDpLi7gtee
func ReduceBy[T any, U any](slice []T, initial U, reducer func(index int, item T, agg U) U) U {
	accumulator := initial

	for i, v := range slice {
		accumulator = reducer(i, v, accumulator)
	}

	return accumulator
}

// ReduceRight is like ReduceBy, but it iterates over elements of slice from right to left.
// Play: https://go.dev/play/p/qT9dZC03A1K
func ReduceRight[T any, U any](slice []T, initial U, reducer func(index int, item T, agg U) U) U {
	accumulator := initial

	for i := len(slice) - 1; i >= 0; i-- {
		accumulator = reducer(i, slice[i], accumulator)
	}

	return accumulator
}

// Replace returns a copy of the slice with the first n non-overlapping instances of old replaced by new.
// Play: https://go.dev/play/p/P5mZp7IhOFo
func Replace[T comparable](slice []T, old T, new T, n int) []T {
	result := make([]T, len(slice))
	copy(result, slice)

	for i := range result {
		if result[i] == old && n != 0 {
			result[i] = new
			n--
		}
	}

	return result
}

// ReplaceAll returns a copy of the slice with all non-overlapping instances of old replaced by new.
// Play: https://go.dev/play/p/CzqXMsuYUrx
func ReplaceAll[T comparable](slice []T, old T, new T) []T {
	return Replace(slice, old, new, -1)
}

// Repeat creates a slice with length n whose elements are param `item`.
// Play: https://go.dev/play/p/1CbOmtgILUU
func Repeat[T any](item T, n int) []T {
	result := make([]T, n)

	for i := range result {
		result[i] = item
	}

	return result
}

// InterfaceSlice convert param to slice of interface.
// This function is deprecated, use generics feature of go1.18+ for replacement.
// Play: https://go.dev/play/p/FdQXF0Vvqs-
func InterfaceSlice(slice any) []any {
	sv := sliceValue(slice)
	if sv.IsNil() {
		return nil
	}

	result := make([]any, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		result[i] = sv.Index(i).Interface()
	}

	return result
}

// StringSlice convert param to slice of string.
// This function is deprecated, use generics feature of go1.18+ for replacement.
// Play: https://go.dev/play/p/W0TZDWCPFcI
func StringSlice(slice any) []string {
	v := sliceValue(slice)

	result := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		v, ok := v.Index(i).Interface().(string)
		if !ok {
			panic("invalid element type")
		}
		result[i] = v
	}

	return result
}

// IntSlice convert param to slice of int.
// This function is deprecated, use generics feature of go1.18+ for replacement.
// Play: https://go.dev/play/p/UQDj-on9TGN
func IntSlice(slice any) []int {
	sv := sliceValue(slice)

	result := make([]int, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		v, ok := sv.Index(i).Interface().(int)
		if !ok {
			panic("invalid element type")
		}
		result[i] = v
	}

	return result
}

// DeleteAt delete the element of slice at index.
// Play: https://go.dev/play/p/800B1dPBYyd
func DeleteAt[T any](slice []T, index int) []T {
	if index >= len(slice) {
		index = len(slice) - 1
	}

	result := make([]T, len(slice)-1)
	copy(result, slice[:index])
	copy(result[index:], slice[index+1:])

	return result
}

// DeleteRange delete the element of slice from start index to end index（exclude).
// Play: https://go.dev/play/p/945HwiNrnle
func DeleteRange[T any](slice []T, start, end int) []T {
	result := make([]T, 0, len(slice)-(end-start))

	for i := 0; i < start; i++ {
		result = append(result, slice[i])
	}

	for i := end; i < len(slice); i++ {
		result = append(result, slice[i])
	}

	return result
}

// Drop drop n elements from the start of a slice.
// Play: https://go.dev/play/p/jnPO2yQsT8H
func Drop[T any](slice []T, n int) []T {
	size := len(slice)

	if size <= n {
		return []T{}
	}

	if n <= 0 {
		return slice
	}

	result := make([]T, 0, size-n)

	return append(result, slice[n:]...)
}

// DropRight drop n elements from the end of a slice.
// Play: https://go.dev/play/p/8bcXvywZezG
func DropRight[T any](slice []T, n int) []T {
	size := len(slice)

	if size <= n {
		return []T{}
	}

	if n <= 0 {
		return slice
	}

	result := make([]T, 0, size-n)

	return append(result, slice[:size-n]...)
}

// DropWhile drop n elements from the start of a slice while predicate function returns true.
// Play: https://go.dev/play/p/4rt252UV_qs
func DropWhile[T any](slice []T, predicate func(item T) bool) []T {
	i := 0

	for ; i < len(slice); i++ {
		if !predicate(slice[i]) {
			break
		}
	}

	result := make([]T, 0, len(slice)-i)

	return append(result, slice[i:]...)
}

// DropRightWhile drop n elements from the end of a slice while predicate function returns true.
// Play: https://go.dev/play/p/6wyK3zMY56e
func DropRightWhile[T any](slice []T, predicate func(item T) bool) []T {
	i := len(slice) - 1

	for ; i >= 0; i-- {
		if !predicate(slice[i]) {
			break
		}
	}

	result := make([]T, 0, i+1)

	return append(result, slice[:i+1]...)
}

// InsertAt insert the value or other slice into slice at index.
// Play: https://go.dev/play/p/hMLNxPEGJVE
func InsertAt[T any](slice []T, index int, value any) []T {
	size := len(slice)

	if index < 0 || index > size {
		return slice
	}

	if v, ok := value.(T); ok {
		slice = append(slice[:index], append([]T{v}, slice[index:]...)...)
		return slice
	}

	if v, ok := value.([]T); ok {
		slice = append(slice[:index], append(v, slice[index:]...)...)
		return slice
	}

	return slice
}

// UpdateAt update the slice element at index.
// Play: https://go.dev/play/p/f3mh2KloWVm
func UpdateAt[T any](slice []T, index int, value T) []T {
	size := len(slice)

	if index < 0 || index >= size {
		return slice
	}
	slice = append(slice[:index], append([]T{value}, slice[index+1:]...)...)

	return slice
}

// Unique remove duplicate elements in slice.
// Play: https://go.dev/play/p/AXw0R3ZTE6a
func Unique[T comparable](slice []T) []T {
	result := make([]T, 0, len(slice))
	seen := make(map[T]struct{}, len(slice))

	for i := range slice {
		if _, ok := seen[slice[i]]; ok {
			continue
		}

		seen[slice[i]] = struct{}{}

		result = append(result, slice[i])
	}

	return result
}

// UniqueBy removes duplicate elements from the input slice based on the values returned by the iteratee function.
// The function maintains the order of the elements.
// Play: https://go.dev/play/p/GY7JE4yikrl
func UniqueBy[T any, U comparable](slice []T, iteratee func(item T) U) []T {
	result := make([]T, 0, len(slice))
	seen := make(map[U]struct{}, len(slice))

	for i := range slice {
		key := iteratee(slice[i])
		if _, ok := seen[key]; ok {
			continue
		}

		seen[key] = struct{}{}

		result = append(result, slice[i])
	}

	return result
}

// UniqueByComparator removes duplicate elements from the input slice using the provided comparator function.
// The function maintains the order of the elements.
// Play: https://go.dev/play/p/rwSacr-ZHsR
func UniqueByComparator[T comparable](slice []T, comparator func(item T, other T) bool) []T {
	result := make([]T, 0, len(slice))
	seen := make([]T, 0, len(slice))

	for _, item := range slice {
		duplicate := false
		for _, seenItem := range seen {
			if comparator(item, seenItem) {
				duplicate = true
				break
			}
		}
		if !duplicate {
			seen = append(seen, item)
			result = append(result, item)
		}
	}

	return result
}

// UniqueByField remove duplicate elements in struct slice by struct field.
// Play: https://go.dev/play/p/6cifcZSPIGu
func UniqueByField[T any](slice []T, field string) ([]T, error) {
	seen := map[any]struct{}{}

	var result []T
	for _, item := range slice {
		val, err := getField(item, field)
		if err != nil {
			return nil, fmt.Errorf("get field %s failed: %v", field, err)
		}
		if _, ok := seen[val]; !ok {
			seen[val] = struct{}{}
			result = append(result, item)
		}
	}

	return result, nil
}

func getField[T any](item T, field string) (interface{}, error) {
	v := reflect.ValueOf(item)
	t := reflect.TypeOf(item)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("data type %T not support, shuld be struct or pointer to struct", item)
	}

	f := v.FieldByName(field)
	if !f.IsValid() {
		return nil, fmt.Errorf("field name %s not found", field)
	}

	return v.FieldByName(field).Interface(), nil
}

// Union creates a slice of unique elements, in order, from all given slices.
// Play: https://go.dev/play/p/hfXV1iRIZOf
func Union[T comparable](slices ...[]T) []T {
	result := []T{}
	contain := map[T]struct{}{}

	for _, slice := range slices {
		for _, item := range slice {
			if _, ok := contain[item]; !ok {
				contain[item] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}

// UnionBy is like Union, what's more it accepts iteratee which is invoked for each element of each slice.
// Play: https://go.dev/play/p/HGKHfxKQsFi
func UnionBy[T any, V comparable](predicate func(item T) V, slices ...[]T) []T {
	result := []T{}
	contain := map[V]struct{}{}

	for _, slice := range slices {
		for _, item := range slice {
			val := predicate(item)
			if _, ok := contain[val]; !ok {
				contain[val] = struct{}{}
				result = append(result, item)
			}
		}
	}

	return result
}

// Deprecated:  Please use Concat() function instead.
// Merge all given slices into one slice.
// Play: https://go.dev/play/p/lbjFp784r9N
func Merge[T any](slices ...[]T) []T {
	return Concat(slices...)
}

// Intersection creates a slice of unique elements that included by all slices.
// Play: https://go.dev/play/p/anJXfB5wq_t
func Intersection[T comparable](slices ...[]T) []T {
	result := []T{}
	elementCount := make(map[T]int)

	for _, slice := range slices {
		seen := make(map[T]bool)

		for _, item := range slice {
			if !seen[item] {
				seen[item] = true
				elementCount[item]++
			}
		}
	}

	for _, item := range slices[0] {
		if elementCount[item] == len(slices) {
			result = append(result, item)
			elementCount[item] = 0
		}
	}

	return result
}

// SymmetricDifference oppoiste operation of intersection function.
// Play: https://go.dev/play/p/h42nJX5xMln
func SymmetricDifference[T comparable](slices ...[]T) []T {
	if len(slices) == 0 {
		return []T{}
	}
	if len(slices) == 1 {
		return Unique(slices[0])
	}

	result := make([]T, 0)

	intersectSlice := Intersection(slices...)

	for i := 0; i < len(slices); i++ {
		slice := slices[i]
		for _, v := range slice {
			if !Contain(intersectSlice, v) {
				result = append(result, v)
			}
		}

	}

	return Unique(result)
}

// Reverse return slice of element order is reversed to the given slice.
// Play: https://go.dev/play/p/8uI8f1lwNrQ
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shuffle the slice.
// Play: https://go.dev/play/p/YHvhnWGU3Ge
func Shuffle[T any](slice []T) []T {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return slice
}

// IsAscending checks if a slice is ascending order.
// Play: https://go.dev/play/p/9CtsFjet4SH
func IsAscending[T constraints.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}

	return true
}

// IsDescending checks if a slice is descending order.
// Play: https://go.dev/play/p/U_LljFXma14
func IsDescending[T constraints.Ordered](slice []T) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] < slice[i] {
			return false
		}
	}

	return true
}

// IsSorted checks if a slice is sorted(ascending or descending).
// Play: https://go.dev/play/p/nCE8wPLwSA-
func IsSorted[T constraints.Ordered](slice []T) bool {
	return IsAscending(slice) || IsDescending(slice)
}

// IsSortedByKey checks if a slice is sorted by iteratee function.
// Play: https://go.dev/play/p/tUoGB7DOHI4
func IsSortedByKey[T any, K constraints.Ordered](slice []T, iteratee func(item T) K) bool {
	size := len(slice)

	isAscending := func(data []T) bool {
		for i := 0; i < size-1; i++ {
			if iteratee(data[i]) > iteratee(data[i+1]) {
				return false
			}
		}

		return true
	}

	isDescending := func(data []T) bool {
		for i := 0; i < size-1; i++ {
			if iteratee(data[i]) < iteratee(data[i+1]) {
				return false
			}
		}

		return true
	}

	return isAscending(slice) || isDescending(slice)
}

// Sort sorts a slice of any ordered type(number or string), use quick sort algrithm.
// default sort order is ascending (asc), if want descending order, set param `sortOrder` to `desc`.
// Play: https://go.dev/play/p/V9AVjzf_4Fk
func Sort[T constraints.Ordered](slice []T, sortOrder ...string) {
	if len(sortOrder) > 0 && sortOrder[0] == "desc" {
		quickSort(slice, 0, len(slice)-1, "desc")
	} else {
		quickSort(slice, 0, len(slice)-1, "asc")
	}
}

// SortBy sorts the slice in ascending order as determined by the less function.
// This sort is not guaranteed to be stable.
// Play: https://go.dev/play/p/DAhLQSZEumm
func SortBy[T any](slice []T, less func(a, b T) bool) {
	quickSortBy(slice, 0, len(slice)-1, less)
}

// SortByField return sorted slice by field
// slice element should be struct, field type should be int, uint, string, or bool
// default sortType is ascending (asc), if descending order, set sortType to desc
// This function is deprecated, use Sort and SortBy for replacement.
// Play: https://go.dev/play/p/fU1prOBP9p1
func SortByField[T any](slice []T, field string, sortType ...string) error {
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

// Without creates a slice excluding all given items.
// Play: https://go.dev/play/p/bwhEXEypThg
func Without[T comparable](slice []T, items ...T) []T {
	if len(items) == 0 || len(slice) == 0 {
		return slice
	}

	result := make([]T, 0, len(slice))
	for _, v := range slice {
		if !Contain(items, v) {
			result = append(result, v)
		}
	}

	return result
}

// IndexOf returns the index at which the first occurrence of an item is found in a slice or return -1 if the item cannot be found.
// Play: https://go.dev/play/p/MRN1f0FpABb
func IndexOf[T comparable](arr []T, val T) int {
	limit := 10
	// gets the hash value of the array as the key of the hash table.
	key := fmt.Sprintf("%p", arr)

	muForMemoryHash.RLock()
	// determines whether the hash table is empty. If so, the hash table is created.
	if memoryHashMap[key] == nil {

		muForMemoryHash.RUnlock()
		muForMemoryHash.Lock()

		if memoryHashMap[key] == nil {
			memoryHashMap[key] = make(map[any]int)
			// iterate through the array, adding the value and index of each element to the hash table.
			for i := len(arr) - 1; i >= 0; i-- {
				memoryHashMap[key][arr[i]] = i
			}
		}

		muForMemoryHash.Unlock()
	} else {
		muForMemoryHash.RUnlock()
	}

	muForMemoryHash.Lock()
	// update the hash table counter.
	memoryHashCounter[key]++
	muForMemoryHash.Unlock()

	// use the hash table to find the specified value. If found, the index is returned.
	muForMemoryHash.RLock()
	index, ok := memoryHashMap[key][val]
	muForMemoryHash.RUnlock()

	if ok {
		muForMemoryHash.RLock()
		// calculate the memory usage of the hash table.
		size := len(memoryHashMap)
		muForMemoryHash.RUnlock()

		// If the memory usage of the hash table exceeds the memory limit, the hash table with the lowest counter is cleared.
		if size > limit {
			muForMemoryHash.Lock()
			var minKey string
			var minVal int
			for k, v := range memoryHashCounter {
				if k == key {
					continue
				}
				if minVal == 0 || v < minVal {
					minKey = k
					minVal = v
				}
			}
			delete(memoryHashMap, minKey)
			delete(memoryHashCounter, minKey)
			muForMemoryHash.Unlock()
		}
		return index
	}
	return -1
}

// LastIndexOf returns the index at which the last occurrence of the item is found in a slice or return -1 if the then cannot be found.
// Play: https://go.dev/play/p/DokM4cf1IKH
func LastIndexOf[T comparable](slice []T, item T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if item == slice[i] {
			return i
		}
	}

	return -1
}

// ToSlicePointer returns a pointer to the slices of a variable parameter transformation.
// Play: https://go.dev/play/p/gx4tr6_VXSF
func ToSlicePointer[T any](items ...T) []*T {
	result := make([]*T, len(items))
	for i := range items {
		result[i] = &items[i]
	}

	return result
}

// ToSlice returns a slices of a variable parameter transformation.
// Play: https://go.dev/play/p/YzbzVq5kscN
func ToSlice[T any](items ...T) []T {
	result := make([]T, len(items))
	copy(result, items)

	return result
}

// AppendIfAbsent only absent append the item.
// Play: https://go.dev/play/p/KcC1QXQ-RkL
func AppendIfAbsent[T comparable](slice []T, item T) []T {
	if !Contain(slice, item) {
		slice = append(slice, item)
	}
	return slice
}

// SetToDefaultIf sets elements to their default value if they match the given predicate.
// It retains the positions of the elements in the slice.
// It returns slice of T and the count of modified slice items
// Play: https://go.dev/play/p/9AXGlPRC0-A
func SetToDefaultIf[T any](slice []T, predicate func(T) bool) ([]T, int) {
	var count int
	for i := 0; i < len(slice); i++ {
		if predicate(slice[i]) {
			var zeroValue T
			slice[i] = zeroValue
			count++
		}
	}
	return slice, count
}

// KeyBy converts a slice to a map based on a callback function.
// Play: https://go.dev/play/p/uXod2LWD1Kg
func KeyBy[T any, U comparable](slice []T, iteratee func(item T) U) map[U]T {
	result := make(map[U]T, len(slice))

	for _, v := range slice {
		k := iteratee(v)
		result[k] = v
	}

	return result
}

// Join the slice item with specify separator.
// Play: https://go.dev/play/p/huKzqwNDD7V
func Join[T any](slice []T, separator string) string {
	str := Map(slice, func(_ int, item T) string {
		return fmt.Sprint(item)
	})

	return strings.Join(str, separator)
}

// Partition all slice elements with the evaluation of the given predicate functions.
// Play: https://go.dev/play/p/lkQ3Ri2NQhV
func Partition[T any](slice []T, predicates ...func(item T) bool) [][]T {
	l := len(predicates)

	result := make([][]T, l+1)

	for _, item := range slice {
		processed := false

		for i, f := range predicates {
			if f == nil {
				panic("predicate function must not be nill")
			}

			if f(item) {
				result[i] = append(result[i], item)
				processed = true
				break
			}
		}

		if !processed {
			result[l] = append(result[l], item)
		}
	}

	return result
}

// Breaks a list into two parts at the point where the predicate for the first time is true.
// Play: https://go.dev/play/p/yLYcBTyeQIz
func Break[T any](values []T, predicate func(T) bool) ([]T, []T) {
	a := make([]T, 0)
	b := make([]T, 0)
	if len(values) == 0 {
		return a, b
	}
	matched := false
	for _, value := range values {

		if !matched && predicate(value) {
			matched = true
		}

		if matched {
			b = append(b, value)
		} else {
			a = append(a, value)
		}
	}
	return a, b
}

// Random get a random item of slice, return idx=-1 when slice is empty
// Play: https://go.dev/play/p/UzpGQptWppw
func Random[T any](slice []T) (val T, idx int) {
	if len(slice) == 0 {
		return val, -1
	}

	idx = random.RandInt(0, len(slice))
	return slice[idx], idx
}

// RightPadding adds padding to the right end of a slice.
// Play: https://go.dev/play/p/0_2rlLEMBXL
func RightPadding[T any](slice []T, paddingValue T, paddingLength int) []T {
	if paddingLength == 0 {
		return slice
	}
	for i := 0; i < paddingLength; i++ {
		slice = append(slice, paddingValue)
	}
	return slice
}

// LeftPadding adds padding to the left begin of a slice.
// Play: https://go.dev/play/p/jlQVoelLl2k
func LeftPadding[T any](slice []T, paddingValue T, paddingLength int) []T {
	if paddingLength == 0 {
		return slice
	}

	paddedSlice := make([]T, len(slice)+paddingLength)
	i := 0
	for ; i < paddingLength; i++ {
		paddedSlice[i] = paddingValue
	}
	for j := 0; j < len(slice); j++ {
		paddedSlice[i] = slice[j]
		i++
	}

	return paddedSlice
}

// Frequency counts the frequency of each element in the slice.
// Play: https://go.dev/play/p/CW3UVNdUZOq
func Frequency[T comparable](slice []T) map[T]int {
	result := make(map[T]int)

	for _, v := range slice {
		result[v]++
	}

	return result
}

// JoinFunc joins the slice elements into a single string with the given separator.
// Play: https://go.dev/play/p/55ib3SB5fM2
func JoinFunc[T any](slice []T, sep string, transform func(T) T) string {
	var buf strings.Builder
	for i, v := range slice {
		if i > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(fmt.Sprint(transform(v)))
	}
	return buf.String()
}

// ConcatBy concats the elements of a slice into a single value using the provided separator and connector function.
// Play: https://go.dev/play/p/6QcUpcY4UMW
func ConcatBy[T any](slice []T, sep T, connector func(T, T) T) T {
	var result T

	if len(slice) == 0 {
		return result
	}

	for i, v := range slice {
		result = connector(result, v)
		if i < len(slice)-1 {
			result = connector(result, sep)
		}
	}

	return result
}

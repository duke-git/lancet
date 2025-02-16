// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package maputil includes some functions to manipulate map.
package maputil

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"

	"github.com/duke-git/lancet/v2/slice"
)

// Keys returns a slice of the map's keys.
// Play: https://go.dev/play/p/xNB5bTb97Wd
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	var i int
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// Values returns a slice of the map's values.
// Play: https://go.dev/play/p/CBKdUc5FTW6
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	var i int
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}

// KeysBy creates a slice whose element is the result of function mapper invoked by every map's key.
// Play: https://go.dev/play/p/hI371iB8Up8
func KeysBy[K comparable, V any, T any](m map[K]V, mapper func(item K) T) []T {
	keys := make([]T, 0, len(m))

	for k := range m {
		keys = append(keys, mapper(k))
	}

	return keys
}

// ValuesBy creates a slice whose element is the result of function mapper invoked by every map's value.
// Play: https://go.dev/play/p/sg9-oRidh8f
func ValuesBy[K comparable, V any, T any](m map[K]V, mapper func(item V) T) []T {
	keys := make([]T, 0, len(m))

	for _, v := range m {
		keys = append(keys, mapper(v))
	}

	return keys
}

// Merge maps, next key will overwrite previous key.
// Play: https://go.dev/play/p/H95LENF1uB-
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	size := 0
	for i := range maps {
		size += len(maps[i])
	}

	result := make(map[K]V, size)

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

// ForEach executes iteratee funcation for every key and value pair in map.
// Play: https://go.dev/play/p/OaThj6iNVXK
func ForEach[K comparable, V any](m map[K]V, iteratee func(key K, value V)) {
	for k, v := range m {
		iteratee(k, v)
	}
}

// Filter iterates over map, return a new map contains all key and value pairs pass the predicate function.
// Play: https://go.dev/play/p/fSvF3wxuNG7
func Filter[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// FilterByKeys iterates over map, return a new map whose keys are all given keys.
// Play: https://go.dev/play/p/7ov6BJHbVqh
func FilterByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if slice.Contain(keys, k) {
			result[k] = v
		}
	}
	return result
}

// FilterByValues iterates over map, return a new map whose values are all given values.
// Play: https://go.dev/play/p/P3-9MdcXegR
func FilterByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if slice.Contain(values, v) {
			result[k] = v
		}
	}
	return result
}

// OmitBy is the opposite of Filter, removes all the map elements for which the predicate function returns true.
// Play: https://go.dev/play/p/YJM4Hj5hNwm
func OmitBy[K comparable, V any](m map[K]V, predicate func(key K, value V) bool) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if !predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// OmitByKeys the opposite of FilterByKeys, extracts all the map elements which keys are not omitted.
// Play: https://go.dev/play/p/jXGrWDBfSRp
func OmitByKeys[K comparable, V any](m map[K]V, keys []K) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if !slice.Contain(keys, k) {
			result[k] = v
		}
	}
	return result
}

// OmitByValues the opposite of FilterByValues. remov all elements whose value are in the give slice.
// Play: https://go.dev/play/p/XB7Y10uw20_U
func OmitByValues[K comparable, V comparable](m map[K]V, values []V) map[K]V {
	result := make(map[K]V)

	for k, v := range m {
		if !slice.Contain(values, v) {
			result[k] = v
		}
	}
	return result
}

// Intersect iterates over maps, return a new map of key and value pairs in all given maps.
// Play: https://go.dev/play/p/Zld0oj3sjcC
func Intersect[K comparable, V any](maps ...map[K]V) map[K]V {
	if len(maps) == 0 {
		return map[K]V{}
	}
	if len(maps) == 1 {
		return maps[0]
	}

	var result map[K]V

	reducer := func(m1, m2 map[K]V) map[K]V {
		m := make(map[K]V)
		for k, v1 := range m1 {
			if v2, ok := m2[k]; ok && reflect.DeepEqual(v1, v2) {
				m[k] = v1
			}
		}
		return m
	}

	reduceMaps := make([]map[K]V, 2)
	result = reducer(maps[0], maps[1])

	for i := 2; i < len(maps); i++ {
		reduceMaps[0] = result
		reduceMaps[1] = maps[i]
		result = reducer(reduceMaps[0], reduceMaps[1])
	}

	return result
}

// Minus creates a map of whose key in mapA but not in mapB.
// Play: https://go.dev/play/p/3u5U9K7YZ9m
func Minus[K comparable, V any](mapA, mapB map[K]V) map[K]V {
	result := make(map[K]V)

	for k, v := range mapA {
		if _, ok := mapB[k]; !ok {
			result[k] = v
		}
	}
	return result
}

// IsDisjoint two map are disjoint if they have no keys in common.
// Play: https://go.dev/play/p/N9qgYg_Ho6f
func IsDisjoint[K comparable, V any](mapA, mapB map[K]V) bool {
	for k := range mapA {
		if _, ok := mapB[k]; ok {
			return false
		}
	}
	return true
}

// Entry is a key/value pairs.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Entries transforms a map into array of key/value pairs.
// Play: https://go.dev/play/p/Ltb11LNcElY
func Entries[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(m))

	for k, v := range m {
		entries = append(entries, Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}

	return entries
}

// FromEntries creates a map based on a slice of key/value pairs
// Play: https://go.dev/play/p/fTdu4sCNjQO
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	result := make(map[K]V, len(entries))

	for _, v := range entries {
		result[v.Key] = v.Value
	}

	return result
}

// Transform a map to another type map.
// Play: https://go.dev/play/p/P6ovfToM3zj
func Transform[K1 comparable, V1 any, K2 comparable, V2 any](m map[K1]V1, iteratee func(key K1, value V1) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(m))

	for k1, v1 := range m {
		k2, v2 := iteratee(k1, v1)
		result[k2] = v2
	}

	return result
}

// MapKeys transforms a map to other type map by manipulating it's keys.
// Play: https://go.dev/play/p/8scDxWeBDKd
func MapKeys[K comparable, V any, T comparable](m map[K]V, iteratee func(key K, value V) T) map[T]V {
	result := make(map[T]V, len(m))

	for k, v := range m {
		result[iteratee(k, v)] = v
	}

	return result
}

// MapValues transforms a map to other type map by manipulating it's values.
// Play: https://go.dev/play/p/g92aY3fc7Iw
func MapValues[K comparable, V any, T any](m map[K]V, iteratee func(key K, value V) T) map[K]T {
	result := make(map[K]T, len(m))

	for k, v := range m {
		result[k] = iteratee(k, v)
	}

	return result
}

// HasKey checks if map has key or not.
// This function is used to replace the following boilerplate code:
// _, haskey := amap["baz"];
//
//	if haskey {
//	   fmt.Println("map has key baz")
//	}
//
// Play: https://go.dev/play/p/isZZHOsDhFc
func HasKey[K comparable, V any](m map[K]V, key K) bool {
	_, haskey := m[key]
	return haskey
}

// MapToStruct converts map to struct
// Play: https://go.dev/play/p/7wYyVfX38Dp
func MapToStruct(m map[string]any, structObj any) error {
	for k, v := range m {
		err := setStructField(structObj, k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func setStructField(structObj any, fieldName string, fieldValue any) error {
	structVal := reflect.ValueOf(structObj).Elem()

	fName := getFieldNameByJsonTag(structObj, fieldName)
	if fName == "" {
		return fmt.Errorf("Struct field json tag don't match map key : %s in obj", fieldName)
	}

	fieldVal := structVal.FieldByName(fName)

	if !fieldVal.IsValid() {
		return fmt.Errorf("No such field: %s in obj", fieldName)
	}

	if !fieldVal.CanSet() {
		return fmt.Errorf("Cannot set %s field value", fieldName)
	}

	val := reflect.ValueOf(fieldValue)

	if fieldVal.Type() != val.Type() {

		if val.CanConvert(fieldVal.Type()) {
			fieldVal.Set(val.Convert(fieldVal.Type()))
			return nil
		}

		if m, ok := fieldValue.(map[string]any); ok {

			if fieldVal.Kind() == reflect.Struct {
				return MapToStruct(m, fieldVal.Addr().Interface())
			}

			if fieldVal.Kind() == reflect.Ptr && fieldVal.Type().Elem().Kind() == reflect.Struct {
				if fieldVal.IsNil() {
					fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
				}

				return MapToStruct(m, fieldVal.Interface())
			}

		}

		return fmt.Errorf("Map value type don't match struct field type")
	}

	fieldVal.Set(val)

	return nil
}

func getFieldNameByJsonTag(structObj any, jsonTag string) string {
	s := reflect.TypeOf(structObj).Elem()

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		tag := field.Tag
		name, _, _ := strings.Cut(tag.Get("json"), ",")
		if name == jsonTag {
			return field.Name
		}
	}

	return ""
}

// ToSortedSlicesDefault converts a map to two slices sorted by key: one for the keys and another for the values.
// Play: https://go.dev/play/p/43gEM2po-qy
func ToSortedSlicesDefault[K constraints.Ordered, V any](m map[K]V) ([]K, []V) {
	keys := make([]K, 0, len(m))

	// store the map’s keys into a slice
	for k := range m {
		keys = append(keys, k)
	}

	// sort the slice of keys
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// adjust the order of values according to the sorted keys
	sortedValues := make([]V, len(keys))
	for i, k := range keys {
		sortedValues[i] = m[k]
	}

	return keys, sortedValues
}

// ToSortedSlicesWithComparator converts a map to two slices sorted by key and using a custom comparison function:
// one for the keys and another for the values.
// Play: https://go.dev/play/p/0nlPo6YLdt3
func ToSortedSlicesWithComparator[K comparable, V any](m map[K]V, comparator func(a, b K) bool) ([]K, []V) {
	keys := make([]K, 0, len(m))

	// store the map’s keys into a slice
	for k := range m {
		keys = append(keys, k)
	}

	// sort the key slice using the provided comparison function
	sort.Slice(keys, func(i, j int) bool {
		return comparator(keys[i], keys[j])
	})

	// adjust the order of values according to the sorted keys
	sortedValues := make([]V, len(keys))
	for i, k := range keys {
		sortedValues[i] = m[k]
	}

	return keys, sortedValues
}

// GetOrSet returns value of the given key or set the given value value if not present.
// Play: https://go.dev/play/p/IVQwO1OkEJC
func GetOrSet[K comparable, V any](m map[K]V, key K, value V) V {
	if v, ok := m[key]; ok {
		return v
	}

	m[key] = value

	return value
}

// SortByKey sorts the map by its keys and returns a new map with sorted keys.
// Play: https://go.dev/play/p/PVdmBSnm6P_W
func SortByKey[K constraints.Ordered, V any](m map[K]V, less func(a, b K) bool) (sortedKeysMap map[K]V) {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return less(keys[i], keys[j])
	})

	sortedKeysMap = make(map[K]V, len(m))
	for _, k := range keys {
		sortedKeysMap[k] = m[k]
	}

	return
}

var mapHandlers = map[reflect.Kind]func(reflect.Value, reflect.Value) error{
	reflect.String:     convertNormal,
	reflect.Int:        convertNormal,
	reflect.Int16:      convertNormal,
	reflect.Int32:      convertNormal,
	reflect.Int64:      convertNormal,
	reflect.Uint:       convertNormal,
	reflect.Uint16:     convertNormal,
	reflect.Uint32:     convertNormal,
	reflect.Uint64:     convertNormal,
	reflect.Float32:    convertNormal,
	reflect.Float64:    convertNormal,
	reflect.Uint8:      convertNormal,
	reflect.Int8:       convertNormal,
	reflect.Struct:     convertNormal,
	reflect.Complex64:  convertNormal,
	reflect.Complex128: convertNormal,
}

var _ = func() struct{} {
	mapHandlers[reflect.Map] = convertMap
	mapHandlers[reflect.Array] = convertSlice
	mapHandlers[reflect.Slice] = convertSlice
	return struct{}{}
}()

// MapTo try to map any interface to struct or base type
/*
	Eg:
		v := map[string]interface{}{
			"service":map[string]interface{}{
				"ip":"127.0.0.1",
				"port":1234,
			},
			version:"v1.0.01"
		}

		type Target struct {
			Service struct {
				Ip string `json:"ip"`
				Port int `json:"port"`
			} `json:"service"`
			Ver string `json:"version"`
		}

		var dist Target
		if err := maputil.MapTo(v,&dist); err != nil {
			log.Println(err)
			return
		}

		log.Println(dist)

*/
// Play: https://go.dev/play/p/4K7KBEPgS5M
func MapTo(src any, dst any) error {
	dstRef := reflect.ValueOf(dst)

	if dstRef.Kind() != reflect.Ptr {
		return fmt.Errorf("dst is not ptr")
	}

	dstElem := dstRef.Type().Elem()
	if dstElem.Kind() == reflect.Struct {
		srcMap := src.(map[string]interface{})
		return MapToStruct(srcMap, dst)
	}

	dstRef = reflect.Indirect(dstRef)

	srcRef := reflect.ValueOf(src)
	if srcRef.Kind() == reflect.Ptr || srcRef.Kind() == reflect.Interface {
		srcRef = srcRef.Elem()
	}

	if f, ok := mapHandlers[srcRef.Kind()]; ok {
		return f(srcRef, dstRef)
	}

	return fmt.Errorf("no implemented:%s", srcRef.Type())
}

func convertNormal(src reflect.Value, dst reflect.Value) error {
	if dst.CanSet() {
		if src.Type() == dst.Type() {
			dst.Set(src)
		} else if src.CanConvert(dst.Type()) {
			dst.Set(src.Convert(dst.Type()))
		} else {
			return fmt.Errorf("can not convert:%s:%s", src.Type().String(), dst.Type().String())
		}
	}
	return nil
}

func convertSlice(src reflect.Value, dst reflect.Value) error {
	if dst.Kind() != reflect.Array && dst.Kind() != reflect.Slice {
		return fmt.Errorf("error type:%s", dst.Type().String())
	}
	l := src.Len()
	target := reflect.MakeSlice(dst.Type(), l, l)
	if dst.CanSet() {
		dst.Set(target)
	}
	for i := 0; i < l; i++ {
		srcValue := src.Index(i)
		if srcValue.Kind() == reflect.Ptr || srcValue.Kind() == reflect.Interface {
			srcValue = srcValue.Elem()
		}
		if f, ok := mapHandlers[srcValue.Kind()]; ok {
			err := f(srcValue, dst.Index(i))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func convertMap(src reflect.Value, dst reflect.Value) error {
	if src.Kind() != reflect.Map || dst.Kind() != reflect.Struct {
		if src.Kind() == reflect.Interface && dst.IsValid() {
			return convertMap(src.Elem(), dst)
		} else {
			return fmt.Errorf("src or dst type error,%s,%s", src.Type().String(), dst.Type().String())
		}
	}
	dstType := dst.Type()
	num := dstType.NumField()

	exist := map[string]int{}

	for i := 0; i < num; i++ {
		k := dstType.Field(i).Tag.Get("json")
		if k == "" {
			k = dstType.Field(i).Name
		}
		if strings.Contains(k, ",") {
			taglist := strings.Split(k, ",")
			if taglist[0] == "" {
				k = dstType.Field(i).Name
			} else {
				k = taglist[0]

			}

		}
		exist[k] = i
	}

	keys := src.MapKeys()

	for _, key := range keys {
		if index, ok := exist[key.String()]; ok {
			v := dst.Field(index)

			if v.Kind() == reflect.Struct {
				err := convertMap(src.MapIndex(key), v)
				if err != nil {
					return err
				}
			} else {
				if v.CanSet() {
					if v.Type() == src.MapIndex(key).Elem().Type() {
						v.Set(src.MapIndex(key).Elem())
					} else if src.MapIndex(key).Elem().CanConvert(v.Type()) {
						v.Set(src.MapIndex(key).Elem().Convert(v.Type()))
					} else if f, ok := mapHandlers[src.MapIndex(key).Elem().Kind()]; ok && f != nil {
						err := f(src.MapIndex(key).Elem(), v)
						if err != nil {
							return err
						}
					} else {
						return fmt.Errorf("error type:d(%s)s(%s)", v.Type(), src.Type())
					}
				}
			}
		}
	}

	return nil
}

// GetOrDefault returns the value of the given key or a default value if the key is not present.
// Play: https://go.dev/play/p/99QjSYSBdiM
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}

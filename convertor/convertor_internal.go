// Copyright 2023 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package convertor implements some functions to convert data.
package convertor

import "reflect"

type cloner struct {
	ptrs map[reflect.Type]map[uintptr]reflect.Value
}

// clone return a duplicate of passed item.
func (c *cloner) clone(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Invalid:
		return reflect.ValueOf(nil)

	// bool
	case reflect.Bool:
		return reflect.ValueOf(v.Bool())

	//int
	case reflect.Int:
		return reflect.ValueOf(int(v.Int()))
	case reflect.Int8:
		return reflect.ValueOf(int8(v.Int()))
	case reflect.Int16:
		return reflect.ValueOf(int16(v.Int()))
	case reflect.Int32:
		return reflect.ValueOf(int32(v.Int()))
	case reflect.Int64:
		return reflect.ValueOf(v.Int())

	// uint
	case reflect.Uint:
		return reflect.ValueOf(uint(v.Uint()))
	case reflect.Uint8:
		return reflect.ValueOf(uint8(v.Uint()))
	case reflect.Uint16:
		return reflect.ValueOf(uint16(v.Uint()))
	case reflect.Uint32:
		return reflect.ValueOf(uint32(v.Uint()))
	case reflect.Uint64:
		return reflect.ValueOf(v.Uint())

	// float
	case reflect.Float32:
		return reflect.ValueOf(float32(v.Float()))
	case reflect.Float64:
		return reflect.ValueOf(v.Float())

	// complex
	case reflect.Complex64:
		return reflect.ValueOf(complex64(v.Complex()))
	case reflect.Complex128:
		return reflect.ValueOf(v.Complex())

	// string
	case reflect.String:
		return reflect.ValueOf(v.String())

	// array
	case reflect.Array, reflect.Slice:
		return c.cloneArray(v)

	// map
	case reflect.Map:
		return c.cloneMap(v)

	// Ptr
	case reflect.Ptr:
		return c.clonePtr(v)

	// struct
	case reflect.Struct:
		return c.cloneStruct(v)

	// func
	case reflect.Func:
		return v

	// interface
	case reflect.Interface:
		return c.clone(v.Elem())

	}

	return reflect.Zero(v.Type())
}

func (c *cloner) cloneArray(v reflect.Value) reflect.Value {
	if v.IsNil() {
		return reflect.Zero(v.Type())
	}

	arr := reflect.MakeSlice(v.Type(), v.Len(), v.Len())

	for i := 0; i < v.Len(); i++ {
		val := c.clone(v.Index(i))

		if val.IsValid() {
			continue
		}

		item := arr.Index(i)
		if !item.CanSet() {
			continue
		}

		item.Set(val.Convert(item.Type()))
	}

	return arr
}

func (c *cloner) cloneMap(v reflect.Value) reflect.Value {
	if v.IsNil() {
		return reflect.Zero(v.Type())
	}

	clonedMap := reflect.MakeMap(v.Type())

	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		clonedKey := c.clone(key)
		clonedValue := c.clone(value)

		if !isNillable(clonedKey) || !clonedKey.IsNil() {
			clonedKey = clonedKey.Convert(key.Type())
		}

		if (!isNillable(clonedValue) || !clonedValue.IsNil()) && clonedValue.IsValid() {
			clonedValue = clonedValue.Convert(value.Type())
		}

		if !clonedValue.IsValid() {
			clonedValue = reflect.Zero(clonedMap.Type().Elem())
		}

		clonedMap.SetMapIndex(clonedKey, clonedValue)
	}

	return clonedMap
}

func isNillable(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Interface, reflect.Ptr, reflect.Func:
		return true
	}
	return false
}

func (c *cloner) clonePtr(v reflect.Value) reflect.Value {
	if v.IsNil() {
		return reflect.Zero(v.Type())
	}

	var newVal reflect.Value

	if v.Elem().CanAddr() {
		ptrs, exists := c.ptrs[v.Type()]
		if exists {
			if newVal, exists := ptrs[v.Elem().UnsafeAddr()]; exists {
				return newVal
			}
		}
	}

	newVal = c.clone(v.Elem())

	if v.Elem().CanAddr() {
		ptrs, exists := c.ptrs[v.Type()]
		if exists {
			if newVal, exists := ptrs[v.Elem().UnsafeAddr()]; exists {
				return newVal
			}
		}
	}

	clonedPtr := reflect.New(newVal.Type())
	clonedPtr.Elem().Set(newVal)

	return clonedPtr
}

func (c *cloner) cloneStruct(v reflect.Value) reflect.Value {
	clonedStructPtr := reflect.New(v.Type())
	clonedStruct := clonedStructPtr.Elem()

	if v.CanAddr() {
		ptrs := c.ptrs[clonedStructPtr.Type()]
		if ptrs == nil {
			ptrs = make(map[uintptr]reflect.Value)
			c.ptrs[clonedStructPtr.Type()] = ptrs
		}
		ptrs[v.UnsafeAddr()] = clonedStructPtr
	}

	for i := 0; i < v.NumField(); i++ {
		newStructValue := clonedStruct.Field(i)
		if !newStructValue.CanSet() {
			continue
		}

		clonedVal := c.clone(v.Field(i))
		if !clonedVal.IsValid() {
			continue
		}

		newStructValue.Set(clonedVal.Convert(newStructValue.Type()))
	}

	return clonedStruct
}

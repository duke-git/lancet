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

	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		return v

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
	if v.Kind() == reflect.Slice && v.IsNil() {
		return reflect.Zero(v.Type())
	}

	var arr reflect.Value
	if v.Kind() == reflect.Array {
		arr = reflect.New(v.Type()).Elem()
	} else {
		arr = reflect.MakeSlice(v.Type(), v.Len(), v.Len())
	}

	for i := 0; i < v.Len(); i++ {
		val := c.clone(v.Index(i))

		if !val.IsValid() {
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

	clonedPtr := reflect.New(v.Type().Elem())
	if newVal.IsValid() {
		clonedPtr.Elem().Set(newVal)
	}

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

package slice

import (
	"fmt"
	"reflect"
)

// sliceValue return the reflect value of a slice
func sliceValue(slice interface{}) reflect.Value {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid slice type, value of type %T", slice))
	}
	return v
}

// functionValue return the reflect value of a function
func functionValue(function interface{}) reflect.Value {
	v := reflect.ValueOf(function)
	if v.Kind() != reflect.Func {
		panic(fmt.Sprintf("Invalid function type, value of type %T", function))
	}
	return v
}

// checkSliceCallbackFuncSignature Check func sign :  s :[]type1{} -> func(i int, data type1) type2
// see https://coolshell.cn/articles/21164.html#%E6%B3%9B%E5%9E%8BMap-Reduce
func checkSliceCallbackFuncSignature(fn reflect.Value, types ...reflect.Type) bool {
	//Check it is a function
	if fn.Kind() != reflect.Func {
		return false
	}
	// NumIn() - returns a function type's input parameter count.
	// NumOut() - returns a function type's output parameter count.
	if (fn.Type().NumIn() != len(types)-1) || (fn.Type().NumOut() != 1) {
		return false
	}
	// In() - returns the type of a function type's i'th input parameter.
	// first input param type should be int
	if fn.Type().In(0) != reflect.TypeOf(1) {
		return false
	}
	for i := 0; i < len(types)-1; i++ {
		if fn.Type().In(i) != types[i] {
			return false
		}
	}
	// Out() - returns the type of a function type's i'th output parameter.
	outType := types[len(types)-1]
	if outType != nil && fn.Type().Out(0) != outType {
		return false
	}
	return true
}

// sliceElemType get slice element type
func sliceElemType(reflectType reflect.Type) reflect.Type {
	for {
		if reflectType.Kind() != reflect.Slice {
			return reflectType
		}

		reflectType = reflectType.Elem()
	}
}

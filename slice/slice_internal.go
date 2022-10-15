package slice

import (
	"fmt"
	"reflect"
)

// sliceValue return the reflect value of a slice
func sliceValue(slice any) reflect.Value {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid slice type, value of type %T", slice))
	}
	return v
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

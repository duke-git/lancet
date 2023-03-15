package pointer

import "reflect"

// ExtractPointer returns the underlying value by the given interface type
func ExtractPointer(value any) any {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	if t.Kind() != reflect.Pointer {
		return value
	}
	return ExtractPointer(v.Elem().Interface())
}

package pointer

import "reflect"

func ExtractPointer(value any) any {
	t := reflect.TypeOf(value)
	v := reflect.ValueOf(value)

	if t.Kind() != reflect.Pointer {
		return value
	}
	return ExtractPointer(v.Elem().Interface())
}

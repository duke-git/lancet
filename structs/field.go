package structs

import (
	"reflect"

	"github.com/duke-git/lancet/v2/pointer"
)

// Field is abstract struct field for provide several high level functions
type Field struct {
	Struct
	field reflect.StructField
	tag   *Tag
}

func newField(v reflect.Value, f reflect.StructField, tagName string) *Field {
	tag := f.Tag.Get(tagName)
	field := &Field{
		field: f,
		tag:   newTag(tag),
	}
	field.rvalue = v
	field.rtype = v.Type()
	field.TagName = tagName
	return field
}

// Tag returns the value that the key in the tag string.
func (f *Field) Tag() *Tag {
	return f.tag
}

// Value returns the underlying value of the field.
func (f *Field) Value() any {
	return f.rvalue.Interface()
}

// IsEmbedded returns true if the given field is an embedded field.
func (f *Field) IsEmbedded() bool {
	return len(f.field.Index) > 1
}

// IsExported returns true if the given field is exported.
func (f *Field) IsExported() bool {
	return f.field.IsExported()
}

// IsZero returns true if the given field is zero value.
func (f *Field) IsZero() bool {
	z := reflect.Zero(f.rvalue.Type()).Interface()
	v := f.Value()
	return reflect.DeepEqual(z, v)
}

// IsNil returns true if the given field is nil value.
func (f *Field) IsNil() bool {
	v := f.Value()
	if v == nil || (reflect.ValueOf(v)).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil() {
		return true
	}

	return false
}

// Name returns the name of the given field
func (f *Field) Name() string {
	return f.field.Name
}

// Kind returns the field's kind
func (f *Field) Kind() reflect.Kind {
	return f.rvalue.Kind()
}

// IsSlice check if a struct field type is slice or not
func (f *Field) IsSlice() bool {
	k := f.rvalue.Kind()
	return k == reflect.Slice
}

// IsTargetType check if a struct field type is target type or not
func (f *Field) IsTargetType(targetType reflect.Kind) bool {
	return f.rvalue.Kind() == targetType
}

// mapValue covert field value to map
func (f *Field) mapValue(value any) any {
	val := pointer.ExtractPointer(value)
	v := reflect.ValueOf(val)
	var ret any

	switch v.Kind() {
	case reflect.Struct:
		s := New(val)
		s.TagName = f.TagName
		m, _ := s.ToMap()
		ret = m
	case reflect.Map:
		mapEl := v.Type().Elem()
		switch mapEl.Kind() {
		case reflect.Ptr, reflect.Array, reflect.Map, reflect.Slice, reflect.Chan:
			// iterate the map
			m := make(map[string]any, v.Len())
			for _, key := range v.MapKeys() {
				m[key.String()] = f.mapValue(v.MapIndex(key).Interface())
			}
			ret = m
		default:
			ret = v.Interface()
		}
	case reflect.Slice, reflect.Array:
		sEl := v.Type().Elem()
		switch sEl.Kind() {
		case reflect.Ptr, reflect.Array, reflect.Map, reflect.Slice, reflect.Chan:
			slices := make([]any, v.Len())
			for i := 0; i < v.Len(); i++ {
				slices[i] = f.mapValue(v.Index(i).Interface())
			}
			ret = slices
		default:
			ret = v.Interface()
		}
	default:
		if v.Kind().String() != "invalid" {
			ret = v.Interface()
		}
	}
	return ret
}

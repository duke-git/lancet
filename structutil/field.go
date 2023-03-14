package structutil

import (
	"github.com/duke-git/lancet/v2/pointer"
	"reflect"
)

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
	val := pointer.ExtractPointer(v.Interface())
	rt := reflect.TypeOf(val)
	rv := reflect.ValueOf(val)
	field.rvalue = rv
	field.rtype = rt
	field.TagName = tagName
	return field
}

// Tag returns the Value that the key in the tag string.
func (f *Field) Tag() *Tag {
	return f.tag
}

// Value returns the underlying Value of the field.
func (f *Field) Value() any {
	return f.rvalue.Interface()
}

// IsEmbedded returns true if the given field is an embedded field.
func (f *Field) IsEmbedded() bool {
	return f.field.Anonymous
}

// IsExported returns true if the given field is exported.
func (f *Field) IsExported() bool {
	return f.field.PkgPath == ""
}

// IsZero returns true if the given field is zero Value.
func (f *Field) IsZero() bool {
	z := reflect.Zero(f.rvalue.Type()).Interface()
	v := f.Value()
	return reflect.DeepEqual(z, v)
}

// Name returns the name of the given field
func (f *Field) Name() string {
	return f.field.Name
}

// Kind returns the field's kind
func (f *Field) Kind() reflect.Kind {
	return f.rvalue.Kind()
}

func (f *Field) IsSlice() bool {
	k := f.rvalue.Kind()
	if k == reflect.Invalid {
		return false
	}
	return k == reflect.Slice
}

func (f *Field) MapValue(value any) any {
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
				m[key.String()] = f.MapValue(v.MapIndex(key).Interface())
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
				slices[i] = f.MapValue(v.Index(i).Interface())
			}
			ret = slices
		default:
			ret = v.Interface()
		}
	default:
		ret = v.Interface()
	}
	return ret
}

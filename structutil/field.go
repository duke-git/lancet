package structutil

import "reflect"

type Field struct {
	value reflect.Value
	field reflect.StructField
	tag   *Tag
}

func newField(v reflect.Value, f reflect.StructField, tagName string) *Field {
	tag := f.Tag.Get(tagName)
	return &Field{
		value: v,
		field: f,
		tag:   newTag(tag),
	}
}

// Tag returns the value that the key in the tag string.
func (f *Field) Tag() *Tag {
	return f.tag
}

// Value returns the underlying value of the field.
func (f *Field) Value() any {
	return f.value.Interface()
}

// IsEmbedded returns true if the given field is an embedded field.
func (f *Field) IsEmbedded() bool {
	return f.field.Anonymous
}

// IsExported returns true if the given field is exported.
func (f *Field) IsExported() bool {
	return f.field.PkgPath == ""
}

// IsZero returns true if the given field is zero value.
func (f *Field) IsZero() bool {
	z := reflect.Zero(f.value.Type()).Interface()
	v := f.Value()
	return reflect.DeepEqual(z, v)
}

// Name returns the name of the given field
func (f *Field) Name() string {
	return f.field.Name
}

// Kind returns the field's kind
func (f *Field) Kind() reflect.Kind {
	return f.value.Kind()
}

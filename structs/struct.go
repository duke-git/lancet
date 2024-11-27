// Package structs provide several high level functions to manipulate struct, tag, and field.
package structs

import (
	"fmt"
	"reflect"

	"github.com/duke-git/lancet/v2/pointer"
)

// defaultTagName is the default tag for struct fields to lookup.
var defaultTagName = "json"

// Struct is abstract struct for provide several high level functions
type Struct struct {
	raw     any
	rtype   reflect.Type
	rvalue  reflect.Value
	TagName string
}

// New returns a new *Struct
func New(value any, tagName ...string) *Struct {
	value = pointer.ExtractPointer(value)
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)

	tn := defaultTagName

	if len(tagName) > 0 {
		tn = tagName[0]
	}

	// if need: can also set defaultTagName to tn across structs package level
	// defaultTagName = tn

	return &Struct{
		raw:     value,
		rtype:   t,
		rvalue:  v,
		TagName: tn,
	}
}

// ToMap converts the given struct to a map[string]any, where the keys
// of the keys are the field names and the values of the map are the values
// of the fields. The default map key is the struct field name, but you can
// change it. The `json` key is the default tag key. Example:
//
//		// default
//		Name string `json:"name"`
//
//		// ignore the field
//		Name string							// no tag
//	 	Age string `json:"-"`				// json ignore tag
//		sex string							// unexported field
//	 	Goal int   `json:"goal,omitempty"`  // omitempty if the field is zero value
//
//	 	// custom map key
//	 	Name string `json:"myName"`
//
// ToMap convert the exported fields of a struct to map.
func (s *Struct) ToMap() (map[string]any, error) {
	if !s.IsStruct() {
		return nil, fmt.Errorf("invalid struct %v", s)
	}

	result := make(map[string]any)
	fields := s.Fields()
	for _, f := range fields {
		if !f.IsExported() || f.tag.IsEmpty() || f.tag.Name == "-" {
			continue
		}

		if f.IsZero() && f.tag.HasOption("omitempty") {
			continue
		}

		if f.IsNil() {
			continue
		}

		result[f.tag.Name] = f.mapValue(f.Value())
	}

	return result, nil
}

// Fields returns all the struct fields within a slice
func (s *Struct) Fields() []*Field {
	fieldNum := s.rvalue.NumField()
	fields := make([]*Field, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		v := s.rvalue.Field(i)
		sf := s.rtype.Field(i)
		field := newField(v, sf, s.TagName)
		fields = append(fields, field)
	}
	return fields
}

// Field returns a Field if the given field name was found
func (s *Struct) Field(name string) (*Field, bool) {
	f, ok := s.rtype.FieldByName(name)
	if !ok {
		return nil, false
	}
	return newField(s.rvalue.FieldByName(name), f, s.TagName), true
}

// IsStruct returns true if the given rvalue is a struct
func (s *Struct) IsStruct() bool {
	k := s.rvalue.Kind()
	if k == reflect.Invalid {
		return false
	}
	return k == reflect.Struct
}

// ToMap convert struct to map, only convert exported struct field
// map key is specified same as struct field tag `json` value.
func ToMap(v any) (map[string]any, error) {
	return New(v).ToMap()
}

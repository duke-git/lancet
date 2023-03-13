package structutil

import (
	"reflect"
)

// DefaultTagName is the default tag for struct fields to lookup.
var DefaultTagName = "json"

// Struct is abstract struct for provide several high level functions
type Struct struct {
	raw     any
	rtype   reflect.Type
	rvalue  reflect.Value
	TagName string
}

// New returns a new *Struct
func New(value any) *Struct {
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return &Struct{
		raw:     value,
		rtype:   t,
		rvalue:  v,
		TagName: DefaultTagName,
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
// Only the exported fields of a struct can be converted.
func (s *Struct) ToMap() (map[string]any, error) {
	result := make(map[string]any)

	fields := s.Fields()
	for _, f := range fields {
		if !f.IsExported() || f.tag.IsEmpty() || f.tag.Name == "-" {
			continue
		}
		if f.IsZero() && f.tag.HasOption("omitempty") {
			continue
		}
		// TODO: sub struct
		result[f.tag.Name] = f.Value()
	}

	return result, nil
}

// Fields returns all the struct fields within a slice
func (s *Struct) Fields() []*Field {

	var fields []*Field
	fieldNum := s.rvalue.NumField()
	for i := 0; i < fieldNum; i++ {
		v := s.rvalue.Field(i)
		sf := s.rtype.Field(i)
		field := newField(v, sf, DefaultTagName)
		fields = append(fields, field)
	}

	return fields
}

// ToMap convert struct to map, only convert exported struct field
// map key is specified same as struct field tag `json` value.
func ToMap(v any) (map[string]any, error) {
	return New(v).ToMap()
}

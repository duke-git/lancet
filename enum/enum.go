// Copyright 2025 dudaodong@gmail.com. All rights resulterved.
// Use of this source code is governed by MIT license

// Package enum provides a simple enum implementation.
package enum

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"sync"
)

// Enum defines a common enum interface.
type Enum[T comparable] interface {
	Value() T
	String() string
	Name() string
	Valid() bool
}

// Item defines a common enum item struct implement Enum interface.
type Item[T comparable] struct {
	value T
	name  string
}

func NewItem[T comparable](value T, name string) *Item[T] {
	return &Item[T]{value: value, name: name}
}

// NewWithItems creates enum items from value-name pairs.
// It requires an even number of arguments, where each pair consists of a value of type T and a string name.
// If the number of arguments is odd or if any argument does not match the expected type, it will panic.
func NewItems[T comparable](args ...any) []*Item[T] {
	if len(args) == 0 {
		return []*Item[T]{}
	}
	if len(args)%2 != 0 {
		panic("New requires even number of arguments (value, name pairs)")
	}

	items := make([]*Item[T], 0, len(args)/2)

	for i := 0; i < len(args); i += 2 {
		value, ok := args[i].(T)
		if !ok {
			panic(fmt.Sprintf("argument %d is not of type %T", i, value))
		}

		name, ok := args[i+1].(string)
		if !ok {
			panic(fmt.Sprintf("argument %d is not of type string", i+1))
		}

		items = append(items, &Item[T]{value: value, name: name})
	}

	return items

}

func (e *Item[T]) Value() T {
	return e.value
}

func (e *Item[T]) Name() string {
	return e.name
}

func (e *Item[T]) String() string {
	return e.name
}

// Valid checks if the enum item is valid. If a custom check function is provided, it will be used to validate the value.
func (e *Item[T]) Valid(check ...func(T) bool) bool {
	if len(check) > 0 {
		return check[0](e.value) && e.name != ""
	}
	var zero T
	return e.value != zero && e.name != ""
}

// MarshalJSON implements the json.Marshaler interface.
func (e *Item[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"value": e.value,
		"name":  e.name,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *Item[T]) UnmarshalJSON(data []byte) error {
	type alias struct {
		Value any    `json:"value"`
		Name  string `json:"name"`
	}
	var temp alias
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	var v T
	rv := reflect.TypeOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, ok := temp.Value.(float64)
		if !ok {
			return fmt.Errorf("invalid type for value, want int family")
		}
		converted := reflect.ValueOf(int64(val)).Convert(rv)
		e.value = converted.Interface().(T)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, ok := temp.Value.(float64)
		if !ok {
			return fmt.Errorf("invalid type for value, want uint family")
		}
		converted := reflect.ValueOf(uint64(val)).Convert(rv)
		e.value = converted.Interface().(T)
	case reflect.Float32, reflect.Float64:
		val, ok := temp.Value.(float64)
		if !ok {
			return fmt.Errorf("invalid type for value, want float family")
		}
		converted := reflect.ValueOf(val).Convert(rv)
		e.value = converted.Interface().(T)
	case reflect.String:
		val, ok := temp.Value.(string)
		if !ok {
			return fmt.Errorf("invalid type for value, want string")
		}
		e.value = any(val).(T)
	case reflect.Bool:
		val, ok := temp.Value.(bool)
		if !ok {
			return fmt.Errorf("invalid type for value, want bool")
		}
		e.value = any(val).(T)
	default:
		// 枚举类型底层通常是 int，可以尝试 float64->int64->底层类型
		val, ok := temp.Value.(float64)
		if ok {
			converted := reflect.ValueOf(int64(val)).Convert(rv)
			e.value = converted.Interface().(T)
		} else {
			val2, ok2 := temp.Value.(T)
			if !ok2 {
				return fmt.Errorf("invalid type for value")
			}
			e.value = val2
		}
	}
	e.name = temp.Name
	return nil
}

// Registry defines a common enum registry struct.
type Registry[T comparable] struct {
	mu     sync.RWMutex
	values map[T]*Item[T]
	names  map[string]*Item[T]
	items  []*Item[T]
}

func NewRegistry[T comparable](items ...*Item[T]) *Registry[T] {
	r := &Registry[T]{
		values: make(map[T]*Item[T]),
		names:  make(map[string]*Item[T]),
		items:  make([]*Item[T], 0, len(items)),
	}

	r.Add(items...)

	return r
}

// Add adds enum items to the registry.
func (r *Registry[T]) Add(items ...*Item[T]) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, item := range items {
		if _, exists := r.values[item.value]; exists {
			continue
		}
		if _, exists := r.names[item.name]; exists {
			continue
		}

		r.values[item.value] = item
		r.names[item.name] = item
		r.items = append(r.items, item)
	}
}

// Remove removes an enum item from the registry by its value.
func (r *Registry[T]) Remove(value T) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	item, ok := r.values[value]
	if !ok {
		return false
	}
	delete(r.values, value)
	delete(r.names, item.name)
	for i, it := range r.items {
		if it.value == value {
			r.items = append(r.items[:i], r.items[i+1:]...)
			break
		}
	}
	return true
}

// Update updates the name of an enum item in the registry by its value.
func (r *Registry[T]) Update(value T, newName string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	item, ok := r.values[value]
	if !ok {
		return false
	}
	delete(r.names, item.name)
	item.name = newName
	r.names[newName] = item
	return true
}

// GetByValue retrieves an enum item by its value.
func (r *Registry[T]) GetByValue(value T) (*Item[T], bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.values[value]
	return item, ok
}

// GetByName retrieves an enum item by its name.
func (r *Registry[T]) GetByName(name string) (*Item[T], bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.names[name]
	return item, ok
}

// Items returns a slice of all enum items in the registry.
func (r *Registry[T]) Items() []*Item[T] {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*Item[T], len(r.items))
	copy(result, r.items)
	return result
}

// Contains checks if an enum item with the given value exists in the registry.
func (r *Registry[T]) Contains(value T) bool {
	_, ok := r.GetByValue(value)
	return ok
}

// Validate checks if the given value is a valid enum item in the registry.
func (r *Registry[T]) Validate(value T) error {
	if !r.Contains(value) {
		return fmt.Errorf("invalid enum value: %v", value)
	}
	return nil
}

// ValidateAll checks if all given values are valid enum items in the registry.
func (r *Registry[T]) ValidateAll(values ...T) error {
	for _, value := range values {
		if err := r.Validate(value); err != nil {
			return err
		}
	}
	return nil
}

// Size returns the number of enum items in the registry.
func (r *Registry[T]) Size() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.items)
}

// Range iterates over all enum items in the registry and applies the given function.
func (r *Registry[T]) Range(fn func(*Item[T]) bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, item := range r.items {
		if !fn(item) {
			break
		}
	}
}

// SortedItems returns a slice of all enum items sorted by the given less function.
func (r *Registry[T]) SortedItems(less func(*Item[T], *Item[T]) bool) []*Item[T] {
	items := r.Items()
	sort.Slice(items, func(i, j int) bool {
		return less(items[i], items[j])
	})
	return items
}

// Filter returns a slice of enum items that satisfy the given predicate function.
func (r *Registry[T]) Filter(predicate func(*Item[T]) bool) []*Item[T] {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*Item[T]
	for _, item := range r.items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

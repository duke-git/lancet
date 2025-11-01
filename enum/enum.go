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
	Valid(checker ...func(T) bool) bool
}

// Item defines a common enum item struct implement Enum interface.
type Item[T comparable] struct {
	value T
	name  string
}

// NewItem creates a new enum item.
// Play: https://go.dev/play/p/8qNsLw01HD5
func NewItem[T comparable](value T, name string) *Item[T] {
	return &Item[T]{value: value, name: name}
}

// Pair represents a value-name pair for creating enum items
type Pair[T comparable] struct {
	Value T
	Name  string
}

// NewItemsFromPairs creates enum items from a slice of Pair structs.
// Play: https://go.dev/play/p/xKnoGa7gnev
func NewItemsFromPairs[T comparable](pairs ...Pair[T]) []*Item[T] {
	if len(pairs) == 0 {
		return []*Item[T]{}
	}

	items := make([]*Item[T], 0, len(pairs))
	for _, pair := range pairs {
		items = append(items, &Item[T]{value: pair.Value, name: pair.Name})
	}

	return items
}

// Value returns the value of the enum item.
// Play: https://go.dev/play/p/xKnoGa7gnev
func (it *Item[T]) Value() T {
	return it.value
}

// Name returns the name of the enum item.
// Play: https://go.dev/play/p/xKnoGa7gnev
func (it *Item[T]) Name() string {
	return it.name
}

// String returns the string representation of the enum item.
func (it *Item[T]) String() string {
	return it.name
}

// Valid checks if the enum item is valid. If a custom check function is provided, it will be used to validate the value.
// Play: https://go.dev/play/p/pA3lYY2VSm3
func (it *Item[T]) Valid(checker ...func(T) bool) bool {
	if len(checker) > 0 {
		return checker[0](it.value) && it.name != ""
	}
	var zero T
	return it.value != zero && it.name != ""
}

// MarshalJSON implements the json.Marshaler interface.
// Play: https://go.dev/play/p/zIZEdAnneB5
func (it *Item[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"value": it.value,
		"name":  it.name,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Play: https://go.dev/play/p/zIZEdAnneB5
func (it *Item[T]) UnmarshalJSON(data []byte) error {
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
		it.value = converted.Interface().(T)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, ok := temp.Value.(float64)
		if !ok {
			return fmt.Errorf("invalid type for value, want uint family")
		}
		converted := reflect.ValueOf(uint64(val)).Convert(rv)
		it.value = converted.Interface().(T)
	case reflect.Float32, reflect.Float64:
		val, ok := temp.Value.(float64)
		if !ok {
			return fmt.Errorf("invalid type for value, want float family")
		}
		converted := reflect.ValueOf(val).Convert(rv)
		it.value = converted.Interface().(T)
	case reflect.String:
		val, ok := temp.Value.(string)
		if !ok {
			return fmt.Errorf("invalid type for value, want string")
		}
		it.value = any(val).(T)
	case reflect.Bool:
		val, ok := temp.Value.(bool)
		if !ok {
			return fmt.Errorf("invalid type for value, want bool")
		}
		it.value = any(val).(T)
	default:
		val, ok := temp.Value.(float64)
		if ok {
			converted := reflect.ValueOf(int64(val)).Convert(rv)
			it.value = converted.Interface().(T)
		} else {
			val2, ok2 := temp.Value.(T)
			if !ok2 {
				return fmt.Errorf("invalid type for value")
			}
			it.value = val2
		}
	}
	it.name = temp.Name
	return nil
}

// Registry defines a common enum registry struct.
type Registry[T comparable] struct {
	mu     sync.RWMutex
	values map[T]*Item[T]
	names  map[string]*Item[T]
	items  []*Item[T]
}

// NewRegistry creates a new enum registry.
// Play: https://go.dev/play/p/ABEXsYfJKMo
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
// Play: https://go.dev/play/p/ABEXsYfJKMo
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
// Play: https://go.dev/play/p/dSG84wQ3TuC
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
// Play: https://go.dev/play/p/Ol0moT1J9Xl
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
// Play: https://go.dev/play/p/niJ1U2KlE_m
func (r *Registry[T]) GetByValue(value T) (*Item[T], bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.values[value]
	return item, ok
}

// GetByName retrieves an enum item by its name.
// Play: https://go.dev/play/p/49ie_gpqH0m
func (r *Registry[T]) GetByName(name string) (*Item[T], bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.names[name]
	return item, ok
}

// Items returns a slice of all enum items in the registry.
// Play: https://go.dev/play/p/lAJFAradbvQ
func (r *Registry[T]) Items() []*Item[T] {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*Item[T], len(r.items))
	copy(result, r.items)
	return result
}

// Contains checks if an enum item with the given value exists in the registry.
// Play: https://go.dev/play/p/_T-lPYkZn2j
func (r *Registry[T]) Contains(value T) bool {
	_, ok := r.GetByValue(value)
	return ok
}

// Size returns the number of enum items in the registry.
// Play: https://go.dev/play/p/TeDArWhlQe2
func (r *Registry[T]) Size() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.items)
}

// Range iterates over all enum items in the registry and applies the given function.
// Play: https://go.dev/play/p/GPsZbQbefWN
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
// Play: https://go.dev/play/p/tN9RE_m_WEI
func (r *Registry[T]) SortedItems(less func(*Item[T], *Item[T]) bool) []*Item[T] {
	items := r.Items()
	sort.Slice(items, func(i, j int) bool {
		return less(items[i], items[j])
	})
	return items
}

// Filter returns a slice of enum items that satisfy the given predicate function.
// Play: https://go.dev/play/p/uTUpTdcyoCU
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

package maputil

import (
	"encoding/json"
	"fmt"
	"sync"
)

// BiMap "bidirectional map" is a map that preserves the uniqueness of its values as well as that of its keys
type BiMap[K comparable, V comparable] struct {
	mu      sync.RWMutex
	normal  map[K]V
	reverse map[V]K
}

// Key  returns the key for the given value.
func (m *BiMap[K, V]) Key(v V) K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.reverse[v]
}

// Value  returns the value for the given key.
func (m *BiMap[K, V]) Value(k K) V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.normal[k]
}

// Keys returns a slice of query keys
func (m *BiMap[K, V]) Keys(v ...V) []K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return Keys(FilterByValues(m.normal, v))
}

// Values returns a slice of query values
func (m *BiMap[K, V]) Values(k ...K) []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return Values(FilterByKeys(m.normal, k))
}

// AllKeys returns a slice of all keys
func (m *BiMap[K, V]) AllKeys() []K {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return Keys(m.normal)
}

// AllValues returns a slice of all values
func (m *BiMap[K, V]) AllValues() []V {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return Values(m.normal)
}

// ContainsValue returns true if the given value exists.
func (m *BiMap[K, V]) ContainsValue(v V) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.reverse[v]
	return ok
}

// ContainsKey returns true if the given key exists.
func (m *BiMap[K, V]) ContainsKey(k K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.normal[k]
	return ok
}

// Put insert the given key-value pair. if it already exists return error
func (m *BiMap[K, V]) Put(k K, v V) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.normal[k]; ok {
		return fmt.Errorf("key %v already exists", k)
	}
	if _, ok := m.reverse[v]; ok {
		return fmt.Errorf("value %v already exists", k)
	}
	m.normal[k] = v
	m.reverse[v] = k
	return nil
}

// PutForce force insert the given key-value pair. if the data exists, it will be deleted and inserted again
func (m *BiMap[K, V]) PutForce(k K, v V) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	// check if exists, if exists remove
	if _, ok := m.normal[k]; ok {
		delete(m.reverse, m.normal[k])
		delete(m.normal, k)
	}
	if _, ok := m.reverse[v]; ok {
		delete(m.normal, m.reverse[v])
		delete(m.reverse, v)
	}
	m.normal[k] = v
	m.reverse[v] = k
	return nil
}

// RemoveKey removes a key-value pair from the map by key
func (m *BiMap[K, V]) RemoveKey(k K) {
	if !m.ContainsKey(k) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.reverse, m.normal[k])
	delete(m.normal, k)
}

// RemoveValue removes a key-value pair from the map by Value
func (m *BiMap[K, V]) RemoveValue(v V) {
	if !m.ContainsValue(v) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.normal, m.reverse[v])
	delete(m.reverse, v)
}

// Clear removes all key-value pairs from the map
func (m *BiMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.normal = make(map[K]V)
	m.reverse = make(map[V]K)
}

// Len returns the number of key-value pairs in the map
func (m *BiMap[K, V]) Len() int {
	return len(m.normal)
}

// Inverse returns a new BiMap with the keys and values swapped
func (m *BiMap[K, V]) Inverse() *BiMap[V, K] {
	r, _ := NewBiMapFromMap(m.reverse)
	return r
}

// ToMap returns a map
func (m *BiMap[K, V]) ToMap() map[K]V {
	return m.normal
}

// MarshalJSON implements the json.Marshaler interface.
func (m *BiMap[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.normal)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (m *BiMap[K, V]) UnmarshalJSON(data []byte) error {
	normal := make(map[K]V)
	if err := json.Unmarshal(data, &normal); err != nil {
		return err
	}
	return m.fromMap(normal)
}

func (m *BiMap[K, V]) fromMap(d map[K]V) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.normal == nil || m.reverse == nil {
		m.normal = make(map[K]V)
		m.reverse = make(map[V]K)
	}
	for k, v := range d {
		if _, ok := m.reverse[v]; ok {
			return fmt.Errorf("value %v already exists", v)
		}
		m.normal[k] = v
		m.reverse[v] = k
	}
	return nil
}

// NewBiMap creates a new BiMap
func NewBiMap[K comparable, V comparable]() *BiMap[K, V] {
	return &BiMap[K, V]{
		normal:  make(map[K]V),
		reverse: make(map[V]K),
	}
}

// NewBiMapFromMap creates a new BiMap from a map
func NewBiMapFromMap[K comparable, V comparable](d map[K]V) (*BiMap[K, V], error) {
	biMap := NewBiMap[K, V]()
	if err := biMap.fromMap(d); err != nil {
		return nil, err
	}
	return biMap, nil
}

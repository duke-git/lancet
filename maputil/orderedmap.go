package maputil

import (
	"container/list"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"sync"
)

// OrderedMap is a map that maintains the order of keys.
type OrderedMap[K comparable, V any] struct {
	mu sync.RWMutex

	data  map[K]V
	order *list.List
	index map[K]*list.Element
}

// NewOrderedMap creates a new OrderedMap.
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		data:  make(map[K]V),
		order: list.New(),
		index: make(map[K]*list.Element),
	}
}

// Sets the given key-value pair.
// Play: https://go.dev/play/p/Y4ZJ_oOc1FU
func (om *OrderedMap[K, V]) Set(key K, value V) {
	om.mu.Lock()
	defer om.mu.Unlock()

	if elem, ok := om.index[key]; ok {
		om.data[key] = value
		om.order.MoveToBack(elem)

		return
	}

	om.data[key] = value

	elem := om.order.PushBack(key)
	om.index[key] = elem
}

// Get returns the value for the given key.
// Play: https://go.dev/play/p/Y4ZJ_oOc1FU
func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	om.mu.RLock()
	defer om.mu.RUnlock()

	value, ok := om.data[key]

	return value, ok
}

// Delete deletes the given key.
// Play: https://go.dev/play/p/5bIi4yaZ3K-
func (om *OrderedMap[K, V]) Delete(key K) {
	om.mu.Lock()
	defer om.mu.Unlock()

	if elem, ok := om.index[key]; ok {
		om.order.Remove(elem)
		delete(om.data, key)
		delete(om.index, key)
	}
}

// Clear clears the map.
// Play: https://go.dev/play/p/8LwoJyEfuFr
func (om *OrderedMap[K, V]) Clear() {
	om.mu.Lock()
	defer om.mu.Unlock()

	om.data = make(map[K]V)
	om.order.Init()
	om.index = make(map[K]*list.Element)
}

// Front returns the first key-value pair.
// Play: https://go.dev/play/p/ty57XSimpoe
func (om *OrderedMap[K, V]) Front() (struct {
	Key   K
	Value V
}, bool) {
	om.mu.RLock()
	defer om.mu.RUnlock()

	if elem := om.order.Front(); elem != nil {
		key := elem.Value.(K)
		value := om.data[key]

		return struct {
			Key   K
			Value V
		}{
			Key:   key,
			Value: value,
		}, true
	}

	return struct {
		Key   K
		Value V
	}{}, false
}

// Back returns the last key-value pair.
// Play: https://go.dev/play/p/rQMjp1yQmpa
func (om *OrderedMap[K, V]) Back() (struct {
	Key   K
	Value V
}, bool) {
	om.mu.RLock()
	defer om.mu.RUnlock()

	if elem := om.order.Back(); elem != nil {
		key := elem.Value.(K)
		value := om.data[key]

		return struct {
			Key   K
			Value V
		}{
			Key:   key,
			Value: value,
		}, true
	}

	return struct {
		Key   K
		Value V
	}{}, false
}

// Range calls the given function for each key-value pair.
// Play: https://go.dev/play/p/U-KpORhc7LZ
func (om *OrderedMap[K, V]) Range(iteratee func(key K, value V) bool) {
	om.mu.RLock()
	defer om.mu.RUnlock()

	for elem := om.order.Front(); elem != nil; elem = elem.Next() {
		key := elem.Value.(K)
		value := om.data[key]

		if !iteratee(key, value) {
			break
		}
	}
}

// Keys returns the keys in order.
// Play: https://go.dev/play/p/Vv_y9ExKclA
func (om *OrderedMap[K, V]) Keys() []K {
	om.mu.RLock()
	defer om.mu.RUnlock()

	keys := make([]K, 0, len(om.data))

	for elem := om.order.Front(); elem != nil; elem = elem.Next() {
		keys = append(keys, elem.Value.(K))
	}

	return keys
}

// Values returns the values in order.
// Play: https://go.dev/play/p/TWj5n1-PUfx
func (om *OrderedMap[K, V]) Values() []V {
	om.mu.RLock()
	defer om.mu.RUnlock()

	values := make([]V, 0, len(om.data))

	for elem := om.order.Front(); elem != nil; elem = elem.Next() {
		key := elem.Value.(K)
		values = append(values, om.data[key])
	}

	return values
}

// Len returns the number of key-value pairs.
// Play: https://go.dev/play/p/cLe6z2VX5N-
func (om *OrderedMap[K, V]) Len() int {
	om.mu.RLock()
	defer om.mu.RUnlock()

	return len(om.data)
}

// Contains returns true if the given key exists.
// Play: https://go.dev/play/p/QuwqqnzwDNX
func (om *OrderedMap[K, V]) Contains(key K) bool {
	om.mu.RLock()
	defer om.mu.RUnlock()

	_, ok := om.data[key]

	return ok
}

// Elements returns the key-value pairs in order.
// Play: https://go.dev/play/p/4BHG4kKz6bB
func (om *OrderedMap[K, V]) Elements() []struct {
	Key   K
	Value V
} {
	om.mu.RLock()
	defer om.mu.RUnlock()

	elements := make([]struct {
		Key   K
		Value V
	}, 0, len(om.data))

	for elem := om.order.Front(); elem != nil; elem = elem.Next() {
		key := elem.Value.(K)
		value := om.data[key]
		elements = append(elements, struct {
			Key   K
			Value V
		}{Key: key, Value: value})
	}

	return elements
}

// Iter returns a channel that yields key-value pairs in order.
// Play: https://go.dev/play/p/tlq2tdvicPt
func (om *OrderedMap[K, V]) Iter() <-chan struct {
	Key   K
	Value V
} {
	ch := make(chan struct {
		Key   K
		Value V
	})

	go func() {
		om.mu.RLock()
		defer om.mu.RUnlock()
		defer close(ch)

		for elem := om.order.Front(); elem != nil; elem = elem.Next() {
			key := elem.Value.(K)
			value := om.data[key]

			ch <- struct {
				Key   K
				Value V
			}{Key: key, Value: value}
		}
	}()

	return ch
}

// ReverseIter returns a channel that yields key-value pairs in reverse order.
// Play: https://go.dev/play/p/8Q0ssg6hZzO
func (om *OrderedMap[K, V]) ReverseIter() <-chan struct {
	Key   K
	Value V
} {
	ch := make(chan struct {
		Key   K
		Value V
	})

	go func() {
		om.mu.RLock()
		defer om.mu.RUnlock()
		defer close(ch)

		for elem := om.order.Back(); elem != nil; elem = elem.Prev() {
			key := elem.Value.(K)
			value := om.data[key]

			ch <- struct {
				Key   K
				Value V
			}{Key: key, Value: value}
		}
	}()

	return ch
}

// SortByValue sorts the map by key given less function.
// Play: https://go.dev/play/p/N7hjD_alZPq
func (om *OrderedMap[K, V]) SortByKey(less func(a, b K) bool) {
	om.mu.Lock()
	defer om.mu.Unlock()

	keys := make([]K, 0, om.order.Len())
	for elem := om.order.Front(); elem != nil; elem = elem.Next() {
		keys = append(keys, elem.Value.(K))
	}

	sort.Slice(keys, func(i, j int) bool {
		return less(keys[i], keys[j])
	})

	om.order.Init()
	om.index = make(map[K]*list.Element)
	for _, key := range keys {
		elem := om.order.PushBack(key)
		om.index[key] = elem
	}
}

// MarshalJSON implements the json.Marshaler interface.
// Play: https://go.dev/play/p/C-wAwydIAC7
func (om *OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	om.mu.RLock()
	defer om.mu.RUnlock()

	tempMap := make(map[string]V)
	for e := om.order.Front(); e != nil; e = e.Next() {
		key := e.Value.(K)
		keyStr, err := keyToString(key)
		if err != nil {
			return nil, err
		}
		tempMap[keyStr] = om.data[key]
	}

	return json.Marshal(tempMap)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// Play: https://go.dev/play/p/8C3MvJ3-mut
func (om *OrderedMap[K, V]) UnmarshalJSON(data []byte) error {
	om.mu.Lock()
	defer om.mu.Unlock()

	tempMap := make(map[string]V)
	if err := json.Unmarshal(data, &tempMap); err != nil {
		return err
	}

	om.data = make(map[K]V)
	om.order.Init()
	om.index = make(map[K]*list.Element)

	for keyStr, value := range tempMap {
		key, err := stringToKey[K](keyStr)
		if err != nil {
			return err
		}
		om.data[key] = value
		elem := om.order.PushBack(key)
		om.index[key] = elem
	}

	return nil
}

func keyToString[K any](key K) (string, error) {
	switch v := any(key).(type) {
	case int:
		return strconv.Itoa(v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case string:
		return v, nil
	default:
		// 使用反射将未知类型转换为字符串
		rv := reflect.ValueOf(key)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return strconv.FormatInt(rv.Int(), 10), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return strconv.FormatUint(rv.Uint(), 10), nil
		case reflect.Float32, reflect.Float64:
			return strconv.FormatFloat(rv.Float(), 'f', -1, 64), nil
		case reflect.String:
			return rv.String(), nil
		default:
			return "", fmt.Errorf("unsupported key type: %T", key)
		}
	}
}

func stringToKey[K any](s string) (K, error) {
	var zero K
	switch any(zero).(type) {
	case int:
		value, err := strconv.Atoi(s)
		return any(value).(K), err
	case float64:
		value, err := strconv.ParseFloat(s, 64)
		return any(value).(K), err
	case string:
		return any(s).(K), nil
	default:
		// 使用反射恢复未知类型的键
		rv := reflect.ValueOf(&zero).Elem()
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return zero, err
			}
			rv.SetInt(val)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				return zero, err
			}
			rv.SetUint(val)
		case reflect.Float32, reflect.Float64:
			val, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return zero, err
			}
			rv.SetFloat(val)
		case reflect.String:
			rv.SetString(s)
		default:
			return zero, fmt.Errorf("unsupported key type: %T", zero)
		}

		return rv.Interface().(K), nil
	}
}

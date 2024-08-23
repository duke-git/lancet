// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package maputil includes some functions to manipulate map.
package maputil

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
)

// OrderedMap struct
type OrderedMap[K comparable, V any] struct {
	keys   []K     // key的顺序
	values map[K]V // 核心map
}

// NewOrderedMap create a new OrderedMap
func NewOrderedMap[K comparable, V any]() OrderedMap[K, V] {
	o := OrderedMap[K, V]{}
	o.keys = []K{}
	o.values = map[K]V{}
	return o
}

// Get value by key
func (o *OrderedMap[K, V]) Get(key K) (V, bool) {
	val, exists := o.values[key]
	return val, exists
}

// Set value by key
func (o *OrderedMap[K, V]) Set(key K, value V) {
	_, exists := o.values[key]
	if !exists {
		o.keys = append(o.keys, key)
	}
	o.values[key] = value
}

// Delete value by key
func (o *OrderedMap[K, V]) Delete(key K) {
	_, ok := o.values[key]
	if !ok {
		return
	}
	for i, k := range o.keys {
		if k == key {
			o.keys = append(o.keys[:i], o.keys[i+1:]...)
			break
		}
	}
	delete(o.values, key)
}

// Keys return all keys
func (o *OrderedMap[K, V]) Keys() []K {
	return o.keys
}

// Values return all values
func (o *OrderedMap[K, V]) Values() map[K]V {
	return o.values
}

// MarshalJSON marshal OrderedMap to json
func (o *OrderedMap[K, V]) MarshalJSON() ([]byte, error) {
	// json 不允许数字类型作为key
	for k := range o.values {
		if reflect.ValueOf(k).Kind() != reflect.String {
			return nil, errors.New("json: unsupported type: map key")
		}
		break
	}

	var buf bytes.Buffer
	buf.WriteByte('{')
	encoder := json.NewEncoder(&buf)
	for i, k := range o.keys {
		if i > 0 {
			buf.WriteByte(',')
		}
		if err := encoder.Encode(k); err != nil {
			return nil, err
		}
		buf.WriteByte(':')
		if err := encoder.Encode(o.values[k]); err != nil {
			return nil, err
		}
	}
	buf.WriteByte('}')
	return buf.Bytes(), nil
}

// UnmarshalJSON unmarshal json to OrderedMap
func (o *OrderedMap[K, V]) UnmarshalJSON(b []byte) error {
	if o.values == nil {
		o.values = map[K]V{}
	}
	// 这里我们直接将map解开了
	err := json.Unmarshal(b, &o.values)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(b))
	if _, err = dec.Token(); err != nil {
		return err
	}
	o.keys = make([]K, 0, len(o.values))
	return decodeOrderedMap(dec, o)
}

// decodeOrderedMap internal function to decode json
func decodeOrderedMap[K comparable, V any](dec *json.Decoder, o *OrderedMap[K, V]) error {
	hasKey := make(map[K]bool, len(o.values))
	skip1 := false
	skip2 := false
	for {
		// 这个token是key
		token, err := dec.Token()
		if err != nil {
			return err
		}
		if delim, ok := token.(json.Delim); ok {
			if skip1 && delim == '}' {
				skip1 = false
				continue
			} else if skip2 && delim == ']' {
				skip2 = false
				continue
			}
			if !skip1 && !skip2 {
				if delim == '}' {
					return nil
				}
			}
		}
		if skip1 || skip2 {
			// 跳过key对应的value，直接读取下一个token
			continue
		}
		key := token.(K)
		if hasKey[key] {
			// 重复的key，删除掉
			for j, k := range o.keys {
				if k == key {
					copy(o.keys[j:], o.keys[j+1:])
					break
				}
			}
			o.keys[len(o.keys)-1] = key
		} else {
			hasKey[key] = true
			o.keys = append(o.keys, key)
		}
		// 下一个token是value
		token, err = dec.Token()
		if err != nil {
			return err
		}
		if delim, ok := token.(json.Delim); ok {
			switch delim {
			case '{':
				// 开始进入token提取
				skip1 = true
			case '[':
				// 如果是数组，那么递归解析
				skip2 = true
			}
		}
	}
}

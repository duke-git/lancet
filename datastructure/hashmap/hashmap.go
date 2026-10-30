// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure implements some data structure. hashmap structure.
package datastructure

import (
	"fmt"
	"hash/fnv"
	"reflect"
)

var defaultMapCapacity uint64 = 1 << 10

type mapNode struct {
	key   any
	value any
	next  *mapNode
}

// HashMap implements a hash map
type HashMap struct {
	capacity uint64
	size     uint64
	table    []*mapNode
}

// NewHashMap return a HashMap instance
func NewHashMap() *HashMap {
	return &HashMap{
		capacity: defaultMapCapacity,
		table:    make([]*mapNode, defaultMapCapacity),
	}
}

// NewHashMapWithCapacity return a HashMap instance with given size and capacity
func NewHashMapWithCapacity(size, capacity uint64) *HashMap {
	return &HashMap{
		size:     size,
		capacity: capacity,
		table:    make([]*mapNode, capacity),
	}
}

// Get return the value of given key in hashmap
func (hm *HashMap) Get(key any) any {
	if hm.size == 0 {
		return nil
	}
	hashValue := hm.hash(key)
	node := hm.table[hashValue]
	for node != nil {
		if reflect.DeepEqual(node.key, key) {
			return node.value
		}
		node = node.next
	}
	return nil
}

// GetOrDefault return the value of given key in hashmap, if not found return default value
func (hm *HashMap) GetOrDefault(key any, defaultValue any) any {
	value := hm.Get(key)
	if value == nil {
		return defaultValue
	}
	return value
}

// Put new key value in hashmap
func (hm *HashMap) Put(key any, value any) {
	if hm.capacity == 0 {
		hm.capacity = defaultMapCapacity
		hm.table = make([]*mapNode, defaultMapCapacity)
	}
	hm.putValue(hm.hash(key), key, value)
}

func (hm *HashMap) putValue(hash uint64, key, value any) {
	for node := hm.table[hash]; node != nil; node = node.next {
		if reflect.DeepEqual(node.key, key) {
			node.value = value
			return
		}
	}

	if hm.size >= hm.capacity {
		hm.resize()
		hash = hm.hash(key)
	}
	hm.table[hash] = newMapNodeWithNext(key, value, hm.table[hash])
	hm.size++
}

// Delete item by given key in hashmap
func (hm *HashMap) Delete(key any) {
	if hm.size == 0 {
		return
	}
	hash := hm.hash(key)
	var previous *mapNode
	for node := hm.table[hash]; node != nil; node = node.next {
		if reflect.DeepEqual(node.key, key) {
			if previous == nil {
				hm.table[hash] = node.next
			} else {
				previous.next = node.next
			}
			hm.size--
			return
		}
		previous = node
	}
}

// Contains checks if given key is in hashmap or not
func (hm *HashMap) Contains(key any) bool {
	if hm.size == 0 {
		return false
	}
	node := hm.table[hm.hash(key)]
	for node != nil {
		if reflect.DeepEqual(node.key, key) {
			return true
		}
		node = node.next
	}
	return false
}

// Iterate executes iteratee funcation for every key and value pair of hashmap (random order)
func (hm *HashMap) Iterate(iteratee func(key, value any)) {
	for _, node := range hm.table {
		for node != nil {
			iteratee(node.key, node.value)
			node = node.next
		}
	}
}

// FilterByValue returns a filtered HashMap.
// If any value is not matching the perdicate function then it returns nil
// otherwise it returns the HashMap with selected values.
func (hm *HashMap) FilterByValue(perdicate func(value any) bool) *HashMap {
	var filteredHM *HashMap
	hm.Iterate(func(key, value any) {
		if perdicate(value) {
			if filteredHM == nil {
				filteredHM = NewHashMap()
			}
			filteredHM.Put(key, value)
		}
	})
	return filteredHM
}

// Keys returns a slice of the hashmap's keys (random order)
func (hm *HashMap) Keys() []any {
	keys := make([]any, int(hm.size))
	index := 0
	if hm.size > 0 {
		hm.Iterate(func(key, value any) {
			keys[index] = key
			index++
		})
	}

	return keys
}

// Values returns a slice of the hashmap's keys (random order)
func (hm *HashMap) Values() []any {
	values := make([]any, int(hm.size))
	index := 0
	if hm.size > 0 {
		hm.Iterate(func(key, value any) {
			values[index] = value
			index++
		})
	}

	return values
}

func (hm *HashMap) resize() {
	hm.capacity <<= 1
	tempTable := hm.table
	hm.table = make([]*mapNode, hm.capacity)
	for _, node := range tempTable {
		for node != nil {
			next := node.next
			hash := hm.hash(node.key)
			node.next = hm.table[hash]
			hm.table[hash] = node
			node = next
		}
	}
}

// Size returns current size of Hashmap
func (hm *HashMap) Size() uint64 {
	return hm.size
}

func (hm *HashMap) hash(key any) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

func newMapNodeWithNext(key, value any, next *mapNode) *mapNode {
	return &mapNode{
		key:   key,
		value: value,
		next:  next,
	}
}

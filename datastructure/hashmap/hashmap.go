// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure implements some data structure. eg. list, linklist, stack, queue, tree, graph.
package datastructure

import (
	"fmt"
	"hash/fnv"
)

var defaultMapCapacity uint64 = 1 << 10

type mapNode struct {
	key   any
	value any
	next  *mapNode
}

//HashMap implements a hash map
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
	hashValue := hm.hash(key)
	node := hm.table[hashValue]
	if node != nil {
		return node.value
	}

	return nil
}

// Put new key value in hashmap
func (hm *HashMap) Put(key any, value any) any {
	return hm.putValue(hm.hash(key), key, value)
}

func (hm *HashMap) putValue(hash uint64, key, value any) any {
	if hm.capacity == 0 {
		hm.capacity = defaultMapCapacity
		hm.table = make([]*mapNode, defaultMapCapacity)
	}

	node := hm.table[hash]
	if node == nil {
		hm.table[hash] = newMapNode(key, value)
	} else if node.key == key {
		hm.table[hash] = newMapNodeWithNext(key, value, node)
		return value
	} else {
		hm.resize()
		return hm.putValue(hash, value, value)
	}
	hm.size++

	return value
}

// Delete key value in hashmap
func (hm *HashMap) Delete(key any) {
	hash := hm.hash(key)
	node := hm.table[hash]
	if node == nil {
		return
	}

	hm.table = append(hm.table[:hash], hm.table[hash+1:]...)
	hm.size--
}

// Contains checks if given key is in hashmap or not
func (hm *HashMap) Contains(key any) bool {
	node := hm.table[hm.hash(key)]
	return node != nil
}

func (hm *HashMap) resize() {
	hm.capacity <<= 1

	tempTable := hm.table

	for i := 0; i < len(tempTable); i++ {
		node := tempTable[i]
		if node == nil {
			continue
		}

		hm.table[hm.hash(node.key)] = node
	}
}

func (hm *HashMap) hash(key any) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

func newMapNode(key, value any) *mapNode {
	return &mapNode{
		key:   key,
		value: value,
	}
}

func newMapNodeWithNext(key, value any, next *mapNode) *mapNode {
	return &mapNode{
		key:   key,
		value: value,
		next:  next,
	}
}

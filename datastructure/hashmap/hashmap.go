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

// Get return the value of given key in hash map
func (hm *HashMap) Get(key any) any {
	hashValue := hm.hash(key)
	node := hm.table[hashValue]
	if node != nil {
		return node.value
	}

	return nil
}

func (hm *HashMap) hash(key any) uint64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(fmt.Sprintf("%v", key)))

	hashValue := h.Sum64()

	return (hm.capacity - 1) & (hashValue ^ (hashValue >> 16))
}

func newMapNode(key any, value any) *mapNode {
	return &mapNode{
		key:   key,
		value: value,
	}
}

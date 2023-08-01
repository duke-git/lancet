// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package maputil includes some functions to manipulate map.
package maputil

import (
	"fmt"
	"sync"
)

const defaultShardCount = 32

// ConcurrentMap is like map, but is safe for concurrent use by multiple goroutines.
type ConcurrentMap[K comparable, V any] struct {
	shardCount uint64
	locks      []sync.RWMutex
	maps       []map[K]V
}

// NewConcurrentMap create a ConcurrentMap with specific shard count.
// Play: https://go.dev/play/p/3PenTPETJT0
func NewConcurrentMap[K comparable, V any](shardCount int) *ConcurrentMap[K, V] {
	if shardCount <= 0 {
		shardCount = defaultShardCount
	}

	cm := &ConcurrentMap[K, V]{
		shardCount: uint64(shardCount),
		locks:      make([]sync.RWMutex, shardCount),
		maps:       make([]map[K]V, shardCount),
	}

	for i := range cm.maps {
		cm.maps[i] = make(map[K]V)
	}

	return cm
}

// Set the value for a key.
// Play: https://go.dev/play/p/3PenTPETJT0
func (cm *ConcurrentMap[K, V]) Set(key K, value V) {
	shard := cm.getShard(key)

	cm.locks[shard].Lock()
	cm.maps[shard][key] = value

	cm.locks[shard].Unlock()
}

// Get the value stored in the map for a key, or nil if no.
// Play: https://go.dev/play/p/3PenTPETJT0
func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	shard := cm.getShard(key)

	cm.locks[shard].RLock()
	value, ok := cm.maps[shard][key]
	cm.locks[shard].RUnlock()

	return value, ok
}

// GetOrSet returns the existing value for the key if present.
// Otherwise, it sets and returns the given value.
// Play: https://go.dev/play/p/aDcDApOK01a
func (cm *ConcurrentMap[K, V]) GetOrSet(key K, value V) (actual V, ok bool) {
	shard := cm.getShard(key)

	cm.locks[shard].RLock()
	if actual, ok := cm.maps[shard][key]; ok {
		cm.locks[shard].RUnlock()
		return actual, ok
	}
	cm.locks[shard].RUnlock()

	// lock again
	cm.locks[shard].Lock()
	if actual, ok = cm.maps[shard][key]; ok {
		cm.locks[shard].Unlock()
		return
	}

	cm.maps[shard][key] = value
	cm.locks[shard].Unlock()

	return value, ok
}

// Delete the value for a key.
// Play: https://go.dev/play/p/uTIJZYhpVMS
func (cm *ConcurrentMap[K, V]) Delete(key K) {
	shard := cm.getShard(key)

	cm.locks[shard].Lock()
	delete(cm.maps[shard], key)
	cm.locks[shard].Unlock()
}

// GetAndDelete returns the existing value for the key if present and then delete the value for the key.
// Otherwise, do nothing, just return false
// Play: https://go.dev/play/p/ZyxeIXSZUiM
func (cm *ConcurrentMap[K, V]) GetAndDelete(key K) (actual V, ok bool) {
	shard := cm.getShard(key)

	cm.locks[shard].RLock()
	if actual, ok = cm.maps[shard][key]; ok {
		cm.locks[shard].RUnlock()
		cm.Delete(key)
		return
	}
	cm.locks[shard].RUnlock()

	return actual, false
}

// Has checks if map has the value for a key.
// Play: https://go.dev/play/p/C8L4ul9TVwf
func (cm *ConcurrentMap[K, V]) Has(key K) bool {
	_, ok := cm.Get(key)
	return ok
}

// Range calls iterator sequentially for each key and value present in each of the shards in the map.
// If iterator returns false, range stops the iteration.
// Play: https://go.dev/play/p/iqcy7P8P0Pr
func (cm *ConcurrentMap[K, V]) Range(iterator func(key K, value V) bool) {
	for shard := range cm.locks {
		cm.locks[shard].RLock()

		for k, v := range cm.maps[shard] {
			if !iterator(k, v) {
				cm.locks[shard].RUnlock()
				return
			}
		}
		cm.locks[shard].RUnlock()
	}
}

// getShard get shard by a key.
func (cm *ConcurrentMap[K, V]) getShard(key K) uint64 {
	hash := fnv32(fmt.Sprintf("%v", key))
	return uint64(hash) % cm.shardCount
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

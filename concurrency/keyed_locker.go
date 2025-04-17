// Copyright 2025 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package concurrency contain some functions to support concurrent programming. eg, goroutine, channel, locker.

package concurrency

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// KeyedLocker is a simple implementation of a keyed locker that allows for non-blocking lock acquisition.
type KeyedLocker[K comparable] struct {
	locks sync.Map
	ttl   time.Duration
}

type lockEntry struct {
	mu    sync.Mutex
	ref   int32
	timer atomic.Pointer[time.Timer]
}

// NewKeyedLocker creates a new KeyedLocker with the specified TTL for lock expiration.
// The TTL is used to automatically release locks that are no longer held.
func NewKeyedLocker[K comparable](ttl time.Duration) *KeyedLocker[K] {
	return &KeyedLocker[K]{ttl: ttl}
}

// Do acquires a lock for the specified key and executes the provided function.
// It returns an error if the context is canceled before the function completes.
func (l *KeyedLocker[K]) Do(ctx context.Context, key K, fn func()) error {
	entry := l.acquire(key)
	defer l.release(key, entry, key)

	done := make(chan struct{})

	go func() {
		entry.mu.Lock()
		defer entry.mu.Unlock()

		select {
		case <-ctx.Done():
		default:
			fn()
		}
		close(done)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}
}

func (l *KeyedLocker[K]) acquire(key K) *lockEntry {
	lock, _ := l.locks.LoadOrStore(key, &lockEntry{})
	entry := lock.(*lockEntry)

	atomic.AddInt32(&entry.ref, 1)
	if t := entry.timer.Swap(nil); t != nil {
		t.Stop()
	}

	return entry
}

func (l *KeyedLocker[K]) release(key K, entry *lockEntry, rawKey K) {
	if atomic.AddInt32(&entry.ref, -1) == 0 {
		entry.mu.Lock()
		defer entry.mu.Unlock()

		if entry.ref == 0 {
			if t := entry.timer.Swap(nil); t != nil {
				t.Stop()
			}

			l.locks.Delete(rawKey)
		} else {
			if entry.timer.Load() == nil {
				t := time.AfterFunc(l.ttl, func() {
					l.release(key, entry, rawKey)
				})
				entry.timer.Store(t)
			}
		}
	}
}

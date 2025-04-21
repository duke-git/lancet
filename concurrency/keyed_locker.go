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

// acquire tries to acquire a lock for the specified key.
func (l *KeyedLocker[K]) acquire(key K) *lockEntry {
	lock, _ := l.locks.LoadOrStore(key, &lockEntry{})
	entry := lock.(*lockEntry)

	atomic.AddInt32(&entry.ref, 1)
	if t := entry.timer.Swap(nil); t != nil {
		t.Stop()
	}

	return entry
}

// release releases the lock for the specified key.
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

// RWKeyedLocker is a read-write version of KeyedLocker.
type RWKeyedLocker[K comparable] struct {
	locks sync.Map
	ttl   time.Duration
}

type rwLockEntry struct {
	mu    sync.RWMutex
	ref   int32
	timer atomic.Pointer[time.Timer]
}

// NewRWKeyedLocker creates a new RWKeyedLocker with the specified TTL for lock expiration.
// The TTL is used to automatically release locks that are no longer held.
func NewRWKeyedLocker[K comparable](ttl time.Duration) *RWKeyedLocker[K] {
	return &RWKeyedLocker[K]{ttl: ttl}
}

// RLock acquires a read lock for the specified key and executes the provided function.
// It returns an error if the context is canceled before the function completes.
func (l *RWKeyedLocker[K]) RLock(ctx context.Context, key K, fn func()) error {
	entry := l.acquire(key)
	defer l.release(entry, key)

	done := make(chan struct{})

	go func() {
		entry.mu.RLock()
		defer entry.mu.RUnlock()

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

// Lock acquires a write lock for the specified key and executes the provided function.
// It returns an error if the context is canceled before the function completes.
func (l *RWKeyedLocker[K]) Lock(ctx context.Context, key K, fn func()) error {
	entry := l.acquire(key)
	defer l.release(entry, key)

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

// acquire tries to acquire a read lock for the specified key.
func (l *RWKeyedLocker[K]) acquire(key K) *rwLockEntry {
	actual, _ := l.locks.LoadOrStore(key, &rwLockEntry{})
	entry := actual.(*rwLockEntry)
	atomic.AddInt32(&entry.ref, 1)

	if t := entry.timer.Swap(nil); t != nil {
		t.Stop()
	}
	return entry
}

// release releases the lock for the specified key.
func (l *RWKeyedLocker[K]) release(entry *rwLockEntry, rawKey K) {
	if atomic.AddInt32(&entry.ref, -1) == 0 {
		timer := time.AfterFunc(l.ttl, func() {
			if atomic.LoadInt32(&entry.ref) == 0 {
				l.locks.Delete(rawKey)
			}
		})
		entry.timer.Store(timer)
	}
}

// TryKeyedLocker is a non-blocking version of KeyedLocker.
// It allows for trying to acquire a lock without blocking if the lock is already held.
type TryKeyedLocker[K comparable] struct {
	mu    sync.Mutex
	locks map[K]*casMutex
}

// NewTryKeyedLocker creates a new TryKeyedLocker.
func NewTryKeyedLocker[K comparable]() *TryKeyedLocker[K] {
	return &TryKeyedLocker[K]{locks: make(map[K]*casMutex)}
}

// TryLock tries to acquire a lock for the specified key.
// It returns true if the lock was acquired, false otherwise.
func (l *TryKeyedLocker[K]) TryLock(key K) bool {
	l.mu.Lock()

	lock, ok := l.locks[key]
	if !ok {
		lock = &casMutex{}
		l.locks[key] = lock
	}
	l.mu.Unlock()

	return lock.TryLock()
}

// Unlock releases the lock for the specified key.
func (l *TryKeyedLocker[K]) Unlock(key K) {
	l.mu.Lock()
	defer l.mu.Unlock()

	lock, ok := l.locks[key]
	if ok {
		lock.Unlock()
		if lock.lock == 0 {
			delete(l.locks, key)
		}
	}
}

// casMutex is a simple mutex that uses atomic operations to provide a non-blocking lock.
type casMutex struct {
	lock int32
}

// TryLock tries to acquire the lock without blocking.
// It returns true if the lock was acquired, false otherwise.
func (m *casMutex) TryLock() bool {
	return atomic.CompareAndSwapInt32(&m.lock, 0, 1)
}

// Unlock releases the lock.
func (m *casMutex) Unlock() {
	atomic.StoreInt32(&m.lock, 0)
}

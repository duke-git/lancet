// Copyright 2025 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package eventbus implements a simple event bus.
package eventbus

import (
	"fmt"
	"sort"
	"sync"
)

// Event is the struct that is passed to the event listener, now it directly uses the generic Payload type.
type Event[T any] struct {
	Topic   string
	Payload T
}

// EventBus is the struct that holds the listeners and the error handler.
type EventBus[T any] struct {
	// listeners    map[string][]*EventListener[T]
	listeners    sync.Map
	mu           sync.RWMutex
	errorHandler func(err error)
}

// EventListener is the struct that holds the listener function and its priority.
type EventListener[T any] struct {
	priority int
	listener func(eventData T)
	async    bool
	filter   func(eventData T) bool
}

// NewEventBus creates a new EventBus.
func NewEventBus[T any]() *EventBus[T] {
	return &EventBus[T]{
		listeners: sync.Map{},
	}
}

// Subscribe subscribes to an event with a specific event topic and listener function.
func (eb *EventBus[T]) Subscribe(topic string, listener func(eventData T), async bool, priority int, filter func(eventData T) bool) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	el := &EventListener[T]{
		priority: priority,
		listener: listener,
		async:    async,
		filter:   filter,
	}

	listenersInterface, _ := eb.listeners.LoadOrStore(topic, []*EventListener[T]{})
	listeners := listenersInterface.([]*EventListener[T])

	listeners = append(listeners, el)
	sort.Slice(listeners, func(i, j int) bool {
		return listeners[i].priority > listeners[j].priority
	})

	eb.listeners.Store(topic, listeners)
}

// Unsubscribe unsubscribes from an event with a specific event topic and listener function.
func (eb *EventBus[T]) Unsubscribe(topic string, listener func(eventData T)) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	listenersInterface, ok := eb.listeners.Load(topic)
	if !ok {
		return
	}

	listeners := listenersInterface.([]*EventListener[T])
	listenerPtr := fmt.Sprintf("%p", listener)

	var updatedListeners []*EventListener[T]
	for _, l := range listeners {
		if fmt.Sprintf("%p", l.listener) != listenerPtr {
			updatedListeners = append(updatedListeners, l)
		}
	}

	eb.listeners.Store(topic, updatedListeners)
}

// Publish publishes an event with a specific event topic and data payload.
func (eb *EventBus[T]) Publish(event Event[T]) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	listenersInterface, exists := eb.listeners.Load(event.Topic)
	if !exists {
		return
	}

	listeners := listenersInterface.([]*EventListener[T])

	for _, listener := range listeners {
		if listener.filter != nil && !listener.filter(event.Payload) {
			continue
		}

		if listener.async {
			go eb.publishToListener(listener, event)
		} else {
			eb.publishToListener(listener, event)
		}
	}
}

func (eb *EventBus[T]) publishToListener(listener *EventListener[T], event Event[T]) {
	defer func() {
		if r := recover(); r != nil && eb.errorHandler != nil {
			eb.errorHandler(fmt.Errorf("%v", r))
		}
	}()

	listener.listener(event.Payload)
}

// SetErrorHandler sets the error handler function.
func (eb *EventBus[T]) SetErrorHandler(handler func(err error)) {
	eb.errorHandler = handler
}

// ClearListeners clears all the listeners.
func (eb *EventBus[T]) ClearListeners() {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.listeners = sync.Map{}
}

// ClearListenersByTopic clears all the listeners by topic.
func (eb *EventBus[T]) ClearListenersByTopic(topic string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.listeners.Delete(topic)
}

// GetListenersCount returns the number of listeners for a specific event topic.
func (eb *EventBus[T]) GetListenersCount(topic string) int {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	listenersInterface, ok := eb.listeners.Load(topic)
	if !ok {
		return 0
	}

	listeners := listenersInterface.([]*EventListener[T])
	return len(listeners)
}

// GetAllListenersCount returns the total number of listeners.
func (eb *EventBus[T]) GetAllListenersCount() int {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	count := 0
	eb.listeners.Range(func(key, value interface{}) bool {
		count += len(value.([]*EventListener[T]))
		return true
	})

	return count
}

// GetEvents returns all the events topics.
func (eb *EventBus[T]) GetEvents() []string {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	var events []string

	eb.listeners.Range(func(key, value interface{}) bool {
		events = append(events, key.(string))
		return true
	})

	return events
}

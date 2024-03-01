package optional

import (
	"sync"
)

// Optional is a type that may or may not contain a non-nil value.
type Optional[T any] struct {
	value *T
	mu    *sync.RWMutex
}

// Default returns an default Optional instance.
func Default[T any]() Optional[T] {
	return Optional[T]{mu: &sync.RWMutex{}}
}

// Of returns an Optional with a non-nil value.
func Of[T any](value T) Optional[T] {
	return Optional[T]{value: &value, mu: &sync.RWMutex{}}
}

// FromNillable returns an Optional for a given value, which may be nil.
func FromNillable[T any](value *T) Optional[T] {
	if value == nil {
		return Default[T]()
	}
	return Optional[T]{value: value, mu: &sync.RWMutex{}}
}

// IsNotNil checks if there is a value present.
func (o Optional[T]) IsNotNil() bool {
	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.value != nil
}

// IsNil checks if the Optional is nil.
func (o Optional[T]) IsNil() bool {
	return !o.IsNotNil()
}

// IfNotNil performs the given action with the value if a value is not nil.
func (o Optional[T]) IfNotNil(action func(value T)) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value != nil {
		action(*o.value)
	}
}

// IfNotNilOrElse performs the action with the value if present, otherwise performs the fallback action.
func (o Optional[T]) IfNotNilOrElse(action func(value T), fallbackAction func()) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value != nil {
		action(*o.value)
	} else {
		fallbackAction()
	}
}

// Unwarp returns the value if not nil, otherwise panics.
func (o Optional[T]) Unwarp() T {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value == nil {
		panic("Optional.Get: no value present")
	}
	return *o.value
}

// OrElse returns the value if is not nil, otherwise returns other.
func (o Optional[T]) OrElse(other T) T {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value != nil {
		return *o.value
	}
	return other
}

// OrElseGet returns the value if is not nil, otherwise invokes action and returns the result.
func (o Optional[T]) OrElseGet(action func() T) T {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value != nil {
		return *o.value
	}
	return action()
}

// OrElseTrigger returns the value if present, otherwise returns an error.
func (o Optional[T]) OrElseTrigger(errorHandler func() error) (T, error) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value == nil {
		return *new(T), errorHandler()
	}
	return *o.value, nil
}

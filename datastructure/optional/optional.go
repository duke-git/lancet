package optional

import (
	"sync"
)

// Optional is a type that may or may not contain a non-nil value.
type Optional[T any] struct {
	value *T
	mu    sync.RWMutex 
}

// Empty returns an empty Optional instance.
func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

// Of returns an Optional with a non-nil value.
func Of[T any](value T) Optional[T] {
	return Optional[T]{value: &value}
}

// OfNullable returns an Optional for a given value, which may be nil.
func OfNullable[T any](value *T) Optional[T] {
	if value == nil {
		return Empty[T]()
	}
	return Optional[T]{value: value}
}

// IsPresent checks if there is a value present.
func (o Optional[T]) IsPresent() bool {
	o.mu.RLock()
	defer o.mu.RUnlock()

	return o.value != nil
}

// IsEmpty checks if the Optional is empty.
func (o Optional[T]) IsEmpty() bool {
	return !o.IsPresent()
}

// IfPresent performs the given action with the value if a value is present.
func (o Optional[T]) IfPresent(action func(value T)) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value != nil {
		action(*o.value)
	}
}

// IfPresentOrElse performs the action with the value if present, otherwise performs the empty-based action.
func (o Optional[T]) IfPresentOrElse(action func(value T), emptyAction func()) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value != nil {
		action(*o.value)
	} else {
		emptyAction()
	}
}

// Get returns the value if present, otherwise panics.
func (o Optional[T]) Get() T {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value == nil {
		panic("Optional.Get: no value present")
	}
	return *o.value
}

// OrElse returns the value if present, otherwise returns other.
func (o Optional[T]) OrElse(other T) T {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value != nil {
		return *o.value
	}
	return other
}

// OrElseGet returns the value if present, otherwise invokes supplier and returns the result.
func (o Optional[T]) OrElseGet(supplier func() T) T {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value != nil {
		return *o.value
	}
	return supplier()
}

// OrElseThrow returns the value if present, otherwise returns an error.
func (o Optional[T]) OrElseThrow(errorSupplier func() error) (T, error) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	if o.value == nil {
		return *new(T), errorSupplier()
	}
	return *o.value, nil
}

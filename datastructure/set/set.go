// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package datastructure contains some data structure. Set is a data container, like slice, but element of set is not duplicate.
package datastructure

// Set is a data container, like slice, but element of set is not duplicate.
type Set[T comparable] map[T]struct{}

// New create a instance of set from given values.
func New[T comparable](items ...T) Set[T] {
	set := make(Set[T])
	set.Add(items...)
	return set
}

// FromSlice create a set from given slice.
func FromSlice[T comparable](items []T) Set[T] {
	set := make(Set[T], len(items))
	for _, item := range items {
		set.Add(item)
	}
	return set
}

// Add items to set
func (s Set[T]) Add(items ...T) {
	for _, v := range items {
		s[v] = struct{}{}
	}
}

// AddIfNotExist checks if item exists in the set,
// it adds the item to set and returns true if it does not exist in the set,
// or else it does nothing and returns false.
func (s Set[T]) AddIfNotExist(item T) bool {
	if !s.Contain(item) {
		if _, ok := s[item]; !ok {
			s[item] = struct{}{}
			return true
		}
	}
	return false
}

// AddIfNotExistBy checks if item exists in the set and pass the `checker` function
// it adds the item to set and returns true if it does not exists in the set and
// function `checker` returns true, or else it does nothing and returns false.
func (s Set[T]) AddIfNotExistBy(item T, checker func(element T) bool) bool {
	if !s.Contain(item) {
		if checker(item) {
			if _, ok := s[item]; !ok {
				s[item] = struct{}{}
				return true
			}
		}
	}
	return false
}

// Contain checks if set contains item or not
func (s Set[T]) Contain(item T) bool {
	_, ok := s[item]
	return ok
}

// ContainAll checks if set contains other set
func (s Set[T]) ContainAll(other Set[T]) bool {
	for k := range other {
		_, ok := s[k]
		if !ok {
			return false
		}
	}
	return true
}

// Clone return a copy of set
func (s Set[T]) Clone() Set[T] {
	set := New[T]()
	set.Add(s.Values()...)
	return set
}

// Delete item of set
func (s Set[T]) Delete(items ...T) {
	for _, v := range items {
		delete(s, v)
	}
}

// Equal checks if two set has same elements or not
func (s Set[T]) Equal(other Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}

	return s.ContainAll(other) && other.ContainAll(s)
}

// Iterate call function by every element of set
func (s Set[T]) Iterate(fn func(item T)) {
	for v := range s {
		fn(v)
	}
}

// IsEmpty checks the set is empty or not
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// Size get the number of elements in set
func (s Set[T]) Size() int {
	return len(s)
}

// Values return all values of set
func (s Set[T]) Values() []T {
	result := make([]T, 0, len(s))

	s.Iterate(func(value T) {
		result = append(result, value)
	})

	return result
}

// Union creates a new set contain all element of set s and other
func (s Set[T]) Union(other Set[T]) Set[T] {
	set := s.Clone()
	set.Add(other.Values()...)
	return set
}

// Intersection creates a new set whose element both be contained in set s and other
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	set := New[T]()
	s.Iterate(func(value T) {
		if other.Contain(value) {
			set.Add(value)
		}
	})

	return set
}

// SymmetricDifference creates a new set whose element is in set1 or set2, but not in both sets
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	set := New[T]()
	s.Iterate(func(value T) {
		if !other.Contain(value) {
			set.Add(value)
		}
	})

	other.Iterate(func(value T) {
		if !s.Contain(value) {
			set.Add(value)
		}
	})

	return set
}

// Minus creates an set of whose element in origin set but not in compared set
func (s Set[T]) Minus(comparedSet Set[T]) Set[T] {
	set := New[T]()

	s.Iterate(func(value T) {
		if !comparedSet.Contain(value) {
			set.Add(value)
		}
	})

	return set
}

// EachWithBreak iterates over elements of a set and invokes function for each element,
// when iteratee return false, will break the for each loop.
func (s Set[T]) EachWithBreak(iteratee func(item T) bool) {
	for _, v := range s.Values() {
		if !iteratee(v) {
			break
		}
	}
}

// Pop delete the top element of set then return it, if set is empty, return nil-value of T and false.
func (s Set[T]) Pop() (v T, ok bool) {
	if len(s) > 0 {
		items := s.Values()
		item := items[len(s)-1]
		delete(s, item)
		return item, true
	}

	return v, false
}

// ToSlice convert set to slice
func (s Set[T]) ToSlice() []T {
	if s.IsEmpty() {
		return []T{}
	}

	res := make([]T, 0, s.Size())
	for v := range s {
		res = append(res, v)
	}
	return res
}

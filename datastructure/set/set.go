package datastructure

// Set is a data container, like slice, but element of set is not duplicate
type Set[T comparable] map[T]struct{}

// NewSet return a instance of set
func NewSet[T comparable](items ...T) Set[T] {
	set := make(Set[T])
	set.Add(items...)
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
func (s Set[T]) AddIfNotExist(value T) bool {
	if !s.Contain(value) {
		if _, ok := s[value]; !ok {
			s[value] = struct{}{}
			return true
		}
	}
	return false
}

// Contain checks if set contains value or not
func (s Set[T]) Contain(value T) bool {
	_, ok := s[value]
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
	set := NewSet[T]()
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
func (s Set[T]) Iterate(fn func(value T)) {
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
	set := NewSet[T]()
	s.Iterate(func(value T) {
		if other.Contain(value) {
			set.Add(value)
		}
	})

	return set
}

// SymmetricDifference creates a new set whose element is in set1 or set2, but not in both sets
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	set := NewSet[T]()
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
	set := NewSet[T]()

	s.Iterate(func(value T) {
		if !comparedSet.Contain(value) {
			set.Add(value)
		}
	})

	return set
}

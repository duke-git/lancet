package datastructure

// Set is a data container, like slice, but element of set is not duplicate
type Set[T comparable] map[T]bool

// NewSet return a instance of set
func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T])
	set.Add(values...)
	return set
}

// Add value to set
func (s Set[T]) Add(values ...T) {
	for _, v := range values {
		s[v] = true
	}
}

// Contain checks if set contains value or not
func (s Set[T]) Contain(value T) bool {
	_, ok := s[value]
	return ok
}

// Contain checks if set contains other set
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

// Delete value of set
func (s Set[T]) Delete(values ...T) {
	for _, v := range values {
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
	values := make([]T, 0, 0)

	s.Iterate(func(value T) {
		values = append(values, value)
	})

	return values
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

package function

// And returns a composed predicate that represents the logical AND of a list of predicates.
// It evaluates to true only if all predicates evaluate to true for the given value.
func And[T any](predicates ...func(T) bool) func(T) bool {
	return func(value T) bool {
		for _, predicate := range predicates {
			if !predicate(value) {
				return false // Short-circuit on the first false predicate
			}
		}
		return true // True if all predicates are true
	}
}

// Negate returns a predicate that represents the logical negation of this predicate.
func Negate[T any](predicate func(T) bool) func(T) bool {
	return func(value T) bool {
		return !predicate(value)
	}
}

// Or returns a composed predicate that represents the logical OR of a list of predicates.
// It evaluates to true if at least one of the predicates evaluates to true for the given value.
func Or[T any](predicates ...func(T) bool) func(T) bool {
	return func(value T) bool {
		for _, predicate := range predicates {
			if predicate(value) {
				return true // Short-circuit on the first true predicate
			}
		}
		return false // False if all predicates are false
	}
}

// Nor returns a composed predicate that represents the logical NOR of a list of predicates.
// It evaluates to true only if all predicates evaluate to false for the given value.
func Nor[T any](predicates ...func(T) bool) func(T) bool {
	return func(value T) bool {
		for _, predicate := range predicates {
			if predicate(value) {
				return false // If any predicate evaluates to true, the NOR result is false
			}
		}
		return true // Only returns true if all predicates evaluate to false
	}
}

package function

// And returns a composed predicate that represents the logical AND of a list of predicates.
// It evaluates to true only if all predicates evaluate to true for the given value.
// Play: https://go.dev/play/p/dTBHJMQ0zD2
func And[T any](predicates ...func(T) bool) func(T) bool {
	if len(predicates) < 2 {
		panic("programming error: predicates count must be at least 2")
	}
	return func(value T) bool {
		for _, predicate := range predicates {
			if !predicate(value) {
				return false // Short-circuit on the first false predicate
			}
		}
		return true // True if all predicates are true
	}
}

// Nand returns a composed predicate that represents the logical NAND of a list of predicates.
// It evaluates to false only if all predicates evaluate to true for the given value.
// Play: https://go.dev/play/p/Rb-FdNGpgSO
func Nand[T any](predicates ...func(T) bool) func(T) bool {
	if len(predicates) < 2 {
		panic("programming error: predicates count must be at least 2")
	}
	return func(value T) bool {
		for _, predicate := range predicates {
			if !predicate(value) {
				return true // Short-circuit on the first false predicate
			}
		}
		return false // False if all predicates are true
	}
}

// Negate returns a predicate that represents the logical negation of this predicate.
// Play: https://go.dev/play/p/jbI8BtgFnVE
func Negate[T any](predicate func(T) bool) func(T) bool {
	return func(value T) bool {
		return !predicate(value)
	}
}

// Or returns a composed predicate that represents the logical OR of a list of predicates.
// It evaluates to true if at least one of the predicates evaluates to true for the given value.
// Play: https://go.dev/play/p/LitCIsDFNDA
func Or[T any](predicates ...func(T) bool) func(T) bool {
	if len(predicates) < 2 {
		panic("programming error: predicates count must be at least 2")
	}
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
// Play: https://go.dev/play/p/2KdCoBEOq84
func Nor[T any](predicates ...func(T) bool) func(T) bool {
	if len(predicates) < 2 {
		panic("programming error: predicates count must be at least 2")
	}
	return func(value T) bool {
		for _, predicate := range predicates {
			if predicate(value) {
				return false // If any predicate evaluates to true, the NOR result is false
			}
		}
		return true // Only returns true if all predicates evaluate to false
	}
}

// Xnor returns a composed predicate that represents the logical XNOR of a list of predicates.
// It evaluates to true only if all predicates evaluate to true or false for the given value.
// Play: https://go.dev/play/p/FJxko8SFbqc
func Xnor[T any](predicates ...func(T) bool) func(T) bool {
	if len(predicates) < 2 {
		panic("programming error: predicates count must be at least 2")
	}
	return func(value T) bool {
		trueCount := 0
		for _, predicate := range predicates {
			if predicate(value) {
				trueCount++
			}
		}
		// XNOR is true if either all predicates are true or all are false
		// This is the same as saying trueCount is 0 (all false) or trueCount is len(predicates) (all true)
		return trueCount == 0 || trueCount == len(predicates)
	}
}

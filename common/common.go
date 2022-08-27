package common

// TernaryOperator if true return trueValue else return falseValue
func TernaryOperator[T any](isTrue bool, trueValue T, falseValue T) T {
	if isTrue {
		return trueValue
	} else {
		return falseValue
	}
}

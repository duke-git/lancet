package slice

import "testing"

func TestUniqueByFieldUnreadableInputs(t *testing.T) {
	for _, input := range []any{nil, struct{ hidden string }{"value"}} {
		if _, err := UniqueByField([]any{input}, "hidden"); err == nil {
			t.Errorf("expected an error for %T", input)
		}
	}
}

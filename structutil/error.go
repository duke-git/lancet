package structutil

import "fmt"

func ErrInvalidStruct(v any) error {
	return fmt.Errorf("invalid struct %v", v)
}

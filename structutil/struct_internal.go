package structutil

import "fmt"

func errInvalidStruct(v any) error {
	return fmt.Errorf("invalid struct %v", v)
}

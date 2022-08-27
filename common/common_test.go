package common

import (
	"github.com/duke-git/lancet/v2/internal"
	"testing"
)

func TestTernaryOperator(t *testing.T) {
	assert := internal.NewAssert(t, "TernaryOperator")
	trueValue := "1"
	falseValue := "0"

	assert.Equal(trueValue, TernaryOperator(true, trueValue, falseValue))
}

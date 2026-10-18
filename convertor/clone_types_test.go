package convertor

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestDeepCloneScalarTypes(t *testing.T) {
	assert := internal.NewAssert(t, "TestDeepCloneScalarTypes")
	type label string
	type counter int64
	assert.Equal(label("value"), DeepClone(label("value")))
	assert.Equal(counter(42), DeepClone(counter(42)))
	assert.Equal(uintptr(42), DeepClone(uintptr(42)))
	value := label("pointed")
	cloned := DeepClone(&value)
	assert.Equal(value, *cloned)
	*cloned = "changed"
	assert.Equal(label("pointed"), value)
	var empty any
	assert.Equal(&empty, DeepClone(&empty))
}

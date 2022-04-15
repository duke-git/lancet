package concurrency

import (
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestGenerate(t *testing.T) {
	assert := internal.NewAssert(t, "TestGenerate")

	done := make(chan any)
	defer close(done)

	c := NewChannel()
	intStream := c.Generate(done, 1, 2, 3)

	// for v := range intStream {
	// 	t.Log(v)
	// }
	assert.Equal(1, <-intStream)
	assert.Equal(2, <-intStream)
	assert.Equal(3, <-intStream)
}

func TestRepeat(t *testing.T) {
	// assert := internal.NewAssert(t, "TestRepeat")

	done := make(chan any)
	defer close(done)

	c := NewChannel()
	intStream := c.Repeat(done, 1, 2)

	for v := range intStream {
		t.Log(v)
	}
	// assert.Equal(1, <-intStream)
	// assert.Equal(2, <-intStream)
	// assert.Equal(3, <-intStream)
}

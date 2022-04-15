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
	// 	t.Log(v) //1, 2, 3
	// }
	assert.Equal(1, <-intStream)
	assert.Equal(2, <-intStream)
	assert.Equal(3, <-intStream)
}

func TestRepeat(t *testing.T) {
	assert := internal.NewAssert(t, "TestRepeat")

	done := make(chan any)
	defer close(done)

	c := NewChannel()
	intStream := c.Take(done, c.Repeat(done, 1, 2), 5)

	// for v := range intStream {
	// 	t.Log(v) //1, 2, 1, 2, 1
	// }
	assert.Equal(1, <-intStream)
	assert.Equal(2, <-intStream)
	assert.Equal(1, <-intStream)
	assert.Equal(2, <-intStream)
	assert.Equal(1, <-intStream)
}

func TestTake(t *testing.T) {
	assert := internal.NewAssert(t, "TestRepeat")

	done := make(chan any)
	defer close(done)

	numbers := make(chan any, 5)
	numbers <- 1
	numbers <- 2
	numbers <- 3
	numbers <- 4
	numbers <- 5
	defer close(numbers)

	c := NewChannel()
	intStream := c.Take(done, numbers, 3)

	// for v := range intStream {
	// 	t.Log(v) //1, 2, 3
	// }

	assert.Equal(1, <-intStream)
	assert.Equal(2, <-intStream)
	assert.Equal(3, <-intStream)
}

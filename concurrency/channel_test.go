package concurrency

import (
	"context"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestGenerate(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestGenerate")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	intStream := c.Generate(ctx, 1, 2, 3)

	assert.Equal(1, <-intStream)
	assert.Equal(2, <-intStream)
	assert.Equal(3, <-intStream)
}

func TestRepeat(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRepeat")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	intStream := c.Take(ctx, c.Repeat(ctx, 1, 2), 5)

	assert.Equal(1, <-intStream)
	assert.Equal(2, <-intStream)
	assert.Equal(1, <-intStream)
	assert.Equal(2, <-intStream)
	assert.Equal(1, <-intStream)
}

func TestRepeatFn(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRepeatFn")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fn := func() string {
		s := "a"
		return s
	}
	c := NewChannel[string]()
	dataStream := c.Take(ctx, c.RepeatFn(ctx, fn), 3)

	assert.Equal("a", <-dataStream)
	assert.Equal("a", <-dataStream)
	assert.Equal("a", <-dataStream)
}

func TestTake(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestTake")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numbers := make(chan int, 5)
	numbers <- 1
	numbers <- 2
	numbers <- 3
	numbers <- 4
	numbers <- 5
	defer close(numbers)

	c := NewChannel[int]()
	intStream := c.Take(ctx, numbers, 3)

	assert.Equal(1, <-intStream)
	assert.Equal(2, <-intStream)
	assert.Equal(3, <-intStream)
}

func TestFanIn(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestFanIn")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	channels := make([]<-chan int, 3)

	for i := 0; i < 3; i++ {
		channels[i] = c.Take(ctx, c.Repeat(ctx, i), 3)
	}

	mergedChannel := c.FanIn(ctx, channels...)

	for val := range mergedChannel {
		t.Logf("\t%d\n", val)
	}

	assert.Equal(1, 1)
}

func TestOr(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestOr")

	sig := func(after time.Duration) <-chan any {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()

	c := NewChannel[any]()
	<-c.Or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		sig(5*time.Second),
	)

	assert.Equal(true, time.Since(start).Seconds() < 2)
}

func TestOrDone(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestOrDone")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	intStream := c.Take(ctx, c.Repeat(ctx, 1), 3)

	for val := range c.OrDone(ctx, intStream) {
		assert.Equal(1, val)
	}
}

func TestTee(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestTee")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	inStream := c.Take(ctx, c.Repeat(ctx, 1), 4)

	out1, out2 := c.Tee(ctx, inStream)
	for val := range out1 {
		val1 := val
		val2 := <-out2
		assert.Equal(1, val1)
		assert.Equal(1, val2)
	}
}

func TestBridge(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestBridge")
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	genVals := func() <-chan <-chan int {
		chanStream := make(chan (<-chan int))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan int, 1)
				stream <- i
				m1[i]++
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for val := range c.Bridge(ctx, genVals()) {
		m2[val]++
	}

	for k, v := range m1 {
		assert.Equal(m2[k], v)
	}
}

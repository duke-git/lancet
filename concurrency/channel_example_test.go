package concurrency

import (
	"context"
	"fmt"
	"time"
)

func ExampleChannel_Generate() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	intStream := c.Generate(ctx, 1, 2, 3)

	fmt.Println(<-intStream)
	fmt.Println(<-intStream)
	fmt.Println(<-intStream)

	// Output:
	// 1
	// 2
	// 3
}

func ExampleChannel_Repeat() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	intStream := c.Take(ctx, c.Repeat(ctx, 1, 2), 4)

	for v := range intStream {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 1
	// 2
}

func ExampleChannel_RepeatFn() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fn := func() string {
		return "hello"
	}

	c := NewChannel[string]()
	intStream := c.Take(ctx, c.RepeatFn(ctx, fn), 3)

	for v := range intStream {
		fmt.Println(v)
	}
	// Output:
	// hello
	// hello
	// hello
}

func ExampleChannel_Take() {
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

	for v := range intStream {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleChannel_FanIn() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	channels := make([]<-chan int, 2)

	for i := 0; i < 2; i++ {
		channels[i] = c.Take(ctx, c.Repeat(ctx, i), 2)
	}

	chs := c.FanIn(ctx, channels...)

	for v := range chs {
		fmt.Println(v) //1 1 0 0 or 0 0 1 1
	}

}

func ExampleChannel_Or() {
	sig := func(after time.Duration) <-chan any {
		c := make(chan any)
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
	)

	if time.Since(start).Seconds() < 2 {
		fmt.Println("ok")
	}
	// Output:
	// ok
}

func ExampleChannel_OrDone() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	intStream := c.Take(ctx, c.Repeat(ctx, 1), 3)

	for v := range c.OrDone(ctx, intStream) {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 1
	// 1
}

func ExampleChannel_Tee() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := NewChannel[int]()
	intStream := c.Take(ctx, c.Repeat(ctx, 1), 2)

	ch1, ch2 := c.Tee(ctx, intStream)

	for v := range ch1 {
		fmt.Println(v)
		fmt.Println(<-ch2)
	}
	// Output:
	// 1
	// 1
	// 1
	// 1
}

func ExampleChannel_Bridge() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	c := NewChannel[int]()
	genVals := func() <-chan <-chan int {
		out := make(chan (<-chan int))
		go func() {
			defer close(out)
			for i := 1; i <= 5; i++ {
				stream := make(chan int, 1)
				stream <- i
				m1[i]++
				close(stream)
				out <- stream
			}
		}()
		return out
	}

	for v := range c.Bridge(ctx, genVals()) {
		m2[v]++
	}
	for k, v := range m1 {
		fmt.Println(m2[k] == v)
	}
	// Output:
	// true
	// true
	// true
	// true
	// true
}

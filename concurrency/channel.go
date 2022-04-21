// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package concurrency contain some functions to support concurrent programming. eg, goroutine, channel, async.
package concurrency

import (
	"context"
	"sync"
)

// Channel is a logic object which can generate or manipulate go channel
// all methods of Channel are in the book tilted《Concurrency in Go》
type Channel struct {
}

// NewChannel return a Channel instance
func NewChannel() *Channel {
	return &Channel{}
}

// Generate a data of type any chan, put param `values` into the chan
func (c *Channel) Generate(ctx context.Context, values ...any) <-chan any {
	dataStream := make(chan any)

	go func() {
		defer close(dataStream)

		for _, v := range values {
			select {
			case <-ctx.Done():
				return
			case dataStream <- v:
			}
		}
	}()

	return dataStream
}

// Repeat return a data of type any chan, put param `values` into the chan repeatly,
// until close the `done` chan
func (c *Channel) Repeat(ctx context.Context, values ...any) <-chan any {
	dataStream := make(chan any)

	go func() {
		defer close(dataStream)
		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case dataStream <- v:
				}
			}
		}
	}()
	return dataStream
}

// RepeatFn return a chan, excutes fn repeatly, and put the result into retruned chan
// until close the `done` channel
func (c *Channel) RepeatFn(ctx context.Context, fn func() any) <-chan any {
	dataStream := make(chan any)

	go func() {
		defer close(dataStream)
		for {
			select {
			case <-ctx.Done():
				return
			case dataStream <- fn():
			}
		}
	}()
	return dataStream
}

// Take return a chan whose values are tahken from another chan
func (c *Channel) Take(ctx context.Context, valueStream <-chan any, number int) <-chan any {
	takeStream := make(chan any)

	go func() {
		defer close(takeStream)

		for i := 0; i < number; i++ {
			select {
			case <-ctx.Done():
				return
			case takeStream <- <-valueStream:
			}
		}
	}()

	return takeStream
}

// FanIn merge multiple channels into one channel
func (c *Channel) FanIn(ctx context.Context, channels ...<-chan any) <-chan any {
	var wg sync.WaitGroup
	multiplexedStream := make(chan any)

	multiplex := func(c <-chan any) {
		defer wg.Done()

		for i := range c {
			select {
			case <-ctx.Done():
				return
			case multiplexedStream <- i:
			}
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

// Tee split one chanel into two channels
func (c *Channel) Tee(ctx context.Context, in <-chan any) (<-chan any, <-chan any) {
	out1 := make(chan any)
	out2 := make(chan any)

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range c.OrDone(ctx, in) {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-ctx.Done():
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()

	return out1, out2
}

// Bridge link mutiply channels into one channel
func (c *Channel) Bridge(ctx context.Context, chanStream <-chan <-chan any) <-chan any {
	valStream := make(chan any)

	go func() {
		defer close(valStream)

		for {
			var stream <-chan any
			select {
			case maybeStream, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = maybeStream
			case <-ctx.Done():
				return
			}

			for val := range c.OrDone(ctx, stream) {
				select {
				case valStream <- val:
				case <-ctx.Done():
				}
			}
		}
	}()

	return valStream
}

// Or merge one or more done channels into one done channel, which is closed when any done channel is closed
func (c *Channel) Or(channels ...<-chan any) <-chan any {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan any)

	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-c.Or(append(channels[3:], orDone)...):
			}
		}
	}()

	return orDone
}

// OrDone
func (c *Channel) OrDone(ctx context.Context, channel <-chan any) <-chan any {
	resStream := make(chan any)

	go func() {
		defer close(resStream)

		select {
		case <-ctx.Done():
			return
		case v, ok := <-channel:
			if !ok {
				return
			}
			select {
			case resStream <- v:
			case <-ctx.Done():
			}
		}
	}()

	return resStream
}

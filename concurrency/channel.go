// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package concurrency contain some functions to support concurrent programming. eg, goroutine, channel.
package concurrency

import (
	"context"
	"sync"
)

// Channel is a logic object which can generate or manipulate go channel
// all methods of Channel are in the book tilted《Concurrency in Go》
type Channel[T any] struct {
}

// NewChannel return a Channel instance
func NewChannel[T any]() *Channel[T] {
	return &Channel[T]{}
}

// Generate creates channel, then put values into the channel.
// Play: https://go.dev/play/p/7aB4KyMMp9A
func (c *Channel[T]) Generate(ctx context.Context, values ...T) <-chan T {
	dataStream := make(chan T)

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

// Repeat create channel, put values into the channel repeatly until cancel the context.
// Play: https://go.dev/play/p/k5N_ALVmYjE
func (c *Channel[T]) Repeat(ctx context.Context, values ...T) <-chan T {
	dataStream := make(chan T)

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

// RepeatFn create a channel, excutes fn repeatly, and put the result into the channel
// until close context.
// Play: https://go.dev/play/p/4J1zAWttP85
func (c *Channel[T]) RepeatFn(ctx context.Context, fn func() T) <-chan T {
	dataStream := make(chan T)

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

// Take create a channel whose values are taken from another channel with limit number.
// Play: https://go.dev/play/p/9Utt-1pDr2J
func (c *Channel[T]) Take(ctx context.Context, valueStream <-chan T, number int) <-chan T {
	takeStream := make(chan T)

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

// FanIn merge multiple channels into one channel.
// Play: https://go.dev/play/p/2VYFMexEvTm
func (c *Channel[T]) FanIn(ctx context.Context, channels ...<-chan T) <-chan T {
	out := make(chan T)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(channels))

		for _, c := range channels {
			go func(c <-chan T) {
				defer wg.Done()
				for v := range c {
					select {
					case <-ctx.Done():
						return
					case out <- v:
					}
				}
			}(c)
		}
		wg.Wait()
		close(out)
	}()

	return out
}

// Tee split one chanel into two channels, until cancel the context.
// Play: https://go.dev/play/p/3TQPKnCirrP
func (c *Channel[T]) Tee(ctx context.Context, in <-chan T) (<-chan T, <-chan T) {
	out1 := make(chan T)
	out2 := make(chan T)

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

// Bridge link multiply channels into one channel.
// Play: https://go.dev/play/p/qmWSy1NVF-Y
func (c *Channel[T]) Bridge(ctx context.Context, chanStream <-chan <-chan T) <-chan T {
	valStream := make(chan T)

	go func() {
		defer close(valStream)

		for {
			var stream <-chan T
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			case <-ctx.Done():
				return
			}

			go func() {
				for val := range c.OrDone(ctx, stream) {
					select {
					case valStream <- val:
					case <-ctx.Done():
					}
				}
			}()
		}
	}()

	return valStream
}

// Or read one or more channels into one channel, will close when any readin channel is closed.
// Play: https://go.dev/play/p/Wqz9rwioPww
func (c *Channel[T]) Or(channels ...<-chan T) <-chan T {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan T)

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

// OrDone read a channel into another channel, will close until cancel context.
// Play: https://go.dev/play/p/lm_GoS6aDjo
func (c *Channel[T]) OrDone(ctx context.Context, channel <-chan T) <-chan T {
	valStream := make(chan T)

	go func() {
		defer close(valStream)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-channel:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-ctx.Done():
				}
			}
		}

	}()

	return valStream
}

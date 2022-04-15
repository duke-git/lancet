// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package concurrency contain some functions to support concurrent programming. eg, goroutine, channel, async.
package concurrency

// Channel is a logic object which can generate or manipulate go channel
// all methods of Channel are in the book tilted《Concurrency in Go》
type Channel struct {
}

// NewChannel return a Channel instance
func NewChannel() *Channel {
	return &Channel{}
}

// Generate a data of type any chan, put param `values` into the chan
func (c *Channel) Generate(done <-chan any, values ...any) <-chan any {
	dataStream := make(chan any)

	go func() {
		defer close(dataStream)

		for _, v := range values {
			select {
			case <-done:
				return
			case dataStream <- v:
			}
		}
	}()

	return dataStream
}

// Repeat return a data of type any chan, put param `values` into the chan repeatly,
// until close the `done` chan
func (c *Channel) Repeat(done <-chan any, values ...any) <-chan any {
	dataStream := make(chan any)

	go func() {
		defer close(dataStream)
		for {
			for _, v := range values {
				select {
				case <-done:
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
func (c *Channel) RepeatFn(done <-chan any, fn func() any) <-chan any {
	dataStream := make(chan any)

	go func() {
		defer close(dataStream)
		for {
			select {
			case <-done:
				return
			case dataStream <- fn():
			}
		}
	}()
	return dataStream
}

// Take return a chan whose values are tahken from another chan
func (c *Channel) Take(done <-chan any, valueStream <-chan any, number int) <-chan any {
	takeStream := make(chan any)

	go func() {
		defer close(takeStream)

		for i := 0; i < number; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()

	return takeStream
}

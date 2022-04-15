// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package concurrency contain some functions to support concurrent programming. eg, goroutine, channel, async.
package concurrency

// Channel is a logic object which implemented by go chan
// all methods of channel are in the book tiled《Concurrency in Go》
type Channel struct {
}

// NewChannel return a Channel instance
func NewChannel() *Channel {
	return &Channel{}
}

// Generate a data of type any chan
func (c *Channel) Generate(done <-chan any, datas ...any) <-chan any {
	dataStream := make(chan any)

	go func() {
		defer close(dataStream)

		for _, v := range datas {
			select {
			case <-done:
				return
			case dataStream <- v:
			}
		}
	}()

	return dataStream
}

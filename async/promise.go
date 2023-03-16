// Copyright 2023 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package async contain some features to support async programming. eg, promise, asycn/await, eventbus.
package async

import (
	"errors"
	"fmt"
	"sync"
)

// Promise represents the eventual completion (or failure) of an asynchronous operation and its resulting value.
// see js promise: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise
type Promise[T any] struct {
	runnable func(resolve func(T), reject func(error))
	result   T
	err      error

	pending bool

	mu *sync.Mutex
	wg *sync.WaitGroup
}

// New create a new promise instance.
func New[T any](runnable func(resolve func(T), reject func(error))) *Promise[T] {
	if runnable == nil {
		panic("runnable function should not be nil")
	}

	p := &Promise[T]{
		pending: true,
		mu:      &sync.Mutex{},
		wg:      &sync.WaitGroup{},
	}

	// p.wg.Add(1)
	// go func() {
	// 	defer p.handlePanic()
	// 	runnable(p.resolve, p.reject)
	// }()

	defer p.run()

	return p
}

func (p *Promise[T]) run() {
	p.wg.Add(1)

	go func() {
		defer func() {
			if !p.pending {
				return
			}

			if err := recover(); err != nil {
				p.reject(errors.New(fmt.Sprint(err)))
			}
		}()

		p.runnable(p.resolve, p.reject)
	}()
}

// Resolve returns a Promise that has been resolved with a given value.
func Resolve[T any](resolution T) *Promise[T] {
	return &Promise[T]{
		result:  resolution,
		pending: false,
		mu:      &sync.Mutex{},
		wg:      &sync.WaitGroup{},
	}
}

func (p *Promise[T]) resolve(value T) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.pending {
		return
	}

	p.result = value
	p.pending = false

	p.wg.Done()
}

// Reject returns a Promise that has been rejected with a given error.
func Reject[T any](err error) *Promise[T] {
	return &Promise[T]{
		err:     err,
		pending: false,
		mu:      &sync.Mutex{},
		wg:      &sync.WaitGroup{},
	}
}

func (p *Promise[T]) reject(err error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.pending {
		return
	}

	p.err = err
	p.pending = false

	p.wg.Done()
}

// Then allows chain calls to other promise methods.
func Then[T1, T2 any](promise Promise[T1], resolve1 func(value T1) T2) *Promise[T2] {
	return New(func(resolve2 func(T2), reject func(error)) {
		result, err := promise.Await()
		if err != nil {
			reject(err)
			return
		}
		resolve2(resolve1(result))
	})
}

// Then allows chain calls to other promise methods.
func (p *Promise[T]) Then(resolve func(value T) T) *Promise[T] {
	return New(func(_resolve func(T), reject func(error)) {
		result, err := p.Await()
		if err != nil {
			reject(err)
			return
		}
		_resolve(resolve(result))
	})
}

// Catch allows to chain promises.
func Catch[T any](promise *Promise[T], rejection func(err error) error) *Promise[T] {
	return New(func(resolve func(T), reject func(error)) {
		result, err := promise.Await()
		if err != nil {
			reject(rejection(err))
			return
		}
		resolve(result)
	})
}

// Catch chain an existing promise with an intermediate reject function.
func (p *Promise[T]) Catch(reject func(error) error) *Promise[T] {
	return New(func(resolve func(T), rej func(error)) {
		resutl, err := p.Await()
		if err != nil {
			rej(reject(err))
			return
		}
		resolve(resutl)
	})
}

// Await blocks until the 'runable' to finish execution.
func (p *Promise[T]) Await() (T, error) {
	p.wg.Wait()
	return p.result, p.err
}

// Copyright 2023 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package async contain some featurese to support async programming. eg, promise, asycn/await, eventbus.
package async

import (
	"errors"
	"fmt"
	"sync"

	"github.com/duke-git/lancet/v2/internal"
)

// Promise represents the eventual completion (or failure) of an asynchronous operation and its resulting value.
// ref : chebyrash/promise (https://github.com/chebyrash/promise)
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
		runnable: runnable,
		pending:  true,
		mu:       &sync.Mutex{},
		wg:       &sync.WaitGroup{},
	}

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
func Then[T1, T2 any](promise *Promise[T1], resolve1 func(value T1) T2) *Promise[T2] {
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

type tuple[T1, T2 any] struct {
	_1 T1
	_2 T2
}

// All resolves when all of the promises have resolved, reject immediately upon any of the input promises rejecting.
func All[T any](promises []*Promise[T]) *Promise[[]T] {
	if len(promises) == 0 {
		return nil
	}

	return New(func(resolve func([]T), reject func(error)) {
		valsChan := make(chan tuple[T, int], len(promises))
		errsChan := make(chan error, 1)

		for idx, p := range promises {
			idx := idx
			_ = Then(p, func(data T) T {
				valsChan <- tuple[T, int]{_1: data, _2: idx}
				return data
			})
			_ = Catch(p, func(err error) error {
				errsChan <- err
				return err
			})
		}

		resolutions := make([]T, len(promises))
		for idx := 0; idx < len(promises); idx++ {
			select {
			case val := <-valsChan:
				resolutions[val._2] = val._1
			case err := <-errsChan:
				reject(err)
				return
			}
		}
		resolve(resolutions)
	})
}

// Race will settle the first fullfiled promise among muti promises.
func Race[T any](promises []*Promise[T]) *Promise[T] {
	if len(promises) == 0 {
		return nil
	}

	return New(func(resolve func(T), reject func(error)) {
		valsChan := make(chan T, 1)
		errsChan := make(chan error, 1)

		for _, p := range promises {
			_ = Then(p, func(data T) T {
				valsChan <- data
				return data
			})
			_ = Catch(p, func(err error) error {
				errsChan <- err
				return err
			})
		}

		select {
		case val := <-valsChan:
			resolve(val)
		case err := <-errsChan:
			reject(err)
		}
	})
}

// Any resolves as soon as any of the input's Promises resolve, with the value of the resolved Promise.
// Any rejects if all of the given Promises are rejected with a combination of all errors.
func Any[T any](promises ...*Promise[T]) *Promise[T] {
	if len(promises) == 0 {
		return nil
	}

	return New(func(resolve func(T), reject func(error)) {
		valsChan := make(chan T, 1)
		errsChan := make(chan tuple[error, int], len(promises))

		for idx, p := range promises {
			idx := idx
			_ = Then(p, func(data T) T {
				valsChan <- data
				return data
			})
			_ = Catch(p, func(err error) error {
				errsChan <- tuple[error, int]{_1: err, _2: idx}
				return err
			})
		}

		errs := make([]error, len(promises))
		for idx := 0; idx < len(promises); idx++ {
			select {
			case val := <-valsChan:
				resolve(val)
				return
			case err := <-errsChan:
				errs[err._2] = err._1
			}
		}

		errCombo := errs[0]
		for _, err := range errs[1:] {
			errCombo = internal.JoinError(err)
		}

		reject(errCombo)
	})
}

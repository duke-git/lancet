// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package function implements some functions for control the function execution and some is for functional programming.
package function

import (
	"reflect"
	"sync"
	"time"
)

// After creates a function that invokes func once it's called n or more times
func After(n int, fn interface{}) func(args ...interface{}) []reflect.Value {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	return func(args ...interface{}) []reflect.Value {
		n--
		if n < 1 {
			return unsafeInvokeFunc(fn, args...)
		}
		return nil
	}
}

// Before creates a function that invokes func once it's called less than n times
func Before(n int, fn interface{}) func(args ...interface{}) []reflect.Value {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)
	var res []reflect.Value
	return func(args ...interface{}) []reflect.Value {
		if n > 0 {
			res = unsafeInvokeFunc(fn, args...)
		}
		if n <= 0 {
			fn = nil
		}
		n--
		return res
	}
}

// Fn is for curry function which is func(...interface{}) interface{}
type Fn func(...interface{}) interface{}

// Curry make a curry function
func (f Fn) Curry(i interface{}) func(...interface{}) interface{} {
	return func(values ...interface{}) interface{} {
		v := append([]interface{}{i}, values...)
		return f(v...)
	}
}

// Compose compose the functions from right to left
func Compose(fnList ...func(...interface{}) interface{}) func(...interface{}) interface{} {
	return func(s ...interface{}) interface{} {
		f := fnList[0]
		restFn := fnList[1:]

		if len(fnList) == 1 {
			return f(s...)
		}

		return f(Compose(restFn...)(s...))
	}
}

// Delay make the function execution after delayed time
func Delay(delay time.Duration, fn interface{}, args ...interface{}) {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	time.Sleep(delay)
	invokeFunc(fn, args...)
}

// Debounced creates a debounced function that delays invoking fn until after wait duration have elapsed since the last time the debounced function was invoked.
// Deprecated: Use Debounce function instead.
func Debounced(fn func(), duration time.Duration) func() {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	timer := time.NewTimer(duration)
	timer.Stop()

	go func() {
		for {
			select {
			case <-timer.C:
				go fn()
			}
		}
	}()

	return func() { timer.Reset(duration) }
}

// Debounce creates a debounced version of the provided function.
func Debounce(fn func(), delay time.Duration) (debouncedFn func(), cancelFn func()) {
	var (
		timer *time.Timer
		mu    sync.Mutex
	)

	debouncedFn = func() {
		mu.Lock()
		defer mu.Unlock()

		if timer != nil {
			timer.Stop()
		}

		timer = time.AfterFunc(delay, func() {
			fn()
		})
	}

	cancelFn = func() {
		mu.Lock()
		defer mu.Unlock()

		if timer != nil {
			timer.Stop()
		}
	}

	return debouncedFn, cancelFn
}

// Throttle creates a throttled version of the provided function.
// The returned function guarantees that it will only be invoked at most once per interval.
func Throttle(fn func(), interval time.Duration) func() {
	var (
		timer   *time.Timer
		lastRun time.Time
		mu      sync.Mutex
	)

	return func() {
		mu.Lock()
		defer mu.Unlock()

		now := time.Now()
		if now.Sub(lastRun) >= interval {
			fn()
			lastRun = now
			if timer != nil {
				timer.Stop()
				timer = nil
			}
		} else if timer == nil {
			delay := interval - now.Sub(lastRun)

			timer = time.AfterFunc(delay, func() {
				mu.Lock()
				defer mu.Unlock()

				fn()
				lastRun = time.Now()
				timer = nil
			})
		}
	}
}

// Schedule invoke function every duration time, util close the returned bool chan
func Schedule(d time.Duration, fn interface{}, args ...interface{}) chan bool {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	quit := make(chan bool)
	go func() {
		for {
			unsafeInvokeFunc(fn, args...)
			select {
			case <-time.After(d):
			case <-quit:
				return
			}
		}
	}()

	return quit
}

// Pipeline takes a list of functions and returns a function whose param will be passed into
// the functions one by one.
func Pipeline(funcs ...func(interface{}) interface{}) func(interface{}) interface{} {
	return func(arg interface{}) (result interface{}) {
		result = arg
		for _, fn := range funcs {
			result = fn(result)
		}
		return
	}
}

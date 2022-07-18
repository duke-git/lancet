// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package function implements some functions for control the function execution and some is for functional programming.
package function

import (
	"reflect"
	"time"
)

// After creates a function that invokes func once it's called n or more times
func After(n int, fn any) func(args ...any) []reflect.Value {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	return func(args ...any) []reflect.Value {
		n--
		if n < 1 {
			return unsafeInvokeFunc(fn, args...)
		}
		return nil
	}
}

// Before creates a function that invokes func once it's called less than n times
func Before(n int, fn any) func(args ...any) []reflect.Value {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)
	var result []reflect.Value
	return func(args ...any) []reflect.Value {
		if n > 0 {
			result = unsafeInvokeFunc(fn, args...)
		}
		if n <= 0 {
			fn = nil
		}
		n--
		return result
	}
}

// Fn is for curry function which is func(...any) any
type Fn func(...any) any

// Curry make a curry function
func (f Fn) Curry(i any) func(...any) any {
	return func(values ...any) any {
		v := append([]any{i}, values...)
		return f(v...)
	}
}

// Compose compose the functions from right to left
func Compose(fnList ...func(...any) any) func(...any) any {
	return func(s ...any) any {
		f := fnList[0]
		restFn := fnList[1:]

		if len(fnList) == 1 {
			return f(s...)
		}

		return f(Compose(restFn...)(s...))
	}
}

// Delay make the function execution after delayed time
func Delay(delay time.Duration, fn any, args ...any) {
	// Catch programming error while constructing the closure
	mustBeFunction(fn)

	time.Sleep(delay)
	invokeFunc(fn, args...)
}

// Debounced creates a debounced function that delays invoking fn until after wait duration have elapsed since the last time the debounced function was invoked.
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

// Schedule invoke function every duration time, util close the returned bool chan
func Schedule(d time.Duration, fn any, args ...any) chan bool {
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

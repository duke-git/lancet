// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package function implements some functions for control the function execution and some is for functional programming.
package function

import (
	"reflect"
	"time"
)

// After creates a function that invokes func once it's called n or more times
func After(n int, fn interface{}) func(args ...interface{}) []reflect.Value {
	return func(args ...interface{}) []reflect.Value {
		n--
		if n < 1 {
			return invokeFunc(fn, args...)
		}
		return nil
	}
}

// Before creates a function that invokes func once it's called less than n times
func Before(n int, fn interface{}) func(args ...interface{}) []reflect.Value {
	var res []reflect.Value

	return func(args ...interface{}) []reflect.Value {
		if n > 0 {
			res = invokeFunc(fn, args...)
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

// Curry make a curryed function
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
	time.Sleep(delay)
	invokeFunc(fn, args...)
}

// Schedule invoke function every duration time, util close the returned bool chan
func Schedule(d time.Duration, fn interface{}, args ...interface{}) chan bool {
	quit := make(chan bool)

	go func() {
		for {
			invokeFunc(fn, args...)
			select {
			case <-time.After(d):
			case <-quit:
				return
			}
		}
	}()

	return quit
}

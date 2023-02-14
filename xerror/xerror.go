// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package xerror implements helpers for errors
package xerror

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/duke-git/lancet/v2/random"
)

// XError is to handle error related information.
type XError struct {
	id      string
	message string
	stack   *stack
	cause   error
	values  map[string]any
}

// New creates a new XError with message
func New(format string, args ...any) *XError {
	err := newXError()
	err.message = fmt.Sprintf(format, args...)
	return err
}

// Wrap creates a new XError and add message.
func Wrap(cause error, message ...any) *XError {
	err := newXError()

	if len(message) > 0 {
		var newMsgs []string
		for _, m := range message {
			newMsgs = append(newMsgs, fmt.Sprintf("%v", m))
		}
		err.message = strings.Join(newMsgs, " ")
	}

	err.cause = cause

	return err
}

// Unwrap returns unwrapped XError from err by errors.As. If no XError, returns nil
func Unwrap(err error) *XError {
	var e *XError
	if errors.As(err, &e) {
		return e
	}
	return nil
}

func newXError() *XError {
	id, err := random.UUIdV4()
	if err != nil {
		return nil
	}

	return &XError{
		id:     id,
		stack:  callers(),
		values: make(map[string]any),
	}
}

func (e *XError) copy(dest *XError) {
	dest.message = e.message
	dest.id = e.id
	dest.cause = e.cause

	for k, v := range e.values {
		dest.values[k] = v
	}
}

// Error implements standard error interface.
func (e *XError) Error() string {
	msg := e.message
	cause := e.cause

	if cause == nil {
		return msg
	}

	msg = fmt.Sprintf("%s: %v", msg, cause.Error())

	return msg
}

// Format returns:
// - %v, %s, %q: formatted message
// - %+v: formatted message with stack trace
func (e *XError) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, _ = io.WriteString(s, e.Error())
			e.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

// Wrap creates a new XError and copy message and id to new one.
func (e *XError) Wrap(cause error) *XError {
	err := newXError()
	e.copy(err)
	err.cause = cause
	return err
}

// Unwrap compatible with github.com/pkg/errors
func (e *XError) Unwrap() error {
	return e.cause
}

// With adds key and value related to the error object
func (e *XError) With(key string, value any) *XError {
	e.values[key] = value
	return e
}

// Is checks if target error is XError and Error.id of two errors are matched.
func (e *XError) Is(target error) bool {
	var err *XError

	if errors.As(target, &err) {
		if e.id != "" && e.id == err.id {
			return true
		}
	}

	return e == target
}

// Id sets id to check equality in XError.Is
func (e *XError) Id(id string) *XError {
	e.id = id
	return e
}

// Values returns map of key and value that is set by With. All wrapped xerror.XError key and values will be merged.
// Key and values of wrapped error is overwritten by upper xerror.XError.
func (e *XError) Values() map[string]any {
	var values map[string]any

	if cause := e.Unwrap(); cause != nil {
		if err, ok := cause.(*XError); ok {
			values = err.Values()
		}
	}

	if values == nil {
		values = make(map[string]any)
	}

	for key, value := range e.values {
		values[key] = value
	}

	return values
}

type errInfo struct {
	Message    string         `json:"message"`
	Id         string         `json:"id"`
	StackTrace []*Stack       `json:"stacktrace"`
	Cause      error          `json:"cause"`
	Values     map[string]any `json:"values"`
}

// Info returns information of xerror, which can be printed.
func (e *XError) Info() *errInfo {
	errInfo := &errInfo{
		Message:    e.message,
		Id:         e.id,
		StackTrace: e.Stacks(),
		Cause:      e.cause,
		Values:     make(map[string]any),
	}

	for k, v := range e.values {
		errInfo.Values[k] = v
	}

	return errInfo
}

// TryUnwrap if err is nil then it returns a valid value
// If err is not nil, Unwrap panics with err.
// Play: https://go.dev/play/p/w84d7Mb3Afk
func TryUnwrap[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

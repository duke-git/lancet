package xerror

import (
	"context"
	"fmt"
	"runtime"
)

// TryCatch is a struct to handle try-catch-finally block.
// This implementation is merely a simulation of Java-style try-catch and does not align with Go's error-handling philosophy. It is recommended to use it with caution.
type TryCatch struct {
	ctx         context.Context
	tryFunc     func(ctx context.Context) error
	catchFunc   func(ctx context.Context, err error)
	finallyFunc func(ctx context.Context)
}

// NewTryCatch creates a new TryCatch instance.
func NewTryCatch(ctx context.Context) *TryCatch {
	return &TryCatch{
		ctx: ctx,
	}
}

// Try sets the try function.
func (tc *TryCatch) Try(tryFunc func(ctx context.Context) error) *TryCatch {
	tc.tryFunc = tryFunc
	return tc
}

// Catch sets the catch function.
func (tc *TryCatch) Catch(catchFunc func(ctx context.Context, err error)) *TryCatch {
	tc.catchFunc = catchFunc
	return tc
}

// Finally sets the finally function.
func (tc *TryCatch) Finally(finallyFunc func(ctx context.Context)) *TryCatch {
	tc.finallyFunc = finallyFunc
	return tc
}

// Do executes the try-catch-finally block.
func (tc *TryCatch) Do() {
	defer func() {
		if r := recover(); r != nil {
			if tc.catchFunc != nil {
				err := fmt.Errorf("panic: %v", r)
				tc.catchFunc(tc.ctx, WrapCatchError(err, "Recovered from panic"))
			}
		}

		if tc.finallyFunc != nil {
			tc.finallyFunc(tc.ctx)
		}
	}()

	if tc.ctx.Err() != nil {
		if tc.catchFunc != nil {
			tc.catchFunc(tc.ctx, WrapCatchError(tc.ctx.Err(), "Context cancelled or timed out"))
		}
		return
	}

	if tc.tryFunc != nil {
		if err := tc.tryFunc(tc.ctx); err != nil {
			if tc.catchFunc != nil {
				tc.catchFunc(tc.ctx, WrapCatchError(err, "Error in try block"))
			}
		}
	}
}

// CatchError is an error type to handle try-catch-finally block.
type CatchError struct {
	Msg   string
	File  string
	Line  int
	Cause error
}

// Error returns the error message.
func (e *CatchError) Error() string {
	return fmt.Sprintf("%s at %s:%d - Cause: %v", e.Msg, e.File, e.Line, e.Cause)
}

// WrapCatchError wraps an error with message, file, and line.
func WrapCatchError(err error, msg string) *CatchError {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}
	return &CatchError{
		Msg:   msg,
		File:  file,
		Line:  line,
		Cause: err,
	}
}

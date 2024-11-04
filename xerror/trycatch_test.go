package xerror

import (
	"context"
	"errors"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestTryCatchSuccess(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTryCatchSuccess")

	counter := 0
	calledCatch := false
	calledFinally := false

	tc := NewTryCatch(context.Background())

	tc.Try(func(ctx context.Context) error {
		counter++
		return nil
	}).Catch(func(ctx context.Context, err error) {
		calledCatch = true
		t.Errorf("Catch should not be called")
	}).Finally(func(ctx context.Context) {
		calledFinally = true
	}).Do()

	assert.Equal(1, counter)
	assert.Equal(false, calledCatch)
	assert.Equal(true, calledFinally)
}

func TestTryCatchError(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTryCatchError")

	var catchedError error
	calledFinally := false

	tc := NewTryCatch(context.Background())

	tc.Try(func(ctx context.Context) error {
		return errors.New("error")
	}).Catch(func(ctx context.Context, err error) {
		catchedError = errors.New("catched error")
	}).Finally(func(ctx context.Context) {
		calledFinally = true
	}).Do()

	assert.Equal("catched error", catchedError.Error())
	assert.Equal(true, calledFinally)
}

func TestTryCatchPanic(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTryCatchPanic")

	var catchedError error
	calledFinally := false

	tc := NewTryCatch(context.Background())

	tc.Try(func(ctx context.Context) error {
		panic("panic info")
	}).Catch(func(ctx context.Context, err error) {
		catchedError = errors.New("catched error")
	}).Finally(func(ctx context.Context) {
		calledFinally = true
	}).Do()

	assert.Equal("catched error", catchedError.Error())
	assert.Equal(true, calledFinally)
}

func TestTryCatchContextCancelled(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTryCatchContextCancelled")

	var catchedError error
	calledFinally := false

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	tc := NewTryCatch(ctx)

	tc.Try(func(ctx context.Context) error {
		return nil
	}).Catch(func(ctx context.Context, err error) {
		catchedError = errors.New("catched error")
	}).Finally(func(ctx context.Context) {
		calledFinally = true
	}).Do()

	assert.Equal("catched error", catchedError.Error())
	assert.Equal(true, calledFinally)
}

func TestTryCatchContextTimeout(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTryCatchContextTimeout")

	var catchedError error
	calledFinally := false

	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	tc := NewTryCatch(ctx)

	tc.Try(func(ctx context.Context) error {
		return nil
	}).Catch(func(ctx context.Context, err error) {
		catchedError = errors.New("catched error")
	}).Finally(func(ctx context.Context) {
		calledFinally = true
	}).Do()

	assert.Equal("catched error", catchedError.Error())
	assert.Equal(true, calledFinally)
}

func TestTryCatchContextError(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTryCatchContextError")

	var catchedError error
	calledFinally := false

	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	tc := NewTryCatch(ctx)

	tc.Try(func(ctx context.Context) error {
		return errors.New("error")
	}).Catch(func(ctx context.Context, err error) {
		catchedError = errors.New("catched error")
	}).Finally(func(ctx context.Context) {
		calledFinally = true
	}).Do()

	assert.Equal("catched error", catchedError.Error())
	assert.Equal(true, calledFinally)
}

func TestTryCatchNoCatch(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTryCatchNoCatch")

	calledFinally := false

	tc := NewTryCatch(context.Background())

	tc.Try(func(ctx context.Context) error {
		return errors.New("error")
	}).Finally(func(ctx context.Context) {
		calledFinally = true
	}).Do()

	assert.Equal(true, calledFinally)
}

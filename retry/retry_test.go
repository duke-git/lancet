package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestRetryFailed(t *testing.T) {
	assert := internal.NewAssert(t, "TestRetryFailed")

	var number int
	increaseNumber := func() error {
		number++
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryDuration(time.Microsecond*50))

	assert.IsNotNil(err)
	assert.Equal(DefaultRetryTimes, number)
}

func TestRetrySucceeded(t *testing.T) {
	assert := internal.NewAssert(t, "TestRetrySucceeded")

	var number int
	increaseNumber := func() error {
		number++
		if number == DefaultRetryTimes {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryDuration(time.Microsecond*50))

	assert.IsNil(err)
	assert.Equal(DefaultRetryTimes, number)
}

func TestSetRetryTimes(t *testing.T) {
	assert := internal.NewAssert(t, "TestSetRetryTimes")

	var number int
	increaseNumber := func() error {
		number++
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryDuration(time.Microsecond*50), RetryTimes(3))

	assert.IsNotNil(err)
	assert.Equal(3, number)
}

func TestCancelRetry(t *testing.T) {
	assert := internal.NewAssert(t, "TestCancelRetry")

	ctx, cancel := context.WithCancel(context.TODO())
	var number int
	increaseNumber := func() error {
		number++
		if number > 3 {
			cancel()
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber,
		RetryDuration(time.Microsecond*50),
		Context(ctx),
	)

	assert.IsNotNil(err)
	assert.Equal(4, number)
}

package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestRetryFailed(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryFailed")

	var number int
	increaseNumber := func() error {
		number++
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithLinearBackoff(time.Microsecond*50))

	assert.IsNotNil(err)
	assert.Equal(DefaultRetryTimes, number)
}

func TestRetryShiftExponentialWithJitterFailed(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryShiftExponentialWithJitterFailed")

	var number int
	increaseNumber := func() error {
		number++
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithExponentialWithJitterBackoff(time.Microsecond*50, 2, time.Microsecond*25))

	assert.IsNotNil(err)
	assert.Equal(DefaultRetryTimes, number)
}

func TestRetryExponentialWithJitterFailed(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryExponentialWithJitterFailed")

	var number int
	increaseNumber := func() error {
		number++
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithExponentialWithJitterBackoff(time.Microsecond*50, 3, time.Microsecond*25))

	assert.IsNotNil(err)
	assert.Equal(DefaultRetryTimes, number)
}

func TestRetryWithExponentialSucceeded(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryWithExponentialSucceeded")

	var number int
	increaseNumber := func() error {
		number++
		if number == DefaultRetryTimes {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithExponentialWithJitterBackoff(time.Microsecond*50, 3, time.Microsecond*25))

	assert.IsNil(err)
	assert.Equal(DefaultRetryTimes, number)
}

func TestRetryWithExponentialShiftSucceeded(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryWithExponentialShiftSucceeded")

	var number int
	increaseNumber := func() error {
		number++
		if number == DefaultRetryTimes {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithExponentialWithJitterBackoff(time.Microsecond*50, 4, time.Microsecond*25))

	assert.IsNil(err)
	assert.Equal(DefaultRetryTimes, number)
}

func TestRetrySucceeded(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetrySucceeded")

	var number int
	increaseNumber := func() error {
		number++
		if number == DefaultRetryTimes {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithLinearBackoff(time.Microsecond*50))

	assert.IsNil(err)
	assert.Equal(DefaultRetryTimes, number)
}

func TestRetryOneShotSucceeded(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryOneShotSucceeded")

	var number int
	increaseNumber := func() error {
		number++
		return nil
	}

	err := Retry(increaseNumber, RetryWithLinearBackoff(time.Microsecond*50))

	assert.IsNil(err)
	assert.Equal(1, number)
}

func TestRetryWitCustomBackoffOneShotSucceeded(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryWitCustomBackoffOneShotSucceeded")

	var number int
	increaseNumber := func() error {
		number++
		if number == DefaultRetryTimes {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithCustomBackoff(&TestCustomBackoffStrategy{interval: time.Microsecond * 50}))

	assert.IsNil(err)
	assert.Equal(5, number)
}

type TestCustomBackoffStrategy struct {
	interval time.Duration
}

func (c *TestCustomBackoffStrategy) CalculateInterval() time.Duration {
	return c.interval + 1
}

func TestRetryWithExponentialWithJitterBackoffShiftOneShotSucceeded(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryWithExponentialWithJitterBackoffShiftOneShotSucceeded")

	var number int
	increaseNumber := func() error {
		number++
		return nil
	}

	err := Retry(increaseNumber, RetryWithExponentialWithJitterBackoff(time.Microsecond*50, 2, time.Microsecond*25))

	assert.IsNil(err)
	assert.Equal(1, number)
}

func TestRetryWithExponentialWithJitterBackoffOneShotSucceeded(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryWithExponentialWithJitterBackoffOneShotSucceeded")

	var number int
	increaseNumber := func() error {
		number++
		return nil
	}

	err := Retry(increaseNumber, RetryWithExponentialWithJitterBackoff(time.Microsecond*50, 3, time.Microsecond*25))

	assert.IsNil(err)
	assert.Equal(1, number)
}

func TestRetryWithExponentialWithJitterBackoffNoJitterOneShotSucceeded(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRetryWithExponentialWithJitterBackoffNoJitterOneShotSucceeded")

	var number int
	increaseNumber := func() error {
		number++
		return nil
	}

	err := Retry(increaseNumber, RetryWithExponentialWithJitterBackoff(time.Microsecond*50, 3, 0))

	assert.IsNil(err)
	assert.Equal(1, number)
}

func TestSetRetryTimes(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSetRetryTimes")

	var number int
	increaseNumber := func() error {
		number++
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithLinearBackoff(time.Microsecond*50), RetryTimes(3))

	assert.IsNotNil(err)
	assert.Equal(3, number)
}

func TestCancelRetry(t *testing.T) {
	t.Parallel()

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
		RetryWithLinearBackoff(time.Microsecond*50),
		Context(ctx),
	)

	assert.IsNotNil(err)
	assert.Equal(4, number)
}

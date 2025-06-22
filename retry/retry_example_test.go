package retry

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func ExampleContext() {
	ctx, cancel := context.WithCancel(context.TODO())

	number := 0
	increaseNumber := func() error {
		number++
		if number > 3 {
			cancel()
		}
		return errors.New("error occurs")
	}

	Retry(increaseNumber,
		RetryWithLinearBackoff(time.Microsecond*50),
		Context(ctx),
	)

	fmt.Println(number)

	// Output:
	// 4
}

func ExampleRetryWithLinearBackoff() {
	number := 0
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithLinearBackoff(time.Microsecond*50))
	if err != nil {
		return
	}

	fmt.Println(number)

	// Output:
	// 3
}

type ExampleCustomBackoffStrategy struct {
	interval time.Duration
}

func (c *ExampleCustomBackoffStrategy) CalculateInterval() time.Duration {
	return c.interval + 1
}

func ExampleRetryWithCustomBackoff() {
	number := 0
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithCustomBackoff(&ExampleCustomBackoffStrategy{interval: time.Microsecond * 50}))
	if err != nil {
		return
	}

	fmt.Println(number)

	// Output:
	// 3
}

func ExampleRetryWithExponentialWithJitterBackoff() {
	number := 0
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithExponentialWithJitterBackoff(time.Microsecond*50, 2, time.Microsecond*25))
	if err != nil {
		return
	}

	fmt.Println(number)

	// Output:
	// 3
}

func ExampleRetryTimes() {
	number := 0

	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryTimes(2))
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// function retry.ExampleRetryTimes.func1 run failed after 2 times retry, last error: error occurs
}

func ExampleRetry() {
	number := 0
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryWithLinearBackoff(time.Microsecond*50))
	if err != nil {
		return
	}

	fmt.Println(number)

	// Output:
	// 3
}

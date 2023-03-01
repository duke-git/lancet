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

	err := Retry(increaseNumber,
		RetryDuration(time.Microsecond*50),
		Context(ctx),
	)

	if err != nil {
		return
	}

	fmt.Println(number)

	// Output:
	// 4
}

func ExampleRetryDuration() {
	number := 0
	increaseNumber := func() error {
		number++
		if number == 3 {
			return nil
		}
		return errors.New("error occurs")
	}

	err := Retry(increaseNumber, RetryDuration(time.Microsecond*50))
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
	// function retry.ExampleRetryTimes.func1 run failed after 2 times retry
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

	err := Retry(increaseNumber, RetryDuration(time.Microsecond*50))
	if err != nil {
		return
	}

	fmt.Println(number)

	// Output:
	// 3
}

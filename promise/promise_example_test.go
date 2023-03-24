package promise

import (
	"errors"
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func ExampleNew() {
	p := New(func(resolve func(string), reject func(error)) {
		resolve("hello")
	})

	val, err := p.Await()
	if err != nil {
		return
	}

	fmt.Println(val)

	// Output:
	// hello
}

func ExampleThen() {
	p1 := New(func(resolve func(string), reject func(error)) {
		resolve("hello ")
	})

	p2 := Then(p1, func(s string) string {
		return s + "world"
	})

	result, err := p2.Await()
	if err != nil {
		return
	}

	fmt.Println(result)

	// Output:
	// hello world
}

func ExamplePromise_Then() {
	p1 := New(func(resolve func(string), reject func(error)) {
		resolve("hello ")
	})

	p2 := p1.Then(func(s string) string {
		return s + "world"
	})

	result, err := p2.Await()
	if err != nil {
		return
	}

	fmt.Println(result)

	// Output:
	// hello world
}

func ExampleCatch() {
	p1 := New(func(resolve func(string), reject func(error)) {
		err := errors.New("error1")
		reject(err)
	})

	p2 := Catch(p1, func(err error) error {
		e := errors.New("error2")
		return internal.JoinError(err, e)
	})

	_, err := p1.Await()

	fmt.Println(err.Error())

	result2, err := p2.Await()

	fmt.Println(result2)
	fmt.Println(err.Error())

	// Output:
	// error1
	//
	// error1
	// error2
}

func ExamplePromise_Catch() {
	p1 := New(func(resolve func(string), reject func(error)) {
		err := errors.New("error1")
		reject(err)
	})

	p2 := p1.Catch(func(err error) error {
		e := errors.New("error2")
		return internal.JoinError(err, e)
	})

	_, err := p1.Await()

	fmt.Println(err.Error())

	result2, err := p2.Await()

	fmt.Println(result2)
	fmt.Println(err.Error())

	// Output:
	// error1
	//
	// error1
	// error2
}

func ExampleAll() {
	p1 := New(func(resolve func(string), reject func(error)) {
		resolve("a")
	})
	p2 := New(func(resolve func(string), reject func(error)) {
		resolve("b")
	})
	p3 := New(func(resolve func(string), reject func(error)) {
		resolve("c")
	})

	pms := []*Promise[string]{p1, p2, p3}
	p := All(pms)

	result, err := p.Await()
	if err != nil {
		return
	}

	fmt.Println(result)

	// Output:
	// [a b c]
}

func ExampleAny() {
	p1 := New(func(resolve func(string), reject func(error)) {
		time.Sleep(time.Millisecond * 250)
		resolve("fast")
	})
	p2 := New(func(resolve func(string), reject func(error)) {
		time.Sleep(time.Millisecond * 500)
		resolve("slow")
	})
	p3 := New(func(resolve func(string), reject func(error)) {
		reject(errors.New("error"))
	})

	pms := []*Promise[string]{p1, p2, p3}
	p := Any(pms)

	result, err := p.Await()
	if err != nil {
		return
	}

	fmt.Println(result)

	// Output:
	// fast
}

func ExampleRace() {
	p1 := New(func(resolve func(string), reject func(error)) {
		time.Sleep(time.Millisecond * 100)
		resolve("fast")
	})
	p2 := New(func(resolve func(string), reject func(error)) {
		time.Sleep(time.Millisecond * 300)
		resolve("slow")
	})

	pms := []*Promise[string]{p1, p2}
	p := Race(pms)

	result, err := p.Await()
	if err != nil {
		return
	}

	fmt.Println(result)

	// Output:
	// fast
}

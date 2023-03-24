package promise

import (
	"errors"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestResolve(t *testing.T) {
	assert := internal.NewAssert(t, "TestResolve")

	p := Resolve("abc")

	assert.Equal("abc", p.result)
	assert.Equal(false, p.pending)
}

func TestReject(t *testing.T) {
	assert := internal.NewAssert(t, "TestReject")

	err := errors.New("error")
	p := Reject[string](err)

	assert.Equal("error", p.err.Error())
	assert.Equal(false, p.pending)
}

func TestThen(t *testing.T) {
	assert := internal.NewAssert(t, "TestThen")

	p1 := New(func(resolve func(string), reject func(error)) {
		resolve("abc")
	})

	p2 := Then(p1, func(data string) string {
		return data + "de"
	})

	val, err := p1.Await()
	assert.IsNil(err)
	assert.Equal("abc", val)

	val, err = p2.Await()
	assert.IsNil(err)
	assert.Equal("abcde", val)
}

func TestPromise_Then(t *testing.T) {
	assert := internal.NewAssert(t, "TestPromise_Then")

	p1 := New(func(resolve func(int), reject func(error)) {
		resolve(1)
	})

	p2 := p1.Then(func(n int) int {
		return n + 2
	})

	val, err := p1.Await()
	assert.IsNil(err)
	assert.Equal(1, val)

	val, err = p2.Await()
	assert.IsNil(err)
	assert.Equal(3, val)
}

func TestCatch(t *testing.T) {
	assert := internal.NewAssert(t, "TestCatch")

	p1 := New(func(resolve func(string), reject func(error)) {
		err := errors.New("error1")
		reject(err)
	})

	p2 := Catch(p1, func(err error) error {
		e := errors.New("error2")
		return internal.JoinError(err, e)
	})

	val, err := p1.Await()
	assert.Equal("", val)
	assert.IsNotNil(err)
	assert.Equal("error1", err.Error())

	val, err = p2.Await()

	assert.Equal("", val)
	assert.IsNotNil(err)
	assert.Equal("error1\nerror2", err.Error())
}

func TestPromise_Catch(t *testing.T) {
	assert := internal.NewAssert(t, "TestPromise_Catch")

	p1 := New(func(resolve func(string), reject func(error)) {
		err := errors.New("error1")
		reject(err)
	})

	p2 := p1.Catch(func(err error) error {
		e := errors.New("error2")
		return internal.JoinError(err, e)
	})

	val, err := p1.Await()
	assert.Equal("", val)
	assert.IsNotNil(err)
	assert.Equal("error1", err.Error())

	val, err = p2.Await()

	assert.Equal("", val)
	assert.IsNotNil(err)
	assert.Equal("error1\nerror2", err.Error())
}

func TestAll(t *testing.T) {
	assert := internal.NewAssert(t, "TestPromise_All")

	t.Run("AllPromisesFullfilled", func(_ *testing.T) {
		p1 := New(func(resolve func(string), reject func(error)) {
			resolve("a")
		})
		p2 := New(func(resolve func(string), reject func(error)) {
			resolve("b")
		})
		p3 := New(func(resolve func(string), reject func(error)) {
			resolve("c")
		})

		p := All([]*Promise[string]{p1, p2, p3})

		val, err := p.Await()
		assert.Equal([]string{"a", "b", "c"}, val)
		assert.IsNil(err)
	})

	t.Run("EmptyPromises", func(_ *testing.T) {
		var empty = []*Promise[any]{}
		p := All(empty)
		assert.IsNil(p)
	})

	t.Run("PromisesContainRejected", func(_ *testing.T) {
		p1 := New(func(resolve func(string), reject func(error)) {
			resolve("a")
		})
		p2 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error1"))
		})
		p3 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error2"))
		})

		p := All([]*Promise[string]{p1, p2, p3})

		_, err := p.Await()

		assert.IsNotNil(err)
		// assert.Equal("error1", err.Error())
	})

	t.Run("PromisesOnlyRejected", func(_ *testing.T) {
		p1 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error1"))

		})
		p2 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error2"))
		})
		p3 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error3"))
		})

		p := All([]*Promise[string]{p1, p2, p3})

		_, err := p.Await()

		assert.IsNotNil(err)
		// assert.Equal("error1", err.Error())
	})

}

func TestAny(t *testing.T) {
	assert := internal.NewAssert(t, "TestPromise_Any")

	t.Run("AnyFullfilled", func(_ *testing.T) {
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

		p := Any([]*Promise[string]{p1, p2, p3})

		val, err := p.Await()
		assert.Equal("fast", val)
		assert.IsNil(err)
	})

	t.Run("EmptyPromises", func(_ *testing.T) {
		var empty = []*Promise[any]{}
		p := Any(empty)
		assert.IsNil(p)
	})

	t.Run("OnlyRejected", func(_ *testing.T) {
		p1 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error1"))
		})
		p2 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error2"))
		})

		p := Any([]*Promise[string]{p1, p2})

		_, err := p.Await()

		assert.IsNotNil(err)
		// assert.Equal("error1", err.Error())
	})

}

func TestRace(t *testing.T) {
	assert := internal.NewAssert(t, "TestPromise_Race")

	t.Run("PromisesFullfilled", func(_ *testing.T) {
		p1 := New(func(resolve func(string), reject func(error)) {
			time.Sleep(time.Millisecond * 100)
			resolve("a")
		})
		p2 := New(func(resolve func(string), reject func(error)) {
			time.Sleep(time.Millisecond * 300)
			resolve("b")
		})

		p := Race([]*Promise[string]{p1, p2})

		val, err := p.Await()
		assert.Equal("a", val)
		assert.IsNil(err)
	})

	t.Run("EmptyPromises", func(_ *testing.T) {
		var empty = []*Promise[any]{}
		p := Race(empty)
		assert.IsNil(p)
	})

	t.Run("PromisesContainRejected", func(_ *testing.T) {
		p1 := New(func(resolve func(string), reject func(error)) {
			time.Sleep(time.Millisecond * 100)
			resolve("a")
		})
		p2 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error1"))
		})
		p3 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error2"))
		})

		p := Race([]*Promise[string]{p1, p2, p3})

		val, err := p.Await()

		assert.IsNotNil(err)
		// assert.Equal("error1", err.Error())
		assert.Equal("", val)
	})

	t.Run("PromisesOnlyRejected", func(_ *testing.T) {
		p1 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error1"))

		})
		p2 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error2"))
		})
		p3 := New(func(resolve func(string), reject func(error)) {
			reject(errors.New("error3"))
		})

		p := Race([]*Promise[string]{p1, p2, p3})

		_, err := p.Await()

		assert.IsNotNil(err)
		// assert.Equal("error1", err.Error())
	})

}

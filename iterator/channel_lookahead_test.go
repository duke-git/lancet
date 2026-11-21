package iterator

import (
	"reflect"
	"testing"
)

func TestChannelIteratorLookahead(t *testing.T) {
	for _, tc := range []struct {
		name   string
		values []int
	}{
		{"empty", nil},
		{"one-zero", []int{0}},
		{"several", []int{3, 0, -2}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ch := make(chan int, len(tc.values))
			for _, value := range tc.values {
				ch <- value
			}
			close(ch)
			iter := FromChannel(ch)
			for i, want := range tc.values {
				for peek := 0; peek < 3; peek++ {
					if !iter.HasNext() {
						t.Fatalf("HasNext is false before value %d (peek %d)", i, peek)
					}
				}
				if got, ok := iter.Next(); !ok || got != want {
					t.Fatalf("Next = (%d, %v), want (%d, true)", got, ok, want)
				}
			}
			for i := 0; i < 3; i++ {
				if iter.HasNext() {
					t.Fatal("HasNext is true after the channel is exhausted")
				}
				if got, ok := iter.Next(); ok || got != 0 {
					t.Fatalf("exhausted Next = (%d, %v), want (0, false)", got, ok)
				}
			}
		})
	}
}

func TestChannelIteratorMixedNextAndLookahead(t *testing.T) {
	ch := make(chan *int, 3)
	first, last := 1, 3
	ch <- &first
	ch <- nil
	ch <- &last
	close(ch)
	iter := FromChannel(ch)
	if got, ok := iter.Next(); !ok || got != &first {
		t.Fatal("direct Next did not return the first pointer")
	}
	if !iter.HasNext() || !iter.HasNext() {
		t.Fatal("lookahead did not preserve the nil channel value")
	}
	if got, ok := iter.Next(); !ok || got != nil {
		t.Fatal("a nil value was confused with an exhausted channel")
	}
	if got, ok := iter.Next(); !ok || got != &last {
		t.Fatal("direct Next after lookahead did not return the last pointer")
	}
	if iter.HasNext() {
		t.Fatal("lookahead did not detect the closed channel")
	}
}

func TestChannelIteratorUnbufferedLookahead(t *testing.T) {
	ch := make(chan int)
	stop, done := make(chan struct{}), make(chan struct{})
	go func() {
		defer close(done)
		defer close(ch)
		for _, value := range []int{4, 0, 6} {
			select {
			case ch <- value:
			case <-stop:
				return
			}
		}
	}()
	defer func() { close(stop); <-done }()
	iter := FromChannel(ch)
	var got []int
	// Bound the loop even if an implementation reports true after closure.
	for i := 0; i < 4; i++ {
		if !iter.HasNext() {
			break
		}
		if !iter.HasNext() {
			t.Fatal("repeated lookahead lost an available value")
		}
		value, ok := iter.Next()
		if !ok {
			t.Fatal("HasNext reported a value but Next reported exhaustion")
		}
		got = append(got, value)
	}
	if !reflect.DeepEqual(got, []int{4, 0, 6}) {
		t.Fatalf("iteration = %v, want [4 0 6]", got)
	}
}

func TestChannelIteratorLookaheadComposesWithTakeAndFilter(t *testing.T) {
	ch := make(chan int, 5)
	for _, value := range []int{1, 2, 3, 4, 5} {
		ch <- value
	}
	close(ch)
	iter := Take(Filter[int](FromChannel(ch), func(n int) bool { return n%2 == 0 }), 2)
	var got []int
	for i := 0; i < 3 && iter.HasNext(); i++ {
		if !iter.HasNext() {
			t.Fatal("repeated composed lookahead lost an available value")
		}
		value, ok := iter.Next()
		if !ok {
			t.Fatal("composed Next reported exhaustion after successful lookahead")
		}
		got = append(got, value)
	}
	if !reflect.DeepEqual(got, []int{2, 4}) {
		t.Fatalf("composed iteration = %v, want [2 4]", got)
	}
}

func TestChannelIteratorLookaheadComposesWithMapAndTake(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	iter := Take(Map[int, int](FromChannel(ch), func(n int) int { return n * 2 }), 2)
	var got []int
	for i := 0; i < 3 && iter.HasNext(); i++ {
		if !iter.HasNext() {
			t.Fatal("repeated Map/Take lookahead lost a value")
		}
		value, ok := iter.Next()
		if !ok {
			t.Fatal("Map/Take exhausted after successful lookahead")
		}
		got = append(got, value)
	}
	if !reflect.DeepEqual(got, []int{4, 6}) {
		t.Fatalf("Map/Take iteration = %v, want [4 6]", got)
	}
}

type lookaheadProbe struct{ called bool }

func (p *lookaheadProbe) HasNext() bool     { p.called = true; return false }
func (p *lookaheadProbe) Next() (int, bool) { return 0, false }

func TestJoinLookaheadStopsAtAvailableIterator(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 7
	close(ch)
	later := &lookaheadProbe{}
	iter := Join[int](FromChannel(ch), later)
	if !iter.HasNext() || !iter.HasNext() {
		t.Fatal("Join did not report the first channel value")
	}
	if later.called {
		t.Fatal("Join peeked a later iterator while the current one has a value")
	}
	if got, ok := iter.Next(); !ok || got != 7 {
		t.Fatalf("Join Next = (%d, %v), want (7, true)", got, ok)
	}
	if iter.HasNext() || !later.called {
		t.Fatal("Join did not check the later iterator after the channel was exhausted")
	}
}

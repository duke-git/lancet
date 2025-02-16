package compare

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestEqual(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEqual")

	tests := []struct {
		left  any
		right any
		want  bool
	}{
		{1, 1, true},
		{int64(1), int64(1), true},
		{"a", "a", true},
		{true, true, true},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"}, true},
		{1, 2, false},
		{1, int64(1), false},
		{"a", "b", false},
		{true, false, false},
		{[]int{1, 2}, []int{1, 2, 3}, false},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a"}, false},
		// {time.Now(), time.Now(), true},
		// {time.Now(), time.Now().Add(time.Second), false},
		{[]byte("hello"), []byte("hello"), true},
		{[]byte("hello"), []byte("world"), false},
		{json.Number("123"), json.Number("123"), true},
		{json.Number("123"), json.Number("124"), false},

		{big.NewInt(123), big.NewInt(123), true},
	}

	for _, tt := range tests {
		assert.Equal(tt.want, Equal(tt.left, tt.right))
	}

}

func TestEqualValue(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestEqualValue")

	tests := []struct {
		left  any
		right any
		want  bool
	}{
		{1, 1, true},
		{int64(1), int64(1), true},
		{"a", "a", true},
		{true, true, true},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{map[int]string{1: "a", 2: "b"}, map[int]string{1: "a", 2: "b"}, true},
		{1, 2, false},
		{1, int64(1), true},
		{"a", "b", false},
		{true, false, false},
		{[]int{1, 2}, []int{1, 2, 3}, false},
	}

	for _, tt := range tests {
		assert.Equal(tt.want, EqualValue(tt.left, tt.right))
	}
}

func TestLessThan(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestLessThan")

	tests := []struct {
		left  any
		right any
		want  bool
	}{
		{1, 2, true},
		{1.1, 2.2, true},
		{"a", "b", true},
		{time.Now(), time.Now().Add(time.Second), true},
		{[]byte("hello1"), []byte("hello2"), true},
		{json.Number("123"), json.Number("124"), true},
		{645680099112988673, 645680099112988675, true},
		{1, 1, false},
		{1, int64(1), false},
	}

	for _, tt := range tests {
		assert.Equal(tt.want, LessThan(tt.left, tt.right))
	}
}

func TestGreaterThan(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestGreaterThan")

	tests := []struct {
		left  any
		right any
		want  bool
	}{
		{2, 1, true},
		{2.2, 1.1, true},
		{"b", "a", true},
		{time.Now().Add(time.Second), time.Now(), true},
		{[]byte("hello2"), []byte("hello1"), true},
		{json.Number("124"), json.Number("123"), true},
		{645680099112988675, 645680099112988673, true},
		{1, 1, false},
		{1, int64(1), false},
	}

	for _, tt := range tests {
		assert.Equal(tt.want, GreaterThan(tt.left, tt.right))
	}

}

func TestLessOrEqual(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestLessOrEqual")

	tests := []struct {
		left  any
		right any
		want  bool
	}{
		{1, 2, true},
		{1, 1, true},
		{1.1, 2.2, true},
		{"a", "b", true},
		{time.Now(), time.Now().Add(time.Second), true},
		{[]byte("hello1"), []byte("hello2"), true},
		{json.Number("123"), json.Number("124"), true},
		{645680099112988673, 645680099112988675, true},
		{2, 1, false},
		{1, int64(2), false},
	}

	for _, tt := range tests {
		assert.Equal(tt.want, LessOrEqual(tt.left, tt.right))
	}
}

func TestGreaterOrEqual(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestGreaterThan")

	tests := []struct {
		left  any
		right any
		want  bool
	}{
		{2, 1, true},
		{1, 1, true},
		{2.2, 1.1, true},
		{"b", "b", true},
		{time.Now().Add(time.Second), time.Now(), true},
		{[]byte("hello2"), []byte("hello1"), true},
		{json.Number("124"), json.Number("123"), true},
		{645680099112988675, 645680099112988673, true},
		{1, 2, false},
		{int64(2), 1, false},
		{"b", "c", false},
	}

	for _, tt := range tests {
		assert.Equal(tt.want, GreaterOrEqual(tt.left, tt.right))
	}
}

func TestInDelta(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestInDelta")

	tests := []struct {
		left  float64
		right float64
		delta float64
		want  bool
	}{
		{1, 1, 0, true},
		{1, 2, 0, false},
		{2.0 / 3.0, 0.66667, 0.001, true},
		{2.0 / 3.0, 0.0, 0.001, false},
		{float64(74.96) - float64(20.48), 54.48, 0, false},
		{float64(74.96) - float64(20.48), 54.48, 1e-14, true},
		{float64(float32(80.45)), float64(80.45), 0, false},
		{float64(float32(80.45)), float64(80.45), 1e-5, true},
	}

	for _, tt := range tests {
		assert.Equal(tt.want, InDelta(tt.left, tt.right, tt.delta))
	}
}

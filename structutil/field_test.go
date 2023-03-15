package structutil

import (
	"github.com/duke-git/lancet/v2/internal"
	"reflect"
	"testing"
)

func TestField_Tag(t *testing.T) {
	assert := internal.NewAssert(t, "TestField_Tag")

	type Parent struct {
		Name string `json:"name,omitempty"`
	}
	p1 := &Parent{"111"}

	s := New(p1)
	n, _ := s.Field("Name")
	tag := n.Tag()
	assert.Equal("name", tag.Name)
	assert.Equal(true, tag.HasOption("omitempty"))
}

func TestField_Value(t *testing.T) {
	assert := internal.NewAssert(t, "TestField_Value")

	type Parent struct {
		Name string `json:"name,omitempty"`
	}
	p1 := &Parent{"111"}

	s := New(p1)
	n, _ := s.Field("Name")

	assert.Equal("111", n.Value())
}

func TestField_IsEmbedded(t *testing.T) {
	assert := internal.NewAssert(t, "TestField_IsEmbedded")
	type Parent struct {
		Name string
	}
	type Child struct {
		Parent
		Age int
	}
	c1 := &Child{}
	c1.Name = "111"
	c1.Age = 11

	s := New(c1)
	n, _ := s.Field("Name")
	a, _ := s.Field("Age")
	assert.Equal(true, n.IsEmbedded())
	assert.Equal(false, a.IsEmbedded())
}

func TestField_IsExported(t *testing.T) {
	assert := internal.NewAssert(t, "TestField_IsEmbedded")

	type Parent struct {
		Name string
		age  int
	}
	p1 := &Parent{Name: "11", age: 11}
	s := New(p1)
	n, _ := s.Field("Name")
	a, _ := s.Field("age")
	assert.Equal(true, n.IsExported())
	assert.Equal(false, a.IsExported())
}

func TestField_IsZero(t *testing.T) {
	assert := internal.NewAssert(t, "TestField_IsZero")

	type Parent struct {
		Name string
		Age  int
	}
	p1 := &Parent{Age: 11}
	s := New(p1)
	n, _ := s.Field("Name")
	a, _ := s.Field("Age")
	assert.Equal(true, n.IsZero())
	assert.Equal(false, a.IsZero())
}

func TestField_Name(t *testing.T) {
	assert := internal.NewAssert(t, "TestField_Name")

	type Parent struct {
		Name string
		Age  int
	}
	p1 := &Parent{Age: 11}
	s := New(p1)
	n, _ := s.Field("Name")
	a, _ := s.Field("Age")

	assert.Equal("Name", n.Name())
	assert.Equal("Age", a.Name())
}

func TestField_Kind(t *testing.T) {
	assert := internal.NewAssert(t, "TestField_Kind")

	type Parent struct {
		Name string
		Age  int
	}
	p1 := &Parent{Age: 11}
	s := New(p1)
	n, _ := s.Field("Name")
	a, _ := s.Field("Age")

	assert.Equal(reflect.String, n.Kind())
	assert.Equal(reflect.Int, a.Kind())
}

func TestField_IsSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestField_IsSlice")

	type Parent struct {
		Name string
		arr  []int
	}

	p1 := &Parent{arr: []int{1, 2, 3}}
	s := New(p1)
	a, _ := s.Field("arr")

	assert.Equal(true, a.IsSlice())
}

func TestField_MapValue(t *testing.T) {
	assert := internal.NewAssert(t, "TestField_MapValue")

	t.Run("nested struct", func(t *testing.T) {
		type Child struct {
			Name string `json:"name"`
		}
		type Parent struct {
			Name  string `json:"name"`
			Child *Child `json:"child"`
		}

		c1 := &Child{"11-1"}
		p1 := &Parent{
			Name:  "11",
			Child: c1,
		}

		s := New(p1)
		f, ok := s.Field("Child")
		val := f.MapValue(f.Value())

		assert.Equal(true, ok)
		assert.Equal(map[string]any{"name": "11-1"}, val)
	})

	t.Run("nested ptr struct", func(t *testing.T) {
		type Child struct {
			Name string `json:"name"`
		}
		type Parent struct {
			Name  string `json:"name"`
			Child any    `json:"child"`
		}
		c1 := &Child{"11-1"}
		c2 := &c1
		c3 := &c2
		p1 := &Parent{
			Name:  "11",
			Child: c3,
		}

		s := New(p1)
		f, ok := s.Field("Child")
		val := f.MapValue(f.Value())

		assert.Equal(true, ok)
		assert.Equal(map[string]any{"name": "11-1"}, val)
	})

	t.Run("nested array", func(t *testing.T) {
		type Parent struct {
			Name  string `json:"name"`
			Child []int  `json:"child"`
		}

		p1 := &Parent{
			Name:  "11",
			Child: []int{1, 2, 3},
		}

		s := New(p1)
		f, ok := s.Field("Child")
		val := f.MapValue(f.Value())

		assert.Equal(true, ok)
		assert.Equal([]int{1, 2, 3}, val)
	})

	t.Run("nested array in struct", func(t *testing.T) {
		type Child struct {
			Name string `json:"name"`
		}
		type Parent struct {
			Name  string   `json:"name"`
			Child []*Child `json:"child"`
		}

		c1 := &Child{"11-1"}
		c2 := &Child{"11-2"}

		p1 := &Parent{
			Name:  "11",
			Child: []*Child{c1, c2},
		}

		s := New(p1)
		f, ok := s.Field("Child")
		val := f.MapValue(f.Value())

		assert.Equal(true, ok)
		arr := []any{map[string]any{"name": "11-1"}, map[string]any{"name": "11-2"}}
		assert.Equal(arr, val)
	})

	t.Run("nested ptr array in struct", func(t *testing.T) {
		type Child struct {
			Name string `json:"name"`
		}
		type Parent struct {
			Name  string    `json:"name"`
			Child *[]*Child `json:"child"`
		}

		c1 := &Child{"11-1"}
		c2 := &Child{"11-2"}

		p1 := &Parent{
			Name:  "11",
			Child: &[]*Child{c1, c2},
		}

		s := New(p1)
		f, ok := s.Field("Child")
		val := f.MapValue(f.Value())

		assert.Equal(true, ok)
		arr := []any{map[string]any{"name": "11-1"}, map[string]any{"name": "11-2"}}
		assert.Equal(arr, val)
	})

	t.Run("nested map in struct", func(t *testing.T) {
		type Parent struct {
			Name  string         `json:"name"`
			Child map[string]any `json:"child"`
		}
		p1 := &Parent{
			Name:  "11",
			Child: map[string]any{"a": 1, "b": map[string]any{"name": "11-1"}},
		}

		s := New(p1)
		f, ok := s.Field("Child")
		val := f.MapValue(f.Value())
		assert.Equal(true, ok)
		assert.Equal(map[string]any{"a": 1, "b": map[string]any{"name": "11-1"}}, val)
	})
}

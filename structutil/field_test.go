package structutil

import (
	"fmt"
	"github.com/duke-git/lancet/v2/internal"
	"reflect"
	"testing"
)

func TestField_MapValue(t *testing.T) {
	assert := internal.NewAssert(t, "TestStructToMap")

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

	t.Run("nested array struct", func(t *testing.T) {
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
		v := f.Value()
		val := f.MapValue(v)
		fmt.Println("val", val)

		assert.Equal(true, ok)
		var arr []map[string]string
		arr = append(arr, map[string]string{"name": "11-1"})
		arr = append(arr, map[string]string{"name": "11-2"})
		var arr1 []map[string]string
		arr1 = append(arr, map[string]string{"name": "11-1"})
		arr1 = append(arr, map[string]string{"name": "11-2"})
		fmt.Println("arr", arr)
		fmt.Println(reflect.DeepEqual(arr, arr1))
		assert.Equal(arr, arr1)
	})
}

package structs

import (
	"reflect"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

func TestStruct_ToMap(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStruct_ToMap")
	t.Run("invalid struct", func(t *testing.T) {
		m, _ := ToMap(1)
		var expected map[string]any
		assert.Equal(expected, m)
	})

	t.Run("StructToMap", func(_ *testing.T) {
		type People struct {
			Name string `json:"name"`
			age  int
		}
		p := &People{
			"test",
			100,
		}
		pm, _ := ToMap(p)
		var expected = map[string]any{"name": "test"}
		assert.Equal(expected, pm)
	})

	t.Run("StructToMapWithJsonAttr", func(_ *testing.T) {
		type People struct {
			Name      string `json:"name,omitempty"` // json tag with attribute
			Phone     string `json:"phone"`          // json tag without attribute
			Sex       string `json:"-"`              // ignore by "-"
			Age       int    // ignore by no tag
			email     string // ignore by unexported
			IsWorking bool   `json:"is_working"`
		}
		p1 := People{
			Name:  "AAA", // exist
			Phone: "1111",
			Sex:   "male",
			Age:   100,
			email: "11@gmail.com",
		}
		p1m, _ := ToMap(p1)
		var expect1 = map[string]any{"name": "AAA", "phone": "1111", "is_working": false}
		assert.Equal(expect1, p1m)

		p2 := People{
			Name:      "",
			Phone:     "2222",
			Sex:       "male",
			Age:       0,
			email:     "22@gmail.com",
			IsWorking: true,
		}
		p2m, _ := ToMap(p2)
		var expect2 = map[string]any{"phone": "2222", "is_working": true}
		assert.Equal(expect2, p2m)
	})
}

func Test_ToMap2(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestStruct_ToMap")

	type M struct {
		Name string `json:"name"`
	}

	type S struct {
		ID int `json:"id"`
		M  *M  `json:"m"`
	}

	s1 := &S{
		ID: 1,
	}
	var expect1 = map[string]any{"id": 1}
	map1, err := ToMap(s1)

	assert.IsNil(err)
	assert.Equal(expect1, map1)

	s2 := &S{
		ID: 1,
		M: &M{
			Name: "test",
		},
	}
	var expect2 = map[string]any{
		"id": 1,
		"m": map[string]any{
			"name": "test",
		}}
	map2, err := ToMap(s2)

	assert.IsNil(err)
	assert.Equal(expect2, map2)
}

func TestStruct_Fields(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStruct_Fields")

	type Parent struct {
		A string         `json:"a"`
		B int            `json:"b"`
		C []string       `json:"c"`
		D map[string]any `json:"d"`
	}

	p1 := &Parent{
		A: "1",
		B: 11,
		C: []string{"11", "22"},
		D: map[string]any{"d1": 1, "d2": 2},
	}

	s := New(p1)
	fields := s.Fields()
	assert.Equal(4, len(fields))
	assert.Equal(reflect.String, fields[0].Kind())
	assert.Equal(reflect.Int, fields[1].Kind())
	assert.Equal(reflect.Slice, fields[2].Kind())
	assert.Equal(reflect.Map, fields[3].Kind())
}

func TestStruct_Field(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStruct_Field")

	type Parent struct {
		A string         `json:"a"`
		B int            `json:"b"`
		C []string       `json:"c"`
		D map[string]any `json:"d"`
	}

	p1 := &Parent{
		A: "1",
		B: 11,
		C: []string{"11", "22"},
		D: map[string]any{"d1": 1, "d2": 2},
	}

	s := New(p1)
	a, ok := s.Field("A")

	assert.Equal(true, ok)
	assert.Equal(reflect.String, a.Kind())
	assert.Equal("1", a.Value())
	assert.Equal("a", a.tag.Name)
	assert.Equal(false, a.tag.HasOption("omitempty"))
	assert.Equal(false, a.tag.IsEmpty())
}

func TestStruct_IsStruct(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStruct_Field")

	type Test1 struct{}
	t1 := &Test1{}
	t2 := 1

	s1 := New(t1)
	s2 := New(t2)

	assert.Equal(true, s1.IsStruct())
	assert.Equal(false, s2.IsStruct())
}

func TestStruct_Name(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStruct_Name")

	type Test1 struct{}
	t1 := &Test1{}

	s1 := New(t1)
	assert.Equal("Test1", s1.Name())

	type Test2 struct{}
	t2 := Test2{}

	s2 := New(t2)
	assert.Equal("Test2", s2.Name())
}

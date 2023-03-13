package structutil

import (
	"github.com/duke-git/lancet/v2/internal"
	"testing"
)

func TestToMap(t *testing.T) {
	assert := internal.NewAssert(t, "TestStructToMap")

	t.Run("StructToMap", func(_ *testing.T) {
		type People struct {
			Name string `json:"name"`
			age  int
		}
		p := People{
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

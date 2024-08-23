package maputil

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
)

// TestOrderMap tests the OrderedMap type.
func TestOrderMap(t *testing.T) {

	assert := internal.NewAssert(t, "TestOrderMap")

	m := NewOrderedMap[string, int]()
	m.Set("1", 1)
	m.Set("2", 2)
	m.Set("3", 3)
	assert.Equal(m.keys, []string{"1", "2", "3"})

	m.Delete("2")
	assert.Equal(m.keys, []string{"1", "3"})

	m.Set("2", 2)
	assert.Equal(m.keys, []string{"1", "3", "2"})
	assert.Equal(len(m.Values()), 3)
}

// TestOrderMapJson tests the OrderedMap type with JSON.
func TestOrderMapJson(t *testing.T) {

	assert := internal.NewAssert(t, "TestOrderMapJson")

	type TestData1 struct {
		Data OrderedMap[string, int] `json:"data"`
	}

	// test toJSON
	t1 := TestData1{}
	t1.Data = NewOrderedMap[string, int]()
	for i := 90; i < 100; i++ {
		t1.Data.Set(fmt.Sprintf("t-%d", i), i)
	}
	js, err := json.Marshal(&t1)
	assert.Equal(err, nil)
	assert.Equal(string(js), `{"data":{"t-90":90,"t-91":91,"t-92":92,"t-93":93,`+
		`"t-94":94,"t-95":95,"t-96":96,"t-97":97,"t-98":98,"t-99":99}}`)

	// test fromJSON
	str := `{"data":{"t-90":90,"t-91":91,"t-92":92,"t-93":93,` +
		`"t-94":94,"t-95":95,"t-96":96,"t-97":97,"t-98":98,"t-99":99}}`
	t2 := &TestData1{}
	err = json.Unmarshal([]byte(str), t2)
	assert.Equal(err, nil)
	assert.Equal(t1.Data.Keys(), []string{"t-90", "t-91", "t-92", "t-93", "t-94", "t-95", "t-96", "t-97", "t-98", "t-99"})
	assert.Equal(len(t1.Data.values), 10)

}

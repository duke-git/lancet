package maputil

import (
	"math/cmplx"
	"sort"
	"strconv"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/internal"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestKeys")

	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	keys := Keys(m)
	sort.Ints(keys)

	assert.Equal([]int{1, 2, 3, 4, 5}, keys)
}

func TestValues(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestValues")

	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
		4: "c",
		5: "d",
	}

	values := Values(m)
	sort.Strings(values)

	assert.Equal([]string{"a", "a", "b", "c", "d"}, values)
}

func TestKeysBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestKeysBy")

	m := map[int]string{
		1: "a",
		2: "a",
		3: "b",
	}

	keys := KeysBy(m, func(n int) int {
		return n + 1
	})

	sort.Ints(keys)

	assert.Equal([]int{2, 3, 4}, keys)
}

func TestValuesBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestValuesBy")

	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	values := ValuesBy(m, func(v string) string {
		switch v {
		case "a":
			return "a-1"
		case "b":
			return "b-2"
		case "c":
			return "c-3"
		default:
			return ""
		}
	})

	sort.Strings(values)

	assert.Equal([]string{"a-1", "b-2", "c-3"}, values)
}

func TestMerge(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMerge")

	m1 := map[int]string{
		1: "a",
		2: "b",
	}
	m2 := map[int]string{
		2: "c",
		3: "d",
	}

	expected := map[int]string{
		1: "a",
		2: "c",
		3: "d",
	}
	acturl := Merge(m1, m2)

	assert.Equal(expected, acturl)
}

func TestForEach(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestForEach")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	var sum int

	ForEach(m, func(_ string, value int) {
		sum += value
	})

	assert.Equal(10, sum)
}

func TestFilter(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFilter")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}

	acturl := Filter(m, isEven)

	assert.Equal(map[string]int{
		"b": 2,
		"d": 4,
	}, acturl)
}

func TestFilterByKeys(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFilterByKeys")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	acturl := FilterByKeys(m, []string{"a", "b"})

	assert.Equal(map[string]int{
		"a": 1,
		"b": 2,
	}, acturl)
}

func TestFilterByValues(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFilterByValues")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	acturl := FilterByValues(m, []int{3, 4})

	assert.Equal(map[string]int{
		"c": 3,
		"d": 4,
	}, acturl)
}

func TestOmitBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestOmitBy")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	isEven := func(_ string, value int) bool {
		return value%2 == 0
	}

	acturl := OmitBy(m, isEven)

	assert.Equal(map[string]int{
		"a": 1,
		"c": 3,
		"e": 5,
	}, acturl)
}

func TestOmitByKeys(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestOmitByKeys")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	acturl := OmitByKeys(m, []string{"a", "b"})

	assert.Equal(map[string]int{
		"c": 3,
		"d": 4,
		"e": 5,
	}, acturl)
}

func TestOmitByValues(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestOmitByValues")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}

	acturl := OmitByValues(m, []int{4, 5})

	assert.Equal(map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}, acturl)
}

func TestIntersect(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestIntersect")

	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 6,
		"d": 7,
	}

	m3 := map[string]int{
		"a": 1,
		"b": 9,
		"e": 9,
	}

	assert.Equal(map[string]int{"a": 1, "b": 2, "c": 3}, Intersect(m1))
	assert.Equal(map[string]int{"a": 1, "b": 2}, Intersect(m1, m2))
	assert.Equal(map[string]int{"a": 1}, Intersect(m1, m2, m3))
}

func TestMinus(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMinus")

	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"a": 11,
		"b": 22,
		"d": 33,
	}

	assert.Equal(map[string]int{"c": 3}, Minus(m1, m2))
}

func TestIsDisjoint(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMinus")

	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	m2 := map[string]int{
		"d": 22,
	}

	assert.Equal(true, IsDisjoint(m1, m2))

	m3 := map[string]int{
		"a": 22,
	}

	assert.Equal(false, IsDisjoint(m1, m3))
}

func TestEntries(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestEntries")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	result := Entries(m)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value < result[j].Value
	})

	expected := []Entry[string, int]{{Key: "a", Value: 1}, {Key: "b", Value: 2}, {Key: "c", Value: 3}}

	assert.Equal(expected, result)
}

func TestFromEntries(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFromEntries")

	result := FromEntries([]Entry[string, int]{
		{Key: "a", Value: 1},
		{Key: "b", Value: 2},
		{Key: "c", Value: 3},
	})

	assert.Equal(3, len(result))
	assert.Equal(1, result["a"])
	assert.Equal(2, result["b"])
	assert.Equal(3, result["c"])
}

func TestTransform(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestTransform")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	result := Transform(m, func(k string, v int) (string, string) {
		return k, strconv.Itoa(v)
	})

	expected := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	assert.Equal(expected, result)
}

func TestMapKeys(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMapKeys")

	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	result := MapKeys(m, func(k int, _ string) string {
		return strconv.Itoa(k)
	})

	expected := map[string]string{
		"1": "a",
		"2": "b",
		"3": "c",
	}

	assert.Equal(expected, result)
}

func TestMapValues(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMapKeys")

	m := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	result := MapValues(m, func(k int, v string) string {
		return v + strconv.Itoa(k)
	})

	expected := map[int]string{
		1: "a1",
		2: "b2",
		3: "c3",
	}

	assert.Equal(expected, result)
}

func TestHasKey(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestHasKey")

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	assert.Equal(true, HasKey(m, "a"))
	assert.Equal(false, HasKey(m, "c"))
}

func TestMapToStruct(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestMapToStruct")

	type (
		Person struct {
			Name  string   `json:"name"`
			Age   int      `json:"age"`
			Phone string   `json:"phone"`
			Addr  *Address `json:"address,omitempty"`
		}

		Address struct {
			Street string `json:"street"`
			Number int    `json:"number"`
		}
	)

	m := map[string]interface{}{
		"name":  "Nothin",
		"age":   28,
		"phone": "123456789",
		"address": map[string]interface{}{
			"street": "test",
			"number": 1,
		},
	}

	var p Person
	err := MapToStruct(m, &p)
	assert.IsNil(err)
	assert.Equal(m["name"], p.Name)
	assert.Equal(m["age"], p.Age)
	assert.Equal(m["phone"], p.Phone)
	assert.Equal("test", p.Addr.Street)
	assert.Equal(1, p.Addr.Number)
}

func TestToSortedSlicesDefault(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToSortedSlicesDefault")

	testCases1 := []struct {
		name      string
		inputMap  map[string]any
		expKeys   []string
		expValues []any
	}{
		{
			name:      "Empty Map",
			inputMap:  map[string]any{},
			expKeys:   []string{},
			expValues: []any{},
		},
		{
			name:      "Unsorted Map",
			inputMap:  map[string]any{"c": 3, "a": 1.6, "b": 2},
			expKeys:   []string{"a", "b", "c"},
			expValues: []any{1.6, 2, 3},
		},
	}

	for _, tc := range testCases1 {
		t.Run(tc.name, func(t *testing.T) {
			keys, values := ToSortedSlicesDefault(tc.inputMap)
			assert.Equal(tc.expKeys, keys)
			assert.Equal(tc.expValues, values)
		})
	}

	testCases2 := map[int64]string{
		7: "seven",
		3: "three",
		5: "five",
	}
	actualK2, actualV2 := ToSortedSlicesDefault(testCases2)
	case2Keys := []int64{3, 5, 7}
	case2Values := []string{"three", "five", "seven"}
	assert.Equal(case2Keys, actualK2)
	assert.Equal(case2Values, actualV2)
}

func TestToSortedSlicesWithComparator(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestToSortedSlicesWithComparator")

	type Account struct {
		ID                 int
		Name               string
		CreateTime         time.Time
		FavorComplexNumber complex128
		Signal             chan struct{}
	}

	type testCase struct {
		name      string
		inputMap  map[Account]any
		lessFunc  func(a, b Account) bool
		expKeys   []Account
		expValues []any
	}

	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	yesterday := now.AddDate(0, 0, -1)

	account1 := Account{ID: 1, Name: "cya", CreateTime: now, FavorComplexNumber: complex(1.2, 3),
		Signal: make(chan struct{}, 10)}
	account2 := Account{ID: 2, Name: "Cannian", CreateTime: tomorrow, FavorComplexNumber: complex(1.2, 2),
		Signal: make(chan struct{}, 7)}
	account3 := Account{ID: 3, Name: "Clauviou", CreateTime: yesterday, FavorComplexNumber: complex(3, 4),
		Signal: make(chan struct{}, 3)}
	account1.Signal <- struct{}{}
	account2.Signal <- struct{}{}
	account2.Signal <- struct{}{}

	testCases := []testCase{
		{
			name: "Sorted by Account ID",
			inputMap: map[Account]any{
				account1: 100,
				account2: 200,
				account3: 300,
			},
			lessFunc:  func(a, b Account) bool { return a.ID < b.ID },
			expKeys:   []Account{account1, account2, account3},
			expValues: []any{100, 200, 300},
		},
		{
			name: "Reverse Sorted by Account ID",
			inputMap: map[Account]any{
				account1: 100,
				account2: 200,
				account3: 300,
			},
			lessFunc:  func(a, b Account) bool { return a.ID > b.ID },
			expKeys:   []Account{account3, account2, account1},
			expValues: []any{300, 200, 100},
		},
		{
			name: "Sorted by Account Name",
			inputMap: map[Account]any{
				account1: 100,
				account2: 200,
				account3: 300,
			},
			lessFunc:  func(a, b Account) bool { return a.Name < b.Name },
			expKeys:   []Account{account2, account3, account1},
			expValues: []any{200, 300, 100},
		},
		{
			name:      "Empty Map",
			inputMap:  map[Account]any{},
			lessFunc:  func(a, b Account) bool { return a.ID < b.ID },
			expKeys:   []Account{},
			expValues: []any{},
		},
		{
			name: "Sorted by Account CreateTime",
			inputMap: map[Account]any{
				account1: "now",
				account2: "tomorrow",
				account3: "yesterday",
			},
			lessFunc:  func(a, b Account) bool { return a.CreateTime.Before(b.CreateTime) },
			expKeys:   []Account{account3, account1, account2},
			expValues: []any{"yesterday", "now", "tomorrow"},
		},
		{
			name: "Sorted by Account FavorComplexNumber",
			inputMap: map[Account]any{
				account1: "1.2+3i",
				account2: "1.2+2i",
				account3: "3+4i",
			},
			lessFunc:  func(a, b Account) bool { return cmplx.Abs(a.FavorComplexNumber) < cmplx.Abs(b.FavorComplexNumber) },
			expKeys:   []Account{account2, account1, account3},
			expValues: []any{"1.2+2i", "1.2+3i", "3+4i"},
		},
		{
			name: "Sort by the buffer capacity of the channel",
			inputMap: map[Account]any{
				account1: 10,
				account2: 7,
				account3: 3,
			},
			lessFunc: func(a, b Account) bool {
				return cap(a.Signal) < cap(b.Signal)
			},
			expKeys:   []Account{account3, account2, account1},
			expValues: []any{3, 7, 10},
		},
		{
			name: "Sort by the number of elements in the channel",
			inputMap: map[Account]any{
				account1: 1,
				account2: 2,
				account3: 0,
			},
			lessFunc:  func(a, b Account) bool { return len(a.Signal) < len(b.Signal) },
			expKeys:   []Account{account3, account1, account2},
			expValues: []any{0, 1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keys, values := ToSortedSlicesWithComparator(tc.inputMap, tc.lessFunc)
			assert.Equal(tc.expKeys, keys)
			assert.Equal(tc.expValues, values)
		})
	}
}

func TestGetOrSet(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestGetOrSet")

	m := map[int]string{
		1: "a",
	}

	result1 := GetOrSet(m, 1, "1")
	result2 := GetOrSet(m, 2, "b")

	assert.Equal("a", result1)
	assert.Equal("b", result2)
}

func TestSortByKey(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestSortByKey")

	m1 := map[int]string{
		3: "c",
		1: "a",
		4: "d",
		2: "b",
	}
	expected1 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
		4: "d",
	}

	result1 := SortByKey(m1, func(a, b int) bool {
		return a < b
	})

	assert.Equal(expected1, result1)

	m2 := map[string]int{
		"c": 3,
		"a": 1,
		"d": 4,
		"b": 2,
	}
	expected2 := map[string]int{
		"d": 4,
		"c": 3,
		"b": 2,
		"a": 1,
	}

	result2 := SortByKey(m2, func(a, b string) bool {
		return a > b
	})

	assert.Equal(expected2, result2)
}

type (
	Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Phone   string   `json:"phone"`
		Address *Address `json:"address"`
	}

	Address struct {
		Street string `json:"street"`
		Number int    `json:"number"`
	}
)

func TestStructType(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestStructType")

	src := map[string]interface{}{
		"name":  "Nothin",
		"age":   28,
		"phone": "123456789",
		"address": map[string]interface{}{
			"street": "test",
			"number": 1,
		},
	}

	var p Person
	err := MapTo(src, &p)
	assert.IsNil(err)

	assert.Equal(src["name"], p.Name)
	assert.Equal(src["age"], p.Age)
	assert.Equal(src["phone"], p.Phone)
	assert.Equal("test", p.Address.Street)
	assert.Equal(1, p.Address.Number)
}

func TestBaseType(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestBaseType")

	tc := map[string]interface{}{
		"i32":     -32,
		"i8":      -8,
		"i16":     -16,
		"i64":     -64,
		"i":       -1,
		"u32":     32,
		"u8":      8,
		"u16":     16,
		"u64":     64,
		"u":       1,
		"tf":      true,
		"f32":     1.32,
		"f64":     1.64,
		"str":     "hello mapto",
		"complex": 1 + 3i,
	}

	type BaseType struct {
		I    int        `json:"i"`
		I8   int8       `json:"i8"`
		I16  int16      `json:"i16"`
		I32  int32      `json:"i32"`
		I64  int64      `json:"i64"`
		U    uint       `json:"u"`
		U8   uint8      `json:"u8"`
		U16  uint16     `json:"u16"`
		U32  uint32     `json:"u32"`
		U64  uint64     `json:"u64"`
		F32  float32    `json:"f32"`
		F64  float64    `json:"f64"`
		Tf   bool       `json:"tf"`
		Str  string     `json:"str"`
		Comp complex128 `json:"complex"`
	}

	var dist BaseType
	err := MapTo(tc, &dist)
	assert.IsNil(err)

	var number float64

	MapTo(tc["i"], &number)
	assert.EqualValues(-1, number)

	MapTo(tc["i8"], &number)
	assert.EqualValues(-8, number)

	MapTo(tc["i16"], &number)
	assert.EqualValues(-16, number)

	MapTo(tc["i32"], &number)
	assert.EqualValues(-32, number)

	MapTo(tc["i64"], &number)
	assert.EqualValues(-64, number)

	MapTo(tc["u"], &number)
	assert.EqualValues(1, number)

	MapTo(tc["u8"], &number)
	assert.EqualValues(8, number)

	MapTo(tc["u16"], &number)
	assert.EqualValues(16, number)

	MapTo(tc["u32"], &number)
	assert.EqualValues(32, number)

	MapTo(tc["u64"], &number)
	assert.EqualValues(64, number)
}

func TestGetOrDefault(t *testing.T) {

	t.Parallel()

	assert := internal.NewAssert(t, "GetOrDefault")

	m1 := map[int]string{
		3: "c",
		1: "a",
		4: "d",
		2: "b",
	}
	result1 := GetOrDefault(m1, 1, "123")
	assert.Equal("a", result1)

	result2 := GetOrDefault(m1, 5, "123")
	assert.Equal("123", result2)
}

func TestFindValuesBy(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestFindValuesBy")

	tests := []struct {
		name     string
		inputMap map[string]string
		key      string
		expected []string
	}{
		{
			name:     "Key exists",
			inputMap: map[string]string{"a": "1", "b": "2", "c": "3"},
			key:      "b",
			expected: []string{"2"},
		},
		{
			name:     "Key does not exist",
			inputMap: map[string]string{"a": "1", "b": "2", "c": "3"},
			key:      "d",
			expected: []string{},
		},
		{
			name:     "Empty map",
			inputMap: map[string]string{},
			key:      "a",
			expected: []string{},
		},
	}
	for _, tt := range tests {
		result := FindValuesBy(tt.inputMap, func(key string, value string) bool {
			return key == tt.key
		})
		assert.Equal(tt.expected, result)
	}
}

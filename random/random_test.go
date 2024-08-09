package random

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/duke-git/lancet/v2/internal"
	"github.com/duke-git/lancet/v2/validator"
)

func TestRandString(t *testing.T) {
	t.Parallel()

	pattern := `^[a-zA-Z]+$`
	reg := regexp.MustCompile(pattern)

	randStr := RandString(6)

	assert := internal.NewAssert(t, "TestRandString")
	assert.Equal(6, len(randStr))
	assert.Equal(true, reg.MatchString(randStr))
}

func TestRandUpper(t *testing.T) {
	t.Parallel()

	pattern := `^[A-Z]+$`
	reg := regexp.MustCompile(pattern)

	randStr := RandUpper(6)

	assert := internal.NewAssert(t, "TestRandUpper")
	assert.Equal(6, len(randStr))
	assert.Equal(true, reg.MatchString(randStr))
}

func TestRandLower(t *testing.T) {
	t.Parallel()

	pattern := `^[a-z]+$`
	reg := regexp.MustCompile(pattern)

	randStr := RandLower(6)

	assert := internal.NewAssert(t, "TestRandLower")
	assert.Equal(6, len(randStr))
	assert.Equal(true, reg.MatchString(randStr))
}

func TestRandNumeral(t *testing.T) {
	t.Parallel()

	pattern := `^[0-9]+$`
	reg := regexp.MustCompile(pattern)

	randStr := RandNumeral(12)

	assert := internal.NewAssert(t, "TestRandNumeral")
	assert.Equal(12, len(randStr))
	assert.Equal(true, reg.MatchString(randStr))
}

func TestRandNumeralOrLetter(t *testing.T) {
	t.Parallel()

	pattern := `^[0-9a-zA-Z]+$`
	reg := regexp.MustCompile(pattern)

	randStr := RandNumeralOrLetter(10)

	assert := internal.NewAssert(t, "TestRandNumeralOrLetter")
	assert.Equal(10, len(randStr))
	assert.Equal(true, reg.MatchString(randStr))
}

func TestRandInt(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRandInt")

	r1 := RandInt(1, 10)
	assert.GreaterOrEqual(r1, 1)
	assert.Less(r1, 10)

	r2 := RandInt(1, 1)
	assert.Equal(1, r2)

	r3 := RandInt(10, 1)
	assert.GreaterOrEqual(r1, 1)
	assert.Less(r3, 10)
}

func TestRandBytes(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRandBytes")

	randBytes := RandBytes(4)
	assert.Equal(4, len(randBytes))

	v := reflect.ValueOf(randBytes)
	elemType := v.Type().Elem()
	assert.Equal(reflect.Slice, v.Kind())
	assert.Equal(reflect.Uint8, elemType.Kind())

	assert.Equal([]byte{}, RandBytes(0))
}

func TestUUIdV4(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestUUIdV4")

	uuid, err := UUIdV4()
	if err != nil {
		t.Fail()
	}

	isUUiDV4 := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
	assert.Equal(true, isUUiDV4.MatchString(uuid))
}

func TestRandUniqueIntSlice(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRandUniqueIntSlice")

	r1 := RandUniqueIntSlice(5, 0, 9)
	assert.Equal(len(r1), 5)
	if hasDuplicate(r1) {
		t.Error("hasDuplicate int")
	}

	r2 := RandUniqueIntSlice(20, 0, 10)
	assert.Equal(len(r2), 10)
	if hasDuplicate(r2) {
		t.Error("hasDuplicate int")
	}

	r3 := RandUniqueIntSlice(10, 20, 10)
	assert.Equal(len(r3), 0)

	r4 := RandUniqueIntSlice(0, 20, 10)
	assert.Equal(len(r4), 0)
}

func hasDuplicate(arr []int) bool {
	elements := make(map[int]bool)
	for _, v := range arr {
		if elements[v] {
			return true
		}
		elements[v] = true
	}
	return false
}

func TestRandSymbolChar(t *testing.T) {
	t.Parallel()

	pattern := `^[\W|_]+$`
	reg := regexp.MustCompile(pattern)

	symbolChars := RandSymbolChar(10)

	assert := internal.NewAssert(t, "TestRandSymbolChar")
	assert.Equal(10, len(symbolChars))
	assert.Equal(true, reg.MatchString(symbolChars))
}

func TestRandFloat(t *testing.T) {
	t.Parallel()

	assert := internal.NewAssert(t, "TestRandFloat")

	r1 := RandFloat(1.1, 10.1, 2)
	assert.GreaterOrEqual(r1, 1.1)
	assert.Less(r1, 10.1)

	r2 := RandFloat(1.1, 1.1, 2)
	assert.Equal(1.1, r2)

	r3 := RandFloat(10.1, 1.1, 2)
	assert.GreaterOrEqual(r1, 1.1)
	assert.Less(r3, 10.1)
}

func TestRandFloats(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRandFloats")

	numbers := RandFloats(5, 1.0, 5.0, 2)
	for _, n := range numbers {
		assert.Equal(true, (n >= 1.0 && n < 5.0))
	}

	assert.Equal(len(numbers), 5)
}

func TestRandIntSlice(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRandIntSlice")

	t.Run("empty slice", func(t *testing.T) {
		numbers := RandIntSlice(-1, 1, 5)
		assert.Equal([]int{}, numbers)

		numbers = RandIntSlice(0, 1, 5)
		assert.Equal([]int{}, numbers)

		numbers = RandIntSlice(3, 5, 1)
		assert.Equal([]int{}, numbers)
	})

	t.Run("random int slice", func(t *testing.T) {
		numbers := RandIntSlice(5, 1, 1)
		assert.Equal([]int{1, 1, 1, 1, 1}, numbers)

		numbers = RandIntSlice(5, 1, 5)
		assert.Equal(5, len(numbers))
	})
}

func TestRandStringSlice(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRandStringSlice")

	t.Run("empty slice", func(t *testing.T) {
		strs := RandStringSlice(Letters, -1, -1)
		assert.Equal([]string{}, strs)

		strs = RandStringSlice(Letters, 0, 0)
		assert.Equal([]string{}, strs)

		strs = RandStringSlice(Letters, -1, 0)
		assert.Equal([]string{}, strs)

		strs = RandStringSlice(Letters, 0, -1)
		assert.Equal([]string{}, strs)

		strs = RandStringSlice(Letters, 1, 0)
		assert.Equal([]string{}, strs)

		strs = RandStringSlice(Letters, 0, 1)
		assert.Equal([]string{}, strs)
	})

	t.Run("random string slice", func(t *testing.T) {
		strs := RandStringSlice(Letters, 4, 6)
		assert.Equal(4, len(strs))

		for _, s := range strs {
			assert.Equal(true, validator.IsAlpha(s))
			assert.Equal(6, len(s))
		}
	})

	// fail test: chinese character is not supported for now
	// t.Run("random string slice of chinese ", func(t *testing.T) {
	// 	strs := RandStringSlice("你好你好你好你好你好你好你好你好你好", 4, 6)
	// 	t.Log(strs)

	// 	assert.Equal(4, len(strs))
	// 	for _, s := range strs {
	// 		assert.Equal(true, validator.ContainChinese(s))
	// 		assert.Equal(6, len(s))
	// 	}
	// })
}

func TestRandBool(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRandBool")

	result := RandBool()
	assert.Equal(true, result == true || result == false)
}

func TestRandBoolSlice(t *testing.T) {
	t.Parallel()
	assert := internal.NewAssert(t, "TestRandBoolSlice")

	t.Run("empty slice", func(t *testing.T) {
		bools := RandBoolSlice(-1)
		assert.Equal([]bool{}, bools)

		bools = RandBoolSlice(0)
		assert.Equal([]bool{}, bools)
	})

	t.Run("random bool slice", func(t *testing.T) {
		bools := RandBoolSlice(6)
		assert.Equal(6, len(bools))

		for _, b := range bools {
			assert.Equal(true, b == true || b == false)
		}
	})
}

// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package random implements some basic functions to generate random int and string.
package random

import (
	crand "crypto/rand"
	"fmt"
	"io"
	"math"
	"math/rand"
	"time"
	"unsafe"

	"github.com/duke-git/lancet/mathutil"
)

const (
	MaximumCapacity = math.MaxInt>>1 + 1
	Numeral         = "0123456789"
	LowwerLetters   = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SymbolChars     = "!@#$%^&*()_+-=[]{}|;':\",./<>?"
)

var rn = rand.NewSource(time.Now().UnixNano())

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandBool generates a random boolean value (true or false).
func RandBool() bool {
	return rand.Intn(2) == 1
}

// RandBoolSlice generates a random boolean slice of specified length.
func RandBoolSlice(length int) []bool {
	if length <= 0 {
		return []bool{}
	}

	result := make([]bool, length)
	for i := range result {
		result[i] = RandBool()
	}

	return result
}

// RandInt generate random int between [min, max).
func RandInt(min, max int) int {
	if min == max {
		return min
	}

	if max < min {
		min, max = max, min
	}

	return rand.Intn(max-min) + min
}

// RandIntSlice generates a slice of random integers.
// The generated integers are between min and max (exclusive).
func RandIntSlice(length, min, max int) []int {
	if length <= 0 || min > max {
		return []int{}
	}

	result := make([]int, length)
	for i := range result {
		result[i] = RandInt(min, max)
	}

	return result
}

// RandUniqueIntSlice generate a slice of random int of length n that do not repeat.
func RandUniqueIntSlice(n, min, max int) []int {
	if min > max {
		return []int{}
	}
	if n > max-min {
		n = max - min
	}

	nums := make([]int, n)
	used := make(map[int]struct{}, n)
	for i := 0; i < n; {
		r := RandInt(min, max)
		if _, use := used[r]; use {
			continue
		}
		used[r] = struct{}{}
		nums[i] = r
		i++
	}

	return nums
}

// RandFloat generate random float64 number between [min, max) with specific precision.
func RandFloat(min, max float64, precision int) float64 {
	if min == max {
		return min
	}

	if max < min {
		min, max = max, min
	}

	n := rand.Float64()*(max-min) + min

	return mathutil.RoundToFloat(n, precision)
}

// RandFloats generate a slice of random float64 numbers of length n that do not repeat.
func RandFloats(n int, min, max float64, precision int) []float64 {
	nums := make([]float64, n)
	used := make(map[float64]struct{}, n)
	for i := 0; i < n; {
		r := RandFloat(min, max, precision)
		if _, use := used[r]; use {
			continue
		}
		used[r] = struct{}{}
		nums[i] = r
		i++
	}

	return nums
}

// RandBytes generate random byte slice.
// Play: https://go.dev/play/p/EkiLESeXf8d
func RandBytes(length int) []byte {
	if length < 1 {
		return []byte{}
	}
	b := make([]byte, length)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return nil
	}

	return b
}

// RandString generate random alpha string of specified length.
func RandString(length int) string {
	return random(Letters, length)
}

// RandString generate a slice of random string of length strLen based on charset.
// chartset should be one of the following: random.Numeral, random.LowwerLetters, random.UpperLetters
// random.Letters, random.SymbolChars, random.AllChars. or a combination of them.
func RandStringSlice(charset string, sliceLen, strLen int) []string {
	if sliceLen <= 0 || strLen <= 0 {
		return []string{}
	}

	result := make([]string, sliceLen)

	for i := range result {
		result[i] = random(charset, strLen)
	}

	return result
}

// RandUpper generate a random upper case string of specified length.
func RandUpper(length int) string {
	return random(UpperLetters, length)
}

// RandLower generate a random lower case string of specified length.
func RandLower(length int) string {
	return random(LowwerLetters, length)
}

// RandNumeral generate a random numeral string of specified length.
func RandNumeral(length int) string {
	return random(Numeral, length)
}

// RandNumeralOrLetter generate a random numeral or alpha string of specified length.
func RandNumeralOrLetter(length int) string {
	return random(Numeral+Letters, length)
}

// RandSymbolChar generate a random symbol char of specified length.
// symbol chars: !@#$%^&*()_+-=[]{}|;':\",./<>?.
func RandSymbolChar(length int) string {
	return random(SymbolChars, length)
}

// nearestPowerOfTwo 返回一个大于等于cap的最近的2的整数次幂，参考java8的hashmap的tableSizeFor函数
func nearestPowerOfTwo(cap int) int {
	n := cap - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	if n < 0 {
		return 1
	} else if n >= MaximumCapacity {
		return MaximumCapacity
	}
	return n + 1
}

// random generate a random string based on given string range.
func random(s string, length int) string {
	// 仿照strings.Builder
	// 创建一个长度为 length 的字节切片
	bytes := make([]byte, length)
	strLength := len(s)
	if strLength <= 0 {
		return ""
	} else if strLength == 1 {
		for i := 0; i < length; i++ {
			bytes[i] = s[0]
		}
		return *(*string)(unsafe.Pointer(&bytes))
	}
	// s的字符需要使用多少个比特位数才能表示完
	// letterIdBits := int(math.Ceil(math.Log2(strLength))),下面比上面的代码快
	letterIdBits := int(math.Log2(float64(nearestPowerOfTwo(strLength))))
	// 最大的字母id掩码
	var letterIdMask int64 = 1<<letterIdBits - 1
	// 可用次数的最大值
	letterIdMax := 63 / letterIdBits
	// 循环生成随机字符串
	for i, cache, remain := length-1, rn.Int63(), letterIdMax; i >= 0; {
		// 检查随机数生成器是否用尽所有随机数
		if remain == 0 {
			cache, remain = rn.Int63(), letterIdMax
		}
		// 从可用字符的字符串中随机选择一个字符
		if idx := int(cache & letterIdMask); idx < strLength {
			bytes[i] = s[idx]
			i--
		}
		// 右移比特位数，为下次选择字符做准备
		cache >>= letterIdBits
		remain--
	}
	// 仿照strings.Builder用unsafe包返回一个字符串，避免拷贝
	// 将字节切片转换为字符串并返回
	return *(*string)(unsafe.Pointer(&bytes))
}

// UUIdV4 generate a random UUID of version 4 according to RFC 4122.
func UUIdV4() (string, error) {
	uuid := make([]byte, 16)

	n, err := io.ReadFull(crand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

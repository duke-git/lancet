// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package random implements some basic functions to generate random int and string.
package random

import (
	crand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"time"
)

const (
	Numeral       = "0123456789"
	LowwerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandInt generate random int between min and max, maybe min,  not be max.
// Play: https://go.dev/play/p/pXyyAAI5YxD
func RandInt(min, max int) int {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}

	// fix: https://github.com/duke-git/lancet/issues/75
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// return r.Intn(max-min) + min

	return rand.Intn(max-min) + min
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

// RandString generate random string of specified length.
// Play: https://go.dev/play/p/W2xvRUXA7Mi
func RandString(length int) string {
	return random(Letters, length)
}

// RandUpper generate a random upper case string.
// Play: https://go.dev/play/p/29QfOh0DVuh
func RandUpper(length int) string {
	return random(UpperLetters, length)
}

// RandLower generate a random lower case string.
// Play: https://go.dev/play/p/XJtZ471cmtI
func RandLower(length int) string {
	return random(LowwerLetters, length)
}

// RandNumeral generate a random numeral string of specified length.
// Play: https://go.dev/play/p/g4JWVpHsJcf
func RandNumeral(length int) string {
	return random(Numeral, length)
}

// RandNumeralOrLetter generate a random numeral or letter string.
// Play: https://go.dev/play/p/19CEQvpx2jD
func RandNumeralOrLetter(length int) string {
	return random(Numeral+Letters, length)
}

// random generate a random string based on given string range.
func random(s string, length int) string {
	b := make([]byte, length)

	// fix: https://github.com/duke-git/lancet/issues/75
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range b {
		b[i] = s[rand.Int63()%int64(len(s))]
	}

	return string(b)
}

// UUIdV4 generate a random UUID of version 4 according to RFC 4122.
// Play: https://go.dev/play/p/_Z9SFmr28ft
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

// RandUniqueIntSlice generate a slice of random int of length n that do not repeat
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

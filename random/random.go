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
	NUMERAL       = "0123456789"
	LOWER_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPER_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LETTERS       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandInt generate random int between min and max, maybe min,  not be max
func RandInt(min, max int) int {
	if min == max {
		return min
	}

	if max < min {
		min, max = max, min
	}

	return rand.Intn(max-min) + min
}

// RandBytes generate random byte slice
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

// RandString generate random string
func RandString(length int) string {
	return random(LETTERS, length)
}

// RandUpper generate a random upper case string
func RandUpper(length int) string {
	return random(UPPER_LETTERS, length)
}

// RandLower generate a random lower case string
func RandLower(length int) string {
	return random(LOWER_LETTERS, length)
}

// RandNumeral generate a random numeral string
func RandNumeral(length int) string {
	return random(NUMERAL, length)
}

// RandNumeralOrLetter generate a random numeral or letter string
func RandNumeralOrLetter(length int) string {
	return random(NUMERAL+LETTERS, length)
}

// random generate a random string based on given string range
func random(s string, length int) string {
	b := make([]byte, length)

	// fix: https://github.com/duke-git/lancet/issues/75
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range b {
		b[i] = s[rand.Int63()%int64(len(s))]
	}
	return string(b)
}

// UUIdV4 generate a random UUID of version 4 according to RFC 4122
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

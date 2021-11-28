// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package random implements some basic functions to generate random int and string.
package random

import (
	crand "crypto/rand"
	"io"
	"math/rand"
	"time"
)

// RandString generate random string
// see https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Int63()%int64(len(letters))]
	}
	return string(b)
}

// RandInt generate random int between min and max, maybe min,  not be max
func RandInt(min, max int) int {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// RandBytes generate random byte slice
func RandBytes(length int) []byte {
	if length < 1 {
		return nil
	}
	b := make([]byte, length)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return nil
	}
	return b
}

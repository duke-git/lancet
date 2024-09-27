// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package formatter implements some functions to format string, struct.
package formatter

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/duke-git/lancet/v2/convertor"
	"golang.org/x/exp/constraints"
)

// Comma add comma to a number value by every 3 numbers from right. ahead by symbol char.
// if value is invalid number string eg "aa", return empty string
// Comma("12345", "$") => "$12,345", Comma(12345, "$") => "$12,345"
// Play: https://go.dev/play/p/eRD5k2vzUVX
func Comma[T constraints.Float | constraints.Integer | string](value T, symbol string) string {
	numString := convertor.ToString(value)

	_, err := strconv.ParseFloat(numString, 64)
	if err != nil {
		return ""
	}

	isNegative := strings.HasPrefix(numString, "-")
	if isNegative {
		numString = numString[1:]
	}

	index := strings.Index(numString, ".")
	if index == -1 {
		index = len(numString)
	}

	for index > 3 {
		index -= 3
		numString = numString[:index] + "," + numString[index:]
	}

	if isNegative {
		numString = "-" + numString
	}

	return symbol + numString
}

// Pretty data to JSON string.
// Play: https://go.dev/play/p/YsciGj3FH2x
func Pretty(v any) (string, error) {
	out, err := json.MarshalIndent(v, "", "    ")
	return string(out), err
}

// PrettyToWriter pretty encode data to writer.
// Play: https://go.dev/play/p/LPLZ3lDi5ma
func PrettyToWriter(v any, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")

	if err := enc.Encode(v); err != nil {
		return err
	}

	return nil
}

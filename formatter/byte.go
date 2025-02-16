package formatter

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/duke-git/lancet/v2/strutil"
)

//
// code logic come from:
// https://github.com/docker/go-units/blob/master/size.go

// http://en.wikipedia.org/wiki/Binary_prefix
const (
	// Decimal
	unitB  = 1
	unitKB = 1000
	unitMB = 1000 * unitKB
	unitGB = 1000 * unitMB
	unitTB = 1000 * unitGB
	unitPB = 1000 * unitTB
	unitEB = 1000 * unitPB

	// Binary
	unitBiB = 1
	unitKiB = 1024
	unitMiB = 1024 * unitKiB
	unitGiB = 1024 * unitMiB
	unitTiB = 1024 * unitGiB
	unitPiB = 1024 * unitTiB
	unitEiB = 1024 * unitPiB
)

// type byteUnitMap map[byte]int64

var (
	decimalByteMap = map[string]uint64{
		"b":  unitB,
		"kb": unitKB,
		"mb": unitMB,
		"gb": unitGB,
		"tb": unitTB,
		"pb": unitPB,
		"eb": unitEB,

		// Without suffix
		"":  unitB,
		"k": unitKB,
		"m": unitMB,
		"g": unitGB,
		"t": unitTB,
		"p": unitPB,
		"e": unitEB,
	}

	binaryByteMap = map[string]uint64{
		"bi":  unitBiB,
		"kib": unitKiB,
		"mib": unitMiB,
		"gib": unitGiB,
		"tib": unitTiB,
		"pib": unitPiB,
		"eib": unitEiB,

		// Without suffix
		"":   unitBiB,
		"ki": unitKiB,
		"mi": unitMiB,
		"gi": unitGiB,
		"ti": unitTiB,
		"pi": unitPiB,
		"ei": unitEiB,
	}
)

var (
	decimalByteUnits = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	binaryByteUnits  = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
)

// DecimalBytes returns a human readable byte size under decimal standard (base 1000)
// The precision parameter specifies the number of digits after the decimal point, which defaults to 4.
// Play: https://go.dev/play/p/FPXs1suwRcs
func DecimalBytes(size float64, precision ...int) string {
	pointPosition := 4
	if len(precision) > 0 {
		pointPosition = precision[0]
	}

	size, unit := calculateByteSize(size, 1000.0, decimalByteUnits)

	return roundToToString(size, pointPosition) + unit
}

// BinaryBytes returns a human-readable byte size under binary standard (base 1024)
// The precision parameter specifies the number of digits after the decimal point, which defaults to 4.
// Play: https://go.dev/play/p/G9oHHMCAZxP
func BinaryBytes(size float64, precision ...int) string {
	pointPosition := 4
	if len(precision) > 0 {
		pointPosition = precision[0]
	}

	size, unit := calculateByteSize(size, 1024.0, binaryByteUnits)

	return roundToToString(size, pointPosition) + unit
}

func calculateByteSize(size float64, base float64, byteUnits []string) (float64, string) {
	i := 0
	unitsLimit := len(byteUnits) - 1
	for size >= base && i < unitsLimit {
		size = size / base
		i++
	}
	return size, byteUnits[i]
}

func roundToToString(x float64, max ...int) string {
	pointPosition := 4
	if len(max) > 0 {
		pointPosition = max[0]
	}
	result := mathutil.RoundToString(x, pointPosition)

	// 删除小数位结尾的0
	decimal := strings.TrimRight(strutil.After(result, "."), "0")
	if decimal == "" || pointPosition == 0 {
		// 没有小数位直接返回整数
		return strutil.Before(result, ".")
	}

	// 小数位大于想要设置的位数，按需要截断
	if len(decimal) > pointPosition {
		return strutil.Before(result, ".") + "." + decimal[:pointPosition]
	}

	// 小数位小于等于想要的位数，直接拼接返回
	return strutil.Before(result, ".") + "." + decimal
}

// ParseDecimalBytes return the human readable bytes size string into the amount it represents(base 1000).
// ParseDecimalBytes("42 MB") -> 42000000, nil
// Play: https://go.dev/play/p/Am98ybWjvjj
func ParseDecimalBytes(size string) (uint64, error) {
	return parseBytes(size, "decimal")
}

// ParseBinaryBytes return the human readable bytes size string into the amount it represents(base 1024).
// ParseBinaryBytes("42 mib") -> 44040192, nil
// Play: https://go.dev/play/p/69v1tTT62x8
func ParseBinaryBytes(size string) (uint64, error) {
	return parseBytes(size, "binary")
}

// see https://github.com/dustin/go-humanize/blob/master/bytes.go
func parseBytes(s string, kind string) (uint64, error) {
	lastDigit := 0
	hasComma := false
	for _, r := range s {
		if !(unicode.IsDigit(r) || r == '.' || r == ',') {
			break
		}
		if r == ',' {
			hasComma = true
		}
		lastDigit++
	}

	num := s[:lastDigit]
	if hasComma {
		num = strings.Replace(num, ",", "", -1)
	}

	f, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, err
	}

	extra := strings.ToLower(strings.TrimSpace(s[lastDigit:]))

	if kind == "decimal" {
		if m, ok := decimalByteMap[extra]; ok {
			f *= float64(m)
			if f >= math.MaxUint64 {
				return 0, fmt.Errorf("too large: %v", s)
			}
			return uint64(f), nil
		}
	} else {
		if m, ok := binaryByteMap[extra]; ok {
			f *= float64(m)
			if f >= math.MaxUint64 {
				return 0, fmt.Errorf("too large: %v", s)
			}
			return uint64(f), nil
		}
	}

	return 0, fmt.Errorf("unhandled size name: %v", extra)
}

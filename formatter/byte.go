package formatter

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

//
// code logic come from:
// https://github.com/docker/go-units/blob/master/size.go

// http://en.wikipedia.org/wiki/Binary_prefix
const (
	// Decimal
	UnitB  = 1
	UnitKB = 1000
	UnitMB = 1000 * UnitKB
	UnitGB = 1000 * UnitMB
	UnitTB = 1000 * UnitGB
	UnitPB = 1000 * UnitTB
	UnitEB = 1000 * UnitPB

	// Binary
	UnitBiB = 1
	UnitKiB = 1024
	UnitMiB = 1024 * UnitKiB
	UnitGiB = 1024 * UnitMiB
	UnitTiB = 1024 * UnitGiB
	UnitPiB = 1024 * UnitTiB
	UnitEiB = 1024 * UnitPiB
)

// type byteUnitMap map[byte]int64

var (
	decimalByteMap = map[string]uint64{
		"b":  UnitB,
		"kb": UnitKB,
		"mb": UnitMB,
		"gb": UnitGB,
		"tb": UnitTB,
		"pb": UnitPB,
		"eb": UnitEB,

		// Without suffix
		"":  UnitB,
		"k": UnitKB,
		"m": UnitMB,
		"g": UnitGB,
		"t": UnitTB,
		"p": UnitPB,
		"e": UnitEB,
	}

	binaryByteMap = map[string]uint64{
		"bi":  UnitBiB,
		"kib": UnitKiB,
		"mib": UnitMiB,
		"gib": UnitGiB,
		"tib": UnitTiB,
		"pib": UnitPiB,
		"eib": UnitEiB,

		// Without suffix
		"":   UnitBiB,
		"ki": UnitKiB,
		"mi": UnitMiB,
		"gi": UnitGiB,
		"ti": UnitTiB,
		"pi": UnitPiB,
		"ei": UnitEiB,
	}
)

var (
	decimalByteUnits = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	binaryByteUnits  = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
)

// DecimalBytes returns a human readable byte size under decimal standard (base 1000)
// The precision parameter specifies the number of digits after the decimal point, which defaults to 4.
// Play: todo
func DecimalBytes(size float64, precision ...int) string {
	p := 5
	if len(precision) > 0 {
		p = precision[0] + 1
	}
	size, unit := calculateByteSize(size, 1000.0, decimalByteUnits)

	return fmt.Sprintf("%.*g%s", p, size, unit)
}

// BinaryBytes returns a human-readable byte size under decimal standard (base 1024)
// The precision parameter specifies the number of digits after the decimal point, which defaults to 4.
// Play: todo
func BinaryBytes(size float64, precision ...int) string {
	p := 5
	if len(precision) > 0 {
		p = precision[0] + 1
	}
	size, unit := calculateByteSize(size, 1024.0, binaryByteUnits)

	return fmt.Sprintf("%.*g%s", p, size, unit)
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

// ParseDecimalBytes the human readable bytes size string into the amount it represents(base 1000).
// ParseDecimalBytes("42 MB") -> 42000000, nil
// Play: todo
func ParseDecimalBytes(size string) (uint64, error) {
	return parseBytes(size, "decimal")
}

// ParseBinaryBytes the human readable bytes size string into the amount it represents(base 1024).
// ParseBinaryBytes("42 mib") -> 44040192, nil
// Play: todo
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

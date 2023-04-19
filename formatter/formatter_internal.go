package formatter

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// see https://github.com/dustin/go-humanize/blob/master/comma.go
func commaInt(v int64) string {
	sign := ""

	// Min int64 can't be negated to a usable value, so it has to be special cased.
	if v == math.MinInt64 {
		return "-9,223,372,036,854,775,808"
	}

	if v < 0 {
		sign = "-"
		v = 0 - v
	}

	parts := []string{"", "", "", "", "", "", ""}
	j := len(parts) - 1

	for v > 999 {
		parts[j] = strconv.FormatInt(v%1000, 10)
		switch len(parts[j]) {
		case 2:
			parts[j] = "0" + parts[j]
		case 1:
			parts[j] = "00" + parts[j]
		}
		v = v / 1000
		j--
	}
	parts[j] = strconv.Itoa(int(v))
	return sign + strings.Join(parts[j:], ",")
}

func commaFloat(v float64) string {
	buf := &bytes.Buffer{}
	if v < 0 {
		buf.Write([]byte{'-'})
		v = 0 - v
	}

	comma := []byte{','}

	parts := strings.Split(strconv.FormatFloat(v, 'f', -1, 64), ".")
	pos := 0
	if len(parts[0])%3 != 0 {
		pos += len(parts[0]) % 3
		buf.WriteString(parts[0][:pos])
		buf.Write(comma)
	}
	for ; pos < len(parts[0]); pos += 3 {
		buf.WriteString(parts[0][pos : pos+3])
		buf.Write(comma)
	}
	buf.Truncate(buf.Len() - 1)

	if len(parts) > 1 {
		buf.Write([]byte{'.'})
		buf.WriteString(parts[1])
	}
	return buf.String()
}

func commaStr(s string) string {
	dotIndex := strings.Index(s, ".")
	if dotIndex != -1 {
		return commaStrRecursive(s[:dotIndex]) + s[dotIndex:]
	}

	return commaStrRecursive(s)
}

func commaStrRecursive(s string) string {
	if len(s) <= 3 {
		return s
	}
	return commaStrRecursive(s[:len(s)-3]) + "," + commaStrRecursive(s[len(s)-3:])
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

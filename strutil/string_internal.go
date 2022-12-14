package strutil

import (
	"strings"
	"unicode"
)

// splitWordsToLower split a string into worlds by uppercase char
func splitWordsToLower(s string) []string {
	var result []string

	upperIndexes := upperIndex(s)
	l := len(upperIndexes)
	if upperIndexes == nil || l == 0 {
		if s != "" {
			result = append(result, s)
		}
		return result
	}
	for i := 0; i < l; i++ {
		if i < l-1 {
			result = append(result, strings.ToLower(s[upperIndexes[i]:upperIndexes[i+1]]))
		} else {
			result = append(result, strings.ToLower(s[upperIndexes[i]:]))
		}
	}
	return result
}

// upperIndex get a int slice which elements are all the uppercase char index of a string
func upperIndex(s string) []int {
	var result []int
	for i := 0; i < len(s); i++ {
		if 64 < s[i] && s[i] < 91 {
			result = append(result, i)
		}
	}
	if len(s) > 0 && result != nil && result[0] != 0 {
		result = append([]int{0}, result...)
	}

	return result
}

func splitIntoStrings(s string, upperCase bool) []string {
	var runes [][]rune
	lastCharType := 0
	charType := 0

	// split into fields based on type of unicode character
	for _, r := range s {
		switch true {
		case isLower(r):
			charType = 1
		case isUpper(r):
			charType = 2
		case isDigit(r):
			charType = 3
		default:
			charType = 4
		}

		if charType == lastCharType {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastCharType = charType
	}

	for i := 0; i < len(runes)-1; i++ {
		if isUpper(runes[i][0]) && isLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}

	// filter all none letters and none digit
	var result []string
	for _, rs := range runes {
		if len(rs) > 0 && (unicode.IsLetter(rs[0]) || isDigit(rs[0])) {
			if upperCase {
				result = append(result, string(toUpperAll(rs)))
			} else {
				result = append(result, string(toLowerAll(rs)))

			}
		}
	}

	return result
}

// isDigit checks if a character is digit ('0' to '9')
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// isLower checks if a character is lower case ('a' to 'z')
func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// isUpper checks if a character is upper case ('A' to 'Z')
func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// toLower converts a character  'A' to 'Z' to its lower case
func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

// toLowerAll converts a character  'A' to 'Z' to its lower case
func toLowerAll(rs []rune) []rune {
	for i := range rs {
		rs[i] = toLower(rs[i])
	}
	return rs
}

// toUpper converts a character  'a' to 'z' to its upper case
func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

// toUpperAll converts a character  'a' to 'z' to its upper case
func toUpperAll(rs []rune) []rune {
	for i := range rs {
		rs[i] = toUpper(rs[i])
	}
	return rs
}

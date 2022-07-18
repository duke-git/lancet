package strutil

import "strings"

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

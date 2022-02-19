package strutil

import "strings"

// splitWordsToLower split a string into worlds by uppercase char
func splitWordsToLower(s string) []string {
	var res []string

	upperIndexes := upperIndex(s)
	l := len(upperIndexes)
	if upperIndexes == nil || l == 0 {
		if s != "" {
			res = append(res, s)
		}
		return res
	}
	for i := 0; i < l; i++ {
		if i < l-1 {
			res = append(res, strings.ToLower(s[upperIndexes[i]:upperIndexes[i+1]]))
		} else {
			res = append(res, strings.ToLower(s[upperIndexes[i]:]))
		}
	}
	return res
}

// upperIndex get a int slice which elements are all the uppercase char index of a string
func upperIndex(s string) []int {
	var res []int
	for i := 0; i < len(s); i++ {
		if 64 < s[i] && s[i] < 91 {
			res = append(res, i)
		}
	}
	if len(s) > 0 && res != nil && res[0] != 0 {
		res = append([]int{0}, res...)
	}

	return res
}

package formatter

import (
	"strings"
	"unicode/utf8"
)

// mbStrpos 返回字符串首次出现的位置（UTF-8字符计数）
func mbStrpos(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return -1
	}
	return utf8.RuneCountInString(haystack[:idx])
}

// mbStrripos 返回字符串最后出现的位置（UTF-8字符计数）
func mbStrripos(haystack, needle string) int {
	if needle == "" {
		return utf8.RuneCountInString(haystack)
	}
	idx := strings.LastIndex(haystack, needle)
	if idx == -1 {
		return -1
	}
	return utf8.RuneCountInString(haystack[:idx])
}

// mbStrstr 检查字符串是否包含子串
func mbStrstr(haystack, needle string) bool {
	return strings.Contains(haystack, needle)
}

// mbSubstr 截取字符串（UTF-8字符计数）
// start: 起始位置（从0开始）
// length: 截取长度（字符数）
func mbSubstr(str string, start, length int) string {
	runes := []rune(str)
	strLen := len(runes)

	// 处理负数起始位置
	if start < 0 {
		start = strLen + start
		if start < 0 {
			start = 0
		}
	}

	// 起始位置超出字符串长度
	if start >= strLen {
		return ""
	}

	// 计算结束位置
	end := start + length
	if end > strLen {
		end = strLen
	}
	if end < start {
		return ""
	}

	return string(runes[start:end])
}

// mbSubstrCount 统计子串出现次数
func mbSubstrCount(haystack, needle string) int {
	if needle == "" {
		return 0
	}
	return strings.Count(haystack, needle)
}

//@Author: wulinlin
//@Description: 字符串处理相关的工具类
//@File:  str
//@Version: 1.0.0
//@Date: 2023/12/01 22:56

package utils

import (
	"strings"
	"unicode"
)

func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func IsAnyBlank(strs ...string) bool {
	for _, str := range strs {
		if IsBlank(str) {
			return true
		}
	}
	return false
}

func DefaultIfBlank(str, def string) string {
	if IsBlank(str) {
		return def
	} else {
		return str
	}
}

// IsEmpty checks if a string is empty (""). Returns true if empty, and false otherwise.
func IsEmpty(str string) bool {
	return len(str) == 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

// 截取字符串
func Substr(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

// UUID
//func UUID() string {
//	//u := uuid.Must(uuid.NewV4())
//	uid := uuid.NewV4()
//	var err error
//	u := uuid.Must(uid, err)
//	return strings.ReplaceAll(u.String(), "-", "")
//}

func Equals(a, b string) bool {
	return a == b
}

func EqualsIgnoreCase(a, b string) bool {
	return a == b || strings.ToUpper(a) == strings.ToUpper(b)
}

// RuneLen 字符成长度
func RuneLen(s string) int {
	bt := []rune(s)
	return len(bt)
}

// GetSummary 获取summary
func GetSummary(s string, length int) string {
	s = strings.TrimSpace(s)
	summary := Substr(s, 0, length)
	if RuneLen(s) > length {
		summary += "..."
	}
	return summary
}

// Reverse 反转字符串
func Reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// CountWords 统计字符串中的单词数
func CountWords(input string) int {
	words := strings.Fields(input)
	return len(words)
}

// UppercaseFirst 将字符串的第一个字符转换为大写
func UppercaseFirst(input string) string {
	if len(input) == 0 {
		return input
	}
	return string(unicode.ToUpper(rune(input[0]))) + input[1:]
}

// LowercaseFirst 将字符串的第一个字符转换为小写
func LowercaseFirst(input string) string {
	if len(input) == 0 {
		return input
	}
	return string(unicode.ToLower(rune(input[0]))) + input[1:]
}

// IsPalindrome 判断字符串是否是回文
func IsPalindrome(input string) bool {
	normalized := strings.ToLower(strings.ReplaceAll(input, " ", ""))
	return normalized == Reverse(normalized)
}

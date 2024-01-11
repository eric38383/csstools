package utils

import "strings"

func GetBetweenTwoChars(str string, startChar string, endChar string, result string, separator string) string {
	start := strings.Index(str, startChar)
	if start == -1 {
		return result
	}
	newS := str[start+len(startChar):]
	end := strings.Index(newS, endChar)
	if end == -1 {
		return result
	}
	result += newS[:end] + separator
	return GetBetweenTwoChars(newS[end+1:], startChar, endChar, result, separator)
}

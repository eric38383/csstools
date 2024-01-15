package utils

import "strings"

func GetBetweenTwoChars(str string, startChar string, endChar string) []string {
	start := strings.Index(str, startChar)
	if start == -1 {
		return nil
	}
	newS := str[start+len(startChar):]
	end := strings.Index(newS, endChar)
	if end == -1 {
		return nil
	}

	return append([]string{newS[:end]}, GetBetweenTwoChars(newS[end+1:], startChar, endChar)...)
}

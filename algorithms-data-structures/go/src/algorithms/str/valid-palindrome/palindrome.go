package validpalindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	s = strings.ToLower(s)
	regexp, _ := regexp.Compile("[^a-z0-9]+")
	s = regexp.ReplaceAllString(s, "")

	return reverseString(s) == s
}

func reverseString(s string) string {
	var reversed string
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}

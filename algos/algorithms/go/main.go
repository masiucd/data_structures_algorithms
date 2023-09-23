package main

import "fmt"

func main() {
	res := isPalindrome("abba")

	fmt.Println(res)
}

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

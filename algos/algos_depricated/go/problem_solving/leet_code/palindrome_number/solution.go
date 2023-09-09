package palindromenumber

import "strconv"

func isPalindrome(x int) bool {
	nStr := strconv.Itoa(x)
	return nStr == reverseString(nStr)
}

func reverseString(s string) string {
	var r string
	for _, v := range s {
		r = string(v) + r
	}
	return r
}

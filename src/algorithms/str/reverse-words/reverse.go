package reversewords

import "strings"

func reverseWords(s string) string {
	xs := strings.Fields(s)
	var result string
	for len(xs) > 0 {
		x := xs[len(xs)-1]
		xs = xs[:len(xs)-1]
		result += " " + x
	}
	return strings.TrimSpace(result)
}

func reverseWordsV2(s string) string {
	xs := strings.Fields(s)
	var result string
	for i := len(xs) - 1; i >= 0; i-- {
		result += " " + xs[i]
	}
	return strings.TrimSpace(result)
}

package mixed

import "unicode"

func detectCapitalUse(word string) bool {
	var upperCount int
	var lowerCount int
	for _, v := range word {
		if isUpper(v) {
			upperCount++
		} else if isLower(v) {
			lowerCount++
		}
	}
	// check if the first letter is uppercase but rest are lowercase
	if isUpper(rune(word[0])) && every(word[1:], func(r rune) bool {
		return isLower(r)
	}) {
		return true
	}

	if upperCount == len(word) {
		return true
	}
	if lowerCount == len(word) {
		return true
	}
	return false

}

func every(xs string, fn func(r rune) bool) bool {
	for _, x := range xs {
		if !fn(x) {
			return false
		}
	}
	return true
}

func isUpper(r rune) bool {
	return unicode.IsUpper(r)
}
func isLower(r rune) bool {
	return !isUpper(r)
}

// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".

// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".

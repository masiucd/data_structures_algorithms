package simplestringchars

import "unicode"

func solution(input string) []int {
	var xs [4]int
	for _, v := range input {
		if unicode.IsUpper(v) {
			xs[0]++
		} else if unicode.IsLower(v) {
			xs[1]++
		} else if unicode.IsDigit(v) {
			xs[2]++
		} else if !unicode.IsLetter(v) && !unicode.IsDigit(v) {
			xs[3]++
		}
	}
	var result []int
	for _, v := range xs {
		result = append(result, v)
	}

	return result
}

func solution2(input string) []int {
	r := make([]int, 4)
	for _, c := range input {
		switch {
		case c >= 'A' && c <= 'Z':
			r[0]++
		case c >= 'a' && c <= 'z':
			r[1]++
		case c >= '0' && c <= '9':
			r[2]++
		default:
			r[3]++
		}
	}
	return r
}

package count_encoding

import "fmt"

// Good morning! Here’s your coding interview problem for today.

// This problem was asked by Amazon.

// Run-length encoding is a fast and simple method of encoding strings. The basic idea is to represent repeated successive characters as a single count and character. For example, the string “AAAABBBCCDAA” would be encoded as “4A3B2C1D2A”.

// Implement run-length encoding and decoding. You can assume the string to be encoded have no digits and consists solely of alphabetic characters. You can assume the string to be decoded is valid.
func solution(input string) string {

	var result string
	count := 1
	for i := 1; i <= len(input); i++ {
		currentChar := input[i-1]
		if i <= len(input)-1 && currentChar == input[i] {
			count++
		} else {
			addCharAndAmount(&result, count, currentChar)
			count = 1
		}
	}
	return result
}

func addCharAndAmount(result *string, count int, char uint8) {
	*result += fmt.Sprintf("%d", count)
	*result += string(char)
}

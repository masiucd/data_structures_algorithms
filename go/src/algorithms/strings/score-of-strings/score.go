package strings

import (
	"math"
)

// The score of a string is the sum of the absolute differences between two consecutive characters.
// scoreOfString returns the score of a string.
func scoreOfString(s string) int {
	var result float64
	// var chars []rune
	for i := 0; i < len(s)-1; i++ {
		char, nextChar := s[i], s[i+1]
		result += math.Abs(float64(char) - float64(nextChar))
		// chars = append(chars, char)
	}

	return int(result)
}

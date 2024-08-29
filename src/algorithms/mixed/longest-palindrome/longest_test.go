package mixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	name     string
	input    string
	expected int
}{
	{"should return 7 for abccccdd", "abccccdd", 7},
	{"should return 1 for a", "a", 1},
	{"should return 2 for bb", "bb", 2},
	{"should return 1 for ab", "ab", 1},
	{"should return 3 for ccc", "ccc", 3},
}

func TestLongestPalindrome(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, longestPalindrome(tt.input))
		})
	}
}

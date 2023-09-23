package ispalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	input string
	want  bool
}

var tests = []test{
	{"", true},
	{"a", true},
	{"aa", true},
	{"aba", true},
	{"abba", true},
	{"abca", false},
	{"abcba", true},
	{"abcd", false},
	{"racecar", true},
}

func TestPalindrome(t *testing.T) {
	for _, v := range tests {
		got := isPalindrome(v.input)
		assert.Equal(t, v.want, got)
	}
}

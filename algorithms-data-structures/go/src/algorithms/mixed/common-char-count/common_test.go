package commoncharcount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonCharCount(t *testing.T) {
	var tests = []struct {
		s1 string
		s2 string
		x  int
	}{
		{"aabcc", "adcaa", 3},
		{"abca", "xyzbac", 3},
		{"zzzz", "zzzzzzz", 4},
		{"abca", "xyzbac", 3},
		{"a", "b", 0},
		{"a", "aaa", 1},
	}

	for _, test := range tests {
		assert.Equal(t, test.x, commonCharacterCount(test.s1, test.s2))
	}
}

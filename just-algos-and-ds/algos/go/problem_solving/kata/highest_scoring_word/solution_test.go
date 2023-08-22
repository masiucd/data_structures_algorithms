package highestscoringword

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighestScoringWord(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"abcd", "abcd"},
		{"aa b", "aa"},
		{"a b c d e f g", "g"},
		{"what time are we climbing up the volcano", "volcano"},
		{"b aa", "b"},
	}

	for _, tc := range testCases {
		result := solution(tc.input)
		assert.Equal(t, tc.expected, result, fmt.Sprintf("Expected = %s, got = %s", tc.expected, result))
	}
}

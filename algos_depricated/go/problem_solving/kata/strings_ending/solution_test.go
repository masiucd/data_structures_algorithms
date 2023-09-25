package strings_ending

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []struct {
	input    []string
	expected bool
}{
	{[]string{"", ""}, true},
	{[]string{" ", ""}, true},
	{[]string{"abc", "c"}, true},
	{[]string{"banana", "ana"}, true},
	{[]string{"a", "z"}, false},
	{[]string{"", "t"}, false},
	{[]string{"   ", "   "}, true},
	{[]string{"1oo", "100"}, false},
}

func TestStringsEnding(t *testing.T) {
	for _, tt := range testData {
		actual := solution(tt.input[0], tt.input[1])
		assert.Equal(t, tt.expected, actual, fmt.Sprintf("expected %v, actual %v", tt.expected, actual))
	}

	for _, tt := range testData {
		actual := solutionTwo(tt.input[0], tt.input[1])
		assert.Equal(t, tt.expected, actual, fmt.Sprintf("expected %v, actual %v", tt.expected, actual))
	}
}

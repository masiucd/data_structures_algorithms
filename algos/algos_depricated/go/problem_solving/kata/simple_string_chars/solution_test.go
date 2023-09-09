package simplestringchars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	input    string
	expected []int
}

var testCases = []testData{
	{"a", []int{0, 1, 0, 0}},
	{"aA", []int{1, 1, 0, 0}},
	{"aA1", []int{1, 1, 1, 0}},
	{"aA1!", []int{1, 1, 1, 1}},
	{"aA1!bB2@2", []int{2, 2, 3, 2}},
}

func TestSimpleStringCharsSolutionOne(t *testing.T) {
	for _, tc := range testCases {
		result := solution(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}
func TestSimpleStringCharsSolutionTwo(t *testing.T) {
	for _, tc := range testCases {
		result := solution2(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}

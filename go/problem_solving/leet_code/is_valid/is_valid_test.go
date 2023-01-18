package is_valid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	type test struct {
		input    string
		expected bool
	}

	tests := []test{
		{input: "()", expected: true},
		{input: "()[]{}", expected: true},
		{input: "(]", expected: false},
		{input: "([)]", expected: false},
		{input: "{[]}", expected: true},
		{input: "]", expected: false},
	}

	for _, tc := range tests {
		got := isValid(tc.input)
		assert.Equal(t, tc.expected, got, fmt.Sprintf("input: %s", tc.input))
	}
}

func TestPairs(t *testing.T) {
	got := pairs[')']
	want := '('
	assert.Equal(t, want, got)

	got = pairs['}']
	want = '{'
	assert.Equal(t, want, got)

	got = pairs[']']
	want = '['
	assert.Equal(t, want, got)
}

func TestIsValidTwo(t *testing.T) {
	type test struct {
		input    string
		expected bool
	}

	tests := []test{
		{input: "()", expected: true},
		{input: "()[]{}", expected: true},
		{input: "(]", expected: false},
		{input: "([)]", expected: false},
		{input: "{[]}", expected: true},
		{input: "]", expected: false},
	}

	for _, tc := range tests {
		got := isValidTwo(tc.input)
		assert.Equal(t, tc.expected, got, fmt.Sprintf("input: %s", tc.input))
	}
}

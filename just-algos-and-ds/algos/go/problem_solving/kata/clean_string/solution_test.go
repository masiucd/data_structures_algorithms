package clean_string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCleanString(t *testing.T) {

	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{input: "abc#d##c", expected: "ac"},
		{input: "abc##d######", expected: ""},
		{input: "#6+yqw8hfklsd-=-f", expected: "6+yqw8hfklsd-=-f"},
	}

	for _, tc := range tests {
		got := solution(tc.input)
		assert.Equal(t, tc.expected, got, fmt.Sprintf("input: %s", tc.input))
	}

}

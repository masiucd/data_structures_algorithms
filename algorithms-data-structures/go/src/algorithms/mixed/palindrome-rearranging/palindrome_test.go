package mixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeRearranging(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Test 1", "abbcabb", true},
		{"Test 2", "abbcabbx", false},
		{"Test 3", "abbcab", false},
		{"Test 4", "abbcabbc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := palindromeRearranging(tt.input)
			assert.Equal(t, tt.expected, res)
		})
	}
}

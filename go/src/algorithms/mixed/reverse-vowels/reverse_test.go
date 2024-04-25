package mixed

import "testing"

func TestReverseVowels(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"aA", "Aa"},
		{"leetcode", "leotcede"},
		{"hello", "holle"},
	}

	for _, tc := range testCases {
		output := reverseVowels(tc.input)
		if output != tc.expected {
			t.Errorf("Test failed for input %v, expected: '%v', got:  '%v'", tc.input, tc.expected, output)
		}
	}
}

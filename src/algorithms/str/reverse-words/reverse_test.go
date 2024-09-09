package reversewords

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"should reverse words in a string", "a good   example", "example good a"},
		{"should reverse words in a string", "  hello world  ", "world hello"},
		{"should reverse words in a string", "  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"should reverse words in a string", "Alice does not even like bob", "bob like even not does Alice"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// assert.Equal(t, tt.expected, reverseWords(tt.input))
			if tt.expected != reverseWords(tt.input) {
				t.Errorf("Expected %v but got %v", tt.expected, reverseWords(tt.input))
			}
		})
	}
}

func TestReverseWordsV2(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"should reverse words in a string", "a good   example", "example good a"},
		{"should reverse words in a string", "  hello world  ", "world hello"},
		{"should reverse words in a string", "  Bob    Loves  Alice   ", "Alice Loves Bob"},
		{"should reverse words in a string", "Alice does not even like bob", "bob like even not does Alice"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != reverseWordsV2(tt.input) {
				t.Errorf("Expected %v but got %v", tt.expected, reverseWordsV2(tt.input))
			}
		})
	}
}

package countvowels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountVowels(t *testing.T) {
	tt := []struct {
		name     string
		input    string
		expected int
	}{
		{"empty string", "", 0},
		{"single vowel", "a", 1},
		{"single consonant", "b", 0},
		{"multiple vowels", "abracadabra", 5},
		{"name", "marcell", 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := getVowelsCount(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res := getVowelsCount2(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

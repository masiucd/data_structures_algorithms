package gcdofstrings

import "testing"

func TestGcdOfStrings(t *testing.T) {
	tests := []struct {
		s1, s2, want string
	}{
		{"ABCABC", "ABC", "ABC"},
		{"ABABAB", "ABAB", "AB"},
		{"LEET", "CODE", ""},
		{"ABCDEF", "ABC", ""},
		{"ABC", "ABCDEF", ""},
	}

	for _, tt := range tests {
		if got := gcdOfStrings(tt.s1, tt.s2); got != tt.want {
			t.Errorf("gcdOfStrings(%v, %v) = %v, want %v", tt.s1, tt.s2, got, tt.want)
		}
	}

}

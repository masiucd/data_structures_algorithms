package mixed

import "testing"

var tests = []struct {
	name     string
	word1    string
	word2    string
	expected string
}{
	{"should return apbqcr for abc and pqr", "abc", "pqr", "apbqcr"},
	{"should return apbqrs for ab and pqrs", "ab", "pqrs", "apbqrs"},
	{"should return apbqcd for abcd and pq", "abcd", "pq", "apbqcd"},
}

func TestMergeAlternately(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternately(tt.word1, tt.word2); got != tt.expected {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.expected)
			}
		})
	}
}

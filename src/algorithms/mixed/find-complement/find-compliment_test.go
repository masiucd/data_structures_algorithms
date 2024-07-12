package mixed

import "testing"

var tests = []struct {
	num      int
	expected int
}{
	{num: 5, expected: 2},
	{num: 1, expected: 0},
	{num: 7, expected: 0},
	{num: 10, expected: 5},
}

func TestFindCompliment(t *testing.T) {
	for _, tt := range tests {
		actual := findComplement(tt.num)
		if actual != tt.expected {
			t.Errorf("Expected %d, but got %d", tt.expected, actual)
		}
	}
}

func TestFindComplimentV2(t *testing.T) {
	for _, tt := range tests {
		actual := findComplementV2(tt.num)
		if actual != tt.expected {
			t.Errorf("Expected %d, but got %d", tt.expected, actual)
		}
	}
}

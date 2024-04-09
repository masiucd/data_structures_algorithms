package mixed

import "testing"

func TestFindCompliment(t *testing.T) {
	tests := []struct {
		num      int
		expected int
	}{
		{num: 5, expected: 2},
		{num: 1, expected: 0},
		{num: 7, expected: 0},
		{num: 10, expected: 5},
	}

	for _, tt := range tests {
		actual := findComplement(tt.num)
		if actual != tt.expected {
			t.Errorf("Expected %d, but got %d", tt.expected, actual)
		}
	}
}

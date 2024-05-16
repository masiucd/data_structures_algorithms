package numbers

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{10, 4},
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 2},
		{5, 2},
		{6, 3},
		{7, 3},
		{8, 4},
		{9, 4},
		{11, 4},
		{12, 5},
	}

	for _, tt := range tests {
		result := countPrimes(tt.n)
		if result != tt.expected {
			t.Errorf("expected %d, got %d", tt.expected, result)
		}
	}

}

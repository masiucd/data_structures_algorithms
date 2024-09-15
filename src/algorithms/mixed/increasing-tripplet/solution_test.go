package mixed

import (
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
		{[]int{5, 1, 5, 5, 2, 5, 4}, true},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, false},
	}
	for _, tt := range tests {
		result := increasingTriplet(tt.nums)
		if tt.expected != result {
			t.Errorf("expected %t, got %t", tt.expected, result)
		}
	}
}

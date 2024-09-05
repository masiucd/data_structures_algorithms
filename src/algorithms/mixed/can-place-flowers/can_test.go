package mixed

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{0, 0, 1, 0, 0}, 1, true},
		{[]int{1, 0, 0, 0, 0, 1}, 2, false},
		{[]int{0, 0, 0, 0, 0, 0, 0}, 3, true},
	}

	for _, tt := range tests {
		actual := canPlaceFlowers(tt.flowerbed, tt.n)
		if actual != tt.expected {
			t.Errorf("Expected %v but got %v", tt.expected, actual)
		}
	}
}

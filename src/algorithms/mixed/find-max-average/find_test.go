package mixed

import "testing"

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{0, 1, 1, 3, 3}, 4, 2.0},
		{[]int{3, 2, 4, 1}, 2, 3.0},
		{[]int{5}, 1, 5.0},
	}

	for _, tt := range tests {
		got := findMaxAverage(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}
	}
}

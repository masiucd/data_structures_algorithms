package mixed

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{1, 0, 2, 0, 3}, []int{1, 2, 3, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("moveZeroes(%v)", tt.input), func(t *testing.T) {
			xs := make([]int, len(tt.input))
			copy(xs, tt.input)
			moveZeroes(xs)
			for i := range xs {
				if xs[i] != tt.expected[i] {
					t.Errorf("Expected %d, but got %d", tt.expected[i], xs[i])
				}
			}
		})

		t.Run(fmt.Sprintf("moveZeroesV2(%v)", tt.input), func(t *testing.T) {
			xs := make([]int, len(tt.input))
			copy(xs, tt.input)
			moveZeroesV2(xs)
			for i := range xs {
				if xs[i] != tt.expected[i] {
					t.Errorf("Expected %d, but got %d", tt.expected[i], xs[i])
				}
			}
		})

		t.Run(fmt.Sprintf("moveZeroesV3(%v)", tt.input), func(t *testing.T) {
			xs := make([]int, len(tt.input))
			copy(xs, tt.input)
			moveZeroesV3(xs)
			for i := range xs {
				if xs[i] != tt.expected[i] {
					t.Errorf("Expected %d, but got %d", tt.expected[i], xs[i])
				}
			}
		})
	}

}

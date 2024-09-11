package mixed

import (
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{[]int{}, []int{}},   // Edge case: empty array
		{[]int{1}, []int{1}}, // Edge case: single element
	}

	for _, tt := range tests {
		got := productExceptSelf(tt.input)
		if !equal(got, tt.want) {
			t.Errorf("productExceptSelf(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestProductExceptSelfV2(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{[]int{}, []int{}},   // Edge case: empty array
		{[]int{1}, []int{1}}, // Edge case: single element
	}

	for _, tt := range tests {
		got := productExceptSelfV2(tt.input)
		if !equal(got, tt.want) {
			t.Errorf("productExceptSelfV2(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

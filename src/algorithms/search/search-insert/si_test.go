package search

import (
	"fmt"
	"testing"
)

var tests = []struct {
	xs     []int
	target int
	want   int
}{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 0, 0}, // Additional test case
}

func TestSi(t *testing.T) {

	for _, tt := range tests {
		t.Run(fmt.Sprintf("searchInsert(%v, %d)", tt.xs, tt.target), func(t *testing.T) {
			got := searchInsert(tt.xs, tt.target)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

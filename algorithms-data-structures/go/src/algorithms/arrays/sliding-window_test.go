package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type test struct {
	xs         []int
	windowSize int
	expected   int
}

var tests = []test{
	{[]int{1, 2, 3, 4, 5, 6}, 3, 15},
	{[]int{1, 2, 3, 4, 5, 6}, 2, 11},
	{[]int{1, 2, 3, 4, 5, 6}, 4, 18},
	{[]int{1, 2, 3, 4, 5, 6}, 5, 20},
	{[]int{10, 12, 1, 2, 3}, 2, 22},
}

func TestSlidingWindow(t *testing.T) {
	for _, tc := range tests {
		actual := slidingWindow(tc.xs, tc.windowSize)
		assert.Equal(t, tc.expected, actual)
	}
}

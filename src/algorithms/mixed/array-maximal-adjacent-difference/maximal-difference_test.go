package mixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	name     string
	input    []int
	expected int
}{
	{"Test case 1", []int{2, 4, 1, 0}, 3},
	{"Test case 2", []int{1, 1, 1, 1}, 0},
	{"Test case 3", []int{-1, 4, 10, 3, -2}, 7},
}

func TestArrayMaximalAdjacentDifference(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := arrayMaximalAdjacentDifference(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

func TestArrayMaximalAdjacentDifferenceV2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := arrayMaximalAdjacentDifferenceV2(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

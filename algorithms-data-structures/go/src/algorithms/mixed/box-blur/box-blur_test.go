package mixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoxBlur(t *testing.T) {
	testCases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			input: [][]int{
				{7, 4, 0, 1},
				{5, 6, 2, 2},
				{6, 10, 7, 8},
				{1, 4, 2, 0},
			},
			expected: [][]int{
				{5, 4},
				{4, 4},
			},
		},
		{
			input: [][]int{
				{36, 0, 18, 9},
				{27, 54, 9, 0},
				{81, 63, 72, 45},
			},
			expected: [][]int{
				{40, 30},
				{30, 30},
			},
		},
		{
			input: [][]int{
				{1, 1, 1},
				{1, 7, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1},
			},
		},
	}

	for _, tc := range testCases {
		res := boxBlur(tc.input)
		assertEqualMatrix(t, tc.expected, res)
	}
}

func assertEqualMatrix(t *testing.T, expected, actual [][]int) {
	for i := range actual {
		for j := range actual[i] {
			assert.Equal(t, expected[i][j], actual[i][j])
		}
	}
}

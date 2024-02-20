package mixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAvoidObstacles(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Test case 1", []int{5, 3, 6, 7, 9}, 4},
		{"Test case 2", []int{2, 3}, 4},
		{"Test case 3", []int{1, 4, 10, 6, 2}, 7},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := avoidObstacles(tc.input)
			assert.Equal(t, tc.expected, res)
		})
	}
}

package maxsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSum(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Test #1",
			heights:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Test #2",
			heights:  []int{1, 1},
			expected: 1,
		},
		{
			name:     "Test #3",
			heights:  []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			name:     "Test #4",
			heights:  []int{1, 2, 1},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := maxSum(tt.heights)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

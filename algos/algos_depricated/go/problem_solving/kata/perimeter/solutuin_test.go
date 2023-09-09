package perimeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	// Test cases
	testCases := []struct {
		n        int
		expected int
	}{
		{5, 80},
		{7, 216},
		{20, 114624},
	}

	for _, tc := range testCases {
		result := perimeter(tc.n)
		assert.Equal(t, tc.expected, result, "Expected %d, got %d", tc.expected, result)
	}
}

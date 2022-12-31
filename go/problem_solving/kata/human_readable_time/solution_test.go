package humanreadabletime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHumanReadableTime(t *testing.T) {
	// Test cases
	testCases := []struct {
		seconds  int
		expected string
	}{
		{59, "00:00:59"},
		{60, "00:01:00"},
		{3599, "00:59:59"},
		{3600, "01:00:00"},
		{86399, "23:59:59"},
		{86400, "24:00:00"},
		{359999, "99:59:59"},
		{360000, "100:00:00"},
	}

	for _, tc := range testCases {
		result := solution(tc.seconds)
		assert.Equal(t, tc.expected, result, "Expected %s, got %s", tc.expected, result)
	}

}

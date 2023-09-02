package mili_seconds_past_since_midnight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiliSecondsAfterMidnight(t *testing.T) {
	actual := solution(0, 0, 0)
	expected := 0
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	actual = solution(0, 1, 1)
	expected = 61000
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	actual = solution(1, 1, 1)
	expected = 3661000
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))
}

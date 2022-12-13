package sum_of_sequance

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomOfSequence(t *testing.T) {
	actual := solution(2, 6, 2)
	expected := 12
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	actual = solution(1, 5, 1)
	expected = 15
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	actual = solution(12, 5, 3)
	expected = 0
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

}

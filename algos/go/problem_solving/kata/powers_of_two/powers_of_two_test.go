package powers_of_two

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowersOfTwo(t *testing.T) {
	actual := solution(0)
	expected := []uint64{1}
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %v  actual %v", expected, actual))

	actual = solution(1)
	expected = []uint64{1, 2}
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %v  actual %v", expected, actual))

	actual = solution(4)
	expected = []uint64{1, 2, 4, 8, 16}
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %v  actual %v", expected, actual))
}

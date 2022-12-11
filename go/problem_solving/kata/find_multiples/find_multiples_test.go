package find_multiples

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMultiples(t *testing.T) {
	actual := solution(1, 2)
	expected := []int{1, 2}
	assert.Equal(t, actual, expected)

	actual = solution(5, 25)
	expected = []int{5, 10, 15, 20, 25}
	assert.Equal(t, actual, expected)
}

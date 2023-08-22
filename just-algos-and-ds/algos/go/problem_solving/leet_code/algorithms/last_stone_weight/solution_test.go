package last_stone_weight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastStoneWeight(t *testing.T) {

	expected := lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
	actual := 1
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	expected = lastStoneWeight([]int{1, 3})
	actual = 2
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	expected = lastStoneWeight([]int{1, 5, 3})
	actual = 1
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))
}

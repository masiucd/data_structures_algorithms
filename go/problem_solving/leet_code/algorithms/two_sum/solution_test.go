package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumOne(t *testing.T) {
	var expected = twoSumOne([]int{2, 7, 11, 15}, 9)
	var actual = []int{0, 1}
	for i := range expected {
		assert.Equal(t, expected[i], actual[i], "Didn't match")
	}

	expected = twoSumOne([]int{3, 2, 4}, 6)
	actual = []int{1, 2}
	for i := range expected {
		assert.Equal(t, expected[i], actual[i], "Didn't match")
	}
}

func TestTwoSumTwo(t *testing.T) {
	var expected = twoSumTwo([]int{2, 7, 11, 15}, 9)
	var actual = []int{0, 1}
	for i := range expected {
		assert.Equal(t, expected[i], actual[i], "Didn't match")
	}

	expected = twoSumTwo([]int{3, 2, 4}, 6)
	actual = []int{1, 2}
	for i := range expected {
		assert.Equal(t, expected[i], actual[i], "Didn't match")
	}
}

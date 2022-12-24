package smallestinteger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	input    []int
	expected int
}{
	{[]int{1}, 1},
	{[]int{1, 2}, 1},
	{[]int{1, 2, 3}, 1},
	{[]int{1, 2, 3, 4}, 1},
	{[]int{1, 2, -3, 4}, -3},
}

func TestSmallestInteger(t *testing.T) {
	for _, v := range testData {
		actual := solution(v.input)
		assert.Equal(t, v.expected, actual, fmt.Sprintf("Expected %v, got %v", v.expected, actual))
	}
}

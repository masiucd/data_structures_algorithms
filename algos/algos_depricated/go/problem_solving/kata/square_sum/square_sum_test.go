package square_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []struct {
	input    []int
	expected int
}{{input: []int{1, 2, 3}, expected: 14}, {input: []int{2, 3, 4}, expected: 29}}

func TestSquareSum(t *testing.T) {
	for _, v := range testData {
		got := squareSum(v.input)
		assert.Equal(t, got, v.expected)
	}
}

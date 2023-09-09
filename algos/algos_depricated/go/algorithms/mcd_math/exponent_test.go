package mcd_math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	expected []int
	actual   int
}{
	{[]int{2, 4}, 16},
	{[]int{1, 8}, 1},
	{[]int{5, 0}, 1},
	{[]int{5, 2}, 25},
	{[]int{5, 3}, 125},
}

func TestExponentByRecursion(t *testing.T) {
	for _, tt := range testData {
		expected := exponentByRecursion(tt.expected[0], tt.expected[1])
		assert.Equal(t, expected, tt.actual, fmt.Sprintf("expected %d actual %d ", expected, tt.actual))
	}
}

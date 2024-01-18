package matrixelementsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixElementsSum(t *testing.T) {
	matrix := [][]int{{0, 1, 1, 2},
		{0, 5, 0, 0},
		{2, 0, 3, 3}}
	result := matrixElementsSum(matrix)
	want := 9

	assert.Equal(t, want, result, "The result should be 9")
}

package searchmatrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {

	matrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

	res := searchMatrix(matrix, 12) // true
	assert.Equal(t, res, true)

	res = searchMatrixOpt(matrix, 12) // true
	assert.Equal(t, res, true)
}

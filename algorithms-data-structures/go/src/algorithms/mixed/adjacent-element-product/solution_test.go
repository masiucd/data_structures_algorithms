package adjacentelementproduct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdjacentElementProduct(t *testing.T) {
	input := []int{3, 6, -2, -5, 7, 3}
	expected := 21
	actual := adjacentElementsProduct(input)
	assert.Equal(t, expected, actual)

	input = []int{-1, -2}
	expected = 2
	actual = adjacentElementsProduct(input)
	assert.Equal(t, expected, actual)
}

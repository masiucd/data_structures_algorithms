package recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumRecursive(t *testing.T) {
	res := sumRecursive([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 15, res)

	res = sumRecursive([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, 21, res)

	res = sumRecursive([]int{0})
	assert.Equal(t, 0, res)
}

package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayDiff(t *testing.T) {
	assert.Equal(t, []int{2}, ArrayDiff([]int{1, 2}, []int{1}))
	assert.Equal(t, []int{2, 2}, ArrayDiff([]int{1, 2, 2}, []int{1}))
	assert.Equal(t, []int{1, 2, 2}, ArrayDiff([]int{1, 2, 2}, []int{}))
	assert.Equal(t, []int{}, ArrayDiff([]int{}, []int{1, 2}))
}

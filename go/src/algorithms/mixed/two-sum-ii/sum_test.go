package twosumii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumII(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{1, 2}
	assert.Equal(t, want, got, "Did not get expected result. Wanted %v, got %v", want, got)

	got = twoSum([]int{2, 3, 4}, 6)
	want = []int{1, 3}
	assert.Equal(t, want, got, "Did not get expected result. Wanted %v, got %v", want, got)

	got = twoSum([]int{-1, 0}, -1)
	want = []int{1, 2}
	assert.Equal(t, want, got, "Did not get expected result. Wanted %v, got %v", want, got)
}

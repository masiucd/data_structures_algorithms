package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	got := binarySearch([]int{1, 2, 3, 4, 5}, 3)
	want := 2
	assert.Equal(t, got, want)

	got = binarySearch([]int{1, 2, 3, 4, 5}, 6)
	want = -1
	assert.Equal(t, got, want)

	got = binarySearch([]int{1, 2, 3, 4, 5}, 1)
	want = 0
	assert.Equal(t, got, want)
}

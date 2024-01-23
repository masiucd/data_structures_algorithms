package almostincreasingsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlmostIncreasingSequence(t *testing.T) {

	got := almostIncreasingSequence([]int{1, 3, 2, 1})
	want := false
	assert.Equal(t, got, want)

	got = almostIncreasingSequence([]int{1, 3, 2})
	want = true
	assert.Equal(t, got, want)

	got = almostIncreasingSequence([]int{1, 2, 1, 2})
	want = false
	assert.Equal(t, got, want)
}

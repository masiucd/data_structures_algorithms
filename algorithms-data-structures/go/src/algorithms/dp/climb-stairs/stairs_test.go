package climbstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {

	got := ClimbStairs(5)
	want := 8
	assert.Equal(t, want, got)

	got = ClimbStairs(2)
	want = 2
	assert.Equal(t, want, got)

	got = ClimbStairs(3)
	want = 3
	assert.Equal(t, want, got)

	got = ClimbStairs(4)
	want = 5
	assert.Equal(t, want, got)

	got = ClimbStairs(12)
	want = 233
	assert.Equal(t, want, got)

}
func TestClimbStairsWithFib(t *testing.T) {

	got := ClimbingStairsWithFib(5)
	want := 8
	assert.Equal(t, want, got)

	got = ClimbingStairsWithFib(2)
	want = 2
	assert.Equal(t, want, got)

	got = ClimbingStairsWithFib(3)
	want = 3
	assert.Equal(t, want, got)

	got = ClimbingStairsWithFib(4)
	want = 5
	assert.Equal(t, want, got)

	got = ClimbingStairsWithFib(12)
	want = 233
	assert.Equal(t, want, got)

}

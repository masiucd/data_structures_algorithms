package tictactoechecker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTicTacToeChecker is a test function
func TestTicTacToeChecker(t *testing.T) {
	board := [3][3]int{
		{0, 0, 1},
		{0, 1, 2},
		{2, 2, 2},
	}
	actual := solution(board)
	expected := 2
	assert.Equal(t, expected, actual, "They should be equal")

	board = [3][3]int{
		{0, 0, 1},
		{0, 1, 2},
		{2, 2, 1},
	}
	actual = solution(board)
	expected = -1
	assert.Equal(t, expected, actual, "They should be equal")

	board = [3][3]int{
		{1, 0, 1},
		{0, 1, 2},
		{2, 2, 1},
	}
	actual = solution(board)
	expected = 1
	assert.Equal(t, expected, actual, "They should be equal")
}

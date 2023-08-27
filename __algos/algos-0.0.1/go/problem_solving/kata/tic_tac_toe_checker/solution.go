package tictactoechecker

var possibleWins = [][]int{
	{0, 1, 2}, // left to right
	{3, 4, 5}, // left to right
	{6, 7, 8}, // left to right
	{0, 3, 6}, // top down
	{1, 4, 7}, // top down
	{2, 5, 8}, // top down
	{0, 4, 8}, // left middle down
	{2, 4, 6}, // right middle down
}

func solution(board [3][3]int) int {

	if board[0][0] != 0 && board[0][0] == board[0][1] && board[0][1] == board[0][2] {
		return board[0][0]
	}

	if board[0][0] != 0 && board[0][0] == board[1][0] && board[1][0] == board[2][0] {
		return board[0][0]
	}

	if board[0][0] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}

	if board[2][0] != 0 && board[2][0] == board[2][1] && board[2][1] == board[2][2] {
		return board[2][0]
	}

	if board[0][2] != 0 && board[0][2] == board[1][2] && board[1][2] == board[2][2] {
		return board[0][2]
	}

	if board[1][0] != 0 && board[1][0] == board[1][1] && board[1][1] == board[1][2] {
		return board[1][0]
	}

	if board[0][1] != 0 && board[0][1] == board[1][1] && board[1][1] == board[2][1] {
		return board[0][1]
	}
	if board[2][0] != 0 && board[2][0] == board[1][1] && board[1][1] == board[0][2] {
		return board[2][0]
	}

	for _, v := range board {
		for _, k := range v {
			if k == 0 {
				return -1
			}
		}
	}
	return 0
}

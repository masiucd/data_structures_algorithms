package searchmatrix

// naive solution
func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, col := range row {
			if col == target {
				return true
			}
		}
	}
	return false
}

// optimized solution
// because the matrix is sorted, we can start at the top right corner
func searchMatrixOpt(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			// we going down the row
			row++
		} else {
			// we going left the column
			col--
		}
	}
	return false
}

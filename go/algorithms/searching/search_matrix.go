package searching

func SearchMatrix(matrix [][]int, target int) bool {
	for row := range matrix {
		for _, col := range matrix[row] {
			if col == target {
				return true
			}
		}
	}
	return false
}

func SearchMatrixTwo(matrix [][]int, target int) bool {
	left := 0
	horizontal := len(matrix[0])
	vertical := len(matrix)
	right := horizontal*vertical - 1

	for left <= right {
		mid := left + (right-left)/2
		i := mid / horizontal
		j := mid % horizontal

		if matrix[i][j] < target {
			left = mid + 1
		} else if matrix[i][j] > target {
			right = mid - 1
		} else {
			return true
		}

	}

	return false
}

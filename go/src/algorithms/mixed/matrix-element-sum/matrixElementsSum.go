package matrixelementsum

func matrixElementsSum(matrix [][]int) int {
	sum := 0
	for column := 0; column < len(matrix[0]); column++ {
		for row := 0; row < len(matrix); row++ {
			// Read the value of the room in the current column
			//  we are going column by column
			roomValue := matrix[row][column]
			if roomValue == 0 {
				break
			} else {
				sum += matrix[row][column]
			}
		}
	}
	return sum
}

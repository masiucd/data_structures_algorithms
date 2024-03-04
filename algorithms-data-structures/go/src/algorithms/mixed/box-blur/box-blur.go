package mixed

func boxBlur(image [][]int) [][]int {
	// create a new image
	newImage := make([][]int, len(image)-2)
	// -2 because we are going to skip the last 2 rows
	for row := 0; row < len(image)-2; row++ {
		rowLength := len(image[0]) // can be any row length 0, 1,
		// rowLength -2 = gives us the 2 first values in the row
		xs := make([]int, rowLength-2)
		for col := 0; col < rowLength-2; col++ {

			// top is the first 3 values in the row
			top := []int{image[row][col], image[row][col+1], image[row][col+2]}
			// middle is the second 3 values in the row
			middle := []int{image[row+1][col], image[row+1][col+1], image[row+1][col+2]}
			// bottom is the third 3 values in the row
			bottom := []int{image[row+2][col], image[row+2][col+1], image[row+2][col+2]}
			avg := average(top, middle, bottom)
			xs[col] = avg
		}
		newImage[row] = xs
	}

	return newImage
}

func average(top []int, middle []int, bottom []int) int {
	sum := 0
	for _, v := range top {
		sum += v
	}
	for _, v := range middle {
		sum += v
	}
	for _, v := range bottom {
		sum += v
	}
	return sum / 9
}

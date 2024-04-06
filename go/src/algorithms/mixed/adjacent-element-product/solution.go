package adjacentelementproduct

func adjacentElementsProduct(inputArray []int) int {
	max := inputArray[0] * inputArray[1]
	for i := 1; i < len(inputArray)-1; i++ {
		if inputArray[i]*inputArray[i+1] > max {
			max = inputArray[i] * inputArray[i+1]
		}
	}
	return max
}

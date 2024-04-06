package arrays

func slidingWindow(xs []int, windowSize int) int {
	temp := 0
	for i := 0; i < windowSize; i++ {
		temp += xs[i]
	}
	maxSum := temp
	for i := windowSize; i < len(xs); i++ {
		//Remove the first element of the window
		//Add the next element to the window
		temp = temp - xs[i-windowSize] + xs[i]
		maxSum = max(maxSum, temp)
	}
	return maxSum
}

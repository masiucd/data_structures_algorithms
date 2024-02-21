package mixed

func avoidObstacles(inputArray []int) int {
	i := 1
	for {
		// function to check if every element in the array is not divisible by i (obstacle)
		fn := func(n int) bool {
			return n%i != 0
		}
		// if every element in the array is not divisible by i (obstacle)
		if everyStrict(inputArray, fn) {
			// return the obstacle
			return i
		}
		i++
	}
}

func everyStrict[T any](inputArray []T, fn func(n T) bool) bool {
	i := 0
	for _, value := range inputArray {
		if fn(value) {
			i++
		}
	}
	return i == len(inputArray)
}

func avoidObstaclesV2(inputArray []int) int {
	sum := 1
	for i := 0; i < len(inputArray); i++ {
		// we hit an obstacle
		// so we reset the index and increment the sum
		if inputArray[i]%sum == 0 {
			i = -1
			sum++
		}
	}
	return sum
}

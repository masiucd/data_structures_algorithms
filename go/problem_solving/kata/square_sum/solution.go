package square_sum

func squareSum(numbers []int) (result int) {
	for _, v := range numbers {
		result += v * v
	}

	return result
}

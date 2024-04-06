package mixed

func arrayChange(input []int) int {
	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] <= input[i-1] {
			difference := input[i-1] - input[i]
			count += difference + 1
			input[i] = input[i-1] + 1
		}
	}
	return count
}

package smallestinteger

func solution(numbers []int) int {
	var smallestinteger int = numbers[0]
	for _, n := range numbers {
		if n < smallestinteger {
			smallestinteger = n
		}
	}
	return smallestinteger
}

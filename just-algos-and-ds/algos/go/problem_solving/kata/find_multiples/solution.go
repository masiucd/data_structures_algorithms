package find_multiples

func solution(n, limit int) []int {
	var result []int
	for i := 1; i <= limit; i++ {
		if i%n == 0 {
			result = append(result, i)
		}
	}
	return result
}

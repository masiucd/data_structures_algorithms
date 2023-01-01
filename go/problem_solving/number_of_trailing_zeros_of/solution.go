package numberoftrailingzerosof

func solution(n int) int {
	var count int
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}

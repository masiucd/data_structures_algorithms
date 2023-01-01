package numberoftrailingzerosof

// https://mathworld.wolfram.com/Factorial.html
func solution(n int) int {
	var count int
	for n >= 5 {
		n /= 5
		count += n
	}
	return count
}

func solutionRec(n int) int {
	if n == 0 {
		return 0
	}
	return n/5 + solutionRec(n/5)
}

package sum_of_sequence

func solution(start, end, step int) int {
	if start > end {
		return 0
	}
	var result int
	for i := start; i <= end; i += step {
		result += i
	}
	return result
}

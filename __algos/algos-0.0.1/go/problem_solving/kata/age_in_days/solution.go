package ageindays

func solution(days int) []int {
	years := days / 365
	days -= 365 * years
	months := days / 30
	days = days - months*30
	return []int{years, months, days}
}

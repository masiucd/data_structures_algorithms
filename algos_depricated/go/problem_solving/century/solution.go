package century

func century(year int) int {
	if year%100 > 0 {
		return (year / 100) + 1
	}
	return year / 100
}

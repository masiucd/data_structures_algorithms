package multiple3and5

func multiple3And5(number int) int {
	if number < 0 {
		return 0
	}
	var result int
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	return result
}

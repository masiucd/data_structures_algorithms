package numbers

func sumOfAllSquares(n int) int {
	var result int = 1
	for i := 2; i <= n; i++ {
		result += i * i
	}
	return result
}

func sumOfAllSquaresLevel2(m, n int) int {
	result := 0
	for i := m; i <= n; i++ {
		result += i * i
	}
	return result
}

package recursion

func sumSeries(n int) (result int) {
	for i := 0; i <= n; i++ {
		result += i
	}
	return result
}

func sumSeriesRecursive(n int) (result int) {
	if n == 0 {
		return n
	}
	return n + sumSeries(n-1)
}

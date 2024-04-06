package printsquares

func printSquaresTopDown(n int) string {
	var result string
	for i := 0; i < n; i++ {
		x := ""
		for j := i; j < n; j++ {
			x += "#"
		}
		result += x + "\n"
	}
	return result
}

func printSquaresDownTop(n int) string {
	var result string
	for i := 0; i < n; i++ {
		x := ""
		for j := 0; j < i; j++ {
			x += "#"
		}
		result += x + "\n"
	}
	return result
}

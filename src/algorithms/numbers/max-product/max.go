package numbers

func maxProductLevel1(xs []int) int {
	var index int
	max := 1
	for i := 0; i < len(xs)-1; i++ {
		n := xs[i] * xs[i+1]
		if n > max {
			max = n
			index = i
		}
	}
	return index
}

func maxProductLevel2(xs []int, n int) int {
	var index int
	max := product(xs[:n]) //
	for i := 0; i < len(xs)-n+1; i++ {
		currentProduct := product(xs[i : n+i])
		if currentProduct > max {
			max = currentProduct
			index = i
		}
	}
	return index
}

func product(xs []int) int {
	result := 1
	for _, x := range xs {
		result *= x
	}
	return result
}

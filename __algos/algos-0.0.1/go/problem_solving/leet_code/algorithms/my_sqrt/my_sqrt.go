package my_sqrt

// Using binary search
func mySqrt(x int) int {
	start, end := 0, x
	for start <= end {
		middle := (start + end) / 2
		if middle*middle < x {
			start = middle + 1
		} else if middle*middle > x {
			end = middle - 1
		} else {
			return middle
		}

	}

	return end
}

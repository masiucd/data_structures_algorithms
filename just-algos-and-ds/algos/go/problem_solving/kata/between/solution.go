package between

// a < b
func between(a, b int) []int {
	var res []int
	for i := a; i <= b; i++ {
		res = append(res, i)
	}
	return res
}

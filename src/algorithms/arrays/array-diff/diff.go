package arrays

func listDiff(a, b []int) []int {
	return filter(a, func(n int) bool {
		return !includes(b, func(x int) bool {
			return n == x
		})
	})
}

func listDiffV2(a, b []int) []int {
	store := make(map[int]bool)
	for _, x := range b {
		store[x] = true
	}
	var result []int
	for _, x := range a {
		if !store[x] {
			result = append(result, x)
		}
	}
	return result
}

func filter(xs []int, predicate func(n int) bool) []int {
	var result []int
	for _, x := range xs {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

func includes(xs []int, predicate func(n int) bool) bool {
	for _, x := range xs {
		if predicate(x) {
			return true
		}
	}
	return false
}

package kata

func ArrayDiff(a, b []int) []int {
	var result []int
	for _, aItem := range a {
		if !contains(b, aItem) {
			result = append(result, aItem)
		}
	}

	return result
}

func contains(b []int, aItem int) bool {
	for _, bItem := range b {
		if bItem == aItem {
			return true
		}
	}

	return false
}

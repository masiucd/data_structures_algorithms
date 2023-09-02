package contain_duplicate

func containsDuplicate(nums []int) bool {
	mapper := make(map[int]bool)
	for _, v := range nums {
		if _, ok := mapper[v]; ok {
			return true
		}
		mapper[v] = true
	}
	return false
}

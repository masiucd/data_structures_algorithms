package search

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
		if v > target {
			return i
		}
	}
	return len(nums)
}

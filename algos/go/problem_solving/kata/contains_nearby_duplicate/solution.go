package containsnearbyduplicate

func solution(nums []int, k int) bool {
	store := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if storedIndex, ok := store[nums[i]]; ok && i-storedIndex <= k {
			return true
		}
		store[nums[i]] = i
	}
	return false
}

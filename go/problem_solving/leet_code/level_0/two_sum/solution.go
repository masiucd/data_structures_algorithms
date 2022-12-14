package two_sum

func twoSumOne(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(
					result,
					i,
					j,
				)
				return result
			}
		}
	}
	return result
}

func twoSumTwo(nums []int, target int) []int {
	store := make(map[int]int)
	for currentIndex, v := range nums {
		if foundIndex, ok := store[target-v]; ok {
			return []int{foundIndex, currentIndex}
		}
		store[v] = currentIndex
	}
	return []int{}
}

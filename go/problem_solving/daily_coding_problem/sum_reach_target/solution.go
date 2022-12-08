package sum_reach_target

//naive solution
func sumReachTarget(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == k {
				return true
			}
		}
	}
	return false
}

func sumReachTarget2(nums []int, k int) bool {
	store := make(map[int]int)
	for i, v := range nums {
		total := k - v
		if _, ok := store[total]; ok {
			return true
		}
		store[v] = i
	}
	return false
}

//This problem was recently asked by Google.
//
//Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
//
//For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
//
//Bonus: Can you do this in one pass?

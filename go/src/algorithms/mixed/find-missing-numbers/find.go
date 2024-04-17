package mixed

// Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.
//
//	findDisappearedNumbers function takes in a slice of integers and return a slice of integers
func findDisappearedNumbers(nums []int) []int {

	numCount := make(map[int]bool)
	for _, v := range nums {
		numCount[v] = true
	}

	var result []int

	for i := 1; i <= len(nums); i++ {
		if _, ok := numCount[i]; !ok {
			result = append(result, i)
		}
	}

	return result
}

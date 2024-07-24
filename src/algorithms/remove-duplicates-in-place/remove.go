package algorithms

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

// removeDuplicates removes duplicates in place from a sorted array.
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var index = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

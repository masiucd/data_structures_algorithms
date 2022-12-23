package recursion

func sumDivideAndConcur(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	middle := len(nums) / 2
	first := sumDivideAndConcur(nums[:middle])
	second := sumDivideAndConcur(nums[middle:])
	return first + second
}

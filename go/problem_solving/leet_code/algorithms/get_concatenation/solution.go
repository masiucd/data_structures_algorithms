package getconcatenation

func getConcatenation(nums []int) []int {
	var result []int = make([]int, len(nums)*2)
	for i := 0; i < len(nums)*2; i++ {
		if i <= len(nums)-1 {
			result[i] = nums[i]
		} else {
			result[i] = nums[i-(len(nums))]
		}
	}
	return result
}

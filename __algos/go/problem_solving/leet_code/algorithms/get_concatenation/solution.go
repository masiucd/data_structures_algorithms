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

func getConcatenation2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i := 0; i < n; i++ {
		ans[i], ans[i+n] = nums[i], nums[i]
	}
	return ans
}

func getConcatenation3(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i := 0; i < n*2; i++ {
		ans[i] = nums[i%n]
	}
	return ans
}

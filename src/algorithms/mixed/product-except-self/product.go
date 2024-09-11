package mixed

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	size := len(nums)
	result := make([]int, size)
	left, right := make([]int, size), make([]int, size)
	left[0] = 1
	right[size-1] = 1

	for i := 1; i < size; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := size - 1; i > 0; i-- {
		right[i-1] = right[i] * nums[i]
	}

	for i := range size {
		result[i] = left[i] * right[i]
	}

	return result
}

func productExceptSelfV2(nums []int) []int {
	n := len(nums)
	output := make([]int, n)
	left, right := 1, 1

	for i := 0; i < n; i++ {
		output[i] = left
		left *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}

	return output
}

package mixed

// productExceptSelf returns an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
// It uses two auxiliary arrays to store the product of elements to the left and right of each element.
// nums: input array of integers
// returns: array of integers where each element is the product of all elements in the input array except the element at the same index
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

// productExceptSelfV2 returns an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
// It uses a more space-efficient approach by avoiding the use of auxiliary arrays and instead using two variables to keep track of the product.
// nums: input array of integers
// returns: array of integers where each element is the product of all elements in the input array except the element at the same index
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

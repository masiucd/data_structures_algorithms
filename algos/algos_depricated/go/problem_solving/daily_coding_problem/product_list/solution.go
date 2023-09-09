package product_list

func productOfList(nums []int) int {
	var result = 1
	for _, v := range nums {
		result = result * v
	}
	return result
}

func productList(nums []int) []int {
	var result []int
	var product = productOfList(nums)
	for _, v := range nums {
		result = append(result, product/v)
	}
	return result
}

func productListTwo(nums []int) []int {
	var result []int
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = 1
	right[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := range nums {
		result = append(result, left[i]*right[i])
	}

	return result
}

func productListThree(nums []int) []int {
	var result = make([]int, len(nums))

	left := 1
	for i := 0; i < len(nums); i++ {
		result[i] = left
		left *= nums[i]
	}

	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

//Good morning! Here's your coding interview problem for today.
//
//This problem was asked by Uber.
//
//Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
//
//For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
//
//Follow-up: what if you can't use division?

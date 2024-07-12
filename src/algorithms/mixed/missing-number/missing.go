package mixed

func missingNumber(nums []int) int {
	var store = make(map[int]bool)
	for _, n := range nums {
		store[n] = true
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := store[i]; !ok {
			return i
		}
	}
	return -1
}

// Using the formula for the sum of the first n numbers
// sum = n*(n+1)/2 - sum = missing number
func missingNumberV2(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return len(nums)*(len(nums)+1)/2 - sum
}

func missingNumberV3(nums []int) int {
	missing := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		// missing is the sum of the first n numbers
		missing += i + 1
		// sum is the sum of the numbers in the array
		sum += nums[i]

	}
	// missing - sum is the missing number
	return missing - sum
}

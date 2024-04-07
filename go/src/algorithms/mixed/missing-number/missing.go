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

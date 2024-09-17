package mixed

func findMaxAverage(nums []int, k int) float64 {
	var tempsum float64
	for i := 0; i < k; i++ {
		tempsum += float64(nums[i])
	}
	maxSum := tempsum
	for i := k; i < len(nums); i++ {
		maxSum = maxSum - float64(nums[i-k]) + float64(nums[i])
		tempsum = max(tempsum, maxSum)
	}
	return float64(tempsum) / float64(k)
}

package sorting

func QuickSort(nums []int) []int {
	return quickSortHelper(nums, 0, len(nums)-1)
}

func quickSortHelper(nums []int, i, j int) []int {
	if j-i+1 <= 1 {
		return nums
	}
	pivot := nums[j]
	left := i

	for s := i; s < j; s++ {
		if nums[s] < pivot {
			temp := nums[i]
			nums[left] = nums[s]
			nums[s] = temp
			left++
		}
	}
	nums[j] = nums[left]
	nums[left] = pivot

	quickSortHelper(nums, i, left-1)
	quickSortHelper(nums, left+1, j)
	return nums

}

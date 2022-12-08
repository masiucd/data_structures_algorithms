package sorting

func BucketSort(nums []int) []int {
	// Assuming arr only contains 0, 1 or 2
	var counts = []int{0, 0, 0}

	for _, num := range nums {
		counts[num] += 1
	}

	i := 0
	for index := 0; index < len(counts); index++ {
		for j := 0; j < counts[index]; j++ {
			//	override
			nums[i] = index
			i++
		}
		println(i)
	}

	return nums
}

package mixed

func singleNumber(nums []int) int {
	intCount := make(map[int]int)
	for _, n := range nums {
		intCount[n]++
	}

	for k, v := range intCount {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumberV2(nums []int) int {
	var mask int
	for _, n := range nums {
		mask ^= n
	}
	return mask
}

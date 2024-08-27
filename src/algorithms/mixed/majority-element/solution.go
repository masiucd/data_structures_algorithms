package mixed

func majorityElement(nums []int) int {
	counter := make(map[int]int)
	for _, x := range nums {
		counter[x]++
	}
	max := 0
	num := 0
	for k, v := range counter {
		if v > max {
			max = v
			num = k
		}
	}

	return num
}

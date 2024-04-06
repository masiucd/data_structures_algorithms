package mixed

func findErrorNums(nums []int) []int {
	numCount := make(map[int]int)
	for _, n := range nums {
		numCount[n]++
	}

	var duplicate, missing int
	// since zero will not be in the array, we start from 1
	for i := 1; i <= len(nums); i++ {
		if v, ok := numCount[i]; ok {
			if v == 2 {
				// if the number appears twice it is the duplicate
				duplicate = i
			}
		} else {
			// if it does note exists in the map it is the missing number
			missing = i
		}
	}
	return []int{duplicate, missing}
}

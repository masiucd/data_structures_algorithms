package mixed

func moveZeroes(nums []int) {
	var amountOfZeros int
	for _, n := range nums {
		if n == 0 {
			amountOfZeros++
		}
	}

	var xs []int
	for _, n := range nums {
		if n != 0 {
			xs = append(xs, n)
		}
	}

	for amountOfZeros > 0 {
		xs = append(xs, 0)
		amountOfZeros--
	}

	for i := range nums {
		nums[i] = xs[i]
	}
}

func moveZeroesV2(nums []int) {
	var lastFoundZeroAt int
	for i := range nums {
		if nums[i] != 0 {
			nums[lastFoundZeroAt] = nums[i]
			lastFoundZeroAt++
		}
	}
	for i := lastFoundZeroAt; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroesV3(nums []int) {
	var lastFoundZeroAt int
	for i := range nums {
		// if the current number is not zero we swap it with the last found zero, and increment the lastFoundZeroAt
		// we move the lastFoundZeroAt to the next position
		if nums[i] != 0 {
			// swap
			nums[lastFoundZeroAt], nums[i] = nums[i], nums[lastFoundZeroAt]
			lastFoundZeroAt++
		}
	}
}

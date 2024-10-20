package numbers

import "slices"

func level1ConsecutiveCheck(numberList []int) []int {
	slices.Sort(numberList)
	for i := 0; i < len(numberList)-1; i++ {
		diff := numberList[i+1] - numberList[i]
		if diff != 1 {
			return []int{}
		}
	}
	return numberList
}

func level2ConsecutiveCheck(numberList []int) []int {
	slices.Sort(numberList)
	min, max := numberList[0], numberList[len(numberList)-1]
	xs, missingNums := []int{}, []int{}
	for i := min; i <= max; i++ {
		xs = append(xs, i)
	}
	for _, x := range xs {
		if !includes(numberList, func(n int) bool {
			return x == n
		}) {
			missingNums = append(missingNums, x)
		}
	}
	return missingNums
}

func includes(xs []int, fn func(int) bool) bool {
	for _, x := range xs {
		if fn(x) {
			return true
		}
	}
	return false
}

func level2ConsecutiveCheckV2(numberList []int) []int {
	if len(numberList) == 0 {
		return []int{}
	}

	// Create a map to track the presence of each number
	numberMap := make(map[int]bool)
	for _, num := range numberList {
		numberMap[num] = true
	}

	// Find the minimum and maximum values in the list
	min, max := numberList[0], numberList[0]
	for _, num := range numberList {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	// Collect missing numbers
	missingNums := []int{}
	for i := min; i <= max; i++ {
		if !numberMap[i] {
			missingNums = append(missingNums, i)
		}
	}

	return missingNums
}

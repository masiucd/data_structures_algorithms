package numbers

import (
	"go-ds/src/lib/fp"
	"sort"
)

// https://leetcode.com/problems/minimum-number-game/

// numberGame returns a new array where the minimum number is swapped with the next minimum number.
// It uses getMin to get the minimum number and its index.
func numberGame(nums []int) []int {
	var result []int
	a, b := 0, 0

	for len(nums) > 0 {
		min, index := getMin(nums)
		a = min
		// we can not filter out if it is duplicate instead we need to remove by index
		nums = filterOutNumByIndex(index)(nums)

		min, index = getMin(nums)
		b = min
		// we can not filter out if it is duplicate instead we need to remove by index
		nums = filterOutNumByIndex(index)(nums)

		result = append(result, b)
		result = append(result, a)

	}

	return result
}

func getMin(nums []int) (int, int) {
	min := nums[0]
	index := 0
	for i, n := range nums {
		if n < min {
			min = n
			index = i
		}
	}
	return min, index
}

func filterOutNumByIndex(positionToBeREmoved int) func([]int) []int {
	return fp.Filter(func(i int) bool {
		return i != positionToBeREmoved
	})
}

// numberGameV2 sorts the input array and swaps adjacent elements.
// It then returns the modified array.
func numberGameV2(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
	return nums
}

// numberGameV3 uses modulo to determine if the index is even or odd.
// If the index is even, it appends the next number in the sorted array.
// If the index is odd, it appends the previous number in the sorted array.
func numberGameV3(nums []int) []int {
	sort.Ints(nums)
	var result []int
	for i := 0; i < len(nums); i++ {
		if i == 0 || i%2 == 0 {
			result = append(result, nums[i+1])
		} else {
			result = append(result, nums[i-1])
		}
	}
	return result
}

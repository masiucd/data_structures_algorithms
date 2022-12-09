package findsingledigit

// Good morning! Here's your coding interview problem for today.

// This problem was asked by Google.

// Given an array of integers where every integer occurs three times except for one integer, which only occurs once, find and return the non-duplicated integer.

// For example, given [6, 1, 3, 3, 3, 6, 6], return 1. Given [13, 19, 13, 13], return 19.

func solution(nums []int) int {
	numsMap := make(map[int]int, len(nums))
	for _, num := range nums {
		if _, ok := numsMap[num]; !ok {
			numsMap[num] = 1
		} else {
			numsMap[num]++
		}
	}
	for k, v := range numsMap {
		if v == 1 {
			return k
		}
	}
	return -1
}

package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 4}
	nums = []int{1, 1}
	// nums = []int{2, 2} // 2 1
	res := findErrorNums(nums)
	fmt.Println("res", res) // [2,3]
}

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

// In this corrected version, the findErrorNums function creates a map to count the occurrences of each number in the array. It then iterates from 1 to len(nums), checking the count of each number in the map. If a number appears twice, it's the duplicate. If a number doesn't appear in the map, it's the missing number.

/*
You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.


*/

package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	store := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if storedIndex, ok := store[nums[i]]; ok && i-storedIndex <= k {
			return true
		}
		store[nums[i]] = i
	}
	return false
}

package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums)) // 2
}

func missingNumber(nums []int) int {
	var store = make(map[int]bool)
	for _, n := range nums {
		store[n] = true
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := store[i]; !ok {
			return i
		}
	}
	return -1
}

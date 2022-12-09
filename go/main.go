package main

import "fmt"

func main() {
	fmt.Println(solution([]int{6, 1, 3, 3, 3, 6, 6}))
	fmt.Println(solution([]int{13, 19, 13, 13}))
}
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

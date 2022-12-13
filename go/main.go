package main

import "fmt"

func main() {

	fmt.Println(solution(2, 6, 2)) // 12
	fmt.Println(solution(1, 5, 1)) // 15
	fmt.Println(solution(1, 5, 3)) // 15

}

func solution(start, end, step int) int {
	if start > end {
		return 0
	}
	var result int
	for i := start; i <= end; i += step {
		result += i
	}
	return result
}

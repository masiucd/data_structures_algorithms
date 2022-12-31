package main

import "fmt"

func main() {
	fmt.Println(perimeter(5)) // 80
	fmt.Println(perimeter(7))
	fmt.Println(perimeter(20))
}

func perimeter(n int) int {
	a, b := 1, 1
	for i := 0; i < n+1; i++ {
		a, b = b, a+b
	}
	return 4 * (b - 1)
}

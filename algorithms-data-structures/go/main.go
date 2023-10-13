package main

import (
	"fmt"
)

func main() {

	apa := check(1)
	fmt.Println(apa)
}

func check(n int) int {
	switch n {
	case 1:
		fmt.Println("ENTER")
		return n * 10
	default:
		return n * 100
	}
	return 0
}

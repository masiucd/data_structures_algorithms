package main

import (
	"fmt"
	"time"
)

func main() {
	memo := make(map[int]int)
	start := time.Now()
	fmt.Println(fibFast(100, memo))
	end := time.Since(start)
	fmt.Println("end", end)

	// start = time.Now()
	// fmt.Println(fib(50))
	// end = time.Since(start)
	// fmt.Println("end", end)

}

// 0 1 2 3 4 5 6 7 8 9 10
// 1 2 3 5 7

func fib(n int) int {
	if n <= 2 {
		return n
	}

	return fib(n-2) + fib(n-1)
}


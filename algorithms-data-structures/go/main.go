package main

import "fmt"

func main() {

	res := boxBlur([][]int{
		{7, 4, 0, 1},
		{5, 6, 2, 2},
		{6, 10, 7, 8},
		{1, 4, 2, 0},
	})

	fmt.Println("res", res)
}

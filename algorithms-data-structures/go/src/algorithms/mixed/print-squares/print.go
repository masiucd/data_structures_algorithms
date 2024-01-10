package printsquares

import "fmt"

func printSquaresTopDown(n int) {
	for i := 0; i < n; i++ {
		x := ""
		for j := i; j < n; j++ {
			x += "#"
		}
		fmt.Println(x)
	}
}

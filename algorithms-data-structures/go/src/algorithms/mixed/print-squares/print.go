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

func printSquaresDownTop(n int) {
	for i := 0; i < n; i++ {
		x := ""
		for j := 0; j < i; j++ {
			x += "#"
		}
		fmt.Println(x)
	}
}

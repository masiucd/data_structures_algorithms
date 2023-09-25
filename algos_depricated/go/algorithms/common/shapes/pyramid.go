package shapes

import "fmt"

func MakeShapeUpDown(n int) {
	var s string
	for i := 0; i < n; i++ {
		s += "\n"
		for j := i; j < n; j++ {
			s += "#"
		}
	}
	fmt.Println(s)
}

func MakeShapeDownUp(n int) {
	var s string
	for i := n; i > 0; i-- {
		s += "\n"
		for j := i; j > 0; j-- {
			s += "#"
		}
	}
	fmt.Println(s)
}

package bits

import "fmt"

func printBinaryRepresentation(num int) {
	if num == 0 {
		return
	}
	printBinaryRepresentation(num / 2)
	fmt.Printf("%d", num%2)
}

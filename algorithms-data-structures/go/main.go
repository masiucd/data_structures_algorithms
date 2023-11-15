package main

import "fmt"

func main() {

	res := reverseString("marcell")
	fmt.Println(res)
	//
	//x := "Marcell"
	//fmt.Println(x)
	//fmt.Println(x[1:len(x)] + string(x[0]))
}

func reverseString(input string) string {
	if len(input) == 0 {
		return ""
	}
	return reverseString(input[1:]) + string(input[0])
}

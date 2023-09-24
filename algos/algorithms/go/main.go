package main

import "fmt"

func main() {
	res := isValid("()")
	res = isValidTwo("()[]{}")
	fmt.Println(res)

	// xs := []int{1, 2, 3, 4}
	// fmt.Println(xs)
	// xs = xs[:len(xs)-1]
	// fmt.Println(xs)

}

func isValid(s string) bool {
	var stack []string
	var closed map[string]string = map[string]string{")": "(", "]": "[", "}": "{"}
	for _, v := range s {
		if val, ok := closed[string(v)]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == val {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(v))
		}
	}

	return len(stack) == 0

}

func isValidTwo(s string) bool {
	var stack []string

	for _, c := range s {
		strc := string(c)
		if strc == "(" {
			stack = append(stack, ")")
		} else if strc == "[" {
			stack = append(stack, "]")
		} else if strc == "{" {
			stack = append(stack, "}")
		} else if len(stack) == 0 || stack[len(stack)-1] != strc {
			return false
		} else {
			// pop from stack
			stack = stack[:len(stack)-1]
		}
	}
	// if stack is empty, then it's valid
	return len(stack) == 0
}

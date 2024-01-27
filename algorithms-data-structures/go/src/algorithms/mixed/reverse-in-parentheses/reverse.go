package reverseinparentheses

func reverseInParentheses(inputString string) string {
	stack := []rune{}
	for _, ch := range inputString {
		if ch == ')' {
			temp := []rune{}
			// Pop characters from stack till '(' is found
			for stack[len(stack)-1] != '(' {
				temp = append(temp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// Pop the '('
			stack = stack[:len(stack)-1]
			// Push the reversed string back to stack
			stack = append(stack, temp...)
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}

# Reverse in parentheses

## Description

Write a function that reverses characters in (possibly nested) parentheses in the input string.

## Input/Output

- `[input]` string `inputString`

  A string consisting of lowercase English letters and the characters `(` and `)`. It is guaranteed that all parentheses in `inputString` form a regular bracket sequence.

  Constraints:
  `0 ≤ inputString.length ≤ 50.`

- `[output]` string

  Return `inputString`, with all the characters that were in parentheses reversed.

### About the algorithm

```go
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
```

This function works by iterating over the input string character by character. If the current character is not a closing parenthesis, it is pushed onto the stack. If the current character is a closing parenthesis, the function starts popping characters from the stack and appending them to a temporary slice until it encounters an opening parenthesis. The opening parenthesis is then popped from the stack, and the characters in the temporary slice (which are now in reverse order) are pushed back onto the stack. Finally, the function returns the string representation of the stack, which is the input string with characters in parentheses reversed.

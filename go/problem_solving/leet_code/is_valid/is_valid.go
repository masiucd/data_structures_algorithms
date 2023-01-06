package is_valid

var pairs = map[rune]rune{
	']': '[',
	')': '(',
	'}': '{',
}

func isValid(s string) bool {
	var stack []rune
	for _, v := range s {
		if match, ok := pairs[v]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}

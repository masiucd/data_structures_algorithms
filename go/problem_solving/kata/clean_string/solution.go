package clean_string

func solution(s string) string {
	var stack []rune
	for _, r := range s {
		if r == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}

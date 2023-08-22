package recursion

func reverseString(str string) string {
	if len(str) == 0 {
		return str
	}
	head := string(str[0])
	tail := str[1:]
	return reverseString(tail) + head
}

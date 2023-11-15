package reverse

func reverseString(input string) string {
	if len(input) == 0 {
		return ""
	}
	return reverseString(input[1:]) + string(input[0])
}

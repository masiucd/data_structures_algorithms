package mcd_strings

func reverseString(input string) string {
	reversed := ""
	for _, v := range input {
		reversed = string(v) + reversed
	}
	return reversed
}

func reverseStringRec(input string) string {
	if len(input) == 0 {
		return ""
	}

	return reverseString(input[1:]) + input[:1]
}

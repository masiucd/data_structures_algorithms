package recursion

func isPalindrome(input string) bool {
	if len(input) == 0 || len(input) == 1 {
		return true
	}
	first := input[0]
	last := input[len(input)-1]
	middle := input[1 : len(input)-1]
	return first == last && isPalindrome(middle)
}

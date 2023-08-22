package reverselist

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


// Reverse reverses a slice of any type
func Reverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, ele := range s {
					r[l-i-1] = ele
	}
	return r
}

func reverseList2(nums []int) []int {
	var result []int
	for _, v := range nums {
		result = append([]int{v}, result...)
	}
	return result
}

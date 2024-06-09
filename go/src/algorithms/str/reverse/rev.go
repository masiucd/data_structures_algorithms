package reverse

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// mutates the input slice
func reverseString(s []byte) {
	size := len(s)
	for i := 0; i < size/2; i++ {
		s[i], s[size-i-1] = s[size-i-1], s[i]
	}
}

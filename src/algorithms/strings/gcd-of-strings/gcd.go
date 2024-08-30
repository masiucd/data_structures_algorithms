package gcdofstrings

func gcdOfStrings(s1, s2 string) string {
	if s1+s2 != s2+s1 {
		return ""
	}
	gcdLen := gcd(len(s1), len(s2))
	return s1[:gcdLen]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

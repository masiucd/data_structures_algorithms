package is_triangle

func isTriangle(a, b, c int) bool {
	if a+b <= c {
		return false
	}
	if a+c <= b {
		return false
	}
	if b+c <= a {
		return false
	}
	return true
}

//A triangle is valid if sum of its two sides is greater than the third side. If three sides are a, b and c, then three conditions should be met.

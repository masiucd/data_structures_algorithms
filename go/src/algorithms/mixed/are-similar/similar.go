package aresimilar

func areSimilar(a []int, b []int) bool {
	var c, d []int
	for i := range a {
		// so when we hit a un-match we do some stuff to c and d
		if a[i] != b[i] {
			c = append(c, a[i])
			d = append(d, b[i])
		}
	}
	// reverse c because we want to compare it to d, we could choose d as well
	reverseSlice(c)
	// if c is empty or has 1 or 2 elements and they are equal to d then we are good
	return len(c) <= 2 && slicesAreEqual(c, d)
}

func slicesAreEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func reverseSlice(a []int) {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - i - 1
		a[i], a[j] = a[j], a[i]
	}
}

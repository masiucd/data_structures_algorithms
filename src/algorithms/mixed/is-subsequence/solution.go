package mixed

func isSubsequence(s, t string) bool {
	i, j := 0, 0
	for j < len(t) {
		if i < len(s) && s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

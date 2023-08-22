package recursion

func concat(xs []string) string {
	if len(xs) == 0 {
		return ""
	}
	head := xs[0]
	tail := xs[1:]
	return head + concat(tail)
}

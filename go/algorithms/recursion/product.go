package recursion

func product(xs []int) int {
	if len(xs) == 0 {
		return 1
	}
	head := xs[0]
	tail := xs[1:]
	return head * product(tail)
}

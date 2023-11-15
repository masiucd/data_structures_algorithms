package sum

func sumRecursive(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	head := xs[0]
	tail := xs[1:]
	return head + sumRecursive(tail)
}

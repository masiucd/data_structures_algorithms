package recursion

func sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	head := numbers[0]
	tail := numbers[1:]
	return head + sum(tail)
}

package main

import "strconv"

func main() {
	n := calPoints([]string{"5", "2", "C", "D", "+"})
	println(n)
}

func sumList(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}

func calPoints(operations []string) int {
	stack := make([]int, 0)
	for _, v := range operations {
		switch v {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, 2*stack[len(stack)-1])
		case "C":
			stack = stack[:len(stack)-1]
		default:
			n, _ := strconv.Atoi(v)
			stack = append(stack, n)
		}
	}
	return sumList(stack)
}

// An integer x.
// Record a new score of x.
// '+'.
// Record a new score that is the sum of the previous two scores.
// 'D'.
// Record a new score that is the double of the previous score.
// 'C'.
// Invalidate the previous score, removing it from the record.

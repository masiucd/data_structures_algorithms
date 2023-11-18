package main

func main() {
	res := climbStairs(5)
	println(res)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	first := 1
	second := 2

	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}

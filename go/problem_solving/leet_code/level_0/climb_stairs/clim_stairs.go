package climb_stairs

// https://leetcode.com/problems/climbing-stairs
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a, b, c := 1, 2, 0
	for i := 2; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func climbStairsTwo(n int) int {
	var arr [46]int
	arr[0] = 1
	arr[1] = 2
	for i := 3; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func climbStairsThree(n int) int {
	a, b := 1, 1
	for i := 0; i < n-1; i++ {
		temp := a
		a = a + b
		b = temp
	}
	return a
}

package perimeter

// https://www.codewars.com/kata/559a28007caad2ac4e000083/go
func perimeter(n int) int {
	a, b := 1, 1
	for i := 0; i < n+1; i++ {
		a, b = b, a+b
	}
	return 4 * (b - 1)
}

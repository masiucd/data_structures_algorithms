package fizzbuzz

import "strconv"

func makeFizzBuzList(n int) []string {
	var result []string = make([]string, n)
	for i := 0; i < n; i++ {
		switch {
		case i%15 == 0:
			result[i] = "FizzBuzz"
		case i%3 == 0:
			result[i] = "Fizz"
		case i%5 == 0:
			result[i] = "Buzz"
		default:
			result[i] = strconv.Itoa(i)
		}
	}
	return result
}

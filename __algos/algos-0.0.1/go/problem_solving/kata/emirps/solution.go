package emirps

import (
	"math"
	"strconv"
)

func solution(n int) [3]int {
	if n <= 13 {
		return [3]int{0, 0, 0}
	}
	var emirps []int
	for i := 12; i <= n; i++ {
		if isPrime(i) && isPrime(reverseNumber(i)) && i != reverseNumber(i) {
			emirps = append(emirps, i)
		}
	}
	return [3]int{len(emirps), emirps[len(emirps)-1], sum(emirps)}
}

func sum(nums []int) (result int) {
	for _, v := range nums {
		result += v
	}
	return result
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func reverseNumber(n int) int {
	stringNumber := reverseString(strconv.Itoa(n))
	reversedNumber, _ := strconv.Atoi(stringNumber)
	return reversedNumber
}

func reverseString(s string) string {
	reversed := ""
	for _, v := range s {
		reversed = string(v) + reversed
	}
	return reversed
}

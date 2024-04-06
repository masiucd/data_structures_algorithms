package islucky

import (
	"strconv"
	"strings"
)

func isLucky(n int) bool {
	var s string = strconv.Itoa(n)
	xs := strings.Split(s, "")
	ys, fs := xs[:len(xs)/2], xs[len(xs)/2:]
	sum1, sum2 := 0, 0
	for _, v := range ys {
		n, _ := strconv.Atoi(v)
		sum1 += n
	}
	for _, v := range fs {
		n, _ := strconv.Atoi(v)
		sum2 += n
	}
	return sum1 == sum2
}

func isLuckyV2(n int) bool {
	var sum1, sum2, length int
	temp := n
	// Calculate the length of the number
	for temp != 0 {
		length++
		temp /= 10
	}
	half := length / 2
	// Calculate the sums of the first and second halves of the number
	for i := 0; i < length; i++ {
		digit := n % 10
		if i < half {
			sum1 += digit
		} else {
			sum2 += digit
		}
		n /= 10
	}
	return sum1 == sum2
}

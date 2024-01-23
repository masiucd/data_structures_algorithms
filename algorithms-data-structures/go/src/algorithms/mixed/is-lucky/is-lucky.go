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

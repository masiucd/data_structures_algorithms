package numbers

import (
	"strconv"
	"strings"
)

func Persistence(n int) int {
	persistenceCount := 0
	if isSingleDigit(n) {
		return persistenceCount
	}

	for !isSingleDigit(n) {
		nAsString := strconv.Itoa(n)
		product := 1
		for _, sn := range strings.Split(nAsString, "") {
			nn, _ := strconv.Atoi(sn)
			product *= nn
		}
		n = product
		persistenceCount++
	}
	return persistenceCount
}

func isSingleDigit(n int) bool {
	return n >= 0 && n <= 9
}

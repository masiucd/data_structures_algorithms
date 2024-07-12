package numbers

import (
	"strconv"
	"strings"
)

func Persistence(n int) int {
	persistenceCount := 0
	// loop until n is less than 10
	for n > 9 {
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

func PersistenceV2(n int) int {
	persistenceCount := 0
	// loop until n is less than 10
	for n > 9 {
		product := 1
		for n > 0 {
			digit := n % 10
			// Early termination if product becomes zero
			if digit == 0 {
				persistenceCount++
				return persistenceCount
			}
			product *= digit
			n /= 10
		}
		n = product
		persistenceCount++
	}
	return persistenceCount
}

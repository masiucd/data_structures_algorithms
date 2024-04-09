package mixed

import (
	"fmt"
	"strconv"
	"strings"
)

func findComplement(num int) int {
	flip := func(s rune) rune {
		if s == 49 {
			return 48
		} else {
			return 49
		}
	}
	numBinary := fmt.Sprintf("%b", num)
	r := strings.Map(flip, numBinary)
	i, _ := strconv.ParseInt(r, 2, 64)
	return int(i)
}

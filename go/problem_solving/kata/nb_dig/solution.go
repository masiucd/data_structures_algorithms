package nb_dig

import (
	"strconv"
	"strings"
)

func nbDig(n, d int) (result int) {
	for i := 0; i <= n; i++ {
		result += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}

	return result
}

package romanencoder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanEncoder(t *testing.T) {
	testCases := []struct {
		num      int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{40, "XL"},
		{90, "XC"},
		{400, "CD"},
		{900, "CM"},
		{1990, "MCMXC"},
		{2008, "MMVIII"},
		{1666, "MDCLXVI"},
		{3999, "MMMCMXCIX"},
	}
	for _, test := range testCases {
		result := solution(test.num)
		assert.Equal(t, test.expected, result, fmt.Sprintf("For %d expected %s but got %s", test.num, test.expected, result))
	}
}

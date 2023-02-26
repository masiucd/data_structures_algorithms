package romantoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {

	var tests = []struct {
		input string
		want  int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MMMCMXCIX", 3999},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, romanToInt(test.input))
	}

}

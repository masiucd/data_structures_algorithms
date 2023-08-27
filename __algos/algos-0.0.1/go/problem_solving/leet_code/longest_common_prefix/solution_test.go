package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {

	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, longestCommonPrefix(test.input))
	}

}

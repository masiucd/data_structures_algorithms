package ageindays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgeInDays(t *testing.T) {
	var tests = []struct {
		input int
		want  []int
	}{
		{400, []int{1, 1, 5}},
		{800, []int{2, 2, 10}},
		{30, []int{0, 1, 0}},
	}
	for _, test := range tests {
		got := solution(test.input)
		assert.Equal(t, test.want, got)
	}
}

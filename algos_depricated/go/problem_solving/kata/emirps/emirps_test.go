package emirps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmirps(t *testing.T) {

	var tests = []struct {
		input    int
		expected [3]int
	}{
		{10, [3]int{0, 0, 0}},
		{13, [3]int{0, 0, 0}},
		{50, [3]int{4, 37, 98}},
		{100, [3]int{8, 97, 418}},
	}

	for _, test := range tests {
		got := solution(test.input)
		assert.Equal(t, test.expected, got, fmt.Sprintf("expected =: %v Got = %v ", test.expected, got))
	}

}

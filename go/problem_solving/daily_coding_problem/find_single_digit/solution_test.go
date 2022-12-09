package findsingledigit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	expected []int
	actual   int
}{
	{[]int{6, 1, 3, 3, 3, 6, 6}, 1},
	{[]int{13, 19, 13, 13}, 19}}

func TestFindSingleDigit(t *testing.T) {
	for _, tt := range testData {
		expected := solution(tt.expected)
		assert.Equal(t, expected, tt.actual, fmt.Sprintf("expexted %d actual %d", expected, tt.actual))
	}
}

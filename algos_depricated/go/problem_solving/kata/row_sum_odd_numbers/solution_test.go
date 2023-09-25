package row_sum_odd_numbers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRowSumOddNumbers(t *testing.T) {
	actual := solution(2)
	expected := 8
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	actual = solution(13)
	expected = 2197
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	actual = solution(19)
	expected = 6859
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

}

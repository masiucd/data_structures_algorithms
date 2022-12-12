package longest_consec

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsec(t *testing.T) {
	actual := solution([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2)
	expected := "abigailtheta"
	assert.Equal(t, actual, expected, fmt.Sprintf("Expected %s  got %s ", expected, actual))

	actual = solution([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 3)
	expected = "zoneabigailtheta"
	assert.Equal(t, actual, expected, fmt.Sprintf("Expected %s  got %s ", expected, actual))
}

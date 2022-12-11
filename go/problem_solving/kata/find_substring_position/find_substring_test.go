package find_substring_position

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSubstringPosition(t *testing.T) {
	actual := solution("marcell", "cell")
	expected := 3
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))

	actual = solution("legia warszawa", "war")
	expected = 6
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %d  actual %d", expected, actual))
}

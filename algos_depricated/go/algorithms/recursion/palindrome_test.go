package recursion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	actual := isPalindrome("marcell")
	expected := false
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %v  actual %v", expected, actual))

	actual = isPalindrome("racecar")
	expected = true
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %v  actual %v", expected, actual))

	actual = isPalindrome("apa")
	expected = true
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %v  actual %v", expected, actual))
}

package palindromenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeNumber(t *testing.T) {
	var expected = isPalindrome(121)
	var actual = true
	assert.Equal(t, expected, actual, "Didn't match")

	expected = isPalindrome(-121)
	actual = false
	assert.Equal(t, expected, actual, "Didn't match")

	expected = isPalindrome(10)
	actual = false
	assert.Equal(t, expected, actual, "Didn't match")

	expected = isPalindrome(11)
	actual = true
	assert.Equal(t, expected, actual, "Didn't match")
}

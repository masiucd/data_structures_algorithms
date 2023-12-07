package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {
	got := isPalindrome("A man a plan a canal Panama")
	want := true
	assert.Equal(t, want, got, "should be true")

	got = isPalindrome("race a car")
	want = false
	assert.Equal(t, want, got, "should be false")

}

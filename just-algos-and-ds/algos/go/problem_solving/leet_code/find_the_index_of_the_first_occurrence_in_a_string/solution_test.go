package findtheindexofthefirstoccurrenceinastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	var expected = strStr("hello", "ll")
	var actual = 2
	assert.Equal(t, expected, actual, "Didn't match")

	expected = strStr("aaaaa", "bba")
	actual = -1
	assert.Equal(t, expected, actual, "Didn't match")

	expected = strStr("", "")
	actual = 0
	assert.Equal(t, expected, actual, "Didn't match")

	expected = strStr("a", "a")
	actual = 0
	assert.Equal(t, expected, actual, "Didn't match")

	expected = strStr("mississippi", "issip")
	actual = 4
	assert.Equal(t, expected, actual, "Matched")

	expected = strStr("hello", "lo")
	actual = 3
	assert.Equal(t, expected, actual, "Matched")
}

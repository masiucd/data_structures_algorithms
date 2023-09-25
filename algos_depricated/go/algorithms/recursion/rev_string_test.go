package recursion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRevStringItWorks(t *testing.T) {
	actual := reverseString("")
	expected := ""
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s  actual %s", expected, actual))

	actual = reverseString("abc")
	expected = "cba"
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s  actual %s", expected, actual))

}

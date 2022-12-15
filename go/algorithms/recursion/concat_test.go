package recursion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	actual := concat([]string{""})
	expected := ""
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s  actual %v", expected, actual))

	actual = concat([]string{"foo", "bar", "baz"})
	expected = "foobarbaz"
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected %s  actual %v", expected, actual))
}

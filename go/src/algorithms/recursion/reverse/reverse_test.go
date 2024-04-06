package reverse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	res := reverseString("marcell")
	assert.Equal(t, res, "llecram")

	res = ""
	assert.Equal(t, res, "")

	res = reverseString("m")
	assert.Equal(t, res, "m")

	res = reverseString("Legia")
	assert.Equal(t, res, "aigeL")
}

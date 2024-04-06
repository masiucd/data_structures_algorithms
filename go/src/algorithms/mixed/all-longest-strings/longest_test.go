package alllongeststrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllLongestStrings(t *testing.T) {
	xs := []string{"aba", "aa", "ad", "vcd", "aba"}
	res := allLongestStrings(xs) // ["aba", "vcd", "aba"]
	want := []string{"aba", "vcd", "aba"}

	assert.Equal(t, want, res, "The result should be 9")
}

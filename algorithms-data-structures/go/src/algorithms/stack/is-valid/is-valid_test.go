package isvalid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	input string
	want  bool
}

var tests = []test{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"([)]", false},
	{"{[]}", true},
}

func TestIsValid(t *testing.T) {

	for _, tc := range tests {
		got := isValid(tc.input)
		assert.Equal(t, tc.want, got)
	}
}

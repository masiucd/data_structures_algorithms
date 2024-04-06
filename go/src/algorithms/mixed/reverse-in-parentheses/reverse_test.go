package reverseinparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	inputString string
	want        string
}{
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},

	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
	{"foo(bar(baz))blim", "foobazrabblim"},
}

func TestReverseInParentheses(t *testing.T) {

	for _, tt := range tests {
		if got := reverseInParentheses(tt.inputString); got != tt.want {
			assert.Equal(t, tt.want, got)
		}
	}
}

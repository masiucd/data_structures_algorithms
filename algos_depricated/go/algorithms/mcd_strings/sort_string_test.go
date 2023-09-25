package mcd_strings

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortString(t *testing.T) {
	type test struct {
		input    string
		expected string
	}

	tests := []test{
		{input: "pza", expected: "apz"},
		{input: "pza", expected: "apz"},
		{input: "pza", expected: "apz"},
		{input: "marcell", expected: "acellmr"},
	}

	for _, tc := range tests {
		got := sortString(tc.input)
		assert.Equal(t, tc.expected, got, fmt.Sprintf("input: %s", tc.input))
	}
}

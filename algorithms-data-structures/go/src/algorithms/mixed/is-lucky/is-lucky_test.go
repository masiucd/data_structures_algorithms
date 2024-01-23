package islucky

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLucky(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{1230, true},
		{239017, false},
		{134008, true},
		{10, false},
		{11, true},
		{1010, true},
		{261534, false},
		{100000, false},
		{999999, true},
		{123321, true},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, isLucky(test.input))
	}
}
func TestIsLuckV2y(t *testing.T) {
	var tests = []struct {
		input int
		want  bool
	}{
		{1230, true},
		{239017, false},
		{134008, true},
		{10, false},
		{11, true},
		{1010, true},
		{261534, false},
		{100000, false},
		{999999, true},
		{123321, true},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, isLuckyV2(test.input))
	}
}

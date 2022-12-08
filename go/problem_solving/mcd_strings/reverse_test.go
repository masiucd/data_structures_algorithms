package mcd_strings

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	got := reverseString("hello")
	want := "olleh"
	assert.Equal(t, got, want, fmt.Sprintf("Expected %s  got %s", want, got))

	got = reverseStringRec("hello")
	want = "olleh"
	assert.Equal(t, got, want, fmt.Sprintf("Expected %s  got %s", want, got))

}

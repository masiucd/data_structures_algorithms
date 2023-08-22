package count_encoding

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountEncoding(t *testing.T) {
	got := solution("AAAABBBCCDAA")
	want := "4A3B2C1D2A"
	fmt.Println(got)
	assert.Equalf(t, got, want, fmt.Sprintf("Got %v want %v\n", got, want))
}

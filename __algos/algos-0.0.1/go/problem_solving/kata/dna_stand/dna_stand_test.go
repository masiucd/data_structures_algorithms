package dna_stand

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDnaStandard(t *testing.T) {

	got := dnaStrand("AAAA")
	want := "TTTT"
	assert.Equal(t, got, want, fmt.Sprintf("Got %s want %s", got, want))

	got = dnaStrand("TTTT")
	want = "AAAA"
	assert.Equal(t, got, want, fmt.Sprintf("Got %s want %s", got, want))

	got = dnaStrand("ATCG")
	want = "TAGC"
	assert.Equal(t, got, want, fmt.Sprintf("Got %s want %s", got, want))
}

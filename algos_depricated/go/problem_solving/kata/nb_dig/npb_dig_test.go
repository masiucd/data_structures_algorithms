package nb_dig

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNbDig(t *testing.T) {
	got := nbDig(10, 1)
	want := 4
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	got = nbDig(25, 1)
	want = 11
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))
}

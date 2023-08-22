package list_cardiopackage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnlyBorn1500(t *testing.T) {
	got := len(onlyBorn1500(inventors))
	want := 2
	assert.Equal(t, got, want, fmt.Sprintf("Got %d want %d ", got, want))
}

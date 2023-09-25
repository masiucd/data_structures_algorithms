package list_cardiopackage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstLastNames(t *testing.T) {

	got := len(firstLastNames(inventors))
	want := len(inventors)
	assert.Equal(t, got, want, fmt.Sprintf("Got %d want %d", got, want))

}

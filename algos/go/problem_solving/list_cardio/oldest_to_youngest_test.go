package list_cardiopackage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOldestToYoungest(t *testing.T) {
	got := oldestToYoungest(inventors)
	want := 1473
	assert.Equal(t, got[0].year, want, fmt.Sprintf("Got %d, want %d", got[0].year, want))

	want = 1564
	assert.Equal(t, got[1].year, want, fmt.Sprintf("Got %d, want %d", got[1].year, want))
}

package nb_year

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNbYears(t *testing.T) {

	got := nbYear(1500000, 0.25, -1000, 2000000)
	want := 151
	assert.Equal(t, got, want, fmt.Sprintf("Got %d want %d", got, want))

	got = nbYear(1500000, 0.25, 1000, 2000000)
	want = 94
	assert.Equal(t, got, want, fmt.Sprintf("Got %d want %d", got, want))

	got = nbYear(1500000, 2.5, 10000, 2000000)
	want = 10
	assert.Equal(t, got, want, fmt.Sprintf("Got %d want %d", got, want))
}

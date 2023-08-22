package abbrev_name

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbbrevName(t *testing.T) {
	actual := abbrevName("Bob Smith")
	expected := "B.S"
	assert.Equal(t, expected, actual, fmt.Sprintf("Got %s expected %s", actual, expected))

	actual = abbrevName("Tina Turner")
	expected = "T.T"
	assert.Equal(t, expected, actual, fmt.Sprintf("Got %s expected %s", actual, expected))

	actual = abbrevName("s.h")
	expected = "S.H"
	assert.Equal(t, expected, actual, fmt.Sprintf("Got %s expected %s", actual, expected))
}

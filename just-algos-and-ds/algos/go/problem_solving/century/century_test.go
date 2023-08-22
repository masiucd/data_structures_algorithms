package century

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCentury(t *testing.T) {
	assert.Equal(t, century(1705), 18)
	assert.Equal(t, century(1900), 19)
	assert.Equal(t, century(1601), 17)
	assert.Equal(t, century(2000), 20)
}

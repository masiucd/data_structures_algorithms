package mixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayChange(t *testing.T) {
	res := arrayChange([]int{1, 1, 1})
	assert.Equal(t, 3, res, "Expect 3")

	res = arrayChange([]int{-1000, 0, -2, 0})
	assert.Equal(t, 5, res, "Expect 5")
}

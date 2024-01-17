package makearrayconsecutive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeArrayConsecutive(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		statues := []int{6, 2, 3, 8}
		want := 3
		got := makeArrayConsecutive(statues)
		assert.Equal(t, got, want)
	})
	t.Run("should return 2", func(t *testing.T) {
		statues := []int{0, 3}
		want := 2
		got := makeArrayConsecutive(statues)
		assert.Equal(t, got, want)
	})
	t.Run("should return 0", func(t *testing.T) {
		statues := []int{5, 4, 6}
		want := 0
		got := makeArrayConsecutive(statues)
		assert.Equal(t, got, want)
	})
}

package numbers

import "testing"

func TestMaxProduct(t *testing.T) {
	t.Run("Level 1", func(t *testing.T) {
		got := maxProductLevel1([]int{3, 5, 8, 1, 6, 4, 6, 8, 3, 8, 1})
		want := 5
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Level 2", func(t *testing.T) {
		got := maxProductLevel2([]int{3, 5, 8, 1, 6, 4, 6, 8, 3, 8, 1}, 3)
		want := 5
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

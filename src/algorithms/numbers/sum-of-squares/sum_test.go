package numbers

import "testing"

func TestSumOfSquares(t *testing.T) {

	t.Run("Test sum of squares", func(t *testing.T) {
		got := sumOfAllSquares(3)
		want := 14

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Test sum of squares level 2", func(t *testing.T) {
		got := sumOfAllSquaresLevel2(3, 8)
		want := 199

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func BenchmarkSumOfSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumOfAllSquares(3)
	}
}

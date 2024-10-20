package numbers

import "testing"

func TestConsecutiveNumbers(t *testing.T) {
	t.Run("Level 1", func(t *testing.T) {
		got := level1ConsecutiveCheck([]int{1, 3})
		want := []int{}
		if !slicesAreEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		got = level1ConsecutiveCheck([]int{-1, 0, 1})
		want = []int{-1, 0, 1}
		if !slicesAreEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

	t.Run("Level 2", func(t *testing.T) {
		got := level2ConsecutiveCheck([]int{1, 3})
		want := []int{2}
		if !slicesAreEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		got = level2ConsecutiveCheck([]int{1, 2, 5, 3, 6, 8})
		want = []int{4, 7}
		if !slicesAreEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Level 2 V2", func(t *testing.T) {
		got := level2ConsecutiveCheckV2([]int{1, 3})
		want := []int{2}
		if !slicesAreEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

		got = level2ConsecutiveCheckV2([]int{1, 2, 5, 3, 6, 8})
		want = []int{4, 7}
		if !slicesAreEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

}

func slicesAreEqual(xs, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i, x := range xs {
		if x != ys[i] {
			return false
		}
	}
	return true
}

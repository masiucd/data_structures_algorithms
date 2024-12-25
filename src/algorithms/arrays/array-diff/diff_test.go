package arrays

import "testing"

func TestArrayDiff(t *testing.T) {

	t.Run("should return [3, 3] for listDiff([]int{1, 2, 3, 3}, []int{1, 2})", func(t *testing.T) {
		expected := []int{3, 3}
		got := listDiff([]int{1, 2, 3, 3}, []int{1, 2})

		if len(got) != len(expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}

		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf("got %v, expected %v", got, expected)
			}
		}

		got = listDiffV2([]int{1, 2, 3, 3}, []int{1, 2})
		if len(got) != len(expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("should return [1, 2, 3, 3] for listDiff([]int{1, 2, 3, 3}, []int{})", func(t *testing.T) {
		expected := []int{1, 2, 3, 3}
		got := listDiff([]int{1, 2, 3, 3}, []int{})

		if len(got) != len(expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}

		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf("got %v, expected %v", got, expected)
			}
		}

		got = listDiffV2([]int{1, 2, 3, 3}, []int{})
		if len(got) != len(expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("should return [] for listDiff([]int{}, []int{})", func(t *testing.T) {
		expected := []int{}
		got := listDiff([]int{}, []int{})

		if len(got) != len(expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}

		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf("got %v, expected %v", got, expected)
			}
		}

		got = listDiffV2([]int{}, []int{})
		if len(got) != len(expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("should return [] for listDiff([]int{1, 2, 3, 3}, []int{1, 2, 3, 3})", func(t *testing.T) {
		expected := []int{}
		got := listDiff([]int{1, 2, 3, 3}, []int{1, 2, 3, 3})

		if len(got) != len(expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}

		for i := 0; i < len(expected); i++ {
			if got[i] != expected[i] {
				t.Errorf("got %v, expected %v", got, expected)
			}
		}

		got = listDiffV2([]int{1, 2, 3, 3}, []int{1, 2, 3, 3})
		if len(got) != len(expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

}

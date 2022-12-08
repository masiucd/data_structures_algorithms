package searching

import "testing"

func TestBinarySearch(t *testing.T) {
	target := 4
	got := BinarySearch([]int{1, 2, 3, 4, 5}, target)
	want := target - 1
	if got != want {
		t.Errorf("BinarySearch: got %v, want %v", got, want)
	}

	target = 100
	got = BinarySearch([]int{1, 2, 3, 4, 5}, target)
	want = -1
	if got != want {
		t.Errorf("BinarySearch: got %v, want %v", got, want)
	}
}

func TestLinearSearch(t *testing.T) {
	target := 4
	got := LinearSearch([]int{1, 2, 3, 4, 5}, target)
	want := target - 1
	if got != want {
		t.Errorf("LinearSearch: got %v, want %v", got, want)
	}

	target = 100
	got = LinearSearch([]int{1, 2, 3, 4, 5}, target)
	want = -1
	if got != want {
		t.Errorf("LinearSearch: got %v, want %v", got, want)
	}
}

func TestSearchMatrix(t *testing.T) {
	list := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 4
	got := SearchMatrix(list, target)
	if got {
		t.Errorf("SearchMatrix: got %v, want %v", got, false)
	}

	target = 3
	got = SearchMatrix(list, target)
	if !got {
		t.Errorf("SearchMatrix: got %v, want %v", got, false)
	}

	target = 4
	got = SearchMatrixTwo(list, target)
	if got {
		t.Errorf("SearchMatrix: got %v, want %v", got, false)
	}

	target = 3
	got = SearchMatrixTwo(list, target)
	if !got {
		t.Errorf("SearchMatrix: got %v, want %v", got, false)
	}
}

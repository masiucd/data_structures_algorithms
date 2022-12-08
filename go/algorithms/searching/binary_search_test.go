package searching

import "testing"

func TestPushBack(t *testing.T) {
	got := BinarySearch([]int{1, 2, 3, 4, 5}, 2)
	want := 1
	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}

	got = BinarySearch([]int{1, 2, 3, 4, 5}, 100)
	want = -1
	if got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}

}

package mixed

import "testing"

func TestMajorityElement(t *testing.T) {
	got := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	want := 2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got = majorityElement([]int{3, 2, 3})
	want = 3

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

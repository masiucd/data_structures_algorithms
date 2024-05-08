package minmax

import "testing"

func TestMaxFunc(t *testing.T) {
	xs := []int{1, 2, 3, 10, 4, 5}
	want := 10
	result := Max(xs)
	if result != want {
		t.Errorf("Expected 5 but got %d", result)
	}
}

func TestMinFunc(t *testing.T) {
	xs := []int{1, 2, 3, 10, 4, 5}
	want := 1
	result := Min(xs)
	if result != want {
		t.Errorf("Expected 5 but got %d", result)
	}
}

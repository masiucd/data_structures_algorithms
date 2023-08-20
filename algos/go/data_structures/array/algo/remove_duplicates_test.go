package algo

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	xs := []int{1, 1, 1, 2, 2, 3}
	got := removeDuplicates(xs)
	want := 3
	if got != want {
		t.Errorf("Got = %v want %v ", got, want)
	}

	mutated := []int{1, 2, 3}
	if reflect.DeepEqual(xs, mutated) {
		t.Errorf("Got = %v want %v ", got, mutated)

	}
}

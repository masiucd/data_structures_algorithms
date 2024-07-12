package bits

import (
	"reflect"
	"testing"
)

func TestContBits(t *testing.T) {
	got := countBits(2)
	want := []int{0, 1, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	got = countBits(5)
	want = []int{0, 1, 1, 2, 1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestContBitsV2(t *testing.T) {
	got := countBitsV2(2)
	want := []int{0, 1, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	got = countBitsV2(5)
	want = []int{0, 1, 1, 2, 1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

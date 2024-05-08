package minmax

import "testing"

func TestMaxFunc(t *testing.T) {
	xs := []int{1, 2, 3, 10, 4, 5}
	want := 10
	result := Max(xs)
	if result != want {
		t.Errorf("Expected 5 but got %d", result)
	}

	listWithFloats := []float64{1.1, 2.2, 3.3, 10.1, 4.4, 5.5}
	wantFloat := 10.1
	resultFloat := Max(listWithFloats)
	if resultFloat != wantFloat {
		t.Errorf("Expected 5 but got %f", resultFloat)
	}
}

func TestMinFunc(t *testing.T) {
	xs := []int{1, 2, 3, 10, 4, 5}
	want := 1
	result := Min(xs)
	if result != want {
		t.Errorf("Expected 5 but got %d", result)
	}

	listWithFloats := []float64{1.1, 2.2, 3.3, 10.1, 4.4, 5.5}
	wantFloat := 1.1
	resultFloat := Min(listWithFloats)
	if resultFloat != wantFloat {
		t.Errorf("Expected 5 but got %f", resultFloat)
	}
}

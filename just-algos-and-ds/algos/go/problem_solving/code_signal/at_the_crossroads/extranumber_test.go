package at_the_crossroads

import "testing"

var testData = []struct {
	inputA int
	inputB int
	inputC int
	want   int
}{
	{2, 7, 2, 7},
	{3, 2, 2, 3},
	{5, 3, 3, 5},
}

func TestExtraNumber(t *testing.T) {
	for _, v := range testData {
		got := extraNumber(v.inputA, v.inputB, v.inputC)
		if got != v.want {
			t.Errorf("got %v, want %v", got, v.want)
		}
	}
}
func TestExtraNumberTwo(t *testing.T) {
	for _, v := range testData {
		got := extraNumberTwo(v.inputA, v.inputB, v.inputC)
		if got != v.want {
			t.Errorf("got %v, want %v", got, v.want)
		}
	}
}

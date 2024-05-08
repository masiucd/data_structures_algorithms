package minmax

import "testing"

func TestMaxFunc(t *testing.T) {
	testCases := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"Integers", []float64{1, 2, 3, 10, 4, 5}, 10},
		{"Floats", []float64{1.1, 2.2, 3.3, 10.1, 4.4, 5.5}, 10.1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Max(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}

func TestMinFunc(t *testing.T) {
	testCases := []struct {
		name     string
		input    []float64
		expected float64
	}{
		{"Integers", []float64{1, 2, 3, 10, 4, 5}, 1},
		{"Floats", []float64{1.1, 2.2, 3.3, 10.1, 4.4, 5.5}, 1.1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Min(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}

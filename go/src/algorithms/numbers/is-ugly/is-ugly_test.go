package numbers

import "testing"

var tests = []struct {
	n        int
	expected bool
}{
	{6, true},
	{14, false},
	{8, true},
	{0, false},
	{1, true},
	{2, true},
}

func TestIsUgly(t *testing.T) {

	for _, tt := range tests {
		result := isUgly(tt.n)
		if result != tt.expected {
			t.Errorf("expected %t, got %t", tt.expected, result)
		}
	}
}
func TestIsUglyV2(t *testing.T) {

	for _, tt := range tests {
		result := isUglyV2(tt.n)
		if result != tt.expected {
			t.Errorf("expected %t, got %t", tt.expected, result)
		}
	}
}

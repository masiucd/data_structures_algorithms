package checkperfectnumber

import "testing"

func TestCheckPerfectNumber(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{28, true},
		{27, false},
		{1, false},
		{2, false},
		{3, false},
		{66, false}}

	for _, tt := range tests {
		if got := checkPerfectNumber(tt.num); got != tt.want {
			t.Errorf("checkPerfectNumber(%v) = %v, want %v", tt.num, got, tt.want)
		}
	}
}

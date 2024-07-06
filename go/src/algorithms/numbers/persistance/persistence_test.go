package numbers

import "testing"

func TestPersistence(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"39", 39, 3},
		{"999", 999, 4},
		{"9", 9, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Persistence(tt.input); got != tt.want {
				t.Errorf("Persistence(%d) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

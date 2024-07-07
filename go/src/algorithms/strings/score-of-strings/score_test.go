package strings

import "testing"

func TestScoreOfStrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"hello", "hello", 13},
		{"a", "a", 0},
		{"ab", "ab", 1},
		{"zaz", "zaz", 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreOfString(tt.input); got != tt.want {
				t.Errorf("scoreOfString(%s) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}

}

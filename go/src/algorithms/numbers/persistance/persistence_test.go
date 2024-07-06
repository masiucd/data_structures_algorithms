package numbers

import "testing"

var tests = []struct {
	name        string
	input, want int
}{
	{"39", 39, 3},
	{"999", 999, 4},
	{"9", 9, 0},
}

func TestPersistence(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Persistence(tt.input); got != tt.want {
				t.Errorf("Persistence(%d) = %d, want %d", tt.input, got, tt.want)
			}

		})

	}
}
func TestPersistenceV2(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PersistenceV2(tt.input); got != tt.want {
				t.Errorf("PersistenceV2(%d) = %d, want %d", tt.input, got, tt.want)
			}

		})

	}
}

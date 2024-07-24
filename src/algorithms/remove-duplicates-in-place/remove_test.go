package algorithms

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		want     int
		wantedXs []int
	}{
		{
			name:     "Test with multiple duplicates",
			input:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want:     5,
			wantedXs: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "Test with single duplicate",
			input:    []int{1, 1, 2},
			want:     2,
			wantedXs: []int{1, 2},
		},
		{
			name:     "Test with empty slice",
			input:    []int{},
			want:     0,
			wantedXs: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := removeDuplicates(tc.input)
			if got != tc.want {
				t.Errorf("got %v want %v", got, tc.want)
			}

			for i := 0; i < tc.want; i++ {
				if tc.input[i] != tc.wantedXs[i] {
					t.Errorf("After removeDuplicates, got %v want %v", tc.input, tc.wantedXs)
					break
				}
			}
		})
	}
}

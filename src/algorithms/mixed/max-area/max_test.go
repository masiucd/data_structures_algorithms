package mixed

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "testcase 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "testcase 2",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "testcase 3",
			height: []int{4, 3, 2, 1, 4},
			want:   16,
		},
		{
			name:   "testcase 4",
			height: []int{1, 2, 1},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea(%v) = %v, want %v", tt.height, got, tt.want)
			}
		})
	}
}

package mixed

import "testing"

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		name     string
		g        []int
		s        []int
		expected int
	}{
		{
			name:     "should return 1",
			g:        []int{1, 2, 3},
			s:        []int{1, 1},
			expected: 1,
		},
		{
			name:     "should return 2",
			g:        []int{1, 2},
			s:        []int{1, 2, 3},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findContentChildren(tt.g, tt.s)
			if res != tt.expected {
				t.Errorf("got %v want %v", res, tt.expected)
			}
		})
	}
}

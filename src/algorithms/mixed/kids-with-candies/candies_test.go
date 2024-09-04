package mixed

import "testing"

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			expected:     []bool{true, false, false, false, false},
		},
		{
			candies:      []int{12, 1, 12},
			extraCandies: 10,
			expected:     []bool{true, false, true},
		},
	}

	for _, tt := range tests {
		result := kidsWithCandies(tt.candies, tt.extraCandies)
		for i, v := range result {
			if v != tt.expected[i] {
				t.Errorf("For candies %v and extraCandies %d, expected %v but got %v", tt.candies, tt.extraCandies, tt.expected, result)
			}
		}
	}
}

func TestMaxValue(t *testing.T) {
	tests := []struct {
		candies  []int
		expected int
	}{
		{
			candies:  []int{2, 3, 5, 1, 3},
			expected: 5,
		},
		{
			candies:  []int{4, 2, 1, 1, 2},
			expected: 4,
		},

		{
			candies:  []int{12, 1, 12},
			expected: 12,
		},
	}

	for i, v := range tests {
		got := maxValue(v.candies)
		if got != v.expected {
			t.Errorf("For test %d, expected %d but got %d", i, v.expected, got)
		}
	}
}

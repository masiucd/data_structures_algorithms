package mixed

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, test := range tests {
		got := uniqueOccurrences(test.arr)
		if got != test.want {
			t.Errorf("uniqueOccurrences(%v) = %v, want %v", test.arr, got, test.want)
		}
	}
}

func TestUniqueOccurrences2(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}

	for _, test := range tests {
		got := uniqueOccurrences2(test.arr)
		if got != test.want {
			t.Errorf("uniqueOccurrences2(%v) = %v, want %v", test.arr, got, test.want)
		}
	}
}

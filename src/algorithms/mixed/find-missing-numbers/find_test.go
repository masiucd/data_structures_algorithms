package mixed

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindMissingNumbers(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{4, 3, 2, 7, 8, 2, 3, 1}, expected: []int{5, 6}},
		{input: []int{2, 2}, expected: []int{1}},
		{input: []int{1, 1}, expected: []int{2}},
	}

	for _, tc := range testCases {
		res := findDisappearedNumbers(tc.input)
		if diff := cmp.Diff(tc.expected, res); diff != "" {
			t.Errorf("Mismatch (-want +got):\n%s", diff)
		}
	}
}

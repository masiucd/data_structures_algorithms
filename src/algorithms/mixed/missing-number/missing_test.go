package mixed

import "testing"

var tests = []struct {
	nums []int
	want int
}{
	{[]int{3, 0, 1}, 2},
	{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	{[]int{0}, 1},
}

func TestMissingNumber(t *testing.T) {

	for _, tt := range tests {
		got := missingNumber(tt.nums)
		if got != tt.want {
			t.Errorf("got %v want %v given", got, tt.want)
		}
	}
}
func TestMissingNumberV2(t *testing.T) {
	for _, tt := range tests {
		got := missingNumberV2(tt.nums)
		if got != tt.want {
			t.Errorf("got %v want %v given", got, tt.want)
		}
	}
}

func TestMissingNumberV3(t *testing.T) {
	for _, tt := range tests {
		got := missingNumberV3(tt.nums)
		if got != tt.want {
			t.Errorf("got %v want %v given", got, tt.want)
		}
	}
}

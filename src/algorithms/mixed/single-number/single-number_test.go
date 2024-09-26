package mixed

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, test := range tests {
		if got := singleNumber(test.nums); got != test.want {
			t.Errorf("singleNumber(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}

func TestSingleNumberV2(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, test := range tests {
		if got := singleNumberV2(test.nums); got != test.want {
			t.Errorf("singleNumberV2(%v) = %d, want %d", test.nums, got, test.want)
		}
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumber([]int{2, 2, 1})
	}
}

func BenchmarkSingleNumberV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleNumberV2([]int{2, 2, 1})
	}
}

package sorting

import (
	"reflect"
	"testing"
)

var sortTestData = []struct {
	got  []int
	want []int
}{
	{
		[]int{2, 3, 1},
		[]int{1, 2, 3}},
	{
		[]int{33, 2, 1},
		[]int{1, 2, 33},
	},
	{
		[]int{33, 2, 1, 0, 3},
		[]int{0, 1, 2, 3, 33},
	},
}

func TestInsertionSort(t *testing.T) {
	for _, test := range sortTestData {
		got := InsertionSort(test.got)
		want := InsertionSort(test.want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	}

}

func TestMergeSort(t *testing.T) {

	for _, test := range sortTestData {
		got := MergeSort(test.got)
		want := MergeSort(test.want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	}
}

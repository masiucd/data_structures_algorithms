package two_sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []struct {
	input []string
	want  string
}{
	{input: []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}, want: "b***i***t***c***o***i***n"},
	{input: []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}, want: "a***r***e"},
	{input: []string{"lets", "talk", "about", "javascript", "the", "best", "language"}, want: "a***b***o***u***t"},
}

func TestTwoSortedOne(t *testing.T) {
	for _, v := range testData {
		got := twoSort(v.input)
		assert.Equal(t, got, v.want, fmt.Sprintf("Got %s want %s ", got, v.want))
	}
}

func TestTwoSortedTwo(t *testing.T) {
	for _, v := range testData {
		got := twoSort2(v.input)
		assert.Equal(t, got, v.want, fmt.Sprintf("Got %s want %s ", got, v.want))
	}
}

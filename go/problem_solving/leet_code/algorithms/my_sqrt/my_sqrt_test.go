package my_sqrt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testData = []struct {
	input    int
	expected int
}{
	{1, 1},
	{2, 1},
	{3, 1},
	{4, 2},
	{5, 2},
	{100, 10},
}

func TestMySqrt(t *testing.T) {

	for _, v := range testData {
		t.Run(fmt.Sprintf("%v", v.input), func(t *testing.T) {
			actual := mySqrt(v.input)
			assert.Equal(t, v.expected, actual)
		})
	}

}

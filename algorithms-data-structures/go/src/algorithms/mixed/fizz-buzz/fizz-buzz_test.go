package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	input    int
	expected []string
}{
	{15, []string{"FizzBuzz", "1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz"}},
	{5, []string{"FizzBuzz", "1", "2", "Fizz", "4", "Buzz"}},
	{1, []string{"FizzBuzz"}},
}

func TestFizzBuzz(t *testing.T) {

	for _, data := range testData {
		result := makeFizzBuzList(data.input)
		for i, v := range result {
			assert.Equal(t, v, data.expected[i])
		}
	}
}

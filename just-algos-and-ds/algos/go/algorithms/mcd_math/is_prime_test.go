package mcd_math

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPrime(t *testing.T) {
	var testData = []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{11, true},
	}

	for _, data := range testData {
		actual := isPrime(data.input)
		assert.Equal(t, data.expected, actual, fmt.Sprintf("Expected %t, got %t", data.expected, actual))
	}
}
func TestIsPrimeTwo(t *testing.T) {
	var testData = []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{11, true},
	}

	for _, data := range testData {
		actual := isPrimeTwo(data.input)
		assert.Equal(t, data.expected, actual, fmt.Sprintf("Expected %t, got %t", data.expected, actual))
	}
}

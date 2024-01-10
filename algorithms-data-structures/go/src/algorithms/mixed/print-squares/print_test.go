package printsquares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintSquaresTopDown(t *testing.T) {
	n := 5
	expected := "#####\n####\n###\n##\n#\n"
	actual := printSquaresTopDown(n)
	assert.Equal(t, expected, actual)

	n = 3
	expected = "###\n##\n#\n"
	actual = printSquaresTopDown(n)
	assert.Equal(t, expected, actual)

	n = 1
	expected = "#\n"
	actual = printSquaresTopDown(n)
	assert.Equal(t, expected, actual)
}

func TestPrintSquaresDownTop(t *testing.T) {
	n := 5
	expected := "\n#\n##\n###\n####\n"
	actual := printSquaresDownTop(n)
	assert.Equal(t, expected, actual)

	n = 3
	expected = "\n#\n##\n"
	actual = printSquaresDownTop(n)
	assert.Equal(t, expected, actual)

	n = 1
	expected = "\n"
	actual = printSquaresDownTop(n)
	assert.Equal(t, expected, actual)
}

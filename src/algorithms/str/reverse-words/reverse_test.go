package reversewords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	t.Run("should reverse words in a string", func(t *testing.T) {
		input := "a good   example"
		expected := "example good a"
		assert.Equal(t, expected, reverseWords(input))
	})

	t.Run("should reverse words in a string", func(t *testing.T) {
		input := "  hello world  "
		expected := "world hello"
		assert.Equal(t, expected, reverseWords(input))
	})

	t.Run("should reverse words in a string", func(t *testing.T) {
		input := "  Bob    Loves  Alice   "
		expected := "Alice Loves Bob"
		assert.Equal(t, expected, reverseWords(input))
	})

	t.Run("should reverse words in a string", func(t *testing.T) {
		input := "Alice does not even like bob"
		expected := "bob like even not does Alice"
		assert.Equal(t, expected, reverseWords(input))
	})
}

func TestReverseWordsV2(t *testing.T) {
	t.Run("should reverse words in a string", func(t *testing.T) {
		input := "a good   example"
		expected := "example good a"
		assert.Equal(t, expected, reverseWordsV2(input))
	})

	t.Run("should reverse words in a string", func(t *testing.T) {
		input := "  hello world  "
		expected := "world hello"
		assert.Equal(t, expected, reverseWordsV2(input))
	})

	t.Run("should reverse words in a string", func(t *testing.T) {
		input := "  Bob    Loves  Alice   "
		expected := "Alice Loves Bob"
		assert.Equal(t, expected, reverseWordsV2(input))
	})

	t.Run("should reverse words in a string", func(t *testing.T) {
		input := "Alice does not even like bob"
		expected := "bob like even not does Alice"
		assert.Equal(t, expected, reverseWordsV2(input))
	})
}

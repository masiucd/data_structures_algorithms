package countcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountCharacters(t *testing.T) {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	res := countCharacters(words, chars)
	assert.Equal(t, 6, res)

	words = []string{"hello", "world", "leetcode"}
	chars = "welldonehoneyr"
	res = countCharacters(words, chars)
	assert.Equal(t, 10, res)
}

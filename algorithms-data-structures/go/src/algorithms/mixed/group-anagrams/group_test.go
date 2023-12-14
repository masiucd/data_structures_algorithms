package groupanagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(input)

	assert.Equal(t, 3, len(res))
	assert.Equal(t, []string{"eat", "tea", "ate"}, res[0])

	assert.Equal(t, []string{"tan", "nat"}, res[1])
	assert.Equal(t, []string{"bat"}, res[2])

	input = []string{"a"}
	res = groupAnagrams(input)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, []string{"a"}, res[0])

}

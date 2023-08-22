package hash_set

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeHashSet() *HashSet {
	hs := New()
	hs.Add(10)
	hs.Add(20)
	hs.Add(10)
	hs.Add(30)
	hs.Add(10)
	return &hs
}

func TestHashSet(t *testing.T) {
	hs := makeHashSet()
	assert.Equal(t, hs.Size, 3, fmt.Sprintf("HashSet size: %d", hs.Size))

	hs.Remove(10)
	assert.Equal(t, hs.Size, 2, fmt.Sprintf("HashSet size: %d", hs.Size))

	assert.Equal(t, hs.Contains(20), true, fmt.Sprintf("HashSet contains %v", true))

}

package contain_duplicate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicates(t *testing.T) {

	expected := containsDuplicate([]int{1, 2, 3})
	assert.Equal(t, expected, false, fmt.Sprintf("Expected %v  got %v", expected, false))

	expected = containsDuplicate([]int{1, 2, 3, 1})
	assert.Equal(t, expected, true, fmt.Sprintf("Expected %v  got %v", expected, true))

}

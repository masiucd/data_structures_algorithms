package plusone

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestPlusOne(t *testing.T){
	got := plusOne([]int{1,2,3})
	want := []int{1,2,4}
	assert.Equal(t, got, want, fmt.Sprintf("Expected %v  got %v", want, got))

	got = plusOne([]int{4,3,2,1})
	want = []int{4,3,2,2}
	assert.Equal(t, got, want, fmt.Sprintf("Expected %v  got %v", want, got))

	got = plusOne([]int{9,9,9})
	want = []int{1,0,0,0}
	assert.Equal(t, got, want, fmt.Sprintf("Expected	 %v  got %v", want, got))
}
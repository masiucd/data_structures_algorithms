package sum_reach_target

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumReachForTarget(t *testing.T) {
	got := sumReachTarget([]int{10, 15, 3, 7}, 17)
	assert.Equal(t, got, true, fmt.Sprintf("Expected %t  got %t", true, got))
}

func TestSumReachForTargetFaster(t *testing.T) {
	got := sumReachTarget2([]int{10, 15, 3, 7}, 17)
	fmt.Println(got)
	assert.Equal(t, got, true, fmt.Sprintf("Expected %t  got %t", true, got))
}

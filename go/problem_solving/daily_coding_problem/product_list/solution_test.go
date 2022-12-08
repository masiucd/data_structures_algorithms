package product_list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductListThreeNumbers(t *testing.T) {
	got := productList([]int{3, 2, 1})
	want := []int{2, 3, 6}
	for i, v := range got {
		assert.Equal(t, v, want[i], fmt.Sprintf("Expected %d  got %d", want[i], got))
	}

	got = productListTwo([]int{3, 2, 1})
	for i, v := range got {
		assert.Equal(t, v, want[i], fmt.Sprintf("Expected %d  got %d", want[i], got))
	}

	got = productListThree([]int{3, 2, 1})
	for i, v := range got {
		assert.Equal(t, v, want[i], fmt.Sprintf("Expected %d  got %d", want[i], got))
	}
}

func TestProductListFiveNumbers(t *testing.T) {
	got := productList([]int{1, 2, 3, 4, 5})
	want := []int{120, 60, 40, 30, 24}
	for i, v := range got {
		assert.Equal(t, v, want[i], fmt.Sprintf("Expected %d  got %d", want[i], got))
	}
	got = productListTwo([]int{1, 2, 3, 4, 5})
	for i, v := range got {
		assert.Equal(t, v, want[i], fmt.Sprintf("Expected %d  got %d", want[i], got))
	}

	got = productListThree([]int{1, 2, 3, 4, 5})
	for i, v := range got {
		assert.Equal(t, v, want[i], fmt.Sprintf("Expected %d  got %d", want[i], got))
	}

}

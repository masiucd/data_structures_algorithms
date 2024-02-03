package grouping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cartItems = []CartItem{
	{1, "Item 1", 12.5},
	{2, "Item 2", 15.0},
	{2, "Item 2", 15.0},
	{2, "Item 2", 15.0},
	{3, "Item 3", 20.0},
	{3, "Item 3", 20.0},
	{3, "Item 3", 20.0},
	{4, "Item 4", 25.0},
	{4, "Item 4", 25.0},
}

func TestGroupCartItems(t *testing.T) {
	res := groupCartItems(cartItems)

	assert.Equal(t, 4, len(res))

	assert.Equal(t, 1, res["Item 1"].Qty)
	assert.Equal(t, 3, res["Item 2"].Qty)
	assert.Equal(t, 3, res["Item 3"].Qty)
	assert.Equal(t, 2, res["Item 4"].Qty)

	xs := printTotalPriceForItem(res)
	assert.Equal(t, 4, len(xs))
	assert.Equal(t, "Item 1 ---> 12.50", xs[0])
	assert.Equal(t, "Item 2 ---> 45.00", xs[1])
	assert.Equal(t, "Item 3 ---> 60.00", xs[2])
	assert.Equal(t, "Item 4 ---> 50.00", xs[3])
}

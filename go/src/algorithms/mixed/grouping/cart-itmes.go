package grouping

import "fmt"

type CartItem struct {
	ID    int
	Name  string
	Price float64
}

type GroupedItem struct {
	CartItem
	Qty int
}

func printTotalPriceForItem(items map[string]GroupedItem) []string {
	var result []string
	for k, v := range items {
		fmt.Println(k+" ---> ", (v.Price * float64(v.Qty)))
		result = append(result, fmt.Sprintf("%s ---> %.2f", k, (v.Price*float64(v.Qty))))
	}
	return result
}

func groupCartItems(cartItems []CartItem) map[string]GroupedItem {
	result := map[string]GroupedItem{}
	for _, item := range cartItems {
		if val, ok := result[item.Name]; ok {
			result[item.Name] = GroupedItem{
				CartItem: CartItem{
					ID:    item.ID,
					Name:  item.Name,
					Price: item.Price,
				},
				Qty: val.Qty + 1,
			}
		} else {
			result[item.Name] = GroupedItem{
				CartItem: CartItem{
					ID:    item.ID,
					Name:  item.Name,
					Price: item.Price,
				},
				Qty: 1,
			}
		}
	}
	return result
}

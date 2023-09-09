package unique

type Item struct {
	id   int
	name string
}

//var xs []Item = []Item{
//	{id: 1, name: "abc"},
//	{id: 2, name: "def"},
//	{id: 3, name: "ghi"},
//	{id: 1, name: "abc"},
//	{id: 2, name: "def"},
//}

func getUniqueList(list []Item) []Item {
	var result []Item
	mapper := make(map[int]Item)
	for _, v := range list {
		if _, ok := mapper[v.id]; !ok {
			mapper[v.id] = v
		}
	}
	for _, v := range mapper {
		result = append(result, v)
	}
	return result
}

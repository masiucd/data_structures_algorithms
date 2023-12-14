package topkfrequent

func topKFrequent(nums []int, k int) []int {
	var result []int
	var store = make(map[int]int)
	for _, v := range nums {
		store[v]++
	}
	for i := 0; i < k; i++ {
		var max int
		var maxKey int
		for k, v := range store {
			if v > max {
				max = v
				maxKey = k
			}
		}
		result = append(result, maxKey)
		delete(store, maxKey)
	}
	return result
}

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

func topKFrequentTwo(numbers []int, k int) []int {
	var store = make(map[int]int)
	for _, v := range numbers {
		store[v]++
	}
	// bucket that stores the numbers with the same frequency
	bucket := make([][]int, len(numbers)+1)
	for k, v := range store {
		bucket[v] = append(bucket[v], k)
	}

	var res []int
	// iterate the bucket from the end
	for i := len(bucket) - 1; i >= 0; i-- {
		// if the bucket is not empty and k is not 0
		if len(bucket[i]) > 0 && k > 0 {
			// append the numbers to the result
			res = append(res, bucket[i]...)
			// decrement k by the length of the bucket
			k -= len(bucket[i])
		}
	}
	return res
}

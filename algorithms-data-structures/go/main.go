package main

import (
	"fmt"
)

func main() {

	xs := []int{1, 1, 1, 1, 2, 2, 2, 3}
	k := 2
	res := topKFrequent(xs, k)
	fmt.Println(res)
}

func topKFrequent(numbers []int, k int) []int {
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

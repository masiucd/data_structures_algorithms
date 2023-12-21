package main

<<<<<<< HEAD
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
=======
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	targetSum := 4
	res := hasPathSum(root, targetSum)
	println(res)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	if hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
>>>>>>> ea5f6fd5bc901ff554c38d77e30a75d007512a6f
}

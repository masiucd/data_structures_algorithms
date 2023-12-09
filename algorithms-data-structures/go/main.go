package main

import (
	cc "container/list"
	"fmt"
	"go-ds/src/data-structures/trees/bst"
)

func main() {
	bst := bst.NewBst()
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(20)
	bst.Insert(2)
	bst.Insert(1)
	fmt.Println(bst.BFS())
	fmt.Println(bst.Size())
	fmt.Println(bst.DFS("PRE"))
	fmt.Println(bst.DFS("POST"))
	fmt.Println(bst.DFS("IN"))
	fmt.Println(bst.Contains(222))

	x := cc.New()
	fmt.Println(x)

}

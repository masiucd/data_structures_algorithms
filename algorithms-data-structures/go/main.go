package main

import (
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
	bst.Remove(15)
	fmt.Println(bst.BFS())

}

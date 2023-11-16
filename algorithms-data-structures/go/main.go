package main

import (
	"fmt"
	"go-ds/src/data-structures/queue"
)

func main() {

	q := queue.New()
	q.Enqueue(100)
	q.Enqueue(20)
	q.Enqueue(30)

	res := q.Print()
	fmt.Println(res)

}

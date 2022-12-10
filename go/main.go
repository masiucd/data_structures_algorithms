package main

import (
	"data_structures_algos_go/data_structures/stack"
	"fmt"
	"log"
)

func main() {
	s := stack.NewStack()
	s.Push(5)
	s.Push(2)
	s.Push(3)
	s.Push(312)
	s.Push(1)
	x, err := s.Peek()
	if err != nil {
		log.Fatal(err)
	}
	pop, err := s.Pop()
	if err != nil {
		return
	}
	fmt.Println(x)
	x, _ = s.Peek()
	fmt.Println(x)
}

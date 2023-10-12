package main

import linkedlist "go-ds/src/algorithms/data-structures/linked-list"

func main() {

	ll := linkedlist.New()
	ll.Append(10)
	// ll.Prepend(20)
	ll.Append(30)

	print(ll)
}

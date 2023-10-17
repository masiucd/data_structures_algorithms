package double

import (
	"go-ds/src/data-structures/linked-list"
)

type DoubleList struct {
	Head *linkedlist.DoubleNode
	Tail *linkedlist.DoubleNode
	Size int
}

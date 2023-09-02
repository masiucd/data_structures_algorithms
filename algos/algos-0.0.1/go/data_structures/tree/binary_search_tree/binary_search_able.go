package binary_search_tree

type BinarySearchAble interface {
	Insert(key int)
	InsertRec(key int)
	Search(key int) *Node
	Remove(root *Node, key int) *Node
	MinValue(root *Node) *Node
	MaxValue(root *Node) *Node
	DepthFirstSearch(root *Node) *Node
	BreadthFirstSearch() *Node
	Inorder() []int
	Preorder() []int
	Postorder() []int
}

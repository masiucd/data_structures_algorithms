package hash_set

const SIZE = 1024

func hash(key, numBuckets int) int {
	return key % numBuckets
}

type Node struct {
	Val  int
	next *Node
}

type HashSet struct {
	Blocks [SIZE]*Node
	Size   int
}

func New() HashSet {
	return HashSet{Blocks: [SIZE]*Node{}, Size: 0}
}

func (hs *HashSet) Add(key int) {
	i := hash(key, SIZE)
	if hs.Blocks[i] == nil {
		hs.Blocks[i] = &Node{Val: key}
		hs.Size++
		return
	}

	current := hs.Blocks[i]
	// already exists
	if current.Val == key {
		return
	}
	for current.next != nil && current.Val != key {
		current = current.next
	}
	hs.Blocks[i] = &Node{Val: key}
	hs.Size++
}

func (hs *HashSet) Remove(key int) {
	i := hash(key, SIZE)
	if hs.Blocks[i] == nil {
		return
	}
	hs.Blocks[i] = nil
	hs.Size--

}

func (hs *HashSet) Contains(key int) bool {
	i := hash(key, SIZE)
	return hs.Blocks[i] != nil
}

/**
 * Your HashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

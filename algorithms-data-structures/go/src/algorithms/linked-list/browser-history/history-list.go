package browserhistory

type BrowserNode struct {
	Val  string
	Next *BrowserNode
	Prev *BrowserNode
}

type BrowserHistoryList struct {
	Current *BrowserNode
	Head    *BrowserNode
	Tail    *BrowserNode
}

func New(homepage string) *BrowserHistoryList {
	return &BrowserHistoryList{
		Current: &BrowserNode{Val: homepage},
		Head:    &BrowserNode{Val: homepage},
		Tail:    &BrowserNode{Val: homepage},
	}
}

func (bh *BrowserHistoryList) Visit(url string) {
	// Clear forward history
	bh.Current.Next = &BrowserNode{Val: url, Prev: bh.Current}
	// Move current to new node
	bh.Current = bh.Current.Next
}

func (bh *BrowserHistoryList) Back(steps int) string {
	for i := 0; i < steps; i++ {
		if bh.Current.Prev == nil {
			break
		}
		bh.Current = bh.Current.Prev
	}
	return bh.Current.Val
}

func (bh *BrowserHistoryList) Forward(steps int) string {
	for i := 0; i < steps; i++ {
		if bh.Current.Next == nil {
			break
		}
		bh.Current = bh.Current.Next
	}
	return bh.Current.Val
}

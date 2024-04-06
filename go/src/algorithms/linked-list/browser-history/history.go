package browserhistory

type BrowserHistory struct {
	History []string
	Current int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		History: []string{homepage},
		Current: 0,
	}
}

func (bh *BrowserHistory) Visit(url string) {
	bh.History = bh.History[0 : bh.Current+1] // clear forward history
	bh.History = append(bh.History, url)
	bh.Current++
}

func (bh *BrowserHistory) Back(steps int) string {
	bh.Current = max(bh.Current-steps, 0) // if steps > current, set current to 0
	return bh.History[bh.Current]
}

func (bh *BrowserHistory) Forward(steps int) string {
	bh.Current = min(bh.Current+steps, len(bh.History)-1) // if steps > len(history), set current to len(history)-1
	return bh.History[bh.Current]
}

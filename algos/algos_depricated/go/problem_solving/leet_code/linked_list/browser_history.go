package linked_list

// https://leetcode.com/problems/design-browser-history/
type BrowserHistory struct {
	history []string
	pointer int
}

func newBrowserHistory(homepage string) BrowserHistory {
	return BrowserHistory{
		history: []string{homepage},
		pointer: 0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	if len(this.history) != this.pointer+1 {
		this.history = this.history[:this.pointer+1]
	}
	this.history = append(this.history, url)
	this.pointer += 1
}

func (this *BrowserHistory) Back(steps int) string {
	this.pointer -= steps
	if this.pointer < 0 {
		this.pointer = 0
	}
	return this.history[this.pointer]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.pointer += steps
	if this.pointer+1 > len(this.history) {
		this.pointer = len(this.history) - 1
	}
	return this.history[this.pointer]
}

package browserhistory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBrowserHistory(t *testing.T) {

	bh := Constructor("leetcode.com")
	bh.Visit("google.com")
	bh.Visit("facebook.com")
	bh.Visit("youtube.com")

	assert.Equal(t, "facebook.com", bh.Back(1))
	assert.Equal(t, "google.com", bh.Back(1))
	assert.Equal(t, "leetcode.com", bh.Back(1))

	bh.Visit("ifkgbg.se")
	bh.Visit("legia.com")
	bh.Visit("blabla.com")

	assert.Equal(t, "ifkgbg.se", bh.Back(2))
	assert.Equal(t, "leetcode.com", bh.Back(2))

}

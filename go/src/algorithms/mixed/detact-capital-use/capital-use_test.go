package mixed

import "testing"

func TestCapitalUse(t *testing.T) {

	tests := []struct {
		word string
		want bool
	}{
		{"USA", true},
		{"FlaG", false},
		{"Google", true},
	}

	for _, test := range tests {
		if got := detectCapitalUse(test.word); got != test.want {
			t.Errorf("detectCapitalUse(%q) = %v; want %v", test.word, got, test.want)
		}
		if got := detectCapitalUseV2(test.word); got != test.want {
			t.Errorf("detectCapitalUse(%q) = %v; want %v", test.word, got, test.want)
		}
	}
}

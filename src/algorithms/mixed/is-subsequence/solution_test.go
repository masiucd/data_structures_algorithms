package mixed

import "testing"

func TestIsSubsequence(t *testing.T) {

	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "avbc", true},
		{"abc", "ahbgdc", true},
		{"abc", "ahbgdc", true},
		{"abc", "ahbgdc", true},
	}

	for _, tt := range tests {
		got := isSubsequence(tt.s, tt.t)

		if got != tt.expected {
			t.Errorf("got: %v, want: %v", got, tt.expected)
		}
	}
}

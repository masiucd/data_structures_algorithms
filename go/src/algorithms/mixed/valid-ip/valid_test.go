package mixed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidIp(t *testing.T) {
	testCases := []struct {
		name      string
		ipAddress string
		want      bool
	}{
		{
			name:      "Valid IP",
			ipAddress: "172.16.254.1",
			want:      true,
		},
		{
			name:      "Invalid IP - Out of range",
			ipAddress: "172.316.254.1",
			want:      false,
		},
		{
			name:      "Invalid IP - Contains non-numeric character",
			ipAddress: "1.1.1.1a",
			want:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isValidIp(tc.ipAddress)
			assert.Equal(t, tc.want, got)
		})
	}
}

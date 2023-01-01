package numberoftrailingzerosof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeros(t *testing.T) {

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{0}, 0},
		{"Test 2", args{6}, 1},
		{"Test 3", args{30}, 7},
		{"Test 4", args{100}, 24},
		{"Test 5", args{1000}, 249},
		{"Test 6 (big)", args{1000000000}, 249999998}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, solution(tt.args.n))
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, solutionRec(tt.args.n))
		})
	}

}

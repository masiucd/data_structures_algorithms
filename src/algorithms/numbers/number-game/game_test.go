package numbers

import "testing"

var tests = []struct {
	name        string
	input, want []int
}{
	{"4,4,3,8", []int{4, 4, 3, 8}, []int{4, 3, 8, 4}},
	{"5,4,2,3", []int{5, 4, 2, 3}, []int{3, 2, 5, 4}},
}

func TestNumberGame(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberGame(tt.input); !equal(got, tt.want) {
				t.Errorf("NumberGame(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}

}

func TestNumberGameV2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberGameV2(tt.input); !equal(got, tt.want) {
				t.Errorf("NumberGameV2(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestNumberGameV3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberGameV3(tt.input); !equal(got, tt.want) {
				t.Errorf("NumberGameV3(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

func BenchmarkNumberGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberGame(tests[0].input)
	}
}

func BenchmarkNumberGameV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberGameV2(tests[0].input)
	}
}

func BenchmarkNumberGameV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numberGameV3(tests[0].input)
	}
}

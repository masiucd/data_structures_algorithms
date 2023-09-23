package calPoints

import "testing"

func TestCalP(t *testing.T) {

	type test struct {
		input    []string
		expected int
	}

	tests := []test{
		{[]string{"5", "2", "C", "D", "+"}, 30},
		{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27},
	}

	for _, v := range tests {
		output := calPoints(v.input)
		if output != v.expected {
			t.Errorf("Expected %v, got %v", v.expected, output)
		}
	}

}

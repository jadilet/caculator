package eval

import "testing"

func TestPrecedence(t *testing.T) {
	tests := []struct {
		operator string
		expected int
	}{
		{"+", 1},
		{"-", 1},
		{"*", 2},
		{"/", 2},
		{"^", 0},
	}
	
	for _, test := range tests {
		result := precedence(test.operator)
		if result != test.expected {
			t.Errorf("precedence(%q) = %d, expected = %d", test.operator, result, test.expected)
		}
	}
}

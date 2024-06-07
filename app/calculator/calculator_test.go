package calculator

import (
	"Go_test/calculator/number"
	"testing"
)

func TestCalculator(t *testing.T) {
	cases := []struct {
		left     number.Number
		right    number.Number
		operator string
		expected int
	}{
		{number.ArabicNumber(1), number.ArabicNumber(2), "+", 3},
		{number.ArabicNumber(6), number.ArabicNumber(3), "/", 2},
		{number.ArabicNumber(1), number.ArabicNumber(2), "-", -1},
		{number.RomanNumber("VI"), number.RomanNumber("III"), "/", 2},
	}

	for _, c := range cases {
		calc := Calculator{c.left, c.right, c.operator}
		result, err := calc.Calculate()
		if err != nil {
			if c.expected >= 0 {
				t.Errorf("Expected no error, but got %v", err)
			}
			continue
		}

		resultInt, _ := result.ToInt()
		if resultInt != c.expected {
			t.Errorf("Expected %v, but got %v", c.expected, resultInt)
		}
	}
}

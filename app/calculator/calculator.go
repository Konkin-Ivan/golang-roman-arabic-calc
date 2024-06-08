package calculator

import (
	"Go_test/calculator/number"
	"errors"
)

type Calculator struct {
	Left     number.Number
	Right    number.Number
	Operator string
}

func (c Calculator) Calculate() (number.Number, error) {
	left, err := c.Left.ToInt()
	if err != nil {
		return nil, err
	}
	right, err := c.Right.ToInt()
	if err != nil {
		return nil, err
	}

	switch c.Operator {
	case "+":
		return number.ArabicNumber(left + right), nil
	case "-":
		return number.ArabicNumber(left - right), nil
	case "*":
		return number.ArabicNumber(left * right), nil
	case "/":
		if right == 0 {
			return nil, errors.New("division by zero")
		}
		return number.ArabicNumber(left / right), nil
	default:
		return nil, errors.New("invalid operator")
	}
}

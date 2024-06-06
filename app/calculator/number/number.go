package number

import (
	"strconv"
)

type Number interface {
	ToInt() (int, error)
	ToString() string
}

type ArabicNumber int

func (n ArabicNumber) ToInt() (int, error) {
	return int(n), nil
}

func (n ArabicNumber) ToString() string {
	return strconv.Itoa(int(n))
}

type RomanNumber string

func (n RomanNumber) ToInt() (int, error) {
	romanValues := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := 0
	prev := 0
	for i := len(n) - 1; i >= 0; i-- {
		current := romanValues[rune(n[i])]
		if current < prev {
			result -= current
		} else {
			result += current
		}
		prev = current
	}
	return result, nil
}

func (n RomanNumber) ToString() string {
	return string(n)
}

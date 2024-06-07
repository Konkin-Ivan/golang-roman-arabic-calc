package number

import (
	"strconv"
	"strings"
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

// IntToRoman converts an integer to a Roman numeral string
func IntToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	for i, value := range values {
		for num >= value {
			num -= value
			result.WriteString(symbols[i])
		}
	}
	return result.String()
}

// IsRomanNumberValid Return true if string is a valid Roman numeral and does not exceed 10
func IsRomanNumberValid(num string, maxValue int) bool {
	roman := RomanNumber(num)
	val, err := roman.ToInt()
	if err != nil {
		return false
	}
	return val <= maxValue
}

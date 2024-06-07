package number

import (
	"testing"
)

func TestArabicNumber(t *testing.T) {
	num := ArabicNumber(5)
	if n, _ := num.ToInt(); n != 5 {
		t.Errorf("Expected 5, but got %v", n)
	}
	if s := num.ToString(); s != "5" {
		t.Errorf("Expected '5', but got %v", s)
	}
}

func TestRomanNumber(t *testing.T) {
	num := RomanNumber("X")
	if n, _ := num.ToInt(); n != 10 {
		t.Errorf("Expected 10, but got %v", n)
	}
	if s := num.ToString(); s != "X" {
		t.Errorf("Expected 'X', but got %v", s)
	}
}

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		input    int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{6, "VI"},
		{10, "X"},
		{40, "XL"},
		{90, "XC"},
		{400, "CD"},
		{900, "CM"},
		{1000, "M"},
	}

	for _, c := range cases {
		if s := IntToRoman(c.input); s != c.expected {
			t.Errorf("Expected %v, but got %v", c.expected, s)
		}
	}
}

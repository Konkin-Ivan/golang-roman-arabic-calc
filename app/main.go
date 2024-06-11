package main

import (
	"Go_test/calculator"
	"Go_test/calculator/number"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	firstOperand, mathOperator, lastOperand := getOperandsAndOperator()

	left, leftIsRoman, err := parseNumber(firstOperand)
	if err != nil {
		panic(err)
	}

	right, rightIsRoman, err := parseNumber(lastOperand)
	if err != nil {
		panic(err)
	}

	if !areSameNumericSystem(leftIsRoman, rightIsRoman) {
		panic("Invalid input: mixed numeric systems are not allowed.")
	}

	calc := calculator.Calculator{left, right, mathOperator}
	result, err := calc.Calculate()
	if err != nil {
		panic(err)
	}

	resultInt, _ := result.ToInt()
	if resultInt < 0 && leftIsRoman {
		panic("Invalid result: negative numbers are not allowed in the Roman numeral system.")
	}

	if leftIsRoman {
		fmt.Println(number.IntToRoman(resultInt))
	} else {
		fmt.Println(result.ToString())
	}
}

func getOperandsAndOperator() (string, string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an expression (e.g., 5 + 3 or VII - II): ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Invalid input: the mathematical expression must contain two operands and one operator (+, -, /, *).")
	}

	isValid := validateInput(parts[0], parts[2])
	if isValid == false {
		panic(fmt.Sprintf("Invalid number: %s", parts))
	}

	return parts[0], parts[1], parts[2]
}

func parseNumber(input string) (number.Number, bool, error) {
	if isRomanNumber(input) {
		if !number.IsRomanNumberValid(input, 10) {
			return nil, false, fmt.Errorf("Roman numbers must not exceed X (10).")
		}
		return number.RomanNumber(input), true, nil
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return nil, false, err
	}
	return number.ArabicNumber(num), false, nil
}

func isRomanNumber(num string) bool {
	return strings.ContainsAny(num, "IVXLCDM")
}

func areSameNumericSystem(leftIsRoman, rightIsRoman bool) bool {
	return leftIsRoman == rightIsRoman
}

func validateInput(firstNumber string, lastNumber string) bool {
	validCombinations := map[string]bool{
		"I": true, "II": true, "III": true, "IV": true, "V": true, "VI": true, "VII": true, "VIII": true, "IX": true, "X": true,
		"XI": true, "XII": true, "XIII": true, "XIV": true, "XV": true, "XVI": true, "XVII": true, "XVIII": true, "XIX": true, "XX": true,
		"XXI": true, "XXII": true, "XXIII": true, "XXIV": true, "XXV": true, "XXVI": true, "XXVII": true, "XXVIII": true, "XXIX": true, "XXX": true,
		"XXXI": true, "XXXII": true, "XXXIII": true, "XXXIV": true, "XXXV": true, "XXXVI": true, "XXXVII": true, "XXXVIII": true, "XXXIX": true, "XL": true,
		"XLI": true, "XLII": true, "XLIII": true, "XLIV": true, "XLV": true, "XLVI": true, "XLVII": true, "XLVIII": true, "XLIX": true, "L": true,
		"LI": true, "LII": true, "LIII": true, "LIV": true, "LV": true, "LVI": true, "LVII": true, "LVIII": true, "LIX": true, "LX": true,
		"LXI": true, "LXII": true, "LXIII": true, "LXIV": true, "LXV": true, "LXVI": true, "LXVII": true, "LXVIII": true, "LXIX": true, "LXX": true,
		"LXXI": true, "LXXII": true, "LXXIII": true, "LXXIV": true, "LXXV": true, "LXXVI": true, "LXXVII": true, "LXXVIII": true, "LXXIX": true, "LXXX": true,
		"LXXXI": true, "LXXXII": true, "LXXXIII": true, "LXXXIV": true, "LXXXV": true, "LXXXVI": true, "LXXXVII": true, "LXXXVIII": true, "LXXXIX": true, "XC": true,
		"XCI": true, "XCII": true, "XCIII": true, "XCIV": true, "XCV": true, "XCVI": true, "XCVII": true, "XCVIII": true, "XCIX": true, "C": true,
	}

	if _, ok := validCombinations[firstNumber]; ok {
		if _, ok := validCombinations[lastNumber]; ok {
			return true
		}

		num, err := strconv.Atoi(lastNumber)
		return err == nil && num >= 1 && num <= 10
	}

	firstNum, err := strconv.Atoi(firstNumber)
	if err != nil {
		return false
	}

	return firstNum >= 1 && firstNum <= 10
}

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
		"I": true, "II": true, "III": true, "IV": true, "V": true, "IX": true,
		"X": true, "XL": true, "L": true, "XC": true,
		"C": true, "CD": true, "D": true, "CM": true, "M": true,
	}

	if _, ok := validCombinations[firstNumber]; ok {
		if _, ok := validCombinations[lastNumber]; ok {
			return true
		}

		num, err := strconv.Atoi(lastNumber)
		return err == nil && num >= 0 && num <= 9
	}

	firstNum, err := strconv.Atoi(firstNumber)
	if err != nil {
		return false
	}

	return firstNum >= 0 && firstNum <= 9
}

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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter an expression (e.g., 5 + 3 or VII - II): ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Invalid input: the mathematical expression must contain two operands and one operator (+, -, /, *).")
	}

	left, leftIsRoman, err := parseNumber(parts[0])
	if err != nil {
		panic(err)
	}

	right, rightIsRoman, err := parseNumber(parts[2])
	if err != nil {
		panic(err)
	}

	// Check if both numbers are from the same numeric system
	if !areSameNumericSystem(leftIsRoman, rightIsRoman) {
		panic("Invalid input: mixed numeric systems are not allowed.")
	}

	calc := calculator.Calculator{left, right, parts[1]}
	result, err := calc.Calculate()
	if err != nil {
		panic(err)
	}

	// If result is negative and numbers are Roman, panic
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

func parseNumber(input string) (number.Number, bool, error) {
	maxValue := 10
	if isRomanNumber(input) {
		if !number.IsRomanNumberValid(input, maxValue) {
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

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

	var left, right number.Number
	var leftIsRoman, rightIsRoman bool

	if isRomanNumber(parts[0]) {
		left = number.RomanNumber(parts[0])
		leftIsRoman = true
	} else {
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		left = number.ArabicNumber(num)
	}

	if isRomanNumber(parts[2]) {
		right = number.RomanNumber(parts[2])
		rightIsRoman = true
	} else {
		num, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		right = number.ArabicNumber(num)
	}

	// Check if both numbers are from the same numeric system
	if leftIsRoman != rightIsRoman {
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

func isRomanNumber(num string) bool {
	return strings.ContainsAny(num, "IVXLCDM")
}

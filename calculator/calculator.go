package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// working with golang calculator
/* Basics functionality of a calculator
1. allow user enters two numbers and store into their variables (firstNumber and secondNumber)
2. suggest for operation type(addition, subtraction, multiplication, division)
3. perform operation based on the type of operator */

// addition
func addition(x, y int) int {
	return x + y
}

// subtraction
func subtraction(x, y float64) float64 {
	return x - y
}

// multiplication
func multiplication(x, y int) int {
	return x * y
}

func division(x, y float64) float64 {
	if y != 0 {
		return x / y
	}
	return 1
}

func main() {
	// first number
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter first number: ")
	firstNumberInput, _ := reader.ReadString('\n')
	firstNumber, _ := strconv.Atoi(firstNumberInput)

	// get second number
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter second number: ")
	secondNumberInput, _ := reader2.ReadString('\n')
	secondNumber, _ := strconv.Atoi(secondNumberInput)

	resultAdd := addition(firstNumber, secondNumber)
	fmt.Println("resultAdd", resultAdd)

	resultSubtraction := subtraction(float64(firstNumber), float64(secondNumber))
	fmt.Println("resultSubtraction", resultSubtraction)

	resultMultiplication := multiplication(firstNumber, secondNumber)
	fmt.Println("resultMultiplication", resultMultiplication)

	resultDivision := division(float64(firstNumber), float64(secondNumber))
	fmt.Println("resultDivision", resultDivision)
}
